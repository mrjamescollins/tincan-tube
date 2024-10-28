package vpn

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/nacl/box"
)

// GeneratePrivateKey generates an Ed25519 private key
func GeneratePrivateKey() ed25519.PrivateKey {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	fmt.Println("pub key: %s", publicKey)
	return privateKey
}

// DerivePublicKey derives the public key from a private key
func DerivePublicKey(privateKey ed25519.PrivateKey) ed25519.PublicKey {
	return privateKey.Public().(ed25519.PublicKey)
}

// EncryptMessage encrypts a message using NaCl
func EncryptMessage(message []byte, publicKey ed25519.PublicKey) []byte {
	var out []byte
	box.Seal(out, message, nil, (*[32]byte)(publicKey), nil)
	return out
}

// DecryptMessage decrypts a message
func DecryptMessage(cipherText []byte, privateKey ed25519.PrivateKey) string {
	var out []byte
	box.Open(out, cipherText, nil, (*[32]byte)(privateKey), nil)
	return string(out)
}
