package types

import (
	"github.com/lightmos/lightmos-sdk/codec"
	"github.com/lightmos/lightmos-sdk/codec/legacy"
	"github.com/lightmos/lightmos-sdk/codec/types"
	cryptocodec "github.com/lightmos/lightmos-sdk/crypto/codec"
	sdk "github.com/lightmos/lightmos-sdk/types"
	"github.com/lightmos/lightmos-sdk/types/msgservice"
	authzcodec "github.com/lightmos/lightmos-sdk/x/authz/codec"
	govcodec "github.com/lightmos/lightmos-sdk/x/gov/codec"
	groupcodec "github.com/lightmos/lightmos-sdk/x/group/codec"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the necessary x/consensus interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/consensus/MsgUpdateParams")
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz and gov Amino codec so that this can later be
	// used to properly serialize MsgUpdate instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}
