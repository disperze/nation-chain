package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/disperze/nation-chain/x/nation/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	nationTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nationTxCmd.AddCommand(flags.PostCommands(
		GetCmdRegisterDni(cdc),
	)...)

	return nationTxCmd
}

// GetCmdRegisterDni is the CLI command for doing RegisterDni
func GetCmdRegisterDni(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "register-dni [dni] [names] [first surname] [second surname]",
		Short: "Register new DNI",
		Args:  cobra.ExactArgs(4), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			validator, err := sdk.ValAddressFromBech32(cliCtx.GetFromAddress().String())
			if err != nil {
				return err
			}

			names := strings.Split(args[1], " ")
			msg := types.NewMsgRegisterDni(validator)
			msg.Dni = args[0]
			msg.Name = names[0]
			msg.MiddleName = strings.Join(names[1:], " ")
			msg.Surname1 = args[2]
			msg.Surname2 = args[3]
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
