package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/evan-forbes/stevia/x/stevia/types"
)

func CmdCreateSweetner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-sweetner [calories] [organic] [source]",
		Short: "Creates a new sweetner",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCalories, _ := strconv.ParseInt(args[0], 10, 64)
			argsOrganic, _ := strconv.ParseBool(args[1])
			argsSource := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSweetner(clientCtx.GetFromAddress().String(), uint64(argsCalories), bool(argsOrganic), string(argsSource))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateSweetner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-sweetner [id] [calories] [organic] [source]",
		Short: "Update a sweetner",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsCalories, _ := strconv.ParseInt(args[1], 10, 64)
			argsOrganic, _ := strconv.ParseBool(args[2])
			argsSource := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSweetner(clientCtx.GetFromAddress().String(), id, uint64(argsCalories), bool(argsOrganic), string(argsSource))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteSweetner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-sweetner [id] [calories] [organic] [source]",
		Short: "Delete a sweetner by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSweetner(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
