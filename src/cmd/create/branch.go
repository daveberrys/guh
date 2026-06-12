package createcmd

import (
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var createBranchCmd = &cobra.Command{
	Use:   "branch [branch]",
	Short: "Create and switch to a branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		return utils.RunGitSequence(false, []string{"switch", "-c", args[0]})
	},
}
