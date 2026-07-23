// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package branch

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		utils.RunGitSequence(false, []string{"branch", "-D", args[0]})
		fmt.Printf("Deleted branch: %s\n", args[0])
		return nil
	},
}