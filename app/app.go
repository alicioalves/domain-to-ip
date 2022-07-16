package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// This function will return the terminal application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Domain to IP"
	app.Usage = "Search IPs and Server names online"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP addresses online",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
