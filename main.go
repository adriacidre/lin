package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var err error

func main() {
	app := cli.NewApp()
	app.Name = "lin"
	app.Version = "0.1"
	app.Usage = ""
	app.Commands = []cli.Command{
		yamlLin,
	}
	if err := app.Run(os.Args); err != nil {
		log.Println("Oops, something is broken")
	}
}
