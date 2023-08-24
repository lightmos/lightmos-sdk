package types

import (
	"github.com/lightmos/lightmos-sdk/codec"
	"github.com/lightmos/lightmos-sdk/codec/legacy"
	codectypes "github.com/lightmos/lightmos-sdk/codec/types"
	cryptocodec "github.com/lightmos/lightmos-sdk/crypto/codec"
	sdk "github.com/lightmos/lightmos-sdk/types"
	"github.com/lightmos/lightmos-sdk/types/msgservice"
	authzcodec "github.com/lightmos/lightmos-sdk/x/authz/codec"
	govcodec "github.com/lightmos/lightmos-sdk/x/gov/codec"
	groupcodec "github.com/lightmos/lightmos-sdk/x/group/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/crisis interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgVerifyInvariant{}, "cosmos-sdk/MsgVerifyInvariant")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/crisis/MsgUpdateParams")
}

// RegisterInterfaces registers the interfaces types with the Interface Registry.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifyInvariant{},
		&MsgUpdateParams{},
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

	// Register all Amino interfaces and concrete types on the authz  and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}
