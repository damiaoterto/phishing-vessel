package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:  "phishing-vessel",
		Usage: "A CLI phishing attack tool",
	}

	return app.Run(os.Args)
}
