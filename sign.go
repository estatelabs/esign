package esign

import (
	"golang.org/x/crypto/nacl/sign"
	"golang.org/x/crypto/sha3"
	"io"
	"io/ioutil"
	"os"
)

type Signature struct {
	bytes []byte
}

// Sign a file with ed25519 assuming that the hash from the file and a private key will be provided
func Sign(file string, privateKey *[64]byte) *Signature {
	var out []byte
	signature := &Signature{bytes: sign.Sign(out, hash(file), privateKey)}

	if cap(out) == 0 {
		return signature
	} else {
		panic("Overhead bytes occurred!")
	}
}

// The signature will be stored to provided destination location
func (signature *Signature) Save(destLocation string) {
	if len(destLocation) > 0 && !fileExists(destLocation) {
		err := ioutil.WriteFile(destLocation, signature.bytes, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic("File could not be saved! Caused by: " + destLocation)
	}
}

// Create a SHA3-256 Bit hash from a file
func hash(location string) []byte {
	file, err := os.Open(location)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := sha3.New256()
	if _, err := io.Copy(h, file); err != nil {
		panic(err)
	}

	return h.Sum(nil)
}
