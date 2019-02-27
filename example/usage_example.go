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

	fmt.Println(hex.EncodeToString(hash)) // ad390a98c1c32cdb1f046f6887a4109f12290b690127e6e15da4ca210235510e
}
