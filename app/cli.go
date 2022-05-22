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

var languageMap map[Language]LanguageSet

type Command string

const (
	UPDATE_LANGUAGE Command = "update-language"
)

type TurengCli struct {
	proxy     *TurengProxy
	language  Language `validate:"required"`
	wordCount int
}

func NewCli(language Language, count int) *TurengCli {
	cli := &TurengCli{}
	if !isSupportedLanguage(string(language)) {
		fmt.Println("unsupported language")
		os.Exit(1)
	}

	languageMap = make(map[Language]LanguageSet)
	languageMap[TREN] = LanguageSet{
		SourceLanguageShortForm: "tr",
		SourceLanguage:          "Turkish",
		DestLanguageShortForm:   "en",
		DestLanguage:            "English",
	}
	languageMap[ENDE] = LanguageSet{
		SourceLanguageShortForm: "en",
		SourceLanguage:          "English",
		DestLanguageShortForm:   "de",
		DestLanguage:            "German",
	}
	languageMap[ENES] = LanguageSet{
		SourceLanguageShortForm: "tr",
		SourceLanguage:          "Turkish",
		DestLanguageShortForm:   "es",
		DestLanguage:            "Spanish",
	}
	languageMap[ENFR] = LanguageSet{
		SourceLanguageShortForm: "tr",
		SourceLanguage:          "Turkish",
		DestLanguageShortForm:   "fr",
		DestLanguage:            "French",
	}

	cli.language = language
	cli.wordCount = count
	cli.proxy = NewTurengProxy(string(language))
	return cli
}

func (c *TurengCli) Run() {
	for {
		c.showPrompt()
		input := c.getUserInput()
		if input == "" {
			continue
		}

		if c.isCommandText(input) {
			command := strings.ReplaceAll(input, "-c", "")
			err := c.runCommand(strings.TrimLeft(command, " "))
			if err != nil && err.Error() != fmt.Sprintf(`exec: "%s": executable file not found in $PATH`, input) {
				fmt.Println(err.Error())
				continue
			}
		} else {
			result, err := c.proxy.Query(input, languageMap[c.language])
			if err != nil {
				fmt.Println(err.Error())
			}

			for i, t := range result {
				if i == c.wordCount {
					break
				}

				fmt.Println(t)
			}
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
	return len(strings.Split(text, "-c")) > 1
}

func (c *TurengCli) runCommand(text string) error {
	text = strings.ReplaceAll(text, "-c", "")
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
	if !isSupportedLanguage(language) {
		return errors.New("unsupported language !")
	}

	c.language = Language(language)
	return nil
}

func isSupportedLanguage(language string) bool {
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
		return false
	}

	return true
}
