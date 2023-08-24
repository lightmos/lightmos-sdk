package cli

import (
	lightmos "github.com/lightmos/lightmos-sdk/types"
	"strconv"

	channelutils "github.com/cosmos/ibc-go/v7/modules/core/04-channel/client/utils"
	"github.com/lightmos/lightmos-sdk/client"
	"github.com/lightmos/lightmos-sdk/client/flags"
	"github.com/lightmos/lightmos-sdk/client/tx"
	"github.com/lightmos/lightmos-sdk/x/restaking/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendUndelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-undelegate [src-port] [src-channel] [amount]",
		Short: "Send a undelegate over IBC",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			validatorAddress := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argAmount, err := lightmos.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendUndelegate(validatorAddress, srcPort, srcChannel, timeoutTimestamp, &argAmount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
