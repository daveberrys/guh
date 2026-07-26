// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package branch

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(branchCmd)
	branchCmd.AddCommand(createCmd)
	branchCmd.AddCommand(switchCmd)
	branchCmd.AddCommand(renameCmd)
	branchCmd.AddCommand(deleteCmd)
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage branches - No args given, lists all branches",
	Run: func(c *cobra.Command, args []string) {
		utils.RunGit(false, "branch")
	},
}
