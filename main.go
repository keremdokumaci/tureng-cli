package main

import "github.com/keremdokumaci/tureng-cli/app"

func main() {
	cli := app.NewCli("turkish-english")
	cli.Run()
}
