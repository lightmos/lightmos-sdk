//go:build gofuzz || go1.18

package tests

import (
	"testing"

	"github.com/lightmos/lightmos-sdk/x/bank/types"
)

func FuzzXBankTypesAddressFromBalancesStore(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		types.AddressAndDenomFromBalancesStore(data)
	})
}
