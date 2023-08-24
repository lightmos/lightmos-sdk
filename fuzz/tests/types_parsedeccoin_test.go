//go:build gofuzz || go1.18

package tests

import (
	"testing"

	"github.com/lightmos/lightmos-sdk/types"
)

func FuzzTypesParseDecCoin(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		types.ParseDecCoin(string(data))
	})
}
