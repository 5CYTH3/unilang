package lexer

import (
	"os"
	"strings"

	t "scythe.com/uni/src/tokens"
)

// Split and trim the given string. Return the splitted string
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
	return arr
}

// Parse a file and return the an array of tokens from a splitted string
func LexFile(file string) []t.Tokens {
	var arr []t.Tokens

	// File parsing -> array of Tokens
	f, err := os.ReadFile(file)
	if err == nil {
		t_file := string(f)
		arr = LexString(CleanString(t_file))
	} else {
		panic(err)
	}
	return arr
}
