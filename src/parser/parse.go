package parser

import (
	"os"
	"strings"

	t "scythe.com/uni/tokens"
)

// Parse a file and return the Parse function with the red file as parameter
func ParseFile(file string) t.Tokens {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	for _, i := range trimmed {
		return t.TokenizeWord(i)
	}
}

// Parse a line (string) and return the Parse function with the line passed as parameter
func ParseLine(line string) []t.Tokens {
	trimmed := strings.Split(line, " ")
	return t.Tokenize(trimmed)
}
