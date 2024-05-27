package scanner

import (
	"crypto/tls"
	"log"
	"net"
	"sync"
	"time"
)

const testHost = "cp.cloudflare.com"

var dialer = &net.Dialer{
	Timeout: 5 * time.Second,
}

func ScanTLS(addr string) bool {
	// Establish TCP connection
	conn, err := dialer.Dial("tcp", addr)
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

func ScanIP(ip string) bool {
	return ScanTLS(ip + ":443")
}

func ScanCIDR(cidr *net.IPNet) {
	ones, bits := cidr.Mask.Size()
	numIPs := 1 << (bits - ones)

	var wg sync.WaitGroup
	wg.Add(numIPs)

	for ip := cidr.IP.Mask(cidr.Mask); cidr.Contains(ip); inc(ip) {
		ipStr := ip.String()
		go func() {
			defer wg.Done()
			if ScanIP(ipStr) {
				log.Println("Found:", ipStr)
			} else {
				log.Println("Scanned:", ipStr)
			}
		}()
		// Sleep for a bit to avoid rate limiting
		time.Sleep(500 * time.Microsecond)
	}

	wg.Wait()
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
