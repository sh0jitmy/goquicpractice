package main

import (
	"fmt"
	"log"

	"github.com/quic-go/quic-go"
)

func main() {
	listener, err := quic.ListenAddr("localhost:4242", generateTLSConfig(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		session, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleSession(session)
	}
}

func handleSession(session quic.Session) {
	for {
		data, err := session.ReceiveDatagram()
		if err != nil {
			fmt.Println("Failed to receive datagram:", err)
			return
		}

		fmt.Printf("Received Datagram: %s\n", data)

		_, err = session.SendDatagram([]byte("Hello from server!"))
		if err != nil {
			fmt.Println("Failed to send datagram:", err)
			return
		}
	}
}

func generateTLSConfig() *quic.Config {
	// 任意のTLS設定を行います
	return nil
}

