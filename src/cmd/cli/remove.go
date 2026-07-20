package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Deletes the Guh binary",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		fmt.Println("Deleting Guh binary...")
		fmt.Println("Goodbye, world!")

		sysExec, _ := os.Executable()
		err := os.Remove(sysExec)
		if err != nil { return err }

		return nil
	},
}