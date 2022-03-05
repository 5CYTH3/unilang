package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	t "scythe.com/uni/tokens"
)

func ParseFile(file string) []t.Token {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	return Parse(trimmed)
}

func Parse(data []string) []t.Token {
	var stack []t.Token

	for _, i := range data {
		switch i {
		case "+":
			fmt.Println("Plus")
			stack = append(stack, t.Plus())
		case "-":
			fmt.Println("Min")
			stack = append(stack, t.Min())
		case "dmp":
			fmt.Println("Dumped")
			stack = append(stack, t.Dump())
		default:
			fmt.Println("Num called")
			if num, err := strconv.Atoi(i); err == nil {
				fmt.Println("Num check passed")
				fmt.Println(num)
				stack = append(stack, t.Push(num))
			}
		}
	}
	return stack
}

func ParseLine(line string) []t.Token {
	trimmed := strings.Split(line, " ")
	return Parse(trimmed)
}
