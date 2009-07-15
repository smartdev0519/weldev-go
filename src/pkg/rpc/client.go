// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"bufio";
	"gob";
	"http";
	"io";
	"log";
	"net";
	"os";
	"rpc";
	"strconv";
	"sync";
)

// Call represents an active RPC
type Call struct {
	ServiceMethod	string;	// The name of the service and method to call.
	Args	interface{};	// The argument to the function (*struct).
	Reply	interface{};	// The reply from the function (*struct).
	Error	os.Error;	// After completion, the error status.
	Done	chan *Call;	// Strobes when call is complete; value is the error status.
	seq	uint64;
}

// Client represents an RPC Client.
type Client struct {
	sync.Mutex;	// protects pending, seq
	shutdown	os.Error;	// non-nil if the client is shut down
	sending	sync.Mutex;
	seq	uint64;
	conn io.ReadWriteCloser;
	enc	*gob.Encoder;
	dec	*gob.Decoder;
	pending	map[uint64] *Call;
}

func (client *Client) send(c *Call) {
	// Register this call.
	client.Lock();
	if client.shutdown != nil {
		c.Error = client.shutdown;
		client.Unlock();
		doNotBlock := c.Done <- c;
		return;
	}
	c.seq = client.seq;
	client.seq++;
	client.pending[c.seq] = c;
	client.Unlock();

	// Encode and send the request.
	request := new(Request);
	client.sending.Lock();
	request.Seq = c.seq;
	request.ServiceMethod = c.ServiceMethod;
	client.enc.Encode(request);
	err := client.enc.Encode(c.Args);
	if err != nil {
		panicln("rpc: client encode error:", err);
	}
	client.sending.Unlock();
}

func (client *Client) input() {
	var err os.Error;
	for err == nil {
		response := new(Response);
		err = client.dec.Decode(response);
		if err != nil {
			if err == os.EOF {
				err = io.ErrUnexpectedEOF;
			}
			break
		}
		seq := response.Seq;
		client.Lock();
		c := client.pending[seq];
		client.pending[seq] = c, false;
		client.Unlock();
		err = client.dec.Decode(c.Reply);
		c.Error = os.ErrorString(response.Error);
		// We don't want to block here.  It is the caller's responsibility to make
		// sure the channel has enough buffer space. See comment in Go().
		doNotBlock := c.Done <- c;
	}
	// Terminate pending calls.
	client.Lock();
	client.shutdown = err;
	for seq, call := range client.pending {
		call.Error = err;
		doNotBlock := call.Done <- call;
	}
	client.Unlock();
	log.Stderr("client protocol error:", err);
}

// NewClient returns a new Client to handle requests to the
// set of services at the other end of the connection.
func NewClient(conn io.ReadWriteCloser) *Client {
	client := new(Client);
	client.conn = conn;
	client.enc = gob.NewEncoder(conn);
	client.dec = gob.NewDecoder(conn);
	client.pending = make(map[uint64] *Call);
	go client.input();
	return client;
}

// Dial connects to an HTTP RPC server at the specified network address.
func DialHTTP(network, address string) (*Client, os.Error) {
	conn, err := net.Dial(network, "", address);
	if err != nil {
		return nil, err
	}
	io.WriteString(conn, "CONNECT " + rpcPath + " HTTP/1.0\n\n");

	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn));
	if err == nil && resp.Status == connected {
		return NewClient(conn), nil;
	}
	if err == nil {
		err = os.ErrorString("unexpected HTTP response: " + resp.Status);
	}
	conn.Close();
	return nil, &net.OpError{"dial-http", network, address, err};
}

// Dial connects to an RPC server at the specified network address.
func Dial(network, address string) (*Client, os.Error) {
	conn, err := net.Dial(network, "", address);
	if err != nil {
		return nil, err
	}
	return NewClient(conn), nil;
}

// Go invokes the function asynchronously.  It returns the Call structure representing
// the invocation.
func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
	c := new(Call);
	c.ServiceMethod = serviceMethod;
	c.Args = args;
	c.Reply = reply;
	if done == nil {
		done = make(chan *Call, 1);	// buffered.
	} else {
		// TODO(r): check cap > 0
		// If caller passes done != nil, it must arrange that
		// done has enough buffer for the number of simultaneous
		// RPCs that will be using that channel.
	}
	c.Done = done;
	if client.shutdown != nil {
		c.Error = client.shutdown;
		doNotBlock := c.Done <- c;
		return c;
	}
	client.send(c);
	return c;
}

// Call invokes the named function, waits for it to complete, and returns its error status.
func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) os.Error {
	if client.shutdown != nil {
		return client.shutdown
	}
	call := <-client.Go(serviceMethod, args, reply, nil).Done;
	return call.Error;
}
