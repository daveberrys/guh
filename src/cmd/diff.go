package cmd

import (
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(diffCmd) }

var diffCmd = &cobra.Command{
	Use:   "diff [file:optional]",
	Short: "Show working tree changes",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 { return utils.RunGitSequence(false, []string{"diff", args[0]}) }
		return utils.RunGitSequence(false, []string{"diff"})
	},
}
