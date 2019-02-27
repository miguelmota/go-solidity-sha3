package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/miguelmota/go-solidity-sha3"
)

func main() {
	hash := solsha3.SoliditySHA3(
		solsha3.Address("0x12459c951127e0c374ff9105dda097662a027093"),
		solsha3.Uint256(big.NewInt(100)),
		solsha3.String("foo"),
		solsha3.Bytes32("bar"),
		solsha3.Bool(true),
	)

	fmt.Println(hex.EncodeToString(hash)) // 417a4c44724701ba79bb363151dff48909bc058a2c75a81e9cf5208ae4699369

	hash2 := solsha3.SoliditySHA3WithPrefix(
		solsha3.String("hello"),
	)

	fmt.Println(hex.EncodeToString(hash2)) // 50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750
}
