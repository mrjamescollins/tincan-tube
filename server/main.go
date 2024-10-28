package main

import (
	"fmt"
	"net"
	"os"
	"tincan-tube/server/vpn"
)

func main() {
	fmt.Println("Starting tincan-tube VPN server...")

	privateKey := vpn.GeneratePrivateKey()
	publicKey := vpn.DerivePublicKey(privateKey)
	fmt.Printf("Server public key: %s\n", publicKey)

	listenAddr := ":51820" // the default port to listen for WireGuard protocol
	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("VPN server listening on %s\n", listenAddr)

	for {
		buffer := make([]byte, 1500) // setting MTU size
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Printf("Read error: %v\n", err)
			continue
		}

		decryptedMessage := vpn.DecryptMessage(buffer[:n], privateKey)
		fmt.Printf("Received packet from %s: %s\n", addr, decryptedMessage)

		response := vpn.EncryptMessage([]byte("Hello, there client! Can you here me over the tubes?"), publicKey)
		_, err = conn.WriteTo(response, addr)
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
		}
	}
}
