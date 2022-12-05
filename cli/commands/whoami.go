package commands

import "github.com/urfave/cli/v2"

var WhoamiCommand = &cli.Command{
	Name:   "whoami",
	Usage:  "show current user",
	Action: whoami,
}

func whoami(cCtx *cli.Context) error {
	return nil
}
