package cmd

import (
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Add current changes and commit",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.RunGitSequence([]string{"add", "."})
		utils.RunGitSequence([]string{"commit", "-m", args[0]})
		return nil
	},
}
