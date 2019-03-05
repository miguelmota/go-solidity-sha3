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
		[]interface{}{
			"0x935F7770265D0797B621c49A5215849c333Cc3ce",
			"100000000000000000",
		},
	)

	fmt.Println(hex.EncodeToString(hash)) // 0a3844b522d9e3a837ae56d4c57d668feb26325834bf4ba49e153d84ed7ad53d
}
