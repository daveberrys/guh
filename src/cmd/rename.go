package cmd

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.AddCommand(renameBranchCmd)
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename branches",
}

var renameBranchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "Rename a specific branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return renameBranch(args[0])
	},
}

func renameBranch(branchName string) error {
	utils.RunGitSequence(false, []string{"branch", "-m", branchName})
	fmt.Println("Successfully renamed branch to", branchName)
	return nil
}