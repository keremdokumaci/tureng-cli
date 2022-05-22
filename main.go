package main

import (
	"flag"

	"github.com/keremdokumaci/tureng-cli/app"
)

func main() {
	language := flag.String("language", "", "--language")
	maxNumberOfWord := flag.Int("count", 1, "--count")
	flag.Parse()

	cli := app.NewCli(app.Language(*language), *maxNumberOfWord)
	cli.Run()
}
