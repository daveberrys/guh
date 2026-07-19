package branch

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(branchCmd)
	branchCmd.AddCommand(createCmd)
	branchCmd.AddCommand(switchCmd)
	branchCmd.AddCommand(renameCmd)
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage branches",
}
