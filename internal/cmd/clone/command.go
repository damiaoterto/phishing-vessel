package clone

import "github.com/urfave/cli/v2"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "clone",
		Usage: "Init page clone",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				Usage:   "URL to clone",
			},
		},
		Action: ClonePage,
	}
}
