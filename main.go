package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
	Example usage:
		go build bogus-domains.go
		cat domains.csv|./bogus-domains>output.csv
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Scan stdin line by line
	for scanner.Scan() {
		domainByID := strings.Split(scanner.Text(), ",")
		_, err := net.LookupHost(domainByID[1])
		// Write resolution failure to stdout
		if err != nil {
			fmt.Println(domainByID[0] + "," + domainByID[1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Debug error:", err)
		os.Exit(1)
	}
}
