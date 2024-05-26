package main

import (
	"fmt"
	"net"
	"os"

	"github.com/PeronGH/findcf/internal/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: iscf <ip>")
		os.Exit(1)
	}

	iscf := scanner.ScanIP(net.ParseIP(os.Args[1]))
	fmt.Println(iscf)
}
