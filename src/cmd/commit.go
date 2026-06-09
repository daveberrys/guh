package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use:   "commit [files] [message] [description]",
	Short: "Add specified files and commit",
	Args: cobra.RangeArgs(0, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return utils.RunGit("status", "--short")
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
		if err := utils.RunGit("status", "--short"); err != nil {
			return err
		}
		fmt.Println()

		utils.RunGitSequence(addArgs)

		commitArgs := []string{"commit", "-m", args[1]}
		if len(args) == 3 {
			commitArgs = append(commitArgs, "-m", args[2])
		}
		utils.RunGitSequence(commitArgs)
		return nil
	},
}
