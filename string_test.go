package solsha3

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					[]string{"string"},
					"somedata",
				),
				"fb763c3da6141a6a1464a68583e30d9a77bb999b1f1c491992dcfac7738ecfb4",
			},
			{
				SoliditySHA3(
					[]string{"string[]"},
					[]string{"a", "b", "c"},
				),
				"4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45",
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
