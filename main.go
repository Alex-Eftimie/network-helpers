package networkhelpers

import (
	"log"
	"net"
	"time"
)

// ConnectTCP connects to the specified address and returns a pointer to the TCPConn or nil
func ConnectTCP(addr string) *net.TCPConn {

	d := net.Dialer{Timeout: 5 * time.Second}
	conn, err := d.Dial("tcp", addr)
	if err != nil {
		return nil
	}
	return conn.(*net.TCPConn)
}

// GetIPFromAddr retrieves *net.IP ip from a host:port touple
func GetIPFromAddr(addr string) *net.IP {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Println("ResolveTCPAddr failed:", err.Error())
		return nil
	}
	return &tcpAddr.IP
}

// RemoteAddr returns IP string from a net.Conn
func RemoteAddr(c net.Conn) string {
	return c.RemoteAddr().(*net.TCPAddr).IP.String()
}
