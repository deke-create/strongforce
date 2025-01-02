package types

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxCmd generates the entrypoint for the strongforce module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	strongforceTxCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "Strongforce transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	strongforceTxCmd.AddCommand(client.PostCommands(
		GetCmdExecuteAction(cdc),
	)...)
	return strongforceTxCmd
}

// GetCmdExecuteAction is the CLI command for sending a ExecuteAction transaction
func GetCmdExecuteAction(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "execute-action [name] [value]",
		Short: "execute an action",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.GetClientContextFromCmd(cmd).WithCodec(cdc)

			txBldr := tx.NewFactoryCLI(cliCtx, cmd.Flags())

			msg := NewMsgExecuteAction(cliCtx.GetFromAddress(), []byte(args[0]))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), []sdk.Msg{msg})
		},
	}
}
