package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

const BackSpace = "\b"
const DelChar = "\x7F"
const CarriageReturn = "\r"
const KeyUp = "\033[A"
const KeyDown = "\033[B"
const ClearEntireLine = "\033[2K\r"
const ClearCursorToLineEnd = "\b\033[K"
const KillKey = "\x03"

type UserPrompt struct {
	inputHistory  []string
	tokens        []rune
	cursor        int
	historyCursor int
}

func (p *UserPrompt) getUserInput() ([]string, error) {
	termState, err := term.MakeRaw(0)
	if err != nil {
		os.Exit(1)
	}
	defer term.Restore(0, termState)
	p.tokens = p.tokens[:0]
	p.cursor = 0
	p.historyCursor = -1
	token := make([]byte, 4)
	fmt.Print("Pokedex > ")
	for {
		n, err := os.Stdin.Read(token)
		if err != nil {
			return []string{}, err
		}
		inputKey := string(token[:n])
		switch inputKey {
		case KillKey:
			term.Restore(0, termState)
			os.Exit(0)
		case KeyUp:
			p.setUserPrompt("up")
		case KeyDown:
			p.setUserPrompt("down")
		case BackSpace, DelChar:
			if p.cursor != 0 {
				p.cursor--
				os.Stdout.WriteString(ClearCursorToLineEnd)
			}
			p.tokens = p.tokens[0:p.cursor]
			continue
		case CarriageReturn:
			if len(p.tokens) == 0 {
				continue
			}
			os.Stdout.WriteString(CarriageReturn + "\n")
			p.inputHistory = append(p.inputHistory, string(p.tokens))
			p.historyCursor = -1
			return strings.Fields(string(p.tokens)), nil
		default:
			tokenRune, _ := utf8.DecodeRune(token)
			p.tokens = append(p.tokens, tokenRune)
			fmt.Printf("%v", string(p.tokens[p.cursor]))
			p.cursor++
		}
	}
}

func (p *UserPrompt) setUserPrompt(direction string) bool {
	switch direction {
	case "up":
		if p.historyCursor+1 < len(p.inputHistory) {
			p.historyCursor++
		}
	case "down":
		if p.historyCursor-1 >= 0 {
			p.historyCursor--
		}
	default:
		return false
	}
	os.Stdout.WriteString(ClearEntireLine)
	fmt.Print("Pokedex > ")
	os.Stdout.WriteString((p.inputHistory)[len(p.inputHistory)-1-p.historyCursor])
	p.tokens = []rune((p.inputHistory)[len(p.inputHistory)-1-p.historyCursor])
	p.cursor = len(p.tokens)
	return true
}
