package lexer

import (
	"fmt"
	"strings"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) {
	l := Lexer{input: input}

	result := []string{}

	for l.readChar() {

		if l.currCharIs("\\") {
			result = append(result, string(input[l.readPosition]))
			l.readPosition += 1
			continue
		}

		if l.currCharIs("\"") {
			continue
		}

		result = append(result, string(l.ch))
	}

	fmt.Print(strings.Join(result, ""))
	fmt.Print("\n")
}

func (l *Lexer) currCharIs(char string) bool {
	return char == string(l.input[l.position])
}

func (l *Lexer) peekCharIs(char string) bool {
	if l.readPosition >= len(l.input) {
		return false
	}

	return char == string(l.input[l.readPosition])
}

func (l *Lexer) skipNextToken() {
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) readChar() bool {
	if l.readPosition >= len(l.input) {
		return false
	} else {
		l.ch = l.input[l.readPosition]
		l.position = l.readPosition
		l.readPosition += 1
		return true
	}
}
