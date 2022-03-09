package parser

import (
	"os"
	"strings"

	t "scythe.com/uni/src/tokens"
)

// Parse a file and return the Parse function with the red file as parameter
func ParseFile(file string) []t.Tokens {

	// File reading -> splitted file
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	t_file := string(f)
	t_file = strings.Replace(t_file, "\n", " ", -1)
	splitted := strings.Split(t_file, " ")

	// Array that'll hold all operators
	var arr []t.Tokens

	// Appending to "arr" all the tokens associated to words
	for _, i := range splitted {
		arr = append(arr, t.Tokenize(i))
	}
	return arr
}

// Parse a line (string) and return the Parse function with the line passed as parameter
func ParseLine(line string) []t.Tokens {
	trimmed := strings.Split(line, " ")
	var arr []t.Tokens

	for _, i := range trimmed {
		arr = append(arr, t.Tokenize(i))
	}
	return arr
}
