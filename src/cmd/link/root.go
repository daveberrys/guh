// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package link

import (
    "github.com/daveberrys/guh/src/utils"
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(linkCmd)
	linkCmd.AddCommand(addLinkRepoCmd)
	linkCmd.AddCommand(removeLinkRepoCmd)
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link repositories. Shows all linked repositories - No args given, lists all remotes.",
	RunE: func(c *cobra.Command, args []string) error {
		utils.RunGit(false, "remote", "-v")
		return nil
	},
}