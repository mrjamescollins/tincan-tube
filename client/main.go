package main

import (
	"fmt"
	"net"
	"tincan-tube/server/vpn"
)

func main() {
	serverAddr := "vpn_server:51820" // this is the Docker compose service name
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	privateKey := vpn.GeneratePrivateKey()
	publicKey := vpn.DerivePublicKey(privateKey)

	encryptedMessage := vpn.EncryptMessage([]byte("Hello, VPN Server!"), publicKey)
	_, err = conn.Write(encryptedMessage)
	if err != nil {
		fmt.Println("Error sending packet:", err)
	}

	buffer := make([]byte, 1500)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading packet:", err)
	}

	decryptedResponse := vpn.DecryptMessage(buffer[:n], privateKey)
	fmt.Printf("Response from server: %s\n", decryptedResponse)
}
