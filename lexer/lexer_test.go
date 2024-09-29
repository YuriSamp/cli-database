package lexer

import "testing"

type LexerTestCase struct {
	input  string
	expect []string
}

func TestLexer(t *testing.T) {
	tests := []LexerTestCase{
		{
			input:  "abcd",
			expect: []string{"abcd"},
		},
		{
			input:  "a10",
			expect: []string{"a10"},
		},
		{
			input:  "\"uma string com espacos\"",
			expect: []string{"uma string com espacos"},
		},
		{
			input:  "101",
			expect: []string{"101"},
		},
		{
			input:  "TRUE",
			expect: []string{"TRUE"},
		},
		{
			input:  "\"teste\"",
			expect: []string{"teste"},
		},
		{
			input:  "set teste 1",
			expect: []string{"set", "teste", "1"},
		},
		{
			input:  "commit",
			expect: []string{"commit"},
		},
		{
			input:  "set \"uma string com espacos\" 1",
			expect: []string{"set", "uma string com espacos", "1"},
		},
	}

	for _, tc := range tests {
		result := New(tc.input)

		if len(result) != len(tc.expect) {
			t.Fatalf("expected len of result to be equal to len of expected, but len of result is %d and len of expected is %d", len(result), len(tc.expect))
		}

		for i := range tc.expect {
			if result[i] != tc.expect[i] {
				t.Fatalf("Error on lexer, expected %s to be %s", result[i], tc.expect[i])
			}
		}
	}
}
