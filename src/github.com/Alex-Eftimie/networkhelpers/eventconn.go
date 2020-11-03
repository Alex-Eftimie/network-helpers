package networkhelpers

import "net"

// EventConn is a net.Conn that calls a callback when it's closed
type EventConn struct {
	net.Conn
	Addr          string
	CloseCallback func(ec *EventConn)
	Payload       interface{}
}

func (ec *EventConn) Read(p []byte) (n int, err error) {
	n, err = ec.Conn.Read(p)
	if err != nil && ec.CloseCallback != nil {
		ec.CloseCallback(ec)
		ec.CloseCallback = nil // prevent Double Call
	}
	return n, err
}
func (ec *EventConn) Write(p []byte) (n int, err error) {
	n, err = ec.Conn.Write(p)
	if err != nil && ec.CloseCallback != nil {
		go ec.CloseCallback(ec)
		ec.CloseCallback = nil // prevent Double Call
	}
	return n, err
}
