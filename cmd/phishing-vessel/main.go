package main

import (
	"os"

	"github.com/damiaoterto/phishing-vessel/internal/cmd"
	"github.com/damiaoterto/phishing-vessel/internal/logger"
)

func main() {
	// application entrypoint
	err := cmd.Execute()
	if err != nil {
		logger.Errorf("failed to execute command: %v", err)
		os.Exit(1)
	}
}
