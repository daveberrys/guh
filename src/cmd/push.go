package cmd

import (
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Add current changes and push",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.RunGitSequence(false, []string{"push"})
	},
}
