# eSign

[![Go Report Card](https://goreportcard.com/badge/github.com/estatelabs/esign)](https://goreportcard.com/report/github.com/estatelabs/esign)
[![GitHub license](https://img.shields.io/github/license/estatelabs/esign?style=flat-square)](https://github.com/estatelabs/esign/blob/master/LICENSE)
![GitHub last commit](https://img.shields.io/github/last-commit/estatelabs/esign?style=flat-square)

**eSign** is a powerful library with fluent human readable syntax for the verification of files according to the **SHA3-256 Bit** checksum and the digital signature **ed25519** written in GO. 

## Features
🔑 Create / Load Ed25519 key pairs as a raw byte sequence, base64 string or save / load them directly as a binary file in / from your filesystem

✔️ Verify file against **SHA3-256 Bit** checksum with digital signature **ed25519**

## Getting Started

First of all, [download](https://golang.org/dl/) and install at least Go. `1.11` (with enabled [Go Modules](https://golang.org/doc/go1.11#modules)) or higher is required.

Installation is done using the [`go get`](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get github.com/estatelabs/esign
```

## Usage example

```go
package main

import (
	"fmt"
	"github.com/estatelabs/esign"
	"io/ioutil"
)

const DemoPath = "/Users/fabian/Desktop/test"

func main() {
	createKeyPairAndSaveThem()
	loadKeysAndSign()
	verifyFile()
}

func createKeyPairAndSaveThem() {
	// Generate key pair
	esign.CreateKeyPair().Save(DEMO_PATH + "/keystore/key")
}

func loadKeysAndSign() {
	// Load key from filesystem and use it to sign any file
	privateKey, _ := esign.LoadPrivateKey(DEMO_PATH + "/keystore/key.prv")
	esign.Sign(DEMO_PATH + "/file.txt", privateKey).Save(DEMO_PATH + "/file.sig")
}

func verifyFile() {
	publicKey, _ := esign.LoadPublicKey(DEMO_PATH + "/keystore/key.pub")

	verifiedCase1 := esign.Verify(DEMO_PATH + "/file.txt").
		With(publicKey).
		BySavedSignature(DEMO_PATH + "/file.sig")

	fmt.Printf("%v\n", verifiedCase1)

	rawSignature, _ := ioutil.ReadFile(DEMO_PATH + "/file.sig")

	verifiedCase2 := esign.Verify(DEMO_PATH + "/file.txt").
		With(publicKey).
		ByRawSignature(rawSignature)

	fmt.Printf("%v\n", verifiedCase2)
}
```

## Contributing

Please read [CONTRIBUTING.md](https://github.com/estatelabs/esign/blob/master/.github/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/estatelabs/esign/tags). 

## Authors

* **Fabian Sander**

See also the list of [contributors](https://github.com/estatelabs/esign/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
