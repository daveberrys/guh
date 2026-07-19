package branch

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [name]",
	Short: "Switch to a branch, or list branches with 'list'",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		if args[0] == "list" {
			utils.RunGit(false, "branch")
		} else {
			utils.RunGitSequence(false, []string{"switch", args[0]})
			fmt.Printf("Switched to branch: %s\n", args[0])
		}
		return nil
	},
}
