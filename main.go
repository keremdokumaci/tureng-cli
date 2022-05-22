package main

import (
	"flag"

	"github.com/keremdokumaci/tureng-cli/app"
)

func main() {
	language := flag.String("language", "", "language")
	cli := app.NewCli(app.Language(*language))
	cli.Run()
}
