package clone

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var AppDirName string = ".phishing-vessel"

func ClonePage(ctx *cli.Context) error {
	pageURL, err := url.Parse(ctx.String("url"))
	if err != nil {
		return fmt.Errorf("invalid url format: %w", err)
	}

	log.Printf("Cloning page '%s'", pageURL)
	host := pageURL.Host

	if err := prepareStorageDir(host); err != nil {
		return fmt.Errorf("failed to prepare storage: %w", err)
	}

	return nil
}

func prepareStorageDir(host string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user directory: %w", err)
	}

	pageDir := filepath.Join(homeDir, AppDirName, host)
	exists, err := directoryExists(pageDir)
	if err != nil {
		return fmt.Errorf("failed to check directory: %w", err)
	}

	if !exists {
		createAppDir(pageDir)
	}

	return nil
}

func directoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check directory status")
	}

	return info.IsDir(), nil
}

func createAppDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("error on create application directory %s: %w", path, err)
	}
	return nil
}
