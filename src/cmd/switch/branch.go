package switchcmd

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var switchBranchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "Switch to a specific branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
	    utils.RunGitequence(false, []string{"switch", args[0]})
	    fmt.Printf("Switched to branch: %s\n", args[0])
	    return nil
	},
}