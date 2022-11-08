package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var site = flag.String("site",
	"scanme.nmap.org",
	"site to scan open ports")

func main() {
	flag.Parse() // parse the command line arguments
	wg := sync.WaitGroup{}
	//scan every port and make a connection
	start := time.Now()
	for port := 0; port < 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp",
				fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s:Port %d is open\n", *site, port)
		}(port)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)

}
