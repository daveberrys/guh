// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install's Guh in your system",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		fmt.Println("Installing guh to your PATH...")

		src, _ := os.Executable()

		switch runtime.GOOS {
		case "linux", "darwin":
			destDir := filepath.Join(os.Getenv("HOME"), ".local", "bin")
			os.MkdirAll(destDir, 0755)

			dest := filepath.Join(destDir, "guh")
			data, _ := os.ReadFile(src)
			os.WriteFile(dest, data, 0755)

			fmt.Printf("Installed guh to %s\n", dest)
			fmt.Println("Make sure ~/.local/bin is in your PATH.")

		case "windows":
			destDir := filepath.Join(os.Getenv("APPDATA"), "dev.pages.codedave.guh")
			os.MkdirAll(destDir, 0755)

			dest := filepath.Join(destDir, "guh.exe")
			data, _ := os.ReadFile(src)
			os.WriteFile(dest, data, 0755)

			currentPath := os.Getenv("PATH")
			if !containsPath(currentPath, destDir) {
				newPath := destDir + ";" + currentPath
				exec.Command("setx", "PATH", newPath).Run()
			}

			fmt.Printf("Installed guh to %s\n", dest)
		}

		return nil
	},
}

func containsPath(path, dir string) bool {
	for _, p := range filepath.SplitList(path) {
		if p == dir {
			return true
		}
	}
	return false
}