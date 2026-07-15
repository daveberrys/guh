package cmd

import (
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(initCmd) }

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.RunGitSequence(false, []string{"init"})
	},
}
