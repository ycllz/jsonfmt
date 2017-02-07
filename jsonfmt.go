package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "jsonfmt"
	app.Usage = "Beautify json string data."
	app.UsageText = "jsonfmt [global options] [arg]"
	app.Version = Version("0.1.0")
	app.Authors = []cli.Author{
		cli.Author{Name: "voidnt",
			Email: "voidint@126.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "input source",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output destination",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var writer io.Writer
		var jsonData []byte
		var err error

		input := ctx.String("input")
		if input != "" {
			if jsonData, err = ioutil.ReadFile(input); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
		} else if ctx.NArg() > 0 {
			jsonData = ([]byte)(ctx.Args().First())
		} else {
			if jsonData, err = ioutil.ReadAll(os.Stdin); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
		}

		var buf bytes.Buffer
		if err = json.Indent(&buf, jsonData, "", "    "); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		output := ctx.String("output")
		if len(output) == 0 {
			writer = os.Stdout
		} else {
			var f *os.File
			if f, err = os.Create(output); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
			defer f.Close()
			writer = f
		}

		if _, err = buf.WriteTo(writer); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
