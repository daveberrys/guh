// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package account

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Opens up a code editor, and opens accounts.json file",
	Args:  cobra.ExactArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		path, _ := utils.AccountsFilePath()
		editors := []string{"nano", "code", "zed", "zeditor", "vi", "vim", "emacs", "notepad.exe"}
		// editors := []string{"this", "is", "a", "_test", "for", "fail"}
		for _, editor := range editors {
			if _, err := exec.Command(editor, path).Output(); err == nil {
			    fmt.Printf("Opening editor: %s\n", editor)
			    return nil
			}
		}

		fmt.Println("No supported editors found")
		fmt.Println("Here are the supported editors we have;")
		fmt.Println(strings.Join(editors, ", "))
		// for _, editor := range editors { fmt.Println("- " + editor) }
		return nil
	},
}