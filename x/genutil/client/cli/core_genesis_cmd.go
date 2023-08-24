package cli

import (
	"github.com/lightmos/lightmos-sdk/client"
	"github.com/lightmos/lightmos-sdk/types/module"
	banktypes "github.com/lightmos/lightmos-sdk/x/bank/types"
	"github.com/lightmos/lightmos-sdk/x/genutil"
	genutiltypes "github.com/lightmos/lightmos-sdk/x/genutil/types"

	"github.com/spf13/cobra"
)

// GenesisCoreCommand adds core sdk's sub-commands into genesis command:
// -> gentx, migrate, collect-gentxs, validate-genesis, add-genesis-account
func GenesisCoreCommand(txConfig client.TxConfig, moduleBasics module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "genesis",
		Short:                      "Application's genesis-related subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	gentxModule := moduleBasics[genutiltypes.ModuleName].(genutil.AppModuleBasic)

	cmd.AddCommand(
		GenTxCmd(moduleBasics, txConfig,
			banktypes.GenesisBalancesIterator{}, defaultNodeHome),
		MigrateGenesisCmd(),
		CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, defaultNodeHome,
			gentxModule.GenTxValidator),
		ValidateGenesisCmd(moduleBasics),
		AddGenesisAccountCmd(defaultNodeHome),
	)

	return cmd
}
