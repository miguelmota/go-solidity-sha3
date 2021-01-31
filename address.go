package solsha3

import (
	"encoding/hex"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

// Address address
func Address(input interface{}) []byte {
	switch v := input.(type) {
	case common.Address:
		return v.Bytes()[:]
	case string:
		v = removeHexPrefix(v)
		if v == "" || v == "0" {
			return []byte{0}
		}

		v = evenLengthHex(v)
		decoded, err := hex.DecodeString(v)
		if err != nil {
			panic(err)
		}

		return decoded
	case []byte:
		return v
	}

	if isArray(input) {
		return AddressArray(input)
	}

	return common.HexToAddress("").Bytes()[:]
}

// AddressArray address
func AddressArray(input interface{}) []byte {
	var values []byte
	s := reflect.ValueOf(input)
	for i := 0; i < s.Len(); i++ {
		val := s.Index(i).Interface()
		result := common.LeftPadBytes(Address(val), 32)
		values = append(values, result...)
	}
	return values
}
