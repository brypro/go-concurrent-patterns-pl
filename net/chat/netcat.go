package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	flag.Parse()
	//listen on tcp port
	conn, err := net.Dial("tcp",
		fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Listening on %s:%d", *host, *port)
	done := make(chan struct{})

	//start a goroutine to receive data
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	// Copy what we got to the console line
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

var (
	port = flag.Int("port",
		3090, "port to listen on")
	host = flag.String("host",
		"localhost", "host to listen on")
)

func CopyContent(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy: %v\n", err)
		os.Exit(1)
	}
}
