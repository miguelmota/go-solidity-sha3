package solsha3

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

// Bool bool
func Bool(input interface{}) []byte {
	switch v := input.(type) {
	case bool:
		if v {
			return []byte{0x1}
		}
		return []byte{0x0}
	}

	if isArray(input) {
		return BoolArray(input)
	}

	return []byte{0x0}
}

// BoolArray bool array
func BoolArray(input interface{}) []byte {
	var values []byte
	s := reflect.ValueOf(input)
	for i := 0; i < s.Len(); i++ {
		val := s.Index(i).Interface()
		result := common.LeftPadBytes(Bool(val), 32)
		values = append(values, result...)
	}
	return values
}
