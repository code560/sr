package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/urfave/cli"
)

func makeFilter(script string) func(string) string {
	delim := string([]rune(script)[1:2])
	scripts := regexp.MustCompile(delim).Split(script, 4)
	if len(scripts) >= 3 && scripts[0] == "s" {
		r := regexp.MustCompile(scripts[1])
		return func(in string) string {
			return r.ReplaceAllString(in, scripts[2])
		}
	}
	return func(in string) string { return in }
}

func main() {
	app := &cli.App{
		Usage: "Stream regex",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return nil
			}
			fn := makeFilter(c.Args().Get(0))

			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				res := fn(scanner.Text())
				fmt.Fprintf(os.Stdout, "%s\r\n", res)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
