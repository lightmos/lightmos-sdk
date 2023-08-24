package codec

import (
	"github.com/lightmos/lightmos-sdk/codec"
	cryptocodec "github.com/lightmos/lightmos-sdk/crypto/codec"
	sdk "github.com/lightmos/lightmos-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	cryptocodec.RegisterCrypto(Amino)
	codec.RegisterEvidences(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
