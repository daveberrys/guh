package initcmd

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(initCmd) }

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		return utils.RunGitSequence(false, []string{"init"})
	},
}
