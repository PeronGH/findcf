package main

import (
	"fmt"
	"net"

	"github.com/PeronGH/findcf/internal/scanner"
)

func main() {
	// Placeholders...
	fmt.Println(scanner.ScanIP(net.ParseIP("1.1.1.1")))
	fmt.Println(scanner.ScanIP(net.ParseIP("8.8.8.8")))
}
