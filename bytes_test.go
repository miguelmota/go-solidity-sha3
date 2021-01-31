package solsha3

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	t.Run("bytes1", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Bytes1("0x5"),
				),
				"dbb8d0f4c497851a5043c6363657698cb1387682cac2f786c731f8936109d795",
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

	t.Run("bytes2", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Bytes2("0x1234"),
				),
				"56570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
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

	t.Run("bytes4", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Bytes4("0xe0b6fcfc"),
				),
				"a1204967c3aa63863e35064a313503edf747e6de2721e9dec149014654fddbba",
			},
			{
				SoliditySHA3(
					[4]byte{0xe0, 0xb6, 0xfc, 0xfc},
				),
				"a1204967c3aa63863e35064a313503edf747e6de2721e9dec149014654fddbba",
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

	t.Run("bytes16", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Bytes4("0x12341234123412341234123412341234"),
				),
				"0dd1e9e300f81dfb0638f199ac3dd3d605009475579a3925bb903909bd6ff1e0",
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

	t.Run("bytes32", func(t *testing.T) {
		for i, tt := range []struct {
			in  []byte
			out string
		}{
			{
				SoliditySHA3(
					Bytes32("somedata"),
				),
				"5de179cfd5920200431577c32d9c35e2b0bc5d8eaae3534e8146da8ab45be70b",
			},
			{
				SoliditySHA3(
					Bytes32("a"),
				),
				"294587bf977c4010a60dbad811c63531f90f6ec512975bc6c9a93f8f361cad72",
			},
			{
				SoliditySHA3(
					Bytes32([32]byte{'a'}),
				),
				"294587bf977c4010a60dbad811c63531f90f6ec512975bc6c9a93f8f361cad72",
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

	t.Run("bytes32", func(t *testing.T) {
		for i, tt := range []struct {
			in  func() []byte
			out string
		}{
			{
				func() []byte {
					in := make([]byte, 32)
					copy(in[:], []byte("somedata"))
					return SoliditySHA3(
						[]string{"bytes32"},
						in,
					)
				},
				"5de179cfd5920200431577c32d9c35e2b0bc5d8eaae3534e8146da8ab45be70b",
			},
			{
				func() []byte {
					in := make([]byte, 32)
					copy(in[:], []byte("a"))
					return SoliditySHA3(
						[]string{"bytes32"},
						in,
					)
				},
				"294587bf977c4010a60dbad811c63531f90f6ec512975bc6c9a93f8f361cad72",
			},
			{
				func() []byte {
					return SoliditySHA3(
						[]string{"bytes32"},
						[32]byte{'a'},
					)
				},
				"294587bf977c4010a60dbad811c63531f90f6ec512975bc6c9a93f8f361cad72",
			},
			{
				func() []byte {
					return SoliditySHA3(
						[]string{"bytes32"},
						[32]byte{0x0a},
					)
				},
				"cee931476953956236065da391361912c6a45bc675f9f4e63b8d7bff037b43ae",
			},
			{
				func() []byte {
					return SoliditySHA3(
						[]string{"bytes32"},
						"0xa000000000000000000000000000000000000000000000000000000000000000",
					)
				},
				"2c2171a229511a789fd39fbdea82b2f459f0d6f700e745d01bf69d70230ea3d7",
			},
			{
				func() []byte {
					return SoliditySHA3(
						[]string{"bytes1"},
						"0x04",
					)
				},
				"f343681465b9efe82c933c3e8748c70cb8aa06539c361de20f72eac04e766393",
			},
		} {
			t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
				got := hex.EncodeToString(tt.in())
				if got != tt.out {
					t.Errorf("want %v, got %v", tt.out, got)
				}
			})
		}
	})
}
