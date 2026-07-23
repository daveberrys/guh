// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package pull

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(pullCmd) }

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch and pull updates",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		return utils.RunGitSequence(false,
			[]string{"fetch"},
			[]string{"pull"},
		)
	},
}
