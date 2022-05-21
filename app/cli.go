package app

import (
	"bufio"
	"fmt"
	"os"
)

type Language string

const (
	TREN Language = "turkish-english"
	ENDE Language = "german-english"
	ENES Language = "spanish-english"
	ENFR Language = "french-english"
)

type Command string

const (
	DEFAULT         Command = "search"
	UPDATE_LANGUAGE Command = "update-language"
)

type TurengCli struct {
	language Language `validate:"required"`
}

func NewCli(language Language) *TurengCli {
	cli := &TurengCli{}
	cli.language = language
	return cli
}

func (c *TurengCli) Run() {
	for {
		c.showPrompt()
		input := c.getUserInput()
		fmt.Println(input)
	}
}

func (c *TurengCli) showPrompt() {
	fmt.Printf("tureng (%v) >> ", c.language)
}

func (c *TurengCli) getUserInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	inputText := input.Text()
	return inputText
}
