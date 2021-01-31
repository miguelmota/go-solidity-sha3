package solsha3

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var zeros = "0000000000000000000000000000000000000000000000000000000000000000"

func concatByteSlices(arrays ...[]byte) []byte {
	var result []byte

	for _, b := range arrays {
		result = append(result, b...)
	}

	return result
}

func isArray(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Array ||
		reflect.TypeOf(value).Kind() == reflect.Slice
}

func padZeros(value []byte, width int) []byte {
	return common.LeftPadBytes(value, width)
}

func pack(typ string, value interface{}, _isArray bool) []byte {
	switch typ {
	case "address":
		if _isArray {
			return padZeros(Address(value), 32)
		}

		return Address(value)
	case "string":
		return String(value)
	case "bool":
		if _isArray {
			return padZeros(Bool(value), 32)
		}

		return Bool(value)
	}

	regexNumber := regexp.MustCompile(`^(u?int)([0-9]*)$`)
	matches := regexNumber.FindAllStringSubmatch(typ, -1)
	if len(matches) > 0 {
		match := matches[0]
		var err error
		size := 256
		if len(match) > 1 {
			//signed = match[1] == "int"
		}
		if len(match) > 2 {
			size, err = strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
		}

		_ = size
		if (size%8 != 0) || size == 0 || size > 256 {
			panic("invalid number type " + typ)
		}

		if _isArray {
			size = 256
		}

		var v []byte
		if strings.HasPrefix(typ, "int8") {
			v = Int8(value)
		} else if strings.HasPrefix(typ, "int16") {
			v = Int16(value)
		} else if strings.HasPrefix(typ, "int32") {
			v = Int32(value)
		} else if strings.HasPrefix(typ, "int64") {
			v = Int64(value)
		} else if strings.HasPrefix(typ, "int128") {
			v = Int128(value)
		} else if strings.HasPrefix(typ, "int256") {
			v = Int256(value)
		} else if strings.HasPrefix(typ, "uint8") {
			v = Uint8(value)
		} else if strings.HasPrefix(typ, "uint16") {
			v = Uint16(value)
		} else if strings.HasPrefix(typ, "uint32") {
			v = Uint32(value)
		} else if strings.HasPrefix(typ, "uint128") {
			v = Uint128(value)
		} else if strings.HasPrefix(typ, "uint64") {
			v = Uint64(value)
		} else if strings.HasPrefix(typ, "uint256") {
			v = Uint256(value)
		}
		return padZeros(v, size/8)
	}

	regexBytes := regexp.MustCompile(`^bytes([0-9]+)$`)
	matches = regexBytes.FindAllStringSubmatch(typ, -1)
	if len(matches) > 0 {
		match := matches[0]

		byteLen, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		strSize := strconv.Itoa(byteLen)
		if strSize != match[1] || byteLen == 0 || byteLen > 32 {
			panic("invalid number type " + typ)
		}

		if _isArray {
			s := reflect.ValueOf(value)
			v := s.Index(0).Bytes()
			z := make([]byte, 64)
			copy(z[:], v[:])
			return z[:]
		}

		str, isString := value.(string)
		if isString && isHex(str) {
			s := removeHexPrefix(str)
			if len(s)%2 == 1 {
				s = "0" + s
			}
			hexb, err := hex.DecodeString(s)
			if err != nil {
				panic(err)
			}
			z := make([]byte, byteLen)
			copy(z[:], hexb)
			return z
		} else if isString {
			s := reflect.ValueOf(value)
			z := make([]byte, byteLen)
			copy(z[:], s.Bytes())
			return z
		}

		s := reflect.ValueOf(value)
		z := make([]byte, byteLen)
		b := make([]byte, s.Len())
		for i := 0; i < s.Len(); i++ {
			ifc := s.Index(i).Interface()
			v, ok := ifc.(byte)
			if ok {
				b[i] = v
			} else {
				v, ok := ifc.(string)
				if ok {
					v = removeHexPrefix(v)
					if len(v)%2 == 1 {
						v = "0" + v
					}
					decoded, err := hex.DecodeString(v)
					if err != nil {
						panic(err)
					}
					b[i] = decoded[0]
				}
			}

		}
		copy(z[:], b[:])
		return z
	}

	regexArray := regexp.MustCompile(`^(.*)\[([0-9]*)\]$`)
	matches = regexArray.FindAllStringSubmatch(typ, -1)
	if len(matches) > 0 {
		match := matches[0]

		_ = match
		if isArray(value) {
			baseType := match[1]
			k := reflect.ValueOf(value)
			count := k.Len()
			var err error
			if len(match) > 1 && match[2] != "" {
				count, err = strconv.Atoi(match[2])
				if err != nil {
					panic(err)
				}
			}
			if count != k.Len() {
				panic("invalid value for " + typ)
			}

			var result [][]byte
			for i := 0; i < k.Len(); i++ {
				val := k.Index(i).Interface()

				result = append(result, pack(baseType, val, true))
			}

			return concatByteSlices(result...)
		}
	}
	return nil
}

func removeHexPrefix(str string) string {
	return strings.TrimPrefix(str, "0x")
}

func isHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}

func stringToBytes(str string) []byte {
	if isHex(str) {
		str = evenLengthHex(removeHexPrefix(str))
	} else {
		str = fmt.Sprintf("%0x", str)
	}

	hexb, _ := hex.DecodeString(str)
	return hexb
}

func evenLengthHex(v string) string {
	if len(v)%2 == 1 {
		v = "0" + v
	}
	return v
}
