package main

import (
	"encoding/hex"
	"fmt"

	"github.com/miguelmota/go-solidity-sha3"
)

func main() {
	hash := solsha3.SoliditySHA3(
		solsha3.Uint8Array([3]uint{1, 2, 3}),
	)

	fmt.Println(hex.EncodeToString(hash)) // 6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c
}
