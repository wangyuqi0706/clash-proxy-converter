package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/wangyuqi0706/clash-proxy-converter/converter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var inputFile string
	var outPath string
	outPath, _ = os.Getwd()
	app := &cli.App{
		Name:      "converter",
		Usage:     "convert vmess or ss link to clash config",
		ArgsUsage: "your_subscribe_links",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			var uris []string
			if c.IsSet("input") {
				bytes, err := ioutil.ReadFile(inputFile)
				fileText := string(bytes)
				if err != nil {
					return err
				}
				uris = strings.Fields(fileText)
			}
			m := converter.Convert(uris)
			bytes, err := yaml.Marshal(m)
			if err != nil {
				return err
			}
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			if !c.IsSet("output") {
				outPath = "config.yaml"
			}

			err = ioutil.WriteFile(filepath.Join(wd, outPath), bytes, 0644)
			if err != nil {
				return err
			}
			return nil
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "Load links from `FILE`",
				Destination: &inputFile,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "",
				EnvVars:     nil,
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				TakesFile:   false,
				Value:       "",
				DefaultText: "",
				Destination: &outPath,
				HasBeenSet:  false,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
