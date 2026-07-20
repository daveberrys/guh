package cli

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates Guh to the latest version",
	Args:  cobra.NoArgs,
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		resp, err := http.Get("https://api.github.com/repos/daveberrys/guh/commits")
		if err != nil {
			return err
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var commits []struct {
			Sha string `json:"sha"`
		}
		json.Unmarshal(body, &commits)
		guhGithubHash := commits[0].Sha[:7]

		if guhGithubHash == Version {
			return nil
		}

		fmt.Printf("Updating from %s -> %s\n", Version, guhGithubHash)
		fmt.Println("This may take a long time depending on your internet...")
		return doUpdateSequence()
	},
}

func doUpdateSequence() error {
	var url string
	switch runtime.GOOS {
    	case "windows":
    		url = "https://nightly.link/daveberrys/guh/workflows/compile.yaml/main/GUH-Windows.exe.zip"
    	case "linux":
    		url = "https://nightly.link/daveberrys/guh/workflows/compile.yaml/main/GUH-Linux.zip"
    	case "darwin":
    		url = "https://nightly.link/daveberrys/guh/workflows/compile.yaml/main/GUH-macOS.zip"
    	default:
    		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	resp, err := http.Get(url)
	if err != nil { return err }
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	zipReader, _ := zip.NewReader(bytes.NewReader(body), int64(len(body)))

	var binName string
	switch runtime.GOOS {
		case "windows":
			binName = "GUH-Windows.exe"
		case "linux":
			binName = "GUH-Linux"
		case "darwin":
			binName = "GUH-macOS"
		default:
			return fmt.Errorf("A fatal error occurred! We don't know what happened...")
	}

	var binData []byte
	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}
		if strings.HasSuffix(f.Name, binName) || strings.HasSuffix(f.Name, "/"+binName) {
			rc, _ := f.Open()
			binData, _ = io.ReadAll(rc)
			rc.Close()
			break
		}
	}

	if binData == nil {
		return fmt.Errorf("could not find %s in archive", binName)
	}

	exePath, _ := os.Executable()
	exePath, _ = filepath.EvalSymlinks(exePath)

	tmpPath := exePath + ".tmp"
	os.WriteFile(tmpPath, binData, 0755)
	os.Rename(tmpPath, exePath)

	fmt.Println("Update successful!")
	return nil
}
