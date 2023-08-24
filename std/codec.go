package std

import (
	"github.com/lightmos/lightmos-sdk/codec"
	"github.com/lightmos/lightmos-sdk/codec/types"
	cryptocodec "github.com/lightmos/lightmos-sdk/crypto/codec"
	sdk "github.com/lightmos/lightmos-sdk/types"
	txtypes "github.com/lightmos/lightmos-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
