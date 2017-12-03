package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"math/big"
)

func Address(addr string) []byte {
	return common.HexToAddress(addr).Bytes()[:]
}

func Uint256(n int64) []byte {
	return abi.U256(big.NewInt(n))
}

/*
func Uint32(n uint64) []byte {
	r := abi.U256(new(big.Int).SetUint64(n))
	return r
}
*/

func Bytes32(s string) []byte {
	return common.RightPadBytes([]byte(s), 32)
}

func String(s string) []byte {
	return []byte(s)
}

func Bool(t bool) []byte {
	if t {
		return []byte{0x1}
	}

	return []byte{0x0}
}

func ConcatByteSlices(arrays ...[]byte) []byte {
	var result []byte

	for _, b := range arrays {
		result = append(result, b...)
	}

	return result
}

func SoliditySHA3(data ...[]byte) []byte {
	var result []byte

	hash := sha3.NewKeccak256()
	bs := ConcatByteSlices(data...)

	hash.Write(bs)
	result = hash.Sum(result)

	return result
}
