package cmd

import (
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(stashCmd) }

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Stash current changes",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.RunGitSequence(false, []string{"stash"})
	},
}
