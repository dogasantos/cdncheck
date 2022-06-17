package main

import (
	"log"
	"net"

	"github.com/dogasantos/cdncheck"
)

func main() {
    // uses projectdiscovery endpoint with cached data to avoid ip ban
    // Use cdncheck.New() if you want to scrape each endpoint (don't do it too often or your ip can be blocked)
    client, err := cdncheck.NewWithCache()
    if err != nil {
        log.Fatal(err)
    }
    if found, err := client.Check(net.ParseIP("201.95.254.67")); found && err == nil {
        log.Println("ip is part of cdn")
    }
}
