package createcmd

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createBranchCmd)
	createCmd.AddCommand(createAccountCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create branches or accounts",
}
