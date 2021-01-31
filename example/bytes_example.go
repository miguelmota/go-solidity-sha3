package main

import (
	"encoding/hex"
	"fmt"

	"github.com/miguelmota/go-solidity-sha3"
)

func main() {
	hash := solsha3.SoliditySHA3(
		solsha3.Bytes4([]byte{0xe0, 0xb6, 0xfc, 0xfc}),
	)

	fmt.Println(hex.EncodeToString(hash)) // 0xa1204967c3aa63863e35064a313503edf747e6de2721e9dec149014654fddbba
}
