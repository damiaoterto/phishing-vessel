package cmd

import (
	"os"

	"github.com/damiaoterto/phishing-vessel/internal/cmd/clone"
	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:  "phishing-vessel",
		Usage: "A CLI phishing attack tool",
		Commands: []*cli.Command{
			clone.Command(),
		},
	}

	return app.Run(os.Args)
}
