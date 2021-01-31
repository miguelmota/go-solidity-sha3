package solsha3

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

func TestUint(t *testing.T) {
	t.Run("uint/int", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Uint256(big.NewInt(100)),
				),
				"26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2",
			},
			{
				SoliditySHA3(
					Uint256("100"),
				),
				"26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2",
			},
			{
				SoliditySHA3(
					Uint128(big.NewInt(100)),
				),
				"e676e20e5d612283600da1ff24a86bdbc07d286dfe8e6afba02988c485a6749d",
			},
			{
				SoliditySHA3(
					Uint64(100),
				),
				"dc961d2e5d46532f7b9f65afdb640c5284dfb1a947abf2b23758931400995e0f",
			},
			{
				SoliditySHA3(
					Uint32(100),
				),
				"f01681d2220bfea4bb888a5543db8c0916274ddb1ea93b144c042c01d8164c95",
			},
			{
				SoliditySHA3(
					Uint16(100),
				),
				"2ce80d2bc0bfe54c2499d066ac958c02304ce64ca318ae19a4636c32d583429c",
			},
			{
				SoliditySHA3(
					Uint8(100),
				),
				"f1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3",
			},
			{
				SoliditySHA3(
					Uint8([3]uint8{1, 2, 3}),
				),
				"6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c",
			},
			{
				SoliditySHA3(
					Uint8([]uint8{1, 2, 3}),
				),
				"6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c",
			},
			{
				SoliditySHA3(
					Int256(big.NewInt(100)),
				),
				"26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2",
			},
			{
				SoliditySHA3(
					Int128(big.NewInt(100)),
				),
				"e676e20e5d612283600da1ff24a86bdbc07d286dfe8e6afba02988c485a6749d",
			},
			{
				SoliditySHA3(
					Int64(100),
				),
				"dc961d2e5d46532f7b9f65afdb640c5284dfb1a947abf2b23758931400995e0f",
			},
			{
				SoliditySHA3(
					Int32(100),
				),
				"f01681d2220bfea4bb888a5543db8c0916274ddb1ea93b144c042c01d8164c95",
			},
			{
				SoliditySHA3(
					Int16(100),
				),
				"2ce80d2bc0bfe54c2499d066ac958c02304ce64ca318ae19a4636c32d583429c",
			},
			{
				SoliditySHA3(
					Int16(-1000),
				),
				"779c8f2f6090bb23de632d95b8edb88a9473562dbe3910f0cefdd2b3d969d4da",
			},
			{
				SoliditySHA3(
					Int8(100),
				),
				"f1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3",
			},
			{
				SoliditySHA3(
					Int8([3]int8{1, 2, 3}),
				),
				"6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c",
			},
			{
				SoliditySHA3(
					Int8(-100),
				),
				"26ce63779ce95adee31308d97bf37df2920dcdbb38c8db0620241fe2f08c8ed9",
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

	t.Run("uint/int", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					[]string{"uint256"},
					big.NewInt(100),
				),
				"26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2",
			},
			{
				SoliditySHA3(
					[]string{"uint256"},
					"100",
				),
				"26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2",
			},
			{
				SoliditySHA3(
					[]string{"uint128"},
					big.NewInt(100),
				),
				"e676e20e5d612283600da1ff24a86bdbc07d286dfe8e6afba02988c485a6749d",
			},
			{
				SoliditySHA3(
					[]string{"uint64"},
					100,
				),
				"dc961d2e5d46532f7b9f65afdb640c5284dfb1a947abf2b23758931400995e0f",
			},
			{
				SoliditySHA3(
					[]string{"uint32"},
					100,
				),
				"f01681d2220bfea4bb888a5543db8c0916274ddb1ea93b144c042c01d8164c95",
			},
			{
				SoliditySHA3(
					[]string{"uint16"},
					100,
				),
				"2ce80d2bc0bfe54c2499d066ac958c02304ce64ca318ae19a4636c32d583429c",
			},
			{
				SoliditySHA3(
					[]string{"uint8"},
					100,
				),
				"f1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3",
			},
			{
				SoliditySHA3(
					Uint8([3]uint8{1, 2, 3}),
				),
				"6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c",
			},
			{
				SoliditySHA3(
					Uint8([]uint8{1, 2, 3}),
				),
				"6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c",
			},
			{
				SoliditySHA3(
					[]string{"int256"},
					big.NewInt(100),
				),
				"26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2",
			},
			{
				SoliditySHA3(
					[]string{"int128"},
					big.NewInt(100),
				),
				"e676e20e5d612283600da1ff24a86bdbc07d286dfe8e6afba02988c485a6749d",
			},
			{
				SoliditySHA3(
					[]string{"int64"},
					100,
				),
				"dc961d2e5d46532f7b9f65afdb640c5284dfb1a947abf2b23758931400995e0f",
			},
			{
				SoliditySHA3(
					[]string{"int32"},
					100,
				),
				"f01681d2220bfea4bb888a5543db8c0916274ddb1ea93b144c042c01d8164c95",
			},
			{
				SoliditySHA3(
					[]string{"int16"},
					100,
				),
				"2ce80d2bc0bfe54c2499d066ac958c02304ce64ca318ae19a4636c32d583429c",
			},
			{
				SoliditySHA3(
					[]string{"int16"},
					-1000,
				),
				"779c8f2f6090bb23de632d95b8edb88a9473562dbe3910f0cefdd2b3d969d4da",
			},
			{
				SoliditySHA3(
					[]string{"int8"},
					100,
				),
				"f1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3",
			},
			{
				SoliditySHA3(
					Int8([3]int8{1, 2, 3}),
				),
				"6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c",
			},
			{
				SoliditySHA3(
					[]string{"int8"},
					-100,
				),
				"26ce63779ce95adee31308d97bf37df2920dcdbb38c8db0620241fe2f08c8ed9",
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
