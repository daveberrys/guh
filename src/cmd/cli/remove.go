package cli

import (
	"fmt"
	"os"
	"runtime"

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

		if runtime.GOOS == "windows" {
			fmt.Println("\nThe binary may not be deleted due to Windows' permission issue.")
			fmt.Println("If you wish to delete the binary, please run the following command in PowerShell:")
			fmt.Printf("  del '%s'\n", sysExec)
		}

		return nil
	},
}