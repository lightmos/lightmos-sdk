package types

import (
	"github.com/lightmos/lightmos-sdk/codec"
	"github.com/lightmos/lightmos-sdk/codec/legacy"
	"github.com/lightmos/lightmos-sdk/codec/types"
	cryptocodec "github.com/lightmos/lightmos-sdk/crypto/codec"
	sdk "github.com/lightmos/lightmos-sdk/types"
	"github.com/lightmos/lightmos-sdk/types/msgservice"
	"github.com/lightmos/lightmos-sdk/x/authz"
	authzcodec "github.com/lightmos/lightmos-sdk/x/authz/codec"
	govcodec "github.com/lightmos/lightmos-sdk/x/gov/codec"
	groupcodec "github.com/lightmos/lightmos-sdk/x/group/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgSend{}, "cosmos-sdk/MsgSend")
	legacy.RegisterAminoMsg(cdc, &MsgMultiSend{}, "cosmos-sdk/MsgMultiSend")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/bank/MsgUpdateParams")
	legacy.RegisterAminoMsg(cdc, &MsgSetSendEnabled{}, "cosmos-sdk/MsgSetSendEnabled")

	cdc.RegisterConcrete(&SendAuthorization{}, "cosmos-sdk/SendAuthorization", nil)
	cdc.RegisterConcrete(&Params{}, "cosmos-sdk/x/bank/Params", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
		&MsgUpdateParams{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SendAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
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
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}
