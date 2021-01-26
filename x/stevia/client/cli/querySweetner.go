package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evan-forbes/stevia/x/stevia/types"
	"github.com/spf13/cobra"
)

func CmdListSweetner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-sweetner",
		Short: "list all sweetner",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSweetnerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SweetnerAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSweetner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-sweetner [id]",
		Short: "shows a sweetner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetSweetnerRequest{
				Id: args[0],
			}

			res, err := queryClient.Sweetner(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
