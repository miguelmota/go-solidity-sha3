package solsha3

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestSolSha3Legacy(t *testing.T) {
	t.Run("address", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(Address(0)),
				"5380c7b7ae81a58eb98d9c78de4a1fd7fd9535fc953ed2be602daaa41767312a",
			},
			{
				SoliditySHA3(Address("0x0a")),

				"0ef9d8f8804d174666011a394cab7901679a8944d24249fd148a6a36071151f8",
			},
			{
				SoliditySHA3(Address("0x12459c951127e0c374ff9105dda097662a027092")),
				"4b998b071d7bb74aee1ce2cdcc268cb0f6409b4a3387fc915617ec08415298ad",
			},
			{
				SoliditySHA3(
					Address(
						common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027092"),
					),
				),
				"4b998b071d7bb74aee1ce2cdcc268cb0f6409b4a3387fc915617ec08415298ad",
			},
			{
				SoliditySHA3(Address([]string{"0x0a", "0x0b", "0x0c"})),

				"63f3beae5de2bbda3d06f2c0158ccedcdce66256efcf2f49930ca1c6976979df",
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

	t.Run("string", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					String("somedata"),
				),
				"fb763c3da6141a6a1464a68583e30d9a77bb999b1f1c491992dcfac7738ecfb4",
			},
			{
				SoliditySHA3(
					String([]string{"a", "b", "c"}),
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

	t.Run("bool", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Bool(true),
				),
				"5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2",
			},
			{
				SoliditySHA3(
					Bool(false),
				),
				"bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a",
			},
			{
				SoliditySHA3(
					Bool([3]bool{true, false, true}),
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

	t.Run("prefix", func(t *testing.T) {
		msg := []byte("hello")
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3WithPrefix(msg),
				"50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750",
			},
			{
				SoliditySHA3(
					ConcatByteSlices(
						[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(msg))),
						msg,
					),
				),
				"50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750",
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

func TestSolSha3(t *testing.T) {
	{
		hash := SoliditySHA3(
			[]string{"uint8[]"},
			[3]uint8{1, 2, 3},
		)

		fmt.Println(hex.EncodeToString(hash))
		//6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c

		hash2 := SoliditySHA3(
			Uint8([3]uint8{1, 2, 3}),
		)

		fmt.Println(hex.EncodeToString(hash2))
	}

	{
		hash := SoliditySHA3(
			[]string{"address"},
			"0x0a",
		)

		fmt.Println(hex.EncodeToString(hash))

		hash2 := SoliditySHA3(
			Address("0x0a"),
		)
		fmt.Println(hex.EncodeToString(hash2))
	}

	{
		hash := SoliditySHA3(
			[]string{"uint16"},
			uint16(8),
		)

		fmt.Println(hex.EncodeToString(hash))

		hash2 := SoliditySHA3(
			Uint16(8),
		)

		fmt.Println(hex.EncodeToString(hash2))
	}

	{
		hash := SoliditySHA3(
			[]string{"bytes32"},
			"0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45",
		)

		fmt.Println(hex.EncodeToString(hash))

		hash2 := SoliditySHA3(
			Bytes32("0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45"),
		)

		fmt.Println(hex.EncodeToString(hash2))
	}

	{
		hash := SoliditySHA3(
			[]string{"address", "bytes1", "uint8[]", "bytes32", "uint256", "address[]", "uint32"},
			"0x935F7770265D0797B621c49A5215849c333Cc3ce",
			"0xa",
			[]uint8{128, 255},
			"0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45",
			"100000000000000000",
			[]string{
				"0x13D94859b23AF5F610aEfC2Ae5254D4D7E3F191a",
				"0x473029549e9d898142a169d7234c59068EDcBB33",
			},
			123456789,
		)

		expected := "ad390a98c1c32cdb1f046f6887a4109f12290b690127e6e15da4ca210235510e"
		result := hex.EncodeToString(hash)
		if result != expected {
			t.Error(result, expected)
		}
	}

	{
		hash := SoliditySHA3(
			[]string{"address", "uint256"},
			"0x935F7770265D0797B621c49A5215849c333Cc3ce",
			"100000000000000000",
		)

		expected := "0a3844b522d9e3a837ae56d4c57d668feb26325834bf4ba49e153d84ed7ad53d"
		result := hex.EncodeToString(hash)
		if result != expected {
			t.Error(result, expected)
		}
	}

	{
		hash := SoliditySHA3(
			[]string{"address", "uint256"},
			"0x935F7770265D0797B621c49A5215849c333Cc3ce",
			"100000000000000000",
		)

		expected := "0a3844b522d9e3a837ae56d4c57d668feb26325834bf4ba49e153d84ed7ad53d"
		result := hex.EncodeToString(hash)
		if result != expected {
			t.Error(result, expected)
		}
	}

	{
		types := []string{"address", "uint256"}
		inputs := []interface{}{
			"0x935F7770265D0797B621c49A5215849c333Cc3ce",
			"100000000000000000",
		}

		hash := SoliditySHA3(types, inputs)

		expected := "0a3844b522d9e3a837ae56d4c57d668feb26325834bf4ba49e153d84ed7ad53d"
		result := hex.EncodeToString(hash)
		if result != expected {
			t.Error(result, expected)
		}
	}
}
