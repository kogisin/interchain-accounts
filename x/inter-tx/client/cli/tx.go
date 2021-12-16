package cli

import (
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		getSubmitTxCmd(),
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

func getSubmitTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "submit-tx [interchain_account_address] [path/to/sdk_msg.json]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

			var txMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(args[1]), &txMsg); err != nil {

				// check for file path if JSON input is not provided
				contents, err := ioutil.ReadFile(args[1])
				if err != nil {
					return errors.Wrap(err, "neither JSON input nor path to .json file for sdk msg were provided")
				}

				if err := cdc.UnmarshalInterfaceJSON(contents, &txMsg); err != nil {
					return errors.Wrap(err, "error unmarshalling sdk msg file")
				}
			}

			cmd.Printf("sdk msg: %v\n\n", txMsg)

			protoMsg, ok := txMsg.(proto.Message)
			cmd.Printf("as proto msg: %v\n\n", protoMsg)
			cmd.Printf("as proto msg: %v\n\n", ok)

			delegateMsg, ok := txMsg.(*stakingtypes.MsgDelegate)
			cmd.Printf("as staking delegate msg: %v\n\n", delegateMsg)
			cmd.Printf("as staking delegate msg: %v\n\n", ok)

			cmd.Printf("sdk msg: %v\n\n", txMsg)

			msg, err := types.NewMsgSubmitTx(clientCtx.GetFromAddress(), txMsg, args[0], viper.GetString(FlagConnectionID), viper.GetString(FlagCounterpartyConnectionID))
			if err != nil {
				return err
			}

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
