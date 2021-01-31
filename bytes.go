package solsha3

import (
	"github.com/ethereum/go-ethereum/common"
)

// Byte byte
func Byte(input interface{}) []byte {
  return Bytes1(input)
}

// Bytes1 bytes1
func Bytes1(input interface{}) []byte {
  return parseBytes(input, 1)
}

// Bytes2 bytes2
func Bytes2(input interface{}) []byte {
  return parseBytes(input, 2)
}

// Bytes3 bytes3
func Bytes3(input interface{}) []byte {
  return parseBytes(input, 3)
}

// Bytes4 bytes4
func Bytes4(input interface{}) []byte {
  return parseBytes(input, 4)
}

// Bytes5 bytes5
func Bytes5(input interface{}) []byte {
  return parseBytes(input, 5)
}

// Bytes6 bytes6
func Bytes6(input interface{}) []byte {
  return parseBytes(input, 6)
}

// Bytes7 bytes7
func Bytes7(input interface{}) []byte {
  return parseBytes(input, 7)
}

// Bytes8 bytes8
func Bytes8(input interface{}) []byte {
  return parseBytes(input, 8)
}

// Bytes9 bytes9
func Bytes9(input interface{}) []byte {
  return parseBytes(input, 9)
}

// Bytes10 bytes10
func Bytes10(input interface{}) []byte {
  return parseBytes(input, 10)
}


// Bytes11 bytes11
func Bytes11(input interface{}) []byte {
  return parseBytes(input, 11)
}

// Bytes12 bytes12
func Bytes12(input interface{}) []byte {
  return parseBytes(input, 12)
}

// Bytes13 bytes13
func Bytes13(input interface{}) []byte {
  return parseBytes(input, 13)
}

// Bytes14 bytes14
func Bytes14(input interface{}) []byte {
  return parseBytes(input, 14)
}

// Bytes15 bytes15
func Bytes15(input interface{}) []byte {
  return parseBytes(input, 15)
}

// Bytes16 bytes16
func Bytes16(input interface{}) []byte {
  return parseBytes(input, 16)
}

// Bytes17 bytes17
func Bytes17(input interface{}) []byte {
  return parseBytes(input, 17)
}

// Bytes18 bytes18
func Bytes18(input interface{}) []byte {
  return parseBytes(input, 18)
}

// Bytes19 bytes19
func Bytes19(input interface{}) []byte {
  return parseBytes(input, 19)
}

// Bytes20 bytes20
func Bytes20(input interface{}) []byte {
  return parseBytes(input, 20)
}

// Bytes21 bytes21
func Bytes21(input interface{}) []byte {
  return parseBytes(input, 21)
}

// Bytes22 bytes22
func Bytes22(input interface{}) []byte {
  return parseBytes(input, 22)
}

// Bytes23 bytes23
func Bytes23(input interface{}) []byte {
  return parseBytes(input, 23)
}

// Bytes24 bytes24
func Bytes24(input interface{}) []byte {
  return parseBytes(input, 24)
}

// Bytes25 bytes25
func Bytes25(input interface{}) []byte {
  return parseBytes(input, 25)
}

// Bytes26 bytes26
func Bytes26(input interface{}) []byte {
  return parseBytes(input, 26)
}

// Bytes27 bytes27
func Bytes27(input interface{}) []byte {
  return parseBytes(input, 27)
}

// Bytes28 bytes28
func Bytes28(input interface{}) []byte {
  return parseBytes(input, 28)
}

// Bytes29 bytes29
func Bytes29(input interface{}) []byte {
  return parseBytes(input, 29)
}

// Bytes30 bytes30
func Bytes30(input interface{}) []byte {
  return parseBytes(input, 30)
}

// Bytes31 bytes31
func Bytes31(input interface{}) []byte {
  return parseBytes(input, 31)
}

// Bytes32 bytes32
func Bytes32(input interface{}) []byte {
  return parseBytes(input, 32)
}

func rightPadBytes(v []byte, byteLen int) []byte {
  if byteLen == -1 {
    return v
  }
	return common.RightPadBytes(v[:], byteLen)
}

func parseBytes(input interface{}, byteLen int) []byte {
	switch v := input.(type) {
	case [1]byte:
    return rightPadBytes(v[:], byteLen)
	case [2]byte:
    return rightPadBytes(v[:], byteLen)
	case [3]byte:
    return rightPadBytes(v[:], byteLen)
	case [4]byte:
    return rightPadBytes(v[:], byteLen)
	case [5]byte:
    return rightPadBytes(v[:], byteLen)
	case [6]byte:
    return rightPadBytes(v[:], byteLen)
	case [7]byte:
    return rightPadBytes(v[:], byteLen)
	case [8]byte:
    return rightPadBytes(v[:], byteLen)
	case [9]byte:
    return rightPadBytes(v[:], byteLen)
	case [10]byte:
    return rightPadBytes(v[:], byteLen)
	case [11]byte:
    return rightPadBytes(v[:], byteLen)
	case [12]byte:
    return rightPadBytes(v[:], byteLen)
	case [13]byte:
    return rightPadBytes(v[:], byteLen)
	case [14]byte:
    return rightPadBytes(v[:], byteLen)
	case [15]byte:
    return rightPadBytes(v[:], byteLen)
	case [16]byte:
    return rightPadBytes(v[:], byteLen)
	case [17]byte:
    return rightPadBytes(v[:], byteLen)
	case [18]byte:
    return rightPadBytes(v[:], byteLen)
	case [19]byte:
    return rightPadBytes(v[:], byteLen)
	case [20]byte:
    return rightPadBytes(v[:], byteLen)
	case [21]byte:
    return rightPadBytes(v[:], byteLen)
	case [22]byte:
    return rightPadBytes(v[:], byteLen)
	case [23]byte:
    return rightPadBytes(v[:], byteLen)
	case [24]byte:
    return rightPadBytes(v[:], byteLen)
	case [25]byte:
    return rightPadBytes(v[:], byteLen)
	case [26]byte:
    return rightPadBytes(v[:], byteLen)
	case [27]byte:
    return rightPadBytes(v[:], byteLen)
	case [28]byte:
    return rightPadBytes(v[:], byteLen)
	case [29]byte:
    return rightPadBytes(v[:], byteLen)
	case [30]byte:
    return rightPadBytes(v[:], byteLen)
	case [31]byte:
    return rightPadBytes(v[:], byteLen)
	case [32]byte:
    return rightPadBytes(v[:], byteLen)
	case []byte:
    return rightPadBytes(v, byteLen)
	case string:
    hexb := stringToBytes(v)
    h := common.RightPadBytes(hexb, byteLen)
    return h
	default:
		return rightPadBytes([]byte(""), byteLen)
	}
}
