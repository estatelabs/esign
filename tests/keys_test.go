package tests

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/estatelabs/esign"
	"os"
	"testing"
)

func TestCreatePair(t *testing.T) {
	dir := createTempDir()
	defer os.RemoveAll(dir)

	keyPath := dir + "/key"
	esign.CreateKeyPair().Save(keyPath)
	fmt.Println(keyPath)

	if !fileExists(keyPath + ".pub") {
		t.Error("Could not create and save public key")
	}

	if !fileExists(keyPath + ".prv") {
		t.Error("Could not create and save private key")
	}
}

func TestLoadPrivateKey(t *testing.T) {
	dir := createTempDir()
	defer os.RemoveAll(dir)

	keyPath := dir + "/key"
	esign.CreateKeyPair().Save(keyPath)

	_, err := esign.LoadPrivateKey(keyPath + ".prv")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadPublicKey(t *testing.T) {
	dir := createTempDir()
	defer os.RemoveAll(dir)

	keyPath := dir + "/key"
	esign.CreateKeyPair().Save(keyPath)

	_, err := esign.LoadPublicKey(keyPath + ".prv")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadPublicKeyFromBase64(t *testing.T) {
	keystore := esign.CreateKeyPair()

	publicKeyToBase64 := base64.StdEncoding.EncodeToString(keystore.PublicKey[:])

	publicKeyFromBase64, err := esign.LoadPublicKeyFromBase64(publicKeyToBase64)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(keystore.PublicKey[:], publicKeyFromBase64[:]) {
		t.Error("public key decoding from base64 string doesn't work.")
		t.Errorf("Origin bytes: %v\n", *keystore.PublicKey)
		t.Errorf("Decoded bytes: %v\n", publicKeyFromBase64)
	}
}

func TestLoadPrivateKeyFromBase64(t *testing.T) {
	keystore := esign.CreateKeyPair()

	privateKeyToBase64 := base64.StdEncoding.EncodeToString(keystore.PrivateKey[:])

	privateKeyFromBase64, err := esign.LoadPrivateKeyFromBase64(privateKeyToBase64)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(keystore.PrivateKey[:], privateKeyFromBase64[:]) {
		t.Error("private key decoding from base64 string doesn't work.")
		t.Errorf("Origin bytes: %v\n", *keystore.PrivateKey)
		t.Errorf("Decoded bytes: %v\n", privateKeyFromBase64)
	}
}
