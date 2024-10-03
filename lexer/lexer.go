package lexer

import (
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
	insideQuotes bool
}

func Tokenize(input string) []string {
	l := Lexer{input: input}

	var word strings.Builder
	var result []string

	for l.readChar() {
		switch {
		case l.currCharIs("\\"):
			l.readChar()
			word.WriteRune(l.ch)

		case l.currCharIs("\""):
			l.insideQuotes = !l.insideQuotes

		case l.currCharIs(" ") && !l.insideQuotes:
			if word.Len() > 0 {
				result = append(result, word.String())
				word.Reset()
			}

		default:
			word.WriteRune(l.ch)
		}
	}

	if word.Len() > 0 {
		result = append(result, word.String())
	}

	return result
}

func (l *Lexer) currCharIs(char string) bool {
	return char == string(l.ch)
}

func (l *Lexer) readChar() bool {
	if l.readPosition >= len(l.input) {
		return false
	}

	r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
	l.ch = r
	l.position = l.readPosition
	l.readPosition += size

	return r != utf8.RuneError
}
