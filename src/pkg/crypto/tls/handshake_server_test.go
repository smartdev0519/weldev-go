// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"big"
	"bytes"
	"crypto/rsa"
	"encoding/hex"
	"flag"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
)

type zeroSource struct{}

func (zeroSource) Read(b []byte) (n int, err os.Error) {
	for i := range b {
		b[i] = 0
	}

	return len(b), nil
}

var testConfig *Config

func init() {
	testConfig = new(Config)
	testConfig.Time = func() int64 { return 0 }
	testConfig.Rand = zeroSource{}
	testConfig.Certificates = make([]Certificate, 1)
	testConfig.Certificates[0].Certificate = [][]byte{testCertificate}
	testConfig.Certificates[0].PrivateKey = testPrivateKey
	testConfig.CipherSuites = []uint16{TLS_RSA_WITH_RC4_128_SHA}
}

func testClientHelloFailure(t *testing.T, m handshakeMessage, expected os.Error) {
	// Create in-memory network connection,
	// send message to server.  Should return
	// expected error.
	c, s := net.Pipe()
	go func() {
		cli := Client(c, testConfig)
		if ch, ok := m.(*clientHelloMsg); ok {
			cli.vers = ch.vers
		}
		cli.writeRecord(recordTypeHandshake, m.marshal())
		c.Close()
	}()
	err := Server(s, testConfig).Handshake()
	s.Close()
	if e, ok := err.(*net.OpError); !ok || e.Error != expected {
		t.Errorf("Got error: %s; expected: %s", err, expected)
	}
}

func TestSimpleError(t *testing.T) {
	testClientHelloFailure(t, &serverHelloDoneMsg{}, alertUnexpectedMessage)
}

var badProtocolVersions = []uint16{0x0000, 0x0005, 0x0100, 0x0105, 0x0200, 0x0205}

func TestRejectBadProtocolVersion(t *testing.T) {
	for _, v := range badProtocolVersions {
		testClientHelloFailure(t, &clientHelloMsg{vers: v}, alertProtocolVersion)
	}
}

func TestNoSuiteOverlap(t *testing.T) {
	clientHello := &clientHelloMsg{nil, 0x0301, nil, nil, []uint16{0xff00}, []uint8{0}, false, "", false, nil, nil}
	testClientHelloFailure(t, clientHello, alertHandshakeFailure)

}

func TestNoCompressionOverlap(t *testing.T) {
	clientHello := &clientHelloMsg{nil, 0x0301, nil, nil, []uint16{TLS_RSA_WITH_RC4_128_SHA}, []uint8{0xff}, false, "", false, nil, nil}
	testClientHelloFailure(t, clientHello, alertHandshakeFailure)
}

func TestAlertForwarding(t *testing.T) {
	c, s := net.Pipe()
	go func() {
		Client(c, testConfig).sendAlert(alertUnknownCA)
		c.Close()
	}()

	err := Server(s, testConfig).Handshake()
	s.Close()
	if e, ok := err.(*net.OpError); !ok || e.Error != os.Error(alertUnknownCA) {
		t.Errorf("Got error: %s; expected: %s", err, alertUnknownCA)
	}
}

func TestClose(t *testing.T) {
	c, s := net.Pipe()
	go c.Close()

	err := Server(s, testConfig).Handshake()
	s.Close()
	if err != os.EOF {
		t.Errorf("Got error: %s; expected: %s", err, os.EOF)
	}
}

func testServerScript(t *testing.T, name string, serverScript [][]byte, config *Config) {
	c, s := net.Pipe()
	srv := Server(s, config)
	go func() {
		srv.Write([]byte("hello, world\n"))
		srv.Close()
		s.Close()
	}()

	defer c.Close()
	for i, b := range serverScript {
		if i%2 == 0 {
			c.Write(b)
			continue
		}
		bb := make([]byte, len(b))
		n, err := io.ReadFull(c, bb)
		if err != nil {
			t.Fatalf("%s #%d: %s\nRead %d, wanted %d, got %x, wanted %x\n", name, i, err, n, len(bb), bb[:n], b)
		}
		if !bytes.Equal(b, bb) {
			t.Fatalf("%s #%d: mismatch on read: got:%x want:%x", name, i, bb, b)
		}
	}
}

func TestHandshakeServerRC4(t *testing.T) {
	testServerScript(t, "RC4", rc4ServerScript, testConfig)
}

func TestHandshakeServer3DES(t *testing.T) {
	des3Config := new(Config)
	*des3Config = *testConfig
	des3Config.CipherSuites = []uint16{TLS_RSA_WITH_3DES_EDE_CBC_SHA}
	testServerScript(t, "3DES", des3ServerScript, des3Config)
}

func TestHandshakeServerAES(t *testing.T) {
	aesConfig := new(Config)
	*aesConfig = *testConfig
	aesConfig.CipherSuites = []uint16{TLS_RSA_WITH_AES_128_CBC_SHA}
	testServerScript(t, "AES", aesServerScript, aesConfig)
}

func TestHandshakeServerSSLv3(t *testing.T) {
	testServerScript(t, "SSLv3", sslv3ServerScript, testConfig)
}

var serve = flag.Bool("serve", false, "run a TLS server on :10443")
var testCipherSuites = flag.String("ciphersuites",
	"0x"+strconv.Itob(int(TLS_RSA_WITH_RC4_128_SHA), 16),
	"cipher suites to accept in serving mode")

func TestRunServer(t *testing.T) {
	if !*serve {
		return
	}

	suites := strings.Split(*testCipherSuites, ",")
	testConfig.CipherSuites = make([]uint16, len(suites))
	for i := range suites {
		suite, err := strconv.Btoui64(suites[i], 0)
		if err != nil {
			panic(err)
		}
		testConfig.CipherSuites[i] = uint16(suite)
	}

	l, err := Listen("tcp", ":10443", testConfig)
	if err != nil {
		t.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			break
		}
		_, err = c.Write([]byte("hello, world\n"))
		if err != nil {
			t.Errorf("error from TLS: %s", err)
			break
		}
		c.Close()
	}
}

func bigFromString(s string) *big.Int {
	ret := new(big.Int)
	ret.SetString(s, 10)
	return ret
}

func fromHex(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}

var testCertificate = fromHex("308202b030820219a00302010202090085b0bba48a7fb8ca300d06092a864886f70d01010505003045310b3009060355040613024155311330110603550408130a536f6d652d53746174653121301f060355040a1318496e7465726e6574205769646769747320507479204c7464301e170d3130303432343039303933385a170d3131303432343039303933385a3045310b3009060355040613024155311330110603550408130a536f6d652d53746174653121301f060355040a1318496e7465726e6574205769646769747320507479204c746430819f300d06092a864886f70d010101050003818d0030818902818100bb79d6f517b5e5bf4610d0dc69bee62b07435ad0032d8a7a4385b71452e7a5654c2c78b8238cb5b482e5de1f953b7e62a52ca533d6fe125c7a56fcf506bffa587b263fb5cd04d3d0c921964ac7f4549f5abfef427100fe1899077f7e887d7df10439c4a22edb51c97ce3c04c3b326601cfafb11db8719a1ddbdb896baeda2d790203010001a381a73081a4301d0603551d0e04160414b1ade2855acfcb28db69ce2369ded3268e18883930750603551d23046e306c8014b1ade2855acfcb28db69ce2369ded3268e188839a149a4473045310b3009060355040613024155311330110603550408130a536f6d652d53746174653121301f060355040a1318496e7465726e6574205769646769747320507479204c746482090085b0bba48a7fb8ca300c0603551d13040530030101ff300d06092a864886f70d010105050003818100086c4524c76bb159ab0c52ccf2b014d7879d7a6475b55a9566e4c52b8eae12661feb4f38b36e60d392fdf74108b52513b1187a24fb301dbaed98b917ece7d73159db95d31d78ea50565cd5825a2d5a5f33c4b6d8c97590968c0f5298b5cd981f89205ff2a01ca31b9694dda9fd57e970e8266d71999b266e3850296c90a7bdd9")

var testPrivateKey = &rsa.PrivateKey{
	PublicKey: rsa.PublicKey{
		N: bigFromString("131650079503776001033793877885499001334664249354723305978524647182322416328664556247316495448366990052837680518067798333412266673813370895702118944398081598789828837447552603077848001020611640547221687072142537202428102790818451901395596882588063427854225330436740647715202971973145151161964464812406232198521"),
		E: 65537,
	},
	D: bigFromString("29354450337804273969007277378287027274721892607543397931919078829901848876371746653677097639302788129485893852488285045793268732234230875671682624082413996177431586734171663258657462237320300610850244186316880055243099640544518318093544057213190320837094958164973959123058337475052510833916491060913053867729"),
	Primes: []*big.Int{
		bigFromString("11969277782311800166562047708379380720136961987713178380670422671426759650127150688426177829077494755200794297055316163155755835813760102405344560929062149"),
		bigFromString("10998999429884441391899182616418192492905073053684657075974935218461686523870125521822756579792315215543092255516093840728890783887287417039645833477273829"),
	},
}

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in server mode:
//   % gotest -test.run "TestRunServer" -serve
// and then:
//   % gnutls-cli --insecure --debug 100 -p 10443 localhost > /tmp/log 2>&1
//   % python parse-gnutls-cli-debug-log.py < /tmp/log
var rc4ServerScript = [][]byte{
	{
		0x16, 0x03, 0x02, 0x00, 0x7f, 0x01, 0x00, 0x00,
		0x7b, 0x03, 0x02, 0x4d, 0x08, 0x1f, 0x5a, 0x7a,
		0x0a, 0x92, 0x2f, 0xf0, 0x73, 0x16, 0x3a, 0x88,
		0x14, 0x85, 0x4c, 0x98, 0x15, 0x7b, 0x65, 0xe0,
		0x78, 0xd0, 0xed, 0xd0, 0xf3, 0x65, 0x20, 0xeb,
		0x80, 0xd1, 0x0b, 0x00, 0x00, 0x34, 0x00, 0x33,
		0x00, 0x45, 0x00, 0x39, 0x00, 0x88, 0x00, 0x16,
		0x00, 0x32, 0x00, 0x44, 0x00, 0x38, 0x00, 0x87,
		0x00, 0x13, 0x00, 0x66, 0x00, 0x90, 0x00, 0x91,
		0x00, 0x8f, 0x00, 0x8e, 0x00, 0x2f, 0x00, 0x41,
		0x00, 0x35, 0x00, 0x84, 0x00, 0x0a, 0x00, 0x05,
		0x00, 0x04, 0x00, 0x8c, 0x00, 0x8d, 0x00, 0x8b,
		0x00, 0x8a, 0x01, 0x00, 0x00, 0x1e, 0x00, 0x09,
		0x00, 0x03, 0x02, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x0e, 0x00, 0x0c, 0x00, 0x00, 0x09, 0x6c, 0x6f,
		0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0xff,
		0x01, 0x00, 0x01, 0x00,
	},

	{
		0x16, 0x03, 0x01, 0x00, 0x2a, 0x02, 0x00, 0x00,
		0x26, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x16,
		0x03, 0x01, 0x02, 0xbe, 0x0b, 0x00, 0x02, 0xba,
		0x00, 0x02, 0xb7, 0x00, 0x02, 0xb4, 0x30, 0x82,
		0x02, 0xb0, 0x30, 0x82, 0x02, 0x19, 0xa0, 0x03,
		0x02, 0x01, 0x02, 0x02, 0x09, 0x00, 0x85, 0xb0,
		0xbb, 0xa4, 0x8a, 0x7f, 0xb8, 0xca, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x05, 0x05, 0x00, 0x30, 0x45, 0x31,
		0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06,
		0x13, 0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11,
		0x06, 0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53,
		0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74,
		0x65, 0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55,
		0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65,
		0x72, 0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64,
		0x67, 0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79,
		0x20, 0x4c, 0x74, 0x64, 0x30, 0x1e, 0x17, 0x0d,
		0x31, 0x30, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39,
		0x30, 0x39, 0x33, 0x38, 0x5a, 0x17, 0x0d, 0x31,
		0x31, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39, 0x30,
		0x39, 0x33, 0x38, 0x5a, 0x30, 0x45, 0x31, 0x0b,
		0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06, 0x13,
		0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11, 0x06,
		0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53, 0x6f,
		0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74, 0x65,
		0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55, 0x04,
		0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72,
		0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64, 0x67,
		0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79, 0x20,
		0x4c, 0x74, 0x64, 0x30, 0x81, 0x9f, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 0x81, 0x8d,
		0x00, 0x30, 0x81, 0x89, 0x02, 0x81, 0x81, 0x00,
		0xbb, 0x79, 0xd6, 0xf5, 0x17, 0xb5, 0xe5, 0xbf,
		0x46, 0x10, 0xd0, 0xdc, 0x69, 0xbe, 0xe6, 0x2b,
		0x07, 0x43, 0x5a, 0xd0, 0x03, 0x2d, 0x8a, 0x7a,
		0x43, 0x85, 0xb7, 0x14, 0x52, 0xe7, 0xa5, 0x65,
		0x4c, 0x2c, 0x78, 0xb8, 0x23, 0x8c, 0xb5, 0xb4,
		0x82, 0xe5, 0xde, 0x1f, 0x95, 0x3b, 0x7e, 0x62,
		0xa5, 0x2c, 0xa5, 0x33, 0xd6, 0xfe, 0x12, 0x5c,
		0x7a, 0x56, 0xfc, 0xf5, 0x06, 0xbf, 0xfa, 0x58,
		0x7b, 0x26, 0x3f, 0xb5, 0xcd, 0x04, 0xd3, 0xd0,
		0xc9, 0x21, 0x96, 0x4a, 0xc7, 0xf4, 0x54, 0x9f,
		0x5a, 0xbf, 0xef, 0x42, 0x71, 0x00, 0xfe, 0x18,
		0x99, 0x07, 0x7f, 0x7e, 0x88, 0x7d, 0x7d, 0xf1,
		0x04, 0x39, 0xc4, 0xa2, 0x2e, 0xdb, 0x51, 0xc9,
		0x7c, 0xe3, 0xc0, 0x4c, 0x3b, 0x32, 0x66, 0x01,
		0xcf, 0xaf, 0xb1, 0x1d, 0xb8, 0x71, 0x9a, 0x1d,
		0xdb, 0xdb, 0x89, 0x6b, 0xae, 0xda, 0x2d, 0x79,
		0x02, 0x03, 0x01, 0x00, 0x01, 0xa3, 0x81, 0xa7,
		0x30, 0x81, 0xa4, 0x30, 0x1d, 0x06, 0x03, 0x55,
		0x1d, 0x0e, 0x04, 0x16, 0x04, 0x14, 0xb1, 0xad,
		0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb, 0x69,
		0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e, 0x18,
		0x88, 0x39, 0x30, 0x75, 0x06, 0x03, 0x55, 0x1d,
		0x23, 0x04, 0x6e, 0x30, 0x6c, 0x80, 0x14, 0xb1,
		0xad, 0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb,
		0x69, 0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e,
		0x18, 0x88, 0x39, 0xa1, 0x49, 0xa4, 0x47, 0x30,
		0x45, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
		0x04, 0x06, 0x13, 0x02, 0x41, 0x55, 0x31, 0x13,
		0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x08, 0x13,
		0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74,
		0x61, 0x74, 0x65, 0x31, 0x21, 0x30, 0x1f, 0x06,
		0x03, 0x55, 0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e,
		0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x20, 0x57,
		0x69, 0x64, 0x67, 0x69, 0x74, 0x73, 0x20, 0x50,
		0x74, 0x79, 0x20, 0x4c, 0x74, 0x64, 0x82, 0x09,
		0x00, 0x85, 0xb0, 0xbb, 0xa4, 0x8a, 0x7f, 0xb8,
		0xca, 0x30, 0x0c, 0x06, 0x03, 0x55, 0x1d, 0x13,
		0x04, 0x05, 0x30, 0x03, 0x01, 0x01, 0xff, 0x30,
		0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7,
		0x0d, 0x01, 0x01, 0x05, 0x05, 0x00, 0x03, 0x81,
		0x81, 0x00, 0x08, 0x6c, 0x45, 0x24, 0xc7, 0x6b,
		0xb1, 0x59, 0xab, 0x0c, 0x52, 0xcc, 0xf2, 0xb0,
		0x14, 0xd7, 0x87, 0x9d, 0x7a, 0x64, 0x75, 0xb5,
		0x5a, 0x95, 0x66, 0xe4, 0xc5, 0x2b, 0x8e, 0xae,
		0x12, 0x66, 0x1f, 0xeb, 0x4f, 0x38, 0xb3, 0x6e,
		0x60, 0xd3, 0x92, 0xfd, 0xf7, 0x41, 0x08, 0xb5,
		0x25, 0x13, 0xb1, 0x18, 0x7a, 0x24, 0xfb, 0x30,
		0x1d, 0xba, 0xed, 0x98, 0xb9, 0x17, 0xec, 0xe7,
		0xd7, 0x31, 0x59, 0xdb, 0x95, 0xd3, 0x1d, 0x78,
		0xea, 0x50, 0x56, 0x5c, 0xd5, 0x82, 0x5a, 0x2d,
		0x5a, 0x5f, 0x33, 0xc4, 0xb6, 0xd8, 0xc9, 0x75,
		0x90, 0x96, 0x8c, 0x0f, 0x52, 0x98, 0xb5, 0xcd,
		0x98, 0x1f, 0x89, 0x20, 0x5f, 0xf2, 0xa0, 0x1c,
		0xa3, 0x1b, 0x96, 0x94, 0xdd, 0xa9, 0xfd, 0x57,
		0xe9, 0x70, 0xe8, 0x26, 0x6d, 0x71, 0x99, 0x9b,
		0x26, 0x6e, 0x38, 0x50, 0x29, 0x6c, 0x90, 0xa7,
		0xbd, 0xd9, 0x16, 0x03, 0x01, 0x00, 0x04, 0x0e,
		0x00, 0x00, 0x00,
	},

	{
		0x16, 0x03, 0x01, 0x00, 0x86, 0x10, 0x00, 0x00,
		0x82, 0x00, 0x80, 0x3c, 0x13, 0xd7, 0x12, 0xc1,
		0x6a, 0xf0, 0x3f, 0x8c, 0xa1, 0x35, 0x5d, 0xc5,
		0x89, 0x1e, 0x9e, 0xcd, 0x32, 0xc7, 0x9e, 0xe6,
		0xae, 0xd5, 0xf1, 0xbf, 0x70, 0xd7, 0xa9, 0xef,
		0x2c, 0x4c, 0xf4, 0x22, 0xbc, 0x17, 0x17, 0xaa,
		0x05, 0xf3, 0x9f, 0x80, 0xf2, 0xe9, 0x82, 0x2f,
		0x2a, 0x15, 0x54, 0x0d, 0x16, 0x0e, 0x77, 0x4c,
		0x28, 0x3c, 0x03, 0x2d, 0x2d, 0xd7, 0xc8, 0x64,
		0xd9, 0x59, 0x4b, 0x1c, 0xf4, 0xde, 0xff, 0x2f,
		0xbc, 0x94, 0xaf, 0x18, 0x26, 0x37, 0xce, 0x4f,
		0x84, 0x74, 0x2e, 0x45, 0x66, 0x7c, 0x0c, 0x54,
		0x46, 0x36, 0x5f, 0x65, 0x21, 0x7b, 0x83, 0x8c,
		0x6d, 0x76, 0xcd, 0x0d, 0x9f, 0xda, 0x1c, 0xa4,
		0x6e, 0xfe, 0xb1, 0xf7, 0x09, 0x0d, 0xfb, 0x74,
		0x66, 0x34, 0x99, 0x89, 0x7f, 0x5f, 0x77, 0x87,
		0x4a, 0x66, 0x4b, 0xa9, 0x59, 0x57, 0xe3, 0x56,
		0x0d, 0xdd, 0xd8, 0x14, 0x03, 0x01, 0x00, 0x01,
		0x01, 0x16, 0x03, 0x01, 0x00, 0x24, 0xc0, 0x4e,
		0xd3, 0x0f, 0xb5, 0xc0, 0x57, 0xa6, 0x18, 0x80,
		0x80, 0x6b, 0x49, 0xfe, 0xbd, 0x3a, 0x7a, 0x2c,
		0xef, 0x70, 0xb5, 0x1c, 0xd2, 0xdf, 0x5f, 0x78,
		0x5a, 0xd8, 0x4f, 0xa0, 0x95, 0xb4, 0xb3, 0xb5,
		0xaa, 0x3b,
	},

	{
		0x14, 0x03, 0x01, 0x00, 0x01, 0x01, 0x16, 0x03,
		0x01, 0x00, 0x24, 0x9d, 0xc9, 0xda, 0xdf, 0xeb,
		0xc8, 0xdb, 0xf8, 0x94, 0xa5, 0xef, 0xd5, 0xfc,
		0x89, 0x01, 0x64, 0x30, 0x77, 0x5a, 0x18, 0x4b,
		0x16, 0x79, 0x9c, 0xf6, 0xf5, 0x09, 0x22, 0x12,
		0x4c, 0x3e, 0xa8, 0x8e, 0x91, 0xa5, 0x24,
	},
}

var des3ServerScript = [][]byte{
	{
		0x16, 0x03, 0x02, 0x00, 0x7a, 0x01, 0x00, 0x00,
		0x76, 0x03, 0x02, 0x4e, 0x84, 0xf4, 0x3c, 0xe4,
		0xb8, 0xc7, 0xa0, 0x30, 0x55, 0x2a, 0xbc, 0xb7,
		0x04, 0x6b, 0x6f, 0x87, 0x93, 0x96, 0xbd, 0x1a,
		0x7a, 0x1e, 0xce, 0xd2, 0x0d, 0xf3, 0x01, 0x03,
		0xbe, 0x7b, 0x17, 0x00, 0x00, 0x34, 0x00, 0x33,
		0x00, 0x45, 0x00, 0x39, 0x00, 0x88, 0x00, 0x16,
		0x00, 0x32, 0x00, 0x44, 0x00, 0x38, 0x00, 0x87,
		0x00, 0x13, 0x00, 0x66, 0x00, 0x90, 0x00, 0x91,
		0x00, 0x8f, 0x00, 0x8e, 0x00, 0x2f, 0x00, 0x41,
		0x00, 0x35, 0x00, 0x84, 0x00, 0x0a, 0x00, 0x05,
		0x00, 0x04, 0x00, 0x8c, 0x00, 0x8d, 0x00, 0x8b,
		0x00, 0x8a, 0x01, 0x00, 0x00, 0x19, 0x00, 0x09,
		0x00, 0x03, 0x02, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x0e, 0x00, 0x0c, 0x00, 0x00, 0x09, 0x6c, 0x6f,
		0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74,
	},

	{
		0x16, 0x03, 0x01, 0x00, 0x2a, 0x02, 0x00, 0x00,
		0x26, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x16,
		0x03, 0x01, 0x02, 0xbe, 0x0b, 0x00, 0x02, 0xba,
		0x00, 0x02, 0xb7, 0x00, 0x02, 0xb4, 0x30, 0x82,
		0x02, 0xb0, 0x30, 0x82, 0x02, 0x19, 0xa0, 0x03,
		0x02, 0x01, 0x02, 0x02, 0x09, 0x00, 0x85, 0xb0,
		0xbb, 0xa4, 0x8a, 0x7f, 0xb8, 0xca, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x05, 0x05, 0x00, 0x30, 0x45, 0x31,
		0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06,
		0x13, 0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11,
		0x06, 0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53,
		0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74,
		0x65, 0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55,
		0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65,
		0x72, 0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64,
		0x67, 0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79,
		0x20, 0x4c, 0x74, 0x64, 0x30, 0x1e, 0x17, 0x0d,
		0x31, 0x30, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39,
		0x30, 0x39, 0x33, 0x38, 0x5a, 0x17, 0x0d, 0x31,
		0x31, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39, 0x30,
		0x39, 0x33, 0x38, 0x5a, 0x30, 0x45, 0x31, 0x0b,
		0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06, 0x13,
		0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11, 0x06,
		0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53, 0x6f,
		0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74, 0x65,
		0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55, 0x04,
		0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72,
		0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64, 0x67,
		0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79, 0x20,
		0x4c, 0x74, 0x64, 0x30, 0x81, 0x9f, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 0x81, 0x8d,
		0x00, 0x30, 0x81, 0x89, 0x02, 0x81, 0x81, 0x00,
		0xbb, 0x79, 0xd6, 0xf5, 0x17, 0xb5, 0xe5, 0xbf,
		0x46, 0x10, 0xd0, 0xdc, 0x69, 0xbe, 0xe6, 0x2b,
		0x07, 0x43, 0x5a, 0xd0, 0x03, 0x2d, 0x8a, 0x7a,
		0x43, 0x85, 0xb7, 0x14, 0x52, 0xe7, 0xa5, 0x65,
		0x4c, 0x2c, 0x78, 0xb8, 0x23, 0x8c, 0xb5, 0xb4,
		0x82, 0xe5, 0xde, 0x1f, 0x95, 0x3b, 0x7e, 0x62,
		0xa5, 0x2c, 0xa5, 0x33, 0xd6, 0xfe, 0x12, 0x5c,
		0x7a, 0x56, 0xfc, 0xf5, 0x06, 0xbf, 0xfa, 0x58,
		0x7b, 0x26, 0x3f, 0xb5, 0xcd, 0x04, 0xd3, 0xd0,
		0xc9, 0x21, 0x96, 0x4a, 0xc7, 0xf4, 0x54, 0x9f,
		0x5a, 0xbf, 0xef, 0x42, 0x71, 0x00, 0xfe, 0x18,
		0x99, 0x07, 0x7f, 0x7e, 0x88, 0x7d, 0x7d, 0xf1,
		0x04, 0x39, 0xc4, 0xa2, 0x2e, 0xdb, 0x51, 0xc9,
		0x7c, 0xe3, 0xc0, 0x4c, 0x3b, 0x32, 0x66, 0x01,
		0xcf, 0xaf, 0xb1, 0x1d, 0xb8, 0x71, 0x9a, 0x1d,
		0xdb, 0xdb, 0x89, 0x6b, 0xae, 0xda, 0x2d, 0x79,
		0x02, 0x03, 0x01, 0x00, 0x01, 0xa3, 0x81, 0xa7,
		0x30, 0x81, 0xa4, 0x30, 0x1d, 0x06, 0x03, 0x55,
		0x1d, 0x0e, 0x04, 0x16, 0x04, 0x14, 0xb1, 0xad,
		0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb, 0x69,
		0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e, 0x18,
		0x88, 0x39, 0x30, 0x75, 0x06, 0x03, 0x55, 0x1d,
		0x23, 0x04, 0x6e, 0x30, 0x6c, 0x80, 0x14, 0xb1,
		0xad, 0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb,
		0x69, 0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e,
		0x18, 0x88, 0x39, 0xa1, 0x49, 0xa4, 0x47, 0x30,
		0x45, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
		0x04, 0x06, 0x13, 0x02, 0x41, 0x55, 0x31, 0x13,
		0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x08, 0x13,
		0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74,
		0x61, 0x74, 0x65, 0x31, 0x21, 0x30, 0x1f, 0x06,
		0x03, 0x55, 0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e,
		0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x20, 0x57,
		0x69, 0x64, 0x67, 0x69, 0x74, 0x73, 0x20, 0x50,
		0x74, 0x79, 0x20, 0x4c, 0x74, 0x64, 0x82, 0x09,
		0x00, 0x85, 0xb0, 0xbb, 0xa4, 0x8a, 0x7f, 0xb8,
		0xca, 0x30, 0x0c, 0x06, 0x03, 0x55, 0x1d, 0x13,
		0x04, 0x05, 0x30, 0x03, 0x01, 0x01, 0xff, 0x30,
		0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7,
		0x0d, 0x01, 0x01, 0x05, 0x05, 0x00, 0x03, 0x81,
		0x81, 0x00, 0x08, 0x6c, 0x45, 0x24, 0xc7, 0x6b,
		0xb1, 0x59, 0xab, 0x0c, 0x52, 0xcc, 0xf2, 0xb0,
		0x14, 0xd7, 0x87, 0x9d, 0x7a, 0x64, 0x75, 0xb5,
		0x5a, 0x95, 0x66, 0xe4, 0xc5, 0x2b, 0x8e, 0xae,
		0x12, 0x66, 0x1f, 0xeb, 0x4f, 0x38, 0xb3, 0x6e,
		0x60, 0xd3, 0x92, 0xfd, 0xf7, 0x41, 0x08, 0xb5,
		0x25, 0x13, 0xb1, 0x18, 0x7a, 0x24, 0xfb, 0x30,
		0x1d, 0xba, 0xed, 0x98, 0xb9, 0x17, 0xec, 0xe7,
		0xd7, 0x31, 0x59, 0xdb, 0x95, 0xd3, 0x1d, 0x78,
		0xea, 0x50, 0x56, 0x5c, 0xd5, 0x82, 0x5a, 0x2d,
		0x5a, 0x5f, 0x33, 0xc4, 0xb6, 0xd8, 0xc9, 0x75,
		0x90, 0x96, 0x8c, 0x0f, 0x52, 0x98, 0xb5, 0xcd,
		0x98, 0x1f, 0x89, 0x20, 0x5f, 0xf2, 0xa0, 0x1c,
		0xa3, 0x1b, 0x96, 0x94, 0xdd, 0xa9, 0xfd, 0x57,
		0xe9, 0x70, 0xe8, 0x26, 0x6d, 0x71, 0x99, 0x9b,
		0x26, 0x6e, 0x38, 0x50, 0x29, 0x6c, 0x90, 0xa7,
		0xbd, 0xd9, 0x16, 0x03, 0x01, 0x00, 0x04, 0x0e,
		0x00, 0x00, 0x00,
	},

	{
		0x16, 0x03, 0x01, 0x00, 0x86, 0x10, 0x00, 0x00,
		0x82, 0x00, 0x80, 0xae, 0xcf, 0x4f, 0x70, 0x0e,
		0xe5, 0xe7, 0xba, 0xef, 0x0c, 0x66, 0xe9, 0xae,
		0x76, 0xf4, 0xe0, 0xbc, 0x1c, 0x22, 0x5b, 0x72,
		0xc9, 0x68, 0x63, 0x44, 0xec, 0x72, 0xc2, 0xca,
		0xac, 0xc2, 0xf5, 0x5c, 0x28, 0xa1, 0xaf, 0xd0,
		0xc2, 0xf7, 0x79, 0x71, 0x32, 0x73, 0x86, 0xea,
		0x39, 0xf6, 0x04, 0x26, 0x19, 0x84, 0x1d, 0x7d,
		0xa1, 0x21, 0xa6, 0x88, 0xbf, 0x33, 0x5a, 0x64,
		0xb0, 0xc2, 0xcc, 0x19, 0x7a, 0x8b, 0x6e, 0x94,
		0x9e, 0x2e, 0x20, 0xbe, 0xdc, 0xe9, 0x8e, 0xae,
		0x5c, 0x39, 0xc8, 0xcd, 0x0e, 0x19, 0x9a, 0xa2,
		0xfc, 0x3f, 0x61, 0x9a, 0xca, 0x58, 0x69, 0x0d,
		0xa8, 0x7b, 0xbe, 0x98, 0x8f, 0xb9, 0x9d, 0x8b,
		0x68, 0x65, 0xa9, 0x74, 0xcc, 0x8d, 0x0c, 0xb2,
		0xc4, 0x0f, 0xdc, 0x56, 0x3e, 0x44, 0x61, 0x0a,
		0x26, 0x93, 0x99, 0xef, 0x67, 0xff, 0x6e, 0x73,
		0x01, 0xa1, 0x90, 0x14, 0x03, 0x01, 0x00, 0x01,
		0x01, 0x16, 0x03, 0x01, 0x00, 0x60, 0x49, 0x36,
		0xc8, 0x38, 0x95, 0xe4, 0x5d, 0x8e, 0x80, 0x10,
		0x26, 0x9f, 0x87, 0x7d, 0xcd, 0xb9, 0x32, 0x6c,
		0xff, 0xaa, 0xe0, 0x07, 0xec, 0x33, 0xe2, 0x36,
		0x9d, 0xd5, 0x83, 0x2c, 0xf0, 0x0a, 0xa0, 0xa8,
		0x12, 0x9f, 0xca, 0x72, 0xda, 0x70, 0x7d, 0x76,
		0x80, 0x12, 0x88, 0x07, 0xaa, 0x27, 0x62, 0x33,
		0xab, 0x55, 0xad, 0x3c, 0x2b, 0x54, 0xc4, 0x1c,
		0x91, 0xfd, 0x8f, 0x9c, 0xa7, 0x8b, 0x75, 0x10,
		0xa8, 0x6e, 0xfc, 0x30, 0x52, 0x8a, 0x61, 0x02,
		0xdb, 0x9c, 0x6f, 0xc8, 0x19, 0x93, 0x5d, 0x41,
		0x1d, 0x36, 0x68, 0x0b, 0xec, 0x30, 0xae, 0xfb,
		0x90, 0xdb, 0x6d, 0x83, 0xb0, 0xf2,
	},

	{
		0x14, 0x03, 0x01, 0x00, 0x01, 0x01, 0x16, 0x03,
		0x01, 0x00, 0x28, 0x07, 0xf3, 0x33, 0x84, 0xb1,
		0x5d, 0x2b, 0x52, 0xa4, 0x63, 0x3c, 0x32, 0xe0,
		0x0d, 0x22, 0xf5, 0x23, 0xec, 0xf9, 0xa6, 0xec,
		0xc0, 0x12, 0x69, 0x88, 0xf6, 0x7d, 0x37, 0xcd,
		0xc2, 0x74, 0x2f, 0xef, 0xf6, 0x49, 0x15, 0xea,
		0x88, 0x3f, 0x55, 0x17, 0x03, 0x01, 0x00, 0x28,
		0xaf, 0x00, 0x84, 0xff, 0x11, 0x01, 0x6d, 0xba,
		0x39, 0x5e, 0x45, 0xe1, 0x52, 0x5e, 0xc1, 0xab,
		0xde, 0x5b, 0x16, 0xdd, 0xd6, 0x61, 0x57, 0xb8,
		0x66, 0x8b, 0x2d, 0xde, 0x51, 0x41, 0xc5, 0x09,
		0xb3, 0x6a, 0x06, 0x43, 0xb4, 0x73, 0x5c, 0xf1,
		0x15, 0x03, 0x01, 0x00, 0x18, 0xbd, 0x65, 0xb2,
		0xce, 0x77, 0x2e, 0xf9, 0x11, 0xc4, 0x80, 0x43,
		0x5a, 0x73, 0x8b, 0x73, 0xdd, 0xf0, 0x54, 0x44,
		0x7c, 0x56, 0x19, 0x54, 0xda,
	},
}

var aesServerScript = [][]byte{
	{
		0x16, 0x03, 0x02, 0x00, 0x7f, 0x01, 0x00, 0x00,
		0x7b, 0x03, 0x02, 0x4d, 0x08, 0x2d, 0x0b, 0xb3,
		0x57, 0x85, 0x71, 0x4b, 0xfb, 0x34, 0xab, 0x16,
		0xd4, 0x92, 0x50, 0x81, 0x16, 0x95, 0x11, 0x28,
		0x1a, 0xcb, 0xff, 0x09, 0x4d, 0x23, 0xa6, 0xfe,
		0x2e, 0xbb, 0x78, 0x00, 0x00, 0x34, 0x00, 0x33,
		0x00, 0x45, 0x00, 0x39, 0x00, 0x88, 0x00, 0x16,
		0x00, 0x32, 0x00, 0x44, 0x00, 0x38, 0x00, 0x87,
		0x00, 0x13, 0x00, 0x66, 0x00, 0x90, 0x00, 0x91,
		0x00, 0x8f, 0x00, 0x8e, 0x00, 0x2f, 0x00, 0x41,
		0x00, 0x35, 0x00, 0x84, 0x00, 0x0a, 0x00, 0x05,
		0x00, 0x04, 0x00, 0x8c, 0x00, 0x8d, 0x00, 0x8b,
		0x00, 0x8a, 0x01, 0x00, 0x00, 0x1e, 0x00, 0x09,
		0x00, 0x03, 0x02, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x0e, 0x00, 0x0c, 0x00, 0x00, 0x09, 0x6c, 0x6f,
		0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0xff,
		0x01, 0x00, 0x01, 0x00,
	},

	{
		0x16, 0x03, 0x01, 0x00, 0x2a, 0x02, 0x00, 0x00,
		0x26, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x2f, 0x00, 0x16,
		0x03, 0x01, 0x02, 0xbe, 0x0b, 0x00, 0x02, 0xba,
		0x00, 0x02, 0xb7, 0x00, 0x02, 0xb4, 0x30, 0x82,
		0x02, 0xb0, 0x30, 0x82, 0x02, 0x19, 0xa0, 0x03,
		0x02, 0x01, 0x02, 0x02, 0x09, 0x00, 0x85, 0xb0,
		0xbb, 0xa4, 0x8a, 0x7f, 0xb8, 0xca, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x05, 0x05, 0x00, 0x30, 0x45, 0x31,
		0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06,
		0x13, 0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11,
		0x06, 0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53,
		0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74,
		0x65, 0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55,
		0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65,
		0x72, 0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64,
		0x67, 0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79,
		0x20, 0x4c, 0x74, 0x64, 0x30, 0x1e, 0x17, 0x0d,
		0x31, 0x30, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39,
		0x30, 0x39, 0x33, 0x38, 0x5a, 0x17, 0x0d, 0x31,
		0x31, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39, 0x30,
		0x39, 0x33, 0x38, 0x5a, 0x30, 0x45, 0x31, 0x0b,
		0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06, 0x13,
		0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11, 0x06,
		0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53, 0x6f,
		0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74, 0x65,
		0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55, 0x04,
		0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72,
		0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64, 0x67,
		0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79, 0x20,
		0x4c, 0x74, 0x64, 0x30, 0x81, 0x9f, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 0x81, 0x8d,
		0x00, 0x30, 0x81, 0x89, 0x02, 0x81, 0x81, 0x00,
		0xbb, 0x79, 0xd6, 0xf5, 0x17, 0xb5, 0xe5, 0xbf,
		0x46, 0x10, 0xd0, 0xdc, 0x69, 0xbe, 0xe6, 0x2b,
		0x07, 0x43, 0x5a, 0xd0, 0x03, 0x2d, 0x8a, 0x7a,
		0x43, 0x85, 0xb7, 0x14, 0x52, 0xe7, 0xa5, 0x65,
		0x4c, 0x2c, 0x78, 0xb8, 0x23, 0x8c, 0xb5, 0xb4,
		0x82, 0xe5, 0xde, 0x1f, 0x95, 0x3b, 0x7e, 0x62,
		0xa5, 0x2c, 0xa5, 0x33, 0xd6, 0xfe, 0x12, 0x5c,
		0x7a, 0x56, 0xfc, 0xf5, 0x06, 0xbf, 0xfa, 0x58,
		0x7b, 0x26, 0x3f, 0xb5, 0xcd, 0x04, 0xd3, 0xd0,
		0xc9, 0x21, 0x96, 0x4a, 0xc7, 0xf4, 0x54, 0x9f,
		0x5a, 0xbf, 0xef, 0x42, 0x71, 0x00, 0xfe, 0x18,
		0x99, 0x07, 0x7f, 0x7e, 0x88, 0x7d, 0x7d, 0xf1,
		0x04, 0x39, 0xc4, 0xa2, 0x2e, 0xdb, 0x51, 0xc9,
		0x7c, 0xe3, 0xc0, 0x4c, 0x3b, 0x32, 0x66, 0x01,
		0xcf, 0xaf, 0xb1, 0x1d, 0xb8, 0x71, 0x9a, 0x1d,
		0xdb, 0xdb, 0x89, 0x6b, 0xae, 0xda, 0x2d, 0x79,
		0x02, 0x03, 0x01, 0x00, 0x01, 0xa3, 0x81, 0xa7,
		0x30, 0x81, 0xa4, 0x30, 0x1d, 0x06, 0x03, 0x55,
		0x1d, 0x0e, 0x04, 0x16, 0x04, 0x14, 0xb1, 0xad,
		0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb, 0x69,
		0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e, 0x18,
		0x88, 0x39, 0x30, 0x75, 0x06, 0x03, 0x55, 0x1d,
		0x23, 0x04, 0x6e, 0x30, 0x6c, 0x80, 0x14, 0xb1,
		0xad, 0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb,
		0x69, 0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e,
		0x18, 0x88, 0x39, 0xa1, 0x49, 0xa4, 0x47, 0x30,
		0x45, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
		0x04, 0x06, 0x13, 0x02, 0x41, 0x55, 0x31, 0x13,
		0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x08, 0x13,
		0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74,
		0x61, 0x74, 0x65, 0x31, 0x21, 0x30, 0x1f, 0x06,
		0x03, 0x55, 0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e,
		0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x20, 0x57,
		0x69, 0x64, 0x67, 0x69, 0x74, 0x73, 0x20, 0x50,
		0x74, 0x79, 0x20, 0x4c, 0x74, 0x64, 0x82, 0x09,
		0x00, 0x85, 0xb0, 0xbb, 0xa4, 0x8a, 0x7f, 0xb8,
		0xca, 0x30, 0x0c, 0x06, 0x03, 0x55, 0x1d, 0x13,
		0x04, 0x05, 0x30, 0x03, 0x01, 0x01, 0xff, 0x30,
		0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7,
		0x0d, 0x01, 0x01, 0x05, 0x05, 0x00, 0x03, 0x81,
		0x81, 0x00, 0x08, 0x6c, 0x45, 0x24, 0xc7, 0x6b,
		0xb1, 0x59, 0xab, 0x0c, 0x52, 0xcc, 0xf2, 0xb0,
		0x14, 0xd7, 0x87, 0x9d, 0x7a, 0x64, 0x75, 0xb5,
		0x5a, 0x95, 0x66, 0xe4, 0xc5, 0x2b, 0x8e, 0xae,
		0x12, 0x66, 0x1f, 0xeb, 0x4f, 0x38, 0xb3, 0x6e,
		0x60, 0xd3, 0x92, 0xfd, 0xf7, 0x41, 0x08, 0xb5,
		0x25, 0x13, 0xb1, 0x18, 0x7a, 0x24, 0xfb, 0x30,
		0x1d, 0xba, 0xed, 0x98, 0xb9, 0x17, 0xec, 0xe7,
		0xd7, 0x31, 0x59, 0xdb, 0x95, 0xd3, 0x1d, 0x78,
		0xea, 0x50, 0x56, 0x5c, 0xd5, 0x82, 0x5a, 0x2d,
		0x5a, 0x5f, 0x33, 0xc4, 0xb6, 0xd8, 0xc9, 0x75,
		0x90, 0x96, 0x8c, 0x0f, 0x52, 0x98, 0xb5, 0xcd,
		0x98, 0x1f, 0x89, 0x20, 0x5f, 0xf2, 0xa0, 0x1c,
		0xa3, 0x1b, 0x96, 0x94, 0xdd, 0xa9, 0xfd, 0x57,
		0xe9, 0x70, 0xe8, 0x26, 0x6d, 0x71, 0x99, 0x9b,
		0x26, 0x6e, 0x38, 0x50, 0x29, 0x6c, 0x90, 0xa7,
		0xbd, 0xd9, 0x16, 0x03, 0x01, 0x00, 0x04, 0x0e,
		0x00, 0x00, 0x00,
	},

	{
		0x16, 0x03, 0x01, 0x00, 0x86, 0x10, 0x00, 0x00,
		0x82, 0x00, 0x80, 0x71, 0x9c, 0xe7, 0x23, 0xfc,
		0xb9, 0x19, 0x29, 0x82, 0xbf, 0xef, 0x08, 0xf7,
		0x99, 0x36, 0xc3, 0x4c, 0x6f, 0x05, 0xd2, 0x8b,
		0x62, 0x2b, 0x19, 0x9b, 0x7f, 0xc0, 0xcc, 0x48,
		0x30, 0x5f, 0xcd, 0xc3, 0x70, 0x55, 0x53, 0x73,
		0xfa, 0x79, 0x74, 0xf3, 0xa3, 0x76, 0x9f, 0xa1,
		0x7f, 0x98, 0xc2, 0xc0, 0xe3, 0xc5, 0xa0, 0x31,
		0x2f, 0xa6, 0xe8, 0x1e, 0x61, 0x46, 0xb3, 0x9b,
		0x4b, 0x16, 0xf1, 0x2d, 0xc7, 0x63, 0x7f, 0x79,
		0x22, 0x30, 0xd1, 0xf2, 0xfc, 0x77, 0x98, 0x0a,
		0x16, 0x11, 0x63, 0x71, 0x7f, 0x70, 0xef, 0x16,
		0xbb, 0x39, 0x87, 0x34, 0xac, 0x49, 0xbd, 0x07,
		0x67, 0xcb, 0x9c, 0xcc, 0xde, 0xef, 0xb1, 0xe0,
		0xdb, 0x01, 0xb5, 0x35, 0xa9, 0xb3, 0x10, 0x0c,
		0x4b, 0xee, 0xb3, 0x4e, 0xfd, 0xbe, 0x15, 0x27,
		0xf0, 0x46, 0xb2, 0x38, 0xba, 0x5f, 0xcc, 0x89,
		0xec, 0x29, 0x82, 0x14, 0x03, 0x01, 0x00, 0x01,
		0x01, 0x16, 0x03, 0x01, 0x00, 0x30, 0x3c, 0xfb,
		0xa4, 0x12, 0xcb, 0x00, 0xf9, 0x57, 0x7e, 0x9b,
		0xc9, 0xdc, 0x0c, 0xba, 0x9a, 0x81, 0x62, 0xfb,
		0x26, 0x13, 0x53, 0xfe, 0xaa, 0xcc, 0x82, 0xbb,
		0xb6, 0x67, 0x7f, 0x39, 0xbe, 0x4d, 0xbb, 0xc0,
		0x6c, 0x24, 0x31, 0x83, 0xa5, 0x50, 0x3a, 0x75,
		0x32, 0x64, 0xb5, 0xdb, 0xbe, 0x0a,
	},

	{
		0x14, 0x03, 0x01, 0x00, 0x01, 0x01, 0x16, 0x03,
		0x01, 0x00, 0x30, 0x43, 0x24, 0x42, 0x55, 0x08,
		0xe4, 0xc2, 0x15, 0xc9, 0xdb, 0x71, 0x69, 0xee,
		0x09, 0xc5, 0x1c, 0xfd, 0x46, 0x10, 0xa0, 0x68,
		0x21, 0xf2, 0x48, 0xac, 0x6c, 0xc0, 0x2b, 0x62,
		0x07, 0x8f, 0x48, 0x33, 0x0a, 0x6b, 0x62, 0x28,
		0x2e, 0x2c, 0xad, 0xcb, 0x34, 0x85, 0xca, 0x2e,
		0xcd, 0x84, 0xf0,
	},
}

var sslv3ServerScript = [][]byte{
	{
		0x16, 0x03, 0x00, 0x00, 0x41, 0x01, 0x00, 0x00,
		0x3d, 0x03, 0x00, 0x4e, 0x70, 0xe2, 0x18, 0x86,
		0xd6, 0xc6, 0x6f, 0xf3, 0xc8, 0xf4, 0x02, 0xd6,
		0x4d, 0xee, 0x17, 0x32, 0x4b, 0xd2, 0x78, 0xd8,
		0xa1, 0x03, 0x5d, 0x68, 0x82, 0x89, 0xbe, 0xfd,
		0x12, 0xb9, 0x06, 0x00, 0x00, 0x16, 0x00, 0x33,
		0x00, 0x39, 0x00, 0x16, 0x00, 0x32, 0x00, 0x38,
		0x00, 0x13, 0x00, 0x2f, 0x00, 0x35, 0x00, 0x0a,
		0x00, 0x05, 0x00, 0x04, 0x01, 0x00,
	},

	{
		0x16, 0x03, 0x00, 0x00, 0x2a, 0x02, 0x00, 0x00,
		0x26, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x16,
		0x03, 0x00, 0x02, 0xbe, 0x0b, 0x00, 0x02, 0xba,
		0x00, 0x02, 0xb7, 0x00, 0x02, 0xb4, 0x30, 0x82,
		0x02, 0xb0, 0x30, 0x82, 0x02, 0x19, 0xa0, 0x03,
		0x02, 0x01, 0x02, 0x02, 0x09, 0x00, 0x85, 0xb0,
		0xbb, 0xa4, 0x8a, 0x7f, 0xb8, 0xca, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x05, 0x05, 0x00, 0x30, 0x45, 0x31,
		0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06,
		0x13, 0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11,
		0x06, 0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53,
		0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74,
		0x65, 0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55,
		0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65,
		0x72, 0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64,
		0x67, 0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79,
		0x20, 0x4c, 0x74, 0x64, 0x30, 0x1e, 0x17, 0x0d,
		0x31, 0x30, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39,
		0x30, 0x39, 0x33, 0x38, 0x5a, 0x17, 0x0d, 0x31,
		0x31, 0x30, 0x34, 0x32, 0x34, 0x30, 0x39, 0x30,
		0x39, 0x33, 0x38, 0x5a, 0x30, 0x45, 0x31, 0x0b,
		0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06, 0x13,
		0x02, 0x41, 0x55, 0x31, 0x13, 0x30, 0x11, 0x06,
		0x03, 0x55, 0x04, 0x08, 0x13, 0x0a, 0x53, 0x6f,
		0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74, 0x65,
		0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55, 0x04,
		0x0a, 0x13, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72,
		0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64, 0x67,
		0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79, 0x20,
		0x4c, 0x74, 0x64, 0x30, 0x81, 0x9f, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
		0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 0x81, 0x8d,
		0x00, 0x30, 0x81, 0x89, 0x02, 0x81, 0x81, 0x00,
		0xbb, 0x79, 0xd6, 0xf5, 0x17, 0xb5, 0xe5, 0xbf,
		0x46, 0x10, 0xd0, 0xdc, 0x69, 0xbe, 0xe6, 0x2b,
		0x07, 0x43, 0x5a, 0xd0, 0x03, 0x2d, 0x8a, 0x7a,
		0x43, 0x85, 0xb7, 0x14, 0x52, 0xe7, 0xa5, 0x65,
		0x4c, 0x2c, 0x78, 0xb8, 0x23, 0x8c, 0xb5, 0xb4,
		0x82, 0xe5, 0xde, 0x1f, 0x95, 0x3b, 0x7e, 0x62,
		0xa5, 0x2c, 0xa5, 0x33, 0xd6, 0xfe, 0x12, 0x5c,
		0x7a, 0x56, 0xfc, 0xf5, 0x06, 0xbf, 0xfa, 0x58,
		0x7b, 0x26, 0x3f, 0xb5, 0xcd, 0x04, 0xd3, 0xd0,
		0xc9, 0x21, 0x96, 0x4a, 0xc7, 0xf4, 0x54, 0x9f,
		0x5a, 0xbf, 0xef, 0x42, 0x71, 0x00, 0xfe, 0x18,
		0x99, 0x07, 0x7f, 0x7e, 0x88, 0x7d, 0x7d, 0xf1,
		0x04, 0x39, 0xc4, 0xa2, 0x2e, 0xdb, 0x51, 0xc9,
		0x7c, 0xe3, 0xc0, 0x4c, 0x3b, 0x32, 0x66, 0x01,
		0xcf, 0xaf, 0xb1, 0x1d, 0xb8, 0x71, 0x9a, 0x1d,
		0xdb, 0xdb, 0x89, 0x6b, 0xae, 0xda, 0x2d, 0x79,
		0x02, 0x03, 0x01, 0x00, 0x01, 0xa3, 0x81, 0xa7,
		0x30, 0x81, 0xa4, 0x30, 0x1d, 0x06, 0x03, 0x55,
		0x1d, 0x0e, 0x04, 0x16, 0x04, 0x14, 0xb1, 0xad,
		0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb, 0x69,
		0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e, 0x18,
		0x88, 0x39, 0x30, 0x75, 0x06, 0x03, 0x55, 0x1d,
		0x23, 0x04, 0x6e, 0x30, 0x6c, 0x80, 0x14, 0xb1,
		0xad, 0xe2, 0x85, 0x5a, 0xcf, 0xcb, 0x28, 0xdb,
		0x69, 0xce, 0x23, 0x69, 0xde, 0xd3, 0x26, 0x8e,
		0x18, 0x88, 0x39, 0xa1, 0x49, 0xa4, 0x47, 0x30,
		0x45, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
		0x04, 0x06, 0x13, 0x02, 0x41, 0x55, 0x31, 0x13,
		0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x08, 0x13,
		0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74,
		0x61, 0x74, 0x65, 0x31, 0x21, 0x30, 0x1f, 0x06,
		0x03, 0x55, 0x04, 0x0a, 0x13, 0x18, 0x49, 0x6e,
		0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x20, 0x57,
		0x69, 0x64, 0x67, 0x69, 0x74, 0x73, 0x20, 0x50,
		0x74, 0x79, 0x20, 0x4c, 0x74, 0x64, 0x82, 0x09,
		0x00, 0x85, 0xb0, 0xbb, 0xa4, 0x8a, 0x7f, 0xb8,
		0xca, 0x30, 0x0c, 0x06, 0x03, 0x55, 0x1d, 0x13,
		0x04, 0x05, 0x30, 0x03, 0x01, 0x01, 0xff, 0x30,
		0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7,
		0x0d, 0x01, 0x01, 0x05, 0x05, 0x00, 0x03, 0x81,
		0x81, 0x00, 0x08, 0x6c, 0x45, 0x24, 0xc7, 0x6b,
		0xb1, 0x59, 0xab, 0x0c, 0x52, 0xcc, 0xf2, 0xb0,
		0x14, 0xd7, 0x87, 0x9d, 0x7a, 0x64, 0x75, 0xb5,
		0x5a, 0x95, 0x66, 0xe4, 0xc5, 0x2b, 0x8e, 0xae,
		0x12, 0x66, 0x1f, 0xeb, 0x4f, 0x38, 0xb3, 0x6e,
		0x60, 0xd3, 0x92, 0xfd, 0xf7, 0x41, 0x08, 0xb5,
		0x25, 0x13, 0xb1, 0x18, 0x7a, 0x24, 0xfb, 0x30,
		0x1d, 0xba, 0xed, 0x98, 0xb9, 0x17, 0xec, 0xe7,
		0xd7, 0x31, 0x59, 0xdb, 0x95, 0xd3, 0x1d, 0x78,
		0xea, 0x50, 0x56, 0x5c, 0xd5, 0x82, 0x5a, 0x2d,
		0x5a, 0x5f, 0x33, 0xc4, 0xb6, 0xd8, 0xc9, 0x75,
		0x90, 0x96, 0x8c, 0x0f, 0x52, 0x98, 0xb5, 0xcd,
		0x98, 0x1f, 0x89, 0x20, 0x5f, 0xf2, 0xa0, 0x1c,
		0xa3, 0x1b, 0x96, 0x94, 0xdd, 0xa9, 0xfd, 0x57,
		0xe9, 0x70, 0xe8, 0x26, 0x6d, 0x71, 0x99, 0x9b,
		0x26, 0x6e, 0x38, 0x50, 0x29, 0x6c, 0x90, 0xa7,
		0xbd, 0xd9, 0x16, 0x03, 0x00, 0x00, 0x04, 0x0e,
		0x00, 0x00, 0x00,
	},

	{
		0x16, 0x03, 0x00, 0x00, 0x84, 0x10, 0x00, 0x00,
		0x80, 0x74, 0x0e, 0x3a, 0xcf, 0xba, 0x9f, 0x1a,
		0x9b, 0xb2, 0xa4, 0xc7, 0x5d, 0xf3, 0x0c, 0x80,
		0x06, 0x80, 0xf3, 0x57, 0xb2, 0xd9, 0x36, 0x24,
		0x6a, 0x06, 0x13, 0x40, 0xf9, 0x7c, 0xb9, 0x3e,
		0x4b, 0x68, 0x4f, 0x21, 0x90, 0x2d, 0xbd, 0xca,
		0xd4, 0x83, 0xf0, 0x7a, 0xeb, 0x7a, 0x74, 0x1b,
		0xcd, 0xfe, 0x69, 0xef, 0xc0, 0x86, 0xa0, 0x24,
		0x31, 0x65, 0x40, 0xd2, 0xdd, 0x6f, 0xb9, 0xd7,
		0x8d, 0xc1, 0x69, 0x60, 0x44, 0x7a, 0x75, 0xfb,
		0x42, 0x6a, 0x0f, 0x66, 0x45, 0x10, 0x73, 0xee,
		0x87, 0x28, 0x37, 0x83, 0x86, 0xd8, 0x5a, 0xc8,
		0x60, 0x87, 0xda, 0x33, 0x87, 0xaf, 0x34, 0x8b,
		0xf5, 0x61, 0x63, 0x7a, 0x5c, 0x60, 0x26, 0xb9,
		0xdb, 0xa1, 0xb7, 0xe3, 0x60, 0x38, 0x94, 0x5c,
		0x83, 0x23, 0xd6, 0x8d, 0xc2, 0x14, 0x4a, 0x0f,
		0x0e, 0x4f, 0xf9, 0x4e, 0x7b, 0x15, 0xcd, 0x18,
		0x04, 0x14, 0x03, 0x00, 0x00, 0x01, 0x01, 0x16,
		0x03, 0x00, 0x00, 0x3c, 0xbd, 0xbc, 0xec, 0xdc,
		0x79, 0xb1, 0xae, 0x16, 0xc9, 0x26, 0x9a, 0xc0,
		0xc0, 0x2c, 0x33, 0x36, 0x13, 0x91, 0x58, 0x5d,
		0x7d, 0xee, 0x4e, 0xd8, 0x7e, 0xac, 0x88, 0x87,
		0x0a, 0x75, 0x66, 0xb1, 0x44, 0x79, 0x2f, 0x42,
		0xe8, 0x92, 0x74, 0x4c, 0xab, 0x36, 0xc8, 0x17,
		0x5f, 0x02, 0x8a, 0x20, 0x53, 0xe9, 0x1d, 0xb4,
		0xfe, 0x5c, 0x2b, 0xd9, 0x0a, 0xfb, 0xc6, 0x63,
	},

	{
		0x14, 0x03, 0x00, 0x00, 0x01, 0x01, 0x16, 0x03,
		0x00, 0x00, 0x3c, 0xaa, 0xa1, 0x98, 0xc4, 0x6b,
		0x5a, 0x16, 0x3f, 0x5f, 0xa4, 0x96, 0x3e, 0x78,
		0xe4, 0x6f, 0x49, 0x05, 0x47, 0xc4, 0x05, 0x60,
		0xeb, 0x0b, 0x45, 0xe3, 0xbc, 0x50, 0x11, 0x24,
		0x5f, 0x01, 0xd7, 0xb8, 0x8f, 0x60, 0x63, 0x66,
		0xbd, 0x3e, 0xd9, 0xa8, 0x80, 0x43, 0x9f, 0x0b,
		0x51, 0x61, 0xed, 0x13, 0xc6, 0x21, 0xd0, 0xfe,
		0xbc, 0x17, 0x3c, 0x36, 0xb0, 0x82, 0x7f, 0x17,
		0x03, 0x00, 0x00, 0x21, 0xee, 0x44, 0xf3, 0xa6,
		0x88, 0x9d, 0x78, 0x44, 0xde, 0xdf, 0xeb, 0xc5,
		0xad, 0xc4, 0xcc, 0x56, 0x5c, 0x54, 0x96, 0x52,
		0x3f, 0xd9, 0x40, 0x6e, 0x79, 0xd8, 0x58, 0x78,
		0x4f, 0x5a, 0xe9, 0x06, 0xef, 0x15, 0x03, 0x00,
		0x00, 0x16, 0xd3, 0xc2, 0x52, 0x99, 0x2a, 0x84,
		0xc4, 0x52, 0x5f, 0x3b, 0x19, 0xe7, 0xfc, 0x65,
		0xaf, 0xd3, 0xb7, 0xa3, 0xcc, 0x4a, 0x1d, 0x2e,
	},
}
