package solsha3

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAddress(t *testing.T) {
	t.Run("address", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					[]string{"address"},
					"0x0",
				),
				"bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a",
			},
			{
				SoliditySHA3(
					[]string{"address"},
					"0x0a",
				),

				"0ef9d8f8804d174666011a394cab7901679a8944d24249fd148a6a36071151f8",
			},
			{
				SoliditySHA3(
					[]string{"address"},
					"0x12459c951127e0c374ff9105dda097662a027092",
				),
				"4b998b071d7bb74aee1ce2cdcc268cb0f6409b4a3387fc915617ec08415298ad",
			},
			{
				SoliditySHA3(
					[]string{"address"},
					common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027092"),
				),
				"4b998b071d7bb74aee1ce2cdcc268cb0f6409b4a3387fc915617ec08415298ad",
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
