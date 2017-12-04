package solsha3

import (
	"bytes"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"math/big"
)

func Address(addr string) []byte {
	return common.HexToAddress(addr).Bytes()[:]
}

func Uint256(n *big.Int) []byte {
	return abi.U256(n)
}

func Uint256FromString(s string) []byte {
	n := new(big.Int)
	n.SetString(s, 10)
	return Uint256(n)
}

func Uint128(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 16)
}

func Uint64(n uint64) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

func Uint32(n uint32) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

func Uint16(n uint16) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

func Uint8(n uint8) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, n)
	return b.Bytes()
}

func Int256(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 32)
}

func Int128(n *big.Int) []byte {
	return common.LeftPadBytes(n.Bytes(), 16)
}

func Int64(n int64) []byte {
	i := uint64(n)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func Int32(n int32) []byte {
	i := uint32(n)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func Int16(n int16) []byte {
	i := uint16(n)
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func Int8(n int8) []byte {
	b := make([]byte, 1)
	b[0] = byte(n)
	return b
}

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
