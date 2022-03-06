package parser

import (
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	c "github.com/fatih/color"
	t "scythe.com/uni/tokens"
)

// Parse a file and return the Parse function with the red file as parameter
func ParseFile(file string) []t.Tokens {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	return Parse(trimmed)
}

// Parse a string array and append for each chars an operator to an array. Then, return the array
func Parse(data []string) []t.Tokens {
	var stack []t.Tokens

	for _, i := range data {
		switch i {
		case "+":
			c.Red("Plus")
			stack = append(stack, t.Plus())
		case "-":
			c.Red("Min")
			stack = append(stack, t.Min())
		case "dmp":
			c.Red("Dumped")
			stack = append(stack, t.Dump())
		default:
			c.Red("Num called")
			if num, err := strconv.Atoi(i); err == nil {
				c.Red("Num check passed")
				cyan := c.New(c.FgCyan).Add(color.Underline)
				cyan.Println(num)
				stack = append(stack, t.Push(num))
			}
		}
	}
	return stack
}

// Parse a line (string) and return the Parse function with the line passed as parameter
func ParseLine(line string) []t.Tokens {
	trimmed := strings.Split(line, " ")
	return Parse(trimmed)
}
