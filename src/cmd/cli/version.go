package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of Guh",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		fmt.Printf("Current `guh` hash version; %s\n", Version)
		return nil
	},
}