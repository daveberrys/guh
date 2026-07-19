package push

import (
	"strings"

	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(pushCmd) }

var pushCmd = &cobra.Command{
	Use:   "push [remote]",
	Short: "Add current changes and push",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		branch, _ := utils.RunGit(true, "rev-parse", "--abbrev-ref", "HEAD")

		if len(args) == 0 {
			return utils.RunGitSequence(false, []string{"push", "-u", "origin", branch})
		}
		if args[0] == "all" {
			remotes, _ := utils.RunGit(true, "remote")
			for i, r := range strings.Fields(remotes) {
				if i == 0 {
					utils.RunGitSequence(false, []string{"push", "-u", r, branch})
				} else {
					utils.RunGitSequence(false, []string{"push", r})
				}
			}
			return nil
		}
		return utils.RunGitSequence(false, []string{"push", "-u", args[0], branch})
	},
}
