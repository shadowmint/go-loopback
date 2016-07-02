package loopback

import (
	"fmt"
	"net"
	"ntoolkit/errors"
)

// Loopback is a simple container for a pair of local loopback network sockets
type Loopback struct {
	A net.Conn
	B net.Conn
}

// New returns a loopback pair of net.Conn or an error
func New() (*Loopback, error) {

	// Listen on some random local port
	messages := make(chan net.Conn)
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, errors.Fail(ErrListen, err, "Failed to bind socket")
	}
	go func() {
		conn, err2 := l.Accept()
		if err2 != nil {
			messages <- nil
		} else {
			messages <- conn
		}
	}()

	// Open outgoing connection
	port := l.Addr().(*net.TCPAddr).Port
	conn2, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return nil, errors.Fail(ErrConnect, err, "Failed to connect to local socket")
	}

	// Get active connection
	conn1 := <-messages
	if conn1 == nil {
		return nil, errors.Fail(ErrListen, nil, "Failed to bind socket")
	}

	return &Loopback{conn1, conn2}, nil
}

// Close both sockets
func (loop *Loopback) Close() {
	loop.A.Close()
	loop.B.Close()
}
