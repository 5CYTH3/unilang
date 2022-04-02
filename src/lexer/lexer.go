package lexer

import (
	"os"
	"strings"

	t "scythe.com/uni/src/tokens"
)

// Return a splitted and trimmed given string
func CleanString(str string) []string {
	str = strings.Replace(str, "\n", " ", -1)
	strArr := strings.Split(str, " ")
	return strArr
}

// Takes an array of string as parameter and return an array of tokens.
func LexString(input []string) []t.Tokens {
	arr := make([]t.Tokens, 0)
	for _, i := range input { // For each items of the array, append the Token associated to the current item to an array
		arr = append(arr, t.ParseTokenAsOperator(i))
	}
	return t.InfixToPostfix(arr)
}

// Lex all the tokens of a file
func LexFile(file string) []t.Tokens {
	var arr []t.Tokens

	f, err := os.ReadFile(file) // File reading (return a []byte)

	t_file := string(f)                  // Convert type []byte of the readed file to string
	arr = LexString(CleanString(t_file)) // Lexing of the string from readed file, trimmed and splitted

	if err != nil {
		panic(err)
	}

	return arr
}
