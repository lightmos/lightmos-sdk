package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/lightmos/lightmos-sdk/codec"
	codectypes "github.com/lightmos/lightmos-sdk/codec/types"
	"github.com/lightmos/lightmos-sdk/std"
	"github.com/lightmos/lightmos-sdk/testutil/testdata"
	sdk "github.com/lightmos/lightmos-sdk/types"
	txtestutil "github.com/lightmos/lightmos-sdk/x/auth/tx/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, txtestutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
