package main

import (
	"bufio"
	"log"
	"net"
	"os"

	"github.com/PeronGH/findcf/internal/scanner"
)

func main() {
	// Open the IP list file
	ipList, err := os.Open("ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer ipList.Close()

	// Iterate line by line
	lines := bufio.NewScanner(ipList)
	for lines.Scan() {
		ip := lines.Text()
		_, cidr, err := net.ParseCIDR(ip)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Scanning:", cidr)
		scanner.ScanCIDR(cidr)
	}
}
