package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error:%v \n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecords, dmarcRecords []string
	// check for MXrecords
	MxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	if len(MxRecords) > 0 {
		hasMX = true
	}
	// check for spf records
	textRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	for _, records := range textRecords {
		if strings.HasPrefix(records, "v=spf1") {
			hasSPF = true
			spfRecords = append(spfRecords, records)
			break
		}
	}
	// check for dmarc records

	DMARCRecords, err := net.LookupTXT(domain + "_dmarc")
	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	for _, records := range DMARCRecords {
		if strings.HasPrefix(records, "v=DMARC1") {
			hasDMARC = true
			dmarcRecords = append(dmarcRecords, records)
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecords)

}
