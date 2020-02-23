package tests

import (
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createTempDir() string {
	dir, err := ioutil.TempDir("", "esign")
	if err != nil {
		panic(err)
	}

	return dir
}
