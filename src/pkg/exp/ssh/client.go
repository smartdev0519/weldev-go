// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssh

import (
	"big"
	"crypto"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"net"
	"sync"
)

// clientVersion is the fixed identification string that the client will use.
var clientVersion = []byte("SSH-2.0-Go\r\n")

// ClientConn represents the client side of an SSH connection.
type ClientConn struct {
	*transport
	config *ClientConfig
	chanlist
}

// Client returns a new SSH client connection using c as the underlying transport.
func Client(c net.Conn, config *ClientConfig) (*ClientConn, os.Error) {
	conn := &ClientConn{
		transport: newTransport(c, config.rand()),
		config:    config,
	}
	if err := conn.handshake(); err != nil {
		conn.Close()
		return nil, err
	}
	if err := conn.authenticate(); err != nil {
		conn.Close()
		return nil, err
	}
	go conn.mainLoop()
	return conn, nil
}

// handshake performs the client side key exchange. See RFC 4253 Section 7.
func (c *ClientConn) handshake() os.Error {
	var magics handshakeMagics

	if _, err := c.Write(clientVersion); err != nil {
		return err
	}
	if err := c.Flush(); err != nil {
		return err
	}
	magics.clientVersion = clientVersion[:len(clientVersion)-2]

	// read remote server version
	version, err := readVersion(c)
	if err != nil {
		return err
	}
	magics.serverVersion = version
	clientKexInit := kexInitMsg{
		KexAlgos:                supportedKexAlgos,
		ServerHostKeyAlgos:      supportedHostKeyAlgos,
		CiphersClientServer:     supportedCiphers,
		CiphersServerClient:     supportedCiphers,
		MACsClientServer:        supportedMACs,
		MACsServerClient:        supportedMACs,
		CompressionClientServer: supportedCompressions,
		CompressionServerClient: supportedCompressions,
	}
	kexInitPacket := marshal(msgKexInit, clientKexInit)
	magics.clientKexInit = kexInitPacket

	if err := c.writePacket(kexInitPacket); err != nil {
		return err
	}
	packet, err := c.readPacket()
	if err != nil {
		return err
	}

	magics.serverKexInit = packet

	var serverKexInit kexInitMsg
	if err = unmarshal(&serverKexInit, packet, msgKexInit); err != nil {
		return err
	}

	kexAlgo, hostKeyAlgo, ok := findAgreedAlgorithms(c.transport, &clientKexInit, &serverKexInit)
	if !ok {
		return os.NewError("ssh: no common algorithms")
	}

	if serverKexInit.FirstKexFollows && kexAlgo != serverKexInit.KexAlgos[0] {
		// The server sent a Kex message for the wrong algorithm,
		// which we have to ignore.
		if _, err := c.readPacket(); err != nil {
			return err
		}
	}

	var H, K []byte
	var hashFunc crypto.Hash
	switch kexAlgo {
	case kexAlgoDH14SHA1:
		hashFunc = crypto.SHA1
		dhGroup14Once.Do(initDHGroup14)
		H, K, err = c.kexDH(dhGroup14, hashFunc, &magics, hostKeyAlgo)
	default:
		err = fmt.Errorf("ssh: unexpected key exchange algorithm %v", kexAlgo)
	}
	if err != nil {
		return err
	}

	if err = c.writePacket([]byte{msgNewKeys}); err != nil {
		return err
	}
	if err = c.transport.writer.setupKeys(clientKeys, K, H, H, hashFunc); err != nil {
		return err
	}
	if packet, err = c.readPacket(); err != nil {
		return err
	}
	if packet[0] != msgNewKeys {
		return UnexpectedMessageError{msgNewKeys, packet[0]}
	}
	return c.transport.reader.setupKeys(serverKeys, K, H, H, hashFunc)
}

// authenticate authenticates with the remote server. See RFC 4252. 
// Only "password" authentication is supported.
func (c *ClientConn) authenticate() os.Error {
	if err := c.writePacket(marshal(msgServiceRequest, serviceRequestMsg{serviceUserAuth})); err != nil {
		return err
	}
	packet, err := c.readPacket()
	if err != nil {
		return err
	}

	var serviceAccept serviceAcceptMsg
	if err = unmarshal(&serviceAccept, packet, msgServiceAccept); err != nil {
		return err
	}

	// TODO(dfc) support proper authentication method negotation
	method := "none"
	if c.config.Password != "" {
		method = "password"
	}
	if err := c.sendUserAuthReq(method); err != nil {
		return err
	}

	if packet, err = c.readPacket(); err != nil {
		return err
	}

	if packet[0] != msgUserAuthSuccess {
		return UnexpectedMessageError{msgUserAuthSuccess, packet[0]}
	}
	return nil
}

func (c *ClientConn) sendUserAuthReq(method string) os.Error {
	length := stringLength([]byte(c.config.Password)) + 1
	payload := make([]byte, length)
	// always false for password auth, see RFC 4252 Section 8.
	payload[0] = 0
	marshalString(payload[1:], []byte(c.config.Password))

	return c.writePacket(marshal(msgUserAuthRequest, userAuthRequestMsg{
		User:    c.config.User,
		Service: serviceSSH,
		Method:  method,
		Payload: payload,
	}))
}

// kexDH performs Diffie-Hellman key agreement on a ClientConn. The
// returned values are given the same names as in RFC 4253, section 8.
func (c *ClientConn) kexDH(group *dhGroup, hashFunc crypto.Hash, magics *handshakeMagics, hostKeyAlgo string) ([]byte, []byte, os.Error) {
	x, err := rand.Int(c.config.rand(), group.p)
	if err != nil {
		return nil, nil, err
	}
	X := new(big.Int).Exp(group.g, x, group.p)
	kexDHInit := kexDHInitMsg{
		X: X,
	}
	if err := c.writePacket(marshal(msgKexDHInit, kexDHInit)); err != nil {
		return nil, nil, err
	}

	packet, err := c.readPacket()
	if err != nil {
		return nil, nil, err
	}

	var kexDHReply = new(kexDHReplyMsg)
	if err = unmarshal(kexDHReply, packet, msgKexDHReply); err != nil {
		return nil, nil, err
	}

	if kexDHReply.Y.Sign() == 0 || kexDHReply.Y.Cmp(group.p) >= 0 {
		return nil, nil, os.NewError("server DH parameter out of bounds")
	}

	kInt := new(big.Int).Exp(kexDHReply.Y, x, group.p)
	h := hashFunc.New()
	writeString(h, magics.clientVersion)
	writeString(h, magics.serverVersion)
	writeString(h, magics.clientKexInit)
	writeString(h, magics.serverKexInit)
	writeString(h, kexDHReply.HostKey)
	writeInt(h, X)
	writeInt(h, kexDHReply.Y)
	K := make([]byte, intLength(kInt))
	marshalInt(K, kInt)
	h.Write(K)

	H := h.Sum()

	return H, K, nil
}

// openChan opens a new client channel. The most common session type is "session". 
// The full set of valid session types are listed in RFC 4250 4.9.1.
func (c *ClientConn) openChan(typ string) (*clientChan, os.Error) {
	ch := c.newChan(c.transport)
	if err := c.writePacket(marshal(msgChannelOpen, channelOpenMsg{
		ChanType:      typ,
		PeersId:       ch.id,
		PeersWindow:   1 << 14,
		MaxPacketSize: 1 << 15, // RFC 4253 6.1
	})); err != nil {
		c.chanlist.remove(ch.id)
		return nil, err
	}
	// wait for response
	switch msg := (<-ch.msg).(type) {
	case *channelOpenConfirmMsg:
		ch.peersId = msg.MyId
	case *channelOpenFailureMsg:
		c.chanlist.remove(ch.id)
		return nil, os.NewError(msg.Message)
	default:
		c.chanlist.remove(ch.id)
		return nil, os.NewError("Unexpected packet")
	}
	return ch, nil
}

// mainloop reads incoming messages and routes channel messages
// to their respective ClientChans.
func (c *ClientConn) mainLoop() {
	for {
		packet, err := c.readPacket()
		if err != nil {
			// TODO(dfc) signal the underlying close to all channels
			c.Close()
			return
		}
		// TODO(dfc) A note on blocking channel use. 
		// The msg, win, data and dataExt channels of a clientChan can 
		// cause this loop to block indefinately if the consumer does 
		// not service them. 
		switch msg := decode(packet).(type) {
		case *channelOpenMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *channelOpenConfirmMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *channelOpenFailureMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *channelCloseMsg:
			ch := c.getChan(msg.PeersId)
			close(ch.win)
			close(ch.data)
			close(ch.dataExt)
			c.chanlist.remove(msg.PeersId)
		case *channelEOFMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *channelRequestSuccessMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *channelRequestFailureMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *channelRequestMsg:
			c.getChan(msg.PeersId).msg <- msg
		case *windowAdjustMsg:
			c.getChan(msg.PeersId).win <- int(msg.AdditionalBytes)
		case *channelData:
			c.getChan(msg.PeersId).data <- msg.Payload
		case *channelExtendedData:
			// RFC 4254 5.2 defines data_type_code 1 to be data destined 
			// for stderr on interactive sessions. Other data types are
			// silently discarded.
			if msg.Datatype == 1 {
				c.getChan(msg.PeersId).dataExt <- msg.Payload
			}
		default:
			fmt.Printf("mainLoop: unhandled %#v\n", msg)
		}
	}
}

// Dial connects to the given network address using net.Dial and 
// then initiates a SSH handshake, returning the resulting client connection.
func Dial(network, addr string, config *ClientConfig) (*ClientConn, os.Error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return Client(conn, config)
}

// A ClientConfig structure is used to configure a ClientConn. After one has 
// been passed to an SSH function it must not be modified.
type ClientConfig struct {
	// Rand provides the source of entropy for key exchange. If Rand is 
	// nil, the cryptographic random reader in package crypto/rand will 
	// be used.
	Rand io.Reader

	// The username to authenticate.
	User string

	// Used for "password" method authentication.
	Password string
}

func (c *ClientConfig) rand() io.Reader {
	if c.Rand == nil {
		return rand.Reader
	}
	return c.Rand
}

// A clientChan represents a single RFC 4254 channel that is multiplexed 
// over a single SSH connection.
type clientChan struct {
	packetWriter
	id, peersId uint32
	data        chan []byte      // receives the payload of channelData messages
	dataExt     chan []byte      // receives the payload of channelExtendedData messages
	win         chan int         // receives window adjustments
	msg         chan interface{} // incoming messages
}

func newClientChan(t *transport, id uint32) *clientChan {
	return &clientChan{
		packetWriter: t,
		id:           id,
		data:         make(chan []byte, 16),
		dataExt:      make(chan []byte, 16),
		win:          make(chan int, 16),
		msg:          make(chan interface{}, 16),
	}
}

// Close closes the channel. This does not close the underlying connection.
func (c *clientChan) Close() os.Error {
	return c.writePacket(marshal(msgChannelClose, channelCloseMsg{
		PeersId: c.id,
	}))
}

func (c *clientChan) sendChanReq(req channelRequestMsg) os.Error {
	if err := c.writePacket(marshal(msgChannelRequest, req)); err != nil {
		return err
	}
	msg := <-c.msg
	if _, ok := msg.(*channelRequestSuccessMsg); ok {
		return nil
	}
	return fmt.Errorf("failed to complete request: %s, %#v", req.Request, msg)
}

// Thread safe channel list.
type chanlist struct {
	// protects concurrent access to chans
	sync.Mutex
	// chans are indexed by the local id of the channel, clientChan.id.
	// The PeersId value of messages received by ClientConn.mainloop is
	// used to locate the right local clientChan in this slice.
	chans []*clientChan
}

// Allocate a new ClientChan with the next avail local id.
func (c *chanlist) newChan(t *transport) *clientChan {
	c.Lock()
	defer c.Unlock()
	for i := range c.chans {
		if c.chans[i] == nil {
			ch := newClientChan(t, uint32(i))
			c.chans[i] = ch
			return ch
		}
	}
	i := len(c.chans)
	ch := newClientChan(t, uint32(i))
	c.chans = append(c.chans, ch)
	return ch
}

func (c *chanlist) getChan(id uint32) *clientChan {
	c.Lock()
	defer c.Unlock()
	return c.chans[int(id)]
}

func (c *chanlist) remove(id uint32) {
	c.Lock()
	defer c.Unlock()
	c.chans[int(id)] = nil
}

// A chanWriter represents the stdin of a remote process.
type chanWriter struct {
	win          chan int // receives window adjustments
	id           uint32   // this channel's id
	rwin         int      // current rwin size
	packetWriter          // for sending channelDataMsg
}

// Write writes data to the remote process's standard input.
func (w *chanWriter) Write(data []byte) (n int, err os.Error) {
	for {
		if w.rwin == 0 {
			win, ok := <-w.win
			if !ok {
				return 0, os.EOF
			}
			w.rwin += win
			continue
		}
		n = len(data)
		packet := make([]byte, 0, 9+n)
		packet = append(packet, msgChannelData,
			byte(w.id)>>24, byte(w.id)>>16, byte(w.id)>>8, byte(w.id),
			byte(n)>>24, byte(n)>>16, byte(n)>>8, byte(n))
		err = w.writePacket(append(packet, data...))
		w.rwin -= n
		return
	}
	panic("unreachable")
}

func (w *chanWriter) Close() os.Error {
	return w.writePacket(marshal(msgChannelEOF, channelEOFMsg{w.id}))
}

// A chanReader represents stdout or stderr of a remote process.
type chanReader struct {
	// TODO(dfc) a fixed size channel may not be the right data structure.
	// If writes to this channel block, they will block mainLoop, making
	// it unable to receive new messages from the remote side.
	data         chan []byte // receives data from remote
	id           uint32
	packetWriter // for sending windowAdjustMsg
	buf          []byte
}

// Read reads data from the remote process's stdout or stderr.
func (r *chanReader) Read(data []byte) (int, os.Error) {
	var ok bool
	for {
		if len(r.buf) > 0 {
			n := copy(data, r.buf)
			r.buf = r.buf[n:]
			msg := windowAdjustMsg{
				PeersId:         r.id,
				AdditionalBytes: uint32(n),
			}
			return n, r.writePacket(marshal(msgChannelWindowAdjust, msg))
		}
		r.buf, ok = <-r.data
		if !ok {
			return 0, os.EOF
		}
	}
	panic("unreachable")
}

func (r *chanReader) Close() os.Error {
	return r.writePacket(marshal(msgChannelEOF, channelEOFMsg{r.id}))
}
