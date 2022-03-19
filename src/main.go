package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	b "scythe.com/uni/src/build"
	l "scythe.com/uni/src/lexer"
)

// Interpret the user input from ParseLine function
func sim() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("$uni-> ")
		reader.Scan()
		stack := l.ParseLine(reader.Text())
		fin := l.InfixToRPN(stack)
		b.Simulate(fin)
	}
}

func Usage() {
	fmt.Println(`
Usage: uni <command> [argument]

Commands:
	- run
	- build [file]`)
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		// Build command
		case "build":
			if len(os.Args) >= 3 {
				if strings.HasSuffix(os.Args[2], ".uf") || strings.HasSuffix(os.Args[2], ".uo") {
					b.GenerateAssembly(l.InfixToRPN(l.LexFile(os.Args[2])))
				} else {
					fmt.Println("err: Please provide a valid file. (.uo, .uf)")
					os.Exit(1)
				}
				// Error
			} else {
				fmt.Println("err: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni build <file>")
			}
		// Interpret
		case "run":
			sim()
		default:
			fmt.Println("err: The command " + "\"" + os.Args[1] + "\"" + " is not valid.")
			os.Exit(1)
		}
	} else {
		Usage()
	}
}
