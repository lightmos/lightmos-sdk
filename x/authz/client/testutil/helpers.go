package authz

import (
	"github.com/lightmos/lightmos-sdk/client"
	"github.com/lightmos/lightmos-sdk/testutil"
	clitestutil "github.com/lightmos/lightmos-sdk/testutil/cli"
	"github.com/lightmos/lightmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
