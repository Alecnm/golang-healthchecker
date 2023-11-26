package main

import (
	"github.com/urfave/cli/v2"
	"fmt"
	"log"
	"os"
)


func main(){
	app := &cli.App{
		Name: "Universal Healthchecker",
		Usage: "Healthcheck your services (Whethere they are up or not)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name of the service to be checked",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "Port number of the service to be checked",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == ""{
				port = "80"
			}
			status := healthCheck(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}