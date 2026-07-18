package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(commitCmd) }

var commitCmd = &cobra.Command{
	Use:   "commit [files] [message] [description]",
	Short: "Add specified files and commit",
	Args: cobra.RangeArgs(0, 5),
	RunE: func(cmd *cobra.Command, args []string) error {
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

		// case 1 -> guh commit [files] [message] [desc] push[:optional] [remote_name:optional(origin)]
		// case 2 -> guh commit [files] [message]        push[:optional] [remote_name:optional(origin)]
		// case 3 -> guh commit [files] [message] [desc]
		// case 4 -> guh commit [files] [message]
		switch len(args) {
		case 5:
			// [files] [message] [desc] push [remote]   (case 1)
			commitArgs = append(commitArgs, "-m", args[2])
			afterPush = true
			if args[4] != "" {
				remoteName = args[4]
			}
		case 4:
			if args[2] == "push" {
				// [files] [message] push [remote]      (case 2)
				afterPush = true
				if args[3] != "" {
					remoteName = args[3]
				}
			} else {
				// [files] [message] [desc] push        (case 1)
				commitArgs = append(commitArgs, "-m", args[2])
				afterPush = true
			}
		case 3:
			if args[2] == "push" {
				// [files] [message] push               (case 2)
				afterPush = true
			} else {
				// [files] [message] [desc]             (case 3)
				commitArgs = append(commitArgs, "-m", args[2])
			}
		}

		utils.RunGitSequence(false, commitArgs)
		if afterPush {
			if remoteName == "all" {
				remotes, err := utils.RunGit(true, "remote")
				if err != nil {
					return err
				}
				for _, r := range strings.Fields(remotes) {
					if err := utils.RunGitSequence(false, []string{"push", r}); err != nil {
						return err
					}
				}
			} else {
				utils.RunGitSequence(false, []string{"push", remoteName})
			}
		}
		return nil
	},
}
