package solsha3

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// Address address
func Address(addr string) []byte {
	return common.HexToAddress(addr).Bytes()[:]
}

// Uint256 uint256
func Uint256(n *big.Int) []byte {
	return abi.U256(n)
}

// Uint256FromString uint256
func Uint256FromString(s string) []byte {
	n := new(big.Int)
	n.SetString(s, 10)
	return Uint256(n)
}

// Uint128 uint128
func Uint128(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 16)
}

// Uint64 uint64
func Uint64(n uint64) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

// Uint32 uint32
func Uint32(n uint32) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

// Uint16 uint16
func Uint16(n uint16) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

// Uint8 uint8
func Uint8(n uint8) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

// Int256 int256
func Int256(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 32)
}

// Int128 int128
func Int128(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 16)
}

// Int64 int64
func Int64(n int64) []byte {
	i := uint64(n)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// Int32 int32
func Int32(n int32) []byte {
	i := uint32(n)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

// Int16 int16
func Int16(n int16) []byte {
	i := uint16(n)
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

// Int8 int8
func Int8(n int8) []byte {
	b := make([]byte, 1)
	b[0] = byte(n)
	return b
}

// Bytes32 bytes32
func Bytes32(s string) []byte {
	return common.RightPadBytes([]byte(s), 32)
}

// String string
func String(s string) []byte {
	return []byte(s)
}

// Bool bool
func Bool(t bool) []byte {
	if t {
		return []byte{0x1}
	}

	return []byte{0x0}
}

// ConcatByteSlices concat byte slices
func ConcatByteSlices(arrays ...[]byte) []byte {
	var result []byte

	for _, b := range arrays {
		result = append(result, b...)
	}

	return result
}

// SoliditySHA3 solidity sha3
func SoliditySHA3(data ...[]byte) []byte {
	var result []byte

	hash := sha3.NewKeccak256()
	bs := ConcatByteSlices(data...)

	hash.Write(bs)
	result = hash.Sum(result)

	return result
}

// SoliditySHA3WithPrefix solidity sha3 with prefix
func SoliditySHA3WithPrefix(data []byte) []byte {
	result := SoliditySHA3(
		ConcatByteSlices(
			[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(data))),
			data,
		),
	)

	return result
}
