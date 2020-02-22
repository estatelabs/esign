package esign

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/nacl/sign"
	"io"
	"io/ioutil"
)

type Keystore struct {
	publicKey  *[32]byte
	privateKey *[64]byte
}

// Create Ed25519 key pair
func CreateKeyPair() *Keystore {
	var rand io.Reader
	publicKey, privateKey, _ := sign.GenerateKey(rand)

	keystore := &Keystore{publicKey: publicKey, privateKey: privateKey}

	return keystore
}

// Saves public and private key to desired destination
func (keystore *Keystore) Save(destLocation string) {
	if len(destLocation) > 0 && !fileExists(destLocation) {
		err := ioutil.WriteFile(destLocation+".pub", keystore.publicKey[:], 0644)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(destLocation+".prv", keystore.privateKey[:], 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic("File could not be saved! Caused by: " + destLocation)
	}
}

// Load private key from generic file and provide a pointer to it
func LoadPrivateKey(location string) (*[64]byte, error) {
	privateKeyBytes, err := loadKey(location)
	if err != nil {
		return nil, err
	}
	var privateKey [64]byte
	copy(privateKey[:], *privateKeyBytes)
	return &privateKey, nil
}

// Load and decode private key from base64 string and provide a pointer to it
func LoadPrivateKeyFromBase64(encodedKey string) *[64]byte {
	var privateKey [64]byte
	copy(privateKey[:], convertBase64ToBytes(encodedKey))

	return &privateKey
}

// Load public key from generic file and provide a pointer to it
func LoadPublicKey(location string) (*[32]byte, error) {
	publicKeyBytes, err := loadKey(location)
	if err != nil {
		return nil, err
	}
	var publicKey [32]byte
	copy(publicKey[:], *publicKeyBytes)
	return &publicKey, nil
}

// Load and decode public key from base64 string and provide a pointer to it
func LoadPublicKeyFromBase64(encodedKey string) *[32]byte {
	var publicKey [32]byte
	copy(publicKey[:], convertBase64ToBytes(encodedKey))

	return &publicKey
}

// Load and decode private key from base64 string and provide a pointer to it
func convertBase64ToBytes(encodedKey string) []byte {
	var bytes []byte

	bytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		panic(err)
	}

	return bytes
}

// Internally load a specific file if it exists and is not a directory
// Output provides a pointer of a byte slice
func loadKey(location string) (*[]byte, error) {
	if fileExists(location) {
		file, err := ioutil.ReadFile(location)
		if err != nil {
			panic(err)
		} else {
			return &file, nil
		}
	}
	return nil, errors.New("key could not be loaded")
}
