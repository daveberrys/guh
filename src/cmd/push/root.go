// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package push

import (
	"fmt"
	"strings"

	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(pushCmd) }

var pushCmd = &cobra.Command{
	Use:   "push [force:opt] [remote:opt]",
	Short: "Add current changes and push",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		branch, _ := utils.RunGit(true, "rev-parse", "--abbrev-ref", "HEAD")

		force, remote := false, ""
		for _, arg := range args {
			if arg == "force" {
				force = true
			} else if remote == "" {
				remote = arg
			} else {
				return fmt.Errorf("unexpected argument %q", arg)
			}
		}

		if remote == "all" {
			remotes, _ := utils.RunGit(true, "remote")
			for i, r := range strings.Fields(remotes) {
				if i == 0 {
					utils.RunGitSequence(false, pushArgs(force, "-u", r, branch))
				} else {
					utils.RunGitSequence(false, pushArgs(force, r))
				}
			}
			return nil
		}
		if remote == "" {
			return utils.RunGitSequence(false, pushArgs(force, "-u", "origin", branch))
		}
		return utils.RunGitSequence(false, pushArgs(force, "-u", remote, branch))
	},
}

func pushArgs(force bool, tail ...string) []string {
	if force {
		return append([]string{"push", "--force"}, tail...)
	}
	return append([]string{"push"}, tail...)
}
