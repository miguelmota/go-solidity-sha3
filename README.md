<h3 align="center">
  <br />
  <img src="https://user-images.githubusercontent.com/168240/51433394-10dbe380-1bfe-11e9-86c8-d4d57f77fb11.png" alt="logo" width="700" />
  <br />
  <br />
  <br />
</h3>

# go-solidity-sha3

> Generate Solidity SHA3 (Keccak256) hashes in Go.

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/go-solidity-sha3/master/LICENSE.md) [![Build Status](https://travis-ci.org/miguelmota/go-solidity-sha3.svg?branch=master)](https://travis-ci.org/miguelmota/go-solidity-sha3) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/go-solidity-sha3?)](https://goreportcard.com/report/github.com/miguelmota/go-solidity-sha3) [![GoDoc](https://godoc.org/github.com/miguelmota/go-solidity-sha3?status.svg)](https://godoc.org/github.com/miguelmota/go-solidity-sha3)

This package is the Go equivalent of `require('ethers).utils.solidityKeccak256` [NPM module](https://www.npmjs.com/package/ethers).

## Install

```bash
go get -u github.com/miguelmota/go-solidity-sha3
```

## Documentation

[Documentation on GoDoc](https://godoc.org/github.com/miguelmota/go-solidity-sha3)

## Getting started

Simple example:

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/miguelmota/go-solidity-sha3"
)

func main() {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"address", "uint256"},

		// values
		"0x935F7770265D0797B621c49A5215849c333Cc3ce",
		"100000000000000000",
	)

	fmt.Println(hex.EncodeToString(hash))
}
```

Output:

```bash
0a3844b522d9e3a837ae56d4c57d668feb26325834bf4ba49e153d84ed7ad53d
```

More complex example:

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/miguelmota/go-solidity-sha3"
)

func main() {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"address", "bytes1", "uint8[]", "bytes32", "uint256", "address[]", "uint32"},

		// values
		"0x935F7770265D0797B621c49A5215849c333Cc3ce",
		"0xa",
		[]uint8{128, 255},
		"0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45",
		"100000000000000000",
		[]string{
			"0x13D94859b23AF5F610aEfC2Ae5254D4D7E3F191a",
			"0x473029549e9d898142a169d7234c59068EDcBB33",
		},
		123456789,
	)

	fmt.Println(hex.EncodeToString(hash))
}
```

Output:

```bash
ad390a98c1c32cdb1f046f6887a4109f12290b690127e6e15da4ca210235510e
```

## License

MIT
