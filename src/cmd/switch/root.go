package switchcmd

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(switchCmd)
	switchCmd.AddCommand(switchBranchCmd)
	switchCmd.AddCommand(switchAccountCmd)
}

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch branches or accounts",
}