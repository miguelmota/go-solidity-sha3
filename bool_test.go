package solsha3

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					[]string{"bool"},
					true,
				),
				"5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2",
			},
			{
				SoliditySHA3(
					[]string{"bool"},
					false,
				),
				"bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a",
			},
			{
				SoliditySHA3(
					[]string{"bool[]"},
					[3]bool{true, false, true},
				),
				"5c6090c0461491a2941743bda5c3658bf1ea53bbd3edcde54e16205e18b45792",
			},
		} {
			t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
				got := hex.EncodeToString(tt.in)
				if got != tt.out {
					t.Errorf("want %v, got %v", tt.out, got)
				}
			})
		}
	})
}
