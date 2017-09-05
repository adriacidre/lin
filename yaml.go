package main

import (
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var yamlLin = cli.Command{
	Name:        "yaml",
	Usage:       "$ lin yaml /my/file.yml",
	ArgsUsage:   "Only argument is the file you want to lin",
	Description: ``,
	Action: func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("You must provide the file to be linted")
			return nil
		}
		data, err := ioutil.ReadFile(c.Args()[0])
		if err != nil {
			fmt.Println("error: ", err)
			return cli.NewExitError(string(err.Error()), 1)
		}
		m := make(map[interface{}]interface{})

		err = yaml.Unmarshal([]byte(data), &m)
		if err != nil {
			fmt.Println("error: ", err)
			return nil
		}

		return nil
	},
}
