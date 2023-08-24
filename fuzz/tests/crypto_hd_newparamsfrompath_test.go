//go:build gofuzz || go1.18

package tests

import (
	"testing"

	"github.com/lightmos/lightmos-sdk/crypto/hd"
)

func FuzzCryptoHDNewParamsFromPath(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		hd.NewParamsFromPath(string(data))
	})
}
