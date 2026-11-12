// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package commit

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(commitCmd) }

var commitCmd = &cobra.Command{
	Use:   "commit [files] [message] [description:opt]/push:opt/signoff:opt [push:opt]/signoff:opt [remote_name:opt] [signoff:opt]",
	Short: "Add specified files and commit",
	Args:  cobra.RangeArgs(0, 6),
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return utils.RunGitSequence(false, []string{"status", "--short"})
		}

		if len(args) < 2 {
			return fmt.Errorf("requires at least 2 arguments: files (JSON array or \".\") and commit message")
		}

		var addArgs []string
		if args[0] == "." {
			addArgs = []string{"add", "."}
		} else {
			var files []string
			if err := json.Unmarshal([]byte(args[0]), &files); err != nil {
				return fmt.Errorf("first argument must be \".\" or a JSON array of file paths: %w", err)
			}
			addArgs = append([]string{"add"}, files...)
		}

		fmt.Println("Edited files:")
		if err := utils.RunGitSequence(false, []string{"status", "--short"}); err != nil {
			return err
		}
		fmt.Println()

		utils.RunGitSequence(false, addArgs)

		commitArgs := []string{"commit", "-m", args[1]}
		afterPush, signoff, remoteName := false, false, "origin"
		remoteSet := false

		// The third argument is either a description or an action keyword.
		rest := args[2:]
		if len(rest) > 0 && rest[0] != "push" && rest[0] != "signoff" {
			commitArgs = append(commitArgs, "-m", rest[0])
			rest = rest[1:]
		}

		// case 1 -> guh commit [files] [message] [desc] push[:opt] [remote_name:opt(origin)] [signoff:opt]
		// case 2 -> guh commit [files] [message] [desc] push[:opt] [signoff:opt]
		// case 3 -> guh commit [files] [message]        push[:opt] [remote_name:opt(origin)] [signoff:opt]
		// case 4 -> guh commit [files] [message]        push[:opt] [signoff:opt]
		// case 5 -> guh commit [files] [message] [desc] [signoff:opt]
		// case 6 -> guh commit [files] [message] [signoff:opt]
		for _, arg := range rest {
			switch {
			case arg == "push":
				afterPush = true
			case arg == "signoff":
				signoff = true
			case arg == "":
				// ignore empty optional arguments
			case !remoteSet:
				remoteName, remoteSet = arg, true
			default:
				return fmt.Errorf("unexpected argument %q", arg)
			}
		}

		if signoff {
			commitArgs = append(commitArgs, "--signoff")
		}

		utils.RunGitSequence(false, commitArgs)
		if afterPush {
			branch, _ := utils.RunGit(true, "rev-parse", "--abbrev-ref", "HEAD")
			if remoteName == "all" {
				remotes, _ := utils.RunGit(true, "remote")
				for i, r := range strings.Fields(remotes) {
					if i == 0 {
						utils.RunGitSequence(false, []string{"push", "-u", r, branch})
					} else {
						utils.RunGitSequence(false, []string{"push", r})
					}
				}
			} else {
				utils.RunGitSequence(false, []string{"push", "-u", remoteName, branch})
			}
		}
		return nil
	},
}
