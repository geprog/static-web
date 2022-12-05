package main

import (
	"log"
	"os"

	"github.com/geprog/static-web/cli/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "static-web",
		Usage: "Simply deploy static pages using this cli.",
		Commands: []*cli.Command{
			commands.DeployCommand,
			commands.ListCommand,
			commands.TeardownCommand,
			commands.LoginCommand,
			commands.LogoutCommand,
			commands.WhoamiCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
