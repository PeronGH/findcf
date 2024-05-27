package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/oschwald/geoip2-golang"
	"github.com/oschwald/maxminddb-golang"
)

func main() {
	dbpath := flag.String("db", "", "The filename of the MaxMind GeoLite2 database in mmdb format")
	savepath := flag.String("o", "ip.txt", "The filename to save the extracted IPs")
	countryCode := flag.String("c", "", "The country code to extract")
	ipv4Only := flag.Bool("4", false, "Only extract IPv4 addresses")

	flag.Parse()

	if *dbpath == "" || *countryCode == "" {
		flag.PrintDefaults()
		return
	}
	*countryCode = strings.ToUpper(*countryCode)

	// Open the MaxMind database
	reader, err := maxminddb.Open(*dbpath)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	// Open the save file
	file, err := os.OpenFile(*savepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Extract the IPs
	numIPs := 0
	networks := reader.Networks()

	record := geoip2.Country{}
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Fatalln(err)
		}
		if record.Country.IsoCode == *countryCode {
			if *ipv4Only && subnet.IP.To4() == nil {
				continue
			}
			ones, bits := subnet.Mask.Size()
			numIPs += 1 << (bits - ones)
			file.WriteString(subnet.String() + "\n")
		}
	}

	log.Printf("Extracted %d IPs\n", numIPs)
}
