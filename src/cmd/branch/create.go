package branch

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create and switch to a new branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		utils.RunGitSequence(false, []string{"switch", "-c", args[0]})
		fmt.Printf("Created branch to: %s\n", args[0])
		return nil
	},
}
