package networkhelpers

import (
	"io"
)

// Counter counts upstream and downstream bandwidth
type Counter struct {
	Upstream   int64
	Downstream int64
}

// NewCounterConn returnts a pointer to a new instance of CounterConn
func NewCounterConn(c io.ReadWriteCloser) *CounterConn {
	return &CounterConn{
		c,
		&Counter{},
	}
}

// CounterConn is a net.Conn counts upstream and downstream bandwidth
type CounterConn struct {
	io.ReadWriteCloser
	Counter *Counter
}

// Write writes bytes and counts them
func (c *CounterConn) Write(b []byte) (n int, err error) {
	n, err = c.ReadWriteCloser.Write(b)
	c.Counter.Downstream += int64(n)
	return n, err
}

// Read reads bytes and counts them
func (c *CounterConn) Read(b []byte) (n int, err error) {
	n, err = c.ReadWriteCloser.Read(b)
	c.Counter.Upstream += int64(n)
	return n, err
}
