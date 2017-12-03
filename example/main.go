package main

import (
	"encoding/hex"
	"fmt"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func main() {
	hash := solsha3.SoliditySHA3(
		solsha3.Address("0x12459c951127e0c374ff9105dda097662a027093"),
		solsha3.Uint256(100),
		solsha3.String("foo"),
		solsha3.Bytes32("bar"),
		solsha3.Bool(true),
	)

	fmt.Println(hex.EncodeToString(hash))
}
