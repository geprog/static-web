package commands

import "github.com/urfave/cli/v2"

var LogoutCommand = &cli.Command{
	Name:   "logout",
	Usage:  "logout from the server",
	Action: logout,
}

func logout(cCtx *cli.Context) error {
	return nil
}
