package main

import (
	"fmt"
	"log"

	"github.com/quic-go/quic-go"
)

func main() {
	session, err := quic.DialAddr("localhost:4242", &quic.Config{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err := session.SendDatagram([]byte("Hello from client!"))
		if err != nil {
			fmt.Println("Failed to send datagram:", err)
			return
		}

		data, err := session.ReceiveDatagram()
		if err != nil {
			fmt.Println("Failed to receive datagram:", err)
			return
		}

		fmt.Printf("Received Datagram: %s\n", data)
	}
}

