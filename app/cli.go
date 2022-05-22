package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	UPDATE_LANGUAGE Command = "update-language"
	CLEAR           Command = "clear"
)

type TurengCli struct {
	proxy    *TurengProxy
	language Language `validate:"required"`
}

func NewCli(language Language) *TurengCli {
	cli := &TurengCli{}
	cli.language = language
	cli.proxy = NewTurengProxy(string(language))
	return cli
}

func (c *TurengCli) Run() {
	for {
		c.showPrompt()
		input := c.getUserInput()

		if c.isCommandText(input) {
			err := c.runCommand(input)
			if err != nil && err.Error() != fmt.Sprintf(`exec: "%s": executable file not found in $PATH`, input) {
				fmt.Println(err.Error())
				continue
			}
		} else {
			result, err := c.proxy.Query(input)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(result)
		}
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

func (c *TurengCli) isCommandText(text string) bool {
	return strings.Split(text, "")[0] == "["
}

func (c *TurengCli) runCommand(text string) error {
	text = text[1:]
	text = text[:len(text)-1]

	commands := strings.Split(text, " ")
	cmd := commands[0]

	switch cmd {
	case string(UPDATE_LANGUAGE):
		return c.updateLanguage(commands[1])
	default:
		c := exec.Command(text)
		c.Stdout = os.Stdout
		err := c.Run()
		return err
	}
}

func (c *TurengCli) updateLanguage(language string) error {
	switch language {
	case string(TREN):
		break
	case string(ENDE):
		break
	case string(ENES):
		break
	case string(ENFR):
		break
	default:
		return errors.New("unsupported language")
	}

	c.language = Language(language)
	return nil
}
