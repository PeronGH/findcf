package scanner

import (
	"crypto/tls"
	"net"
)

const (
	testHost = "cp.cloudflare.com"
)

func ScanTLS(addr *net.TCPAddr) bool {
	// Establish TCP connection
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return false
	}
	defer conn.Close()

	// Perform TLS handshake
	tlsConn := tls.Client(conn, &tls.Config{
		ServerName: testHost,
	})
	if err := tlsConn.Handshake(); err != nil {
		return false
	}
	defer tlsConn.Close()

	// So it is on Cloudflare network, now abort
	return true
}

func ScanIP(ip net.IP) bool {
	return ScanTLS(&net.TCPAddr{
		IP:   ip,
		Port: 443,
	})
}
