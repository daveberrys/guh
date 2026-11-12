// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package signoff

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

const sshKeyURL = "https://github.com/settings/ssh/new"

func init() { cmd.RootCmd.AddCommand(signoffCmd) }

var signoffCmd = &cobra.Command{
	Use:   "signoff",
	Short: "Configure SSH signing for Git",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		email, err := utils.RunGit(true, "config", "user.email")
		if err != nil || email == "" {
			return fmt.Errorf("git user.email is not set; run `git config --global user.email <email>` first")
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		sshDir := filepath.Join(homeDir, ".ssh")
		privateKeyPath := filepath.Join(sshDir, "id_ed25519")
		publicKeyPath := filepath.Join(sshDir, "id_ed25519.pub")

		contents, err := os.ReadFile(publicKeyPath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		keyMatchesEmail := err == nil && keyHasEmail(string(contents), email)
		if keyMatchesEmail {
			fmt.Printf("Using existing SSH signing key: %s\n", publicKeyPath)
		} else {
			if err == nil {
				fmt.Printf("SSH key already exists at %s, but its comment does not match %s.\n", publicKeyPath, email)
				fmt.Println("ssh-keygen will ask before overwriting it.")
				fmt.Println()
			}

			if err := os.MkdirAll(sshDir, 0700); err != nil {
				return err
			}

			sshKeyGen := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", privateKeyPath)
			sshKeyGen.Stdin = os.Stdin
			sshKeyGen.Stdout = os.Stdout
			sshKeyGen.Stderr = os.Stderr
			if err := sshKeyGen.Run(); err != nil {
				return fmt.Errorf("ssh-keygen failed: %w", err)
			}
		}

		if err := utils.RunGitSequence(false,
			[]string{"config", "--global", "gpg.format", "ssh"},
			[]string{"config", "--global", "user.signingkey", filepath.ToSlash(publicKeyPath)},
			[]string{"config", "--global", "commit.gpgsign", "true"},
		); err != nil {
			return err
		}

		contents, err = os.ReadFile(publicKeyPath)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Printf("Copy this key below, and add it in GitHub. Yes, this is fine to show.\n  %s\n", strings.TrimSpace(string(contents)))
		fmt.Println(`Make sure to add it as a "Signing Key", not an "Authentication Key".`)
		fmt.Println()
		fmt.Printf("Opening %s\n", sshKeyURL)

		return browser.OpenURL(sshKeyURL)
	},
}

func keyHasEmail(contents, email string) bool {
	fields := strings.Fields(contents)
	if len(fields) < 3 {
		return false
	}

	comment := strings.Join(fields[2:], " ")
	return comment == email || strings.Contains(comment, "<"+email+">")
}
