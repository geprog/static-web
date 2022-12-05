package commands

import "github.com/urfave/cli/v2"

var LoginCommand = &cli.Command{
	Name:   "login",
	Usage:  "login to the server",
	Action: login,
}

func login(cCtx *cli.Context) error {
	return nil
}
