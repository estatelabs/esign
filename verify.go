package esign

import (
	"bytes"
	"golang.org/x/crypto/nacl/sign"
	"io/ioutil"
)

type verifier struct {
	checksum  *[]byte
	signature []byte
	publicKey *[32]byte
}

// Checks if file exists and generates checksum and provide additionally a new verifier object to use func chaining
func Verify(file string) *verifier {
	if fileExists(file) {
		checksum := hash(file)

		return &verifier{checksum: &checksum}
	} else {
		panic("File could not be found!")
	}
}

// Adds public key byte pointer to verifier
func (verifier *verifier) With(publicKey *[32]byte) *verifier {
	verifier.publicKey = publicKey

	return verifier
}

// Adds raw byte signature to verifier
func (verifier *verifier) ByRawSignature(signature []byte) bool {
	verifier.signature = signature
	return check(verifier)
}

// Adds filesystem-based signature if it exists
func (verifier *verifier) BySavedSignature(file string) bool {
	signature, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	verifier.signature = signature

	return check(verifier)
}

// Private function for checking signed and actual file checksum and returns a boolean back to the caller
func check(verifier *verifier) bool {
	var out []byte
	expectedChecksum, _ := sign.Open(out, verifier.signature, verifier.publicKey)

	if cap(out) == 0 && bytes.Equal(expectedChecksum, *verifier.checksum) {
		return true
	} else {
		return false
	}
}
