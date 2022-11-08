package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

type Client chan<- string // an outgoing message channel

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	clientMessages  = make(chan string) // all incoming client messages
)

var (
	serverHost = flag.String("h",
		"localhost", "host to listen on")
	serverPort = flag.Int("p",
		3090, "port to listen on")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	ch := make(chan string) // outgoing client messages
	go ClientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- fmt.Sprintf("[%s|System]:Wellcome!, you are %s", time.Now().Format(time.Kitchen), who)
	clientMessages <- fmt.Sprintf("[%s|System]:%s has arrived", time.Now().Format(time.Kitchen), who)
	incomingClients <- ch

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		clientMessages <- fmt.Sprintf("[%s|%s]: %s", time.Now().Format(time.Kitchen), who, inputMessage.Text())
	}

	leavingClients <- ch
	clientMessages <- fmt.Sprintf("[%s|System]:%s has left", time.Now().Format(time.Kitchen), who)
}

func ClientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func BroadCast() {
	clients := make(map[Client]bool) // all connected clients
	for {                            //multiplexing
		select { //get the message from the channel
		case msg := <-clientMessages:
			// Broadcast incoming message to all clients' outgoing message channels.
			// clients' outgoing message channels.
			for client := range clients {
				client <- msg
			}

		case client := <-incomingClients:
			clients[client] = true

		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", *serverHost, *serverPort)) //listen on tcp port
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Listening on %s:%d", *serverHost, *serverPort)
	go BroadCast() //start a goroutine to receive data
	for {
		conn, err := listener.Accept() //accept the connection
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleConnection(conn) //handle the connection
	}
}
