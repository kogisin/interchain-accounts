package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

// GetTxCmd creates and returns the intertx tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		getRegisterAccountCmd(),
		getSendTxCmd(),
		getDelegateTxCmd(),
	)

	return cmd
}

func getRegisterAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "register",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterAccount(
				clientCtx.GetFromAddress().String(),
				viper.GetString(FlagConnectionID),
				viper.GetString(FlagCounterpartyConnectionID),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(fsConnectionPair)
	_ = cmd.MarkFlagRequired(FlagConnectionID)
	_ = cmd.MarkFlagRequired(FlagCounterpartyConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getSendTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "send [interchain_account_address] [to_address] [amount]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgSend(
				clientCtx.GetFromAddress(),
				amount,
				args[0],
				args[1],
				viper.GetString(FlagConnectionID),
				viper.GetString(FlagCounterpartyConnectionID),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(fsConnectionPair)

	_ = cmd.MarkFlagRequired(FlagConnectionID)
	_ = cmd.MarkFlagRequired(FlagCounterpartyConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getDelegateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delegate [interchain_account_address] [val_address] [amount]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgDelegate(
				clientCtx.GetFromAddress(),
				amount,
				args[0],
				args[1],
				viper.GetString(FlagConnectionID),
				viper.GetString(FlagCounterpartyConnectionID),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(fsConnectionPair)

	_ = cmd.MarkFlagRequired(FlagConnectionID)
	_ = cmd.MarkFlagRequired(FlagCounterpartyConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetCmdSubmitRegisterInterchainAccountProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-ica",
		Args:  cobra.ExactArgs(0),
		Short: "Submit a public farming plan",
		Long:  strings.TrimSpace("Submit a governance proposal to register interchain account"),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewRegisterInterchainAccountProposal(
				"Register Interchain Account",
				"Demo",
				"",
				"",
			)

			from := clientCtx.GetFromAddress()

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")

	return cmd
}
