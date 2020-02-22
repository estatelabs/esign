package tests

import (
	"fmt"
	"github.com/estatelabs/esign"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreatePair(t *testing.T) {
	dir, err := ioutil.TempDir("", "esign")
	if err != nil {
		log.Fatal(err)
	}
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
