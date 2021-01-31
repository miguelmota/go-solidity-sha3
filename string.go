package solsha3

import "reflect"

// String string
func String(input interface{}) []byte {
	switch v := input.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	}

	if isArray(input) {
		return StringArray(input)
	}

	return []byte("")
}

// StringArray string
func StringArray(input interface{}) []byte {
	var values []byte
	s := reflect.ValueOf(input)
	for i := 0; i < s.Len(); i++ {
		val := s.Index(i).Interface()
		result := String(val)
		values = append(values, result...)
	}
	return values
}
