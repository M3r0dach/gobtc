package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Password   string
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, publicKey
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public, ""}
	fmt.Println("Plz input your password:")
	fmt.Scanln(&wallet.Password)
	return &wallet
}

func HashPubKey(publicKey []byte) []byte {
	publicSHA256 := sha256.Sum256(publicKey)
	RIPEMD160Hasher := ripemd160.New()
	RIPEMD160Hasher.Write(publicSHA256[:])
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:4]
}

func DecodeAddress(address string) []byte {
	if ValidateAddress(address) == false {
		panic("address isn't legal")
	}
	pubKeyHash, err := base64.StdEncoding.DecodeString(address)
	if err != nil {
		panic(err)
	}
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	return pubKeyHash
}
func (w Wallet) GetAddress() string {
	pubKeyHash := HashPubKey(w.PublicKey)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	fullPayload := append(versionedPayload, checksum(versionedPayload)...)
	address := base64.StdEncoding.EncodeToString(fullPayload)
	return address
}

func ValidateAddress(address string) bool {
	fullPayload, err := base64.StdEncoding.DecodeString(address)
	if err != nil {
		panic(err)
	}
	actualChecksum := fullPayload[len(fullPayload)-4:]
	versionedPayload := fullPayload[:len(fullPayload)-4]
	if versionedPayload[0] != version {
		panic("address version not support")
	}
	return bytes.Compare(actualChecksum, checksum(versionedPayload)) == 0
}
