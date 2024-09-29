package lexer

import (
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	insideQuotes bool
}

func New(input string) []string {
	l := Lexer{input: input, insideQuotes: false}

	word := []string{}
	result := []string{}

	for l.readChar() {

		if !l.hasNextChar() {
			if !l.currCharIs("\"") {
				word = append(word, string(l.ch))
			}

			result = append(result, strings.Join(word, ""))
			break
		}

		if l.currCharIs("\\") {
			l.readChar()
			word = append(word, string(l.ch))
			continue
		}

		if l.currCharIs("\"") {
			l.insideQuotes = !l.insideQuotes
			continue
		}

		word = append(word, string(l.ch))

		if l.currCharIs(" ") && !l.insideQuotes {
			result = append(result, strings.Join(word, ""))
			word = nil
		}
	}

	return result
}

func (l *Lexer) currCharIs(char string) bool {
	return char == string(l.ch)
}

func (l *Lexer) hasNextChar() bool {
	return l.readPosition < len(l.input)
}

func (l *Lexer) readChar() bool {
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition += 1
	return l.readPosition <= len(l.input)
}
