package clone

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/damiaoterto/phishing-vessel/internal/logger"
	"github.com/go-rod/rod"
	"github.com/urfave/cli/v2"
)

type ResourceType int

var (
	AppDirName   = ".phishing-vessel"
	DefaultIndex = "index.html"
)

func ClonePage(ctx *cli.Context) error {
	pageURL, err := url.Parse(ctx.String("url"))
	if err != nil {
		return fmt.Errorf("invalid url format: %w", err)
	}

	logger.Infof("Cloning page %s", pageURL.String())
	host := pageURL.Host

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user directory: %w", err)
	}

	pageDir := filepath.Join(homeDir, AppDirName, host)
	if err := prepareStorageDir(pageDir); err != nil {
		return fmt.Errorf("failed to prepare storage: %w", err)
	}

	browser := rod.New().
		MustConnect().
		MustIgnoreCertErrors(true).
		SlowMotion(time.Second * 2)
	defer browser.MustClose()

	html, err := getPageHTML(pageURL.String(), browser)
	if err != nil {
		return fmt.Errorf("failed to get page html: %w", err)
	}

	if err := createIndexFile(pageDir, html); err != nil {
		return fmt.Errorf("failed to create index file: %w", err)
	}

	return nil
}

func prepareStorageDir(pageDir string) error {
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

func getPageHTML(url string, browser *rod.Browser) (string, error) {
	page := browser.MustPage(url)
	page.MustWaitLoad()
	html, err := page.HTML()
	if err != nil {
		return "", fmt.Errorf("failed to get page html: %w", err)
	}
	return html, nil
}

func createIndexFile(path string, content string) error {
	file, err := os.Create(filepath.Join(path, DefaultIndex))
	if err != nil {
		return fmt.Errorf("failed to create index file: %w", err)
	}

	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write content to index file: %w", err)
	}

	return nil
}
