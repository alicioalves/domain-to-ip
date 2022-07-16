package app

import "github.com/urfave/cli"

// This function will return the terminal application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Usage = "Search IPs and Server names online"

	return app
}
