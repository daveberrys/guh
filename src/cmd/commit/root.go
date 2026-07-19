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
	Use:   "commit [files] [message] [description]",
	Short: "Add specified files and commit",
	Args:  cobra.RangeArgs(0, 5),
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
		afterPush, remoteName := false, "origin"

		switch len(args) {
		case 5:
			commitArgs = append(commitArgs, "-m", args[2])
			afterPush = true
			if args[4] != "" {
				remoteName = args[4]
			}
		case 4:
			if args[2] == "push" {
				afterPush = true
				if args[3] != "" {
					remoteName = args[3]
				}
			} else {
				commitArgs = append(commitArgs, "-m", args[2])
				afterPush = true
			}
		case 3:
			if args[2] == "push" {
				afterPush = true
			} else {
				commitArgs = append(commitArgs, "-m", args[2])
			}
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
