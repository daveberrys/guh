// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package clone

import (
    "fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(cloneCmd) }

var cloneCmd = &cobra.Command{
	Use:   "clone [repository_url]",
	Short: "Clone a repository",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		utils.RunGitSequence(false, []string{"clone", args[0]})

		repoName := args[0]
		if idx := strings.LastIndex(repoName, "/"); idx != -1 {
			repoName = repoName[idx+1:]
		}
		repoName = strings.TrimSuffix(repoName, ".git")

		shell := os.Getenv("SHELL")
		if shell == "" { shell = "sh" }

		fmt.Println("")
		fmt.Printf("Spawning in a new shell: %s\n", shell)
		fmt.Println("This will create a nested shell.")

		sh := exec.Command(shell)
		sh.Dir = repoName
		sh.Stdin = os.Stdin
		sh.Stdout = os.Stdout
		sh.Stderr = os.Stderr
		sh.Run()

		return nil
	},
}