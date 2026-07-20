package cli

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(cliCmd)
	cliCmd.AddCommand(updateCmd)
	cliCmd.AddCommand(versionCmd)
	cliCmd.AddCommand(installCmd)
}

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Manage the CLI tool",
}