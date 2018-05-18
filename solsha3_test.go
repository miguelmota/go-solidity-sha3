package solsha3

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestMain(t *testing.T) {
	{
		hash := SoliditySHA3(
			Address("0x12459c951127e0c374ff9105dda097662a027092"),
		)

		expected := "4b998b071d7bb74aee1ce2cdcc268cb0f6409b4a3387fc915617ec08415298ad"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Address(common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027092")),
		)

		expected := "4b998b071d7bb74aee1ce2cdcc268cb0f6409b4a3387fc915617ec08415298ad"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Address(0),
		)

		expected := "5380c7b7ae81a58eb98d9c78de4a1fd7fd9535fc953ed2be602daaa41767312a"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Bytes32("somedata"),
		)

		expected := "5de179cfd5920200431577c32d9c35e2b0bc5d8eaae3534e8146da8ab45be70b"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Bytes32("a"),
		)

		expected := "294587bf977c4010a60dbad811c63531f90f6ec512975bc6c9a93f8f361cad72"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		//hash := common.HexToHash("0x0")
		hash := SoliditySHA3(
			Bytes32([32]byte{'a'}),
		)

		expected := "294587bf977c4010a60dbad811c63531f90f6ec512975bc6c9a93f8f361cad72"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hexb, _ := hex.DecodeString("a")
		hash := SoliditySHA3(
			Bytes32(hexb),
		)

		expected := "290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			String("somedata"),
		)

		expected := "fb763c3da6141a6a1464a68583e30d9a77bb999b1f1c491992dcfac7738ecfb4"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Bool(true),
		)

		expected := "5fe7f977e71dba2ea1a68e21057beebb9be2ac30c6410aa38d4f3fbe41dcffd2"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Bool(false),
		)

		expected := "bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint256(big.NewInt(100)),
		)

		expected := "26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint256("100"),
		)

		expected := "26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint128(big.NewInt(100)),
		)

		expected := "e676e20e5d612283600da1ff24a86bdbc07d286dfe8e6afba02988c485a6749d"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint64(100),
		)

		expected := "dc961d2e5d46532f7b9f65afdb640c5284dfb1a947abf2b23758931400995e0f"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint32(100),
		)

		expected := "f01681d2220bfea4bb888a5543db8c0916274ddb1ea93b144c042c01d8164c95"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint16(100),
		)

		expected := "2ce80d2bc0bfe54c2499d066ac958c02304ce64ca318ae19a4636c32d583429c"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Uint8(100),
		)

		expected := "f1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int256(big.NewInt(100)),
		)

		expected := "26700e13983fefbd9cf16da2ed70fa5c6798ac55062a4803121a869731e308d2"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int128(big.NewInt(100)),
		)

		expected := "e676e20e5d612283600da1ff24a86bdbc07d286dfe8e6afba02988c485a6749d"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int64(100),
		)

		expected := "dc961d2e5d46532f7b9f65afdb640c5284dfb1a947abf2b23758931400995e0f"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int32(100),
		)

		expected := "f01681d2220bfea4bb888a5543db8c0916274ddb1ea93b144c042c01d8164c95"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int16(100),
		)

		expected := "2ce80d2bc0bfe54c2499d066ac958c02304ce64ca318ae19a4636c32d583429c"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int16(-1000),
		)

		expected := "779c8f2f6090bb23de632d95b8edb88a9473562dbe3910f0cefdd2b3d969d4da"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int8(100),
		)

		expected := "f1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		hash := SoliditySHA3(
			Int8(-100),
		)

		expected := "26ce63779ce95adee31308d97bf37df2920dcdbb38c8db0620241fe2f08c8ed9"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3 returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		msg := []byte("hello")
		hash := SoliditySHA3(
			ConcatByteSlices(
				[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(msg))),
				msg,
			),
		)

		expected := "50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3WithPrefix returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}

	{
		msg := []byte("hello")
		hash := SoliditySHA3WithPrefix(msg)

		expected := "50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750"

		if got := hex.EncodeToString(hash); got != expected {
			t.Errorf(
				"SoliditySHA3WithPrefix returned unexpected hash: got %v want %v",
				got,
				expected,
			)
		}
	}
}
