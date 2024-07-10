package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type CustomClient struct {
	hostname  string
	portNum   string
	connLimit time.Duration
}

func NewCustomClient(h, p, limit string) *CustomClient {
	connLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatal(err)
	}
	return &CustomClient{
		hostname:  h,
		portNum:   p,
		connLimit: time.Duration(connLimit) * time.Second,
	}
}

func establishConnection(client *CustomClient, ctx context.Context, cancel context.CancelFunc) {
	conn, err := net.DialTimeout("tcp", client.hostname+":"+client.portNum, client.connLimit)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			_, err := fmt.Fprintf(conn, "time is up")
			if err != nil {
				log.Print(err)
			}
			log.Println("time is up...")
			err = conn.Close()
			if err != nil {
				log.Print(err)
			}
			return
		default:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("message: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Print("read error...")
			}
			_, err = fmt.Fprintf(conn, text+"\n")
			if err != nil {
				log.Print(err)
			}
			response, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Println("from server :" + response)
		}
	}
}

func main() {
	timeout := flag.String("timeout", "10", "time on work with server")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("incorrect soccet")
	}
	port := args[1]
	hostname := args[0]

	client := NewCustomClient(hostname, port, *timeout)

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, client.connLimit)

	establishConnection(client, ctx, cancel)

	fmt.Println("Connection closed")
}