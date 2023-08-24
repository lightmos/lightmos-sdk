package cli_test

import (
	"io"
	"testing"

	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"
	"github.com/stretchr/testify/suite"

	"github.com/lightmos/lightmos-sdk/client"
	"github.com/lightmos/lightmos-sdk/crypto/keyring"
	clitestutil "github.com/lightmos/lightmos-sdk/testutil/cli"
	testutilmod "github.com/lightmos/lightmos-sdk/types/module/testutil"
	"github.com/lightmos/lightmos-sdk/x/bank"
)

type CLITestSuite struct {
	suite.Suite

	kr      keyring.Keyring
	encCfg  testutilmod.TestEncodingConfig
	baseCtx client.Context
}

func TestMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}

func (s *CLITestSuite) SetupSuite() {
	s.encCfg = testutilmod.MakeTestEncodingConfig(bank.AppModuleBasic{})
	s.kr = keyring.NewInMemory(s.encCfg.Codec)
	s.baseCtx = client.Context{}.
		WithKeyring(s.kr).
		WithTxConfig(s.encCfg.TxConfig).
		WithCodec(s.encCfg.Codec).
		WithClient(clitestutil.MockTendermintRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard)
}
