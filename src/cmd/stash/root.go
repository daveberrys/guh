// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package stash

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(stashCmd) }

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Stash current changes",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		return utils.RunGitSequence(false, []string{"stash"})
	},
}
