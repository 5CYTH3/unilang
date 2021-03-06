package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	b "scythe.com/uni/src/build"
	l "scythe.com/uni/src/lexer"
)

// Simulate the input, lexed from a string to tokens.
func sim() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Unilang 0.0.1. Development version. Report bugs at https://github.com/5CYTH3/unilang/issues\n")
	for {
		// "\033[32m" is the color code for green. "\033[0m" is the reset code.
		fmt.Printf("\033[32m" + "$uni-> " + "\033[0m")
		reader.Scan()
		input := reader.Text()
		b.Simulate(l.LexString(l.CleanString(input)))
	}
}

func Usage() {
	fmt.Println(`
Usage: uni <command> [argument]

Commands:
	- run (interpreter)
	- build [file] (compile a file)
	- version (get the version of the program)`)
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		// Build command, ASM generation (GenerateAssembly function) and full process of lexing
		case "build":
			if len(os.Args) >= 3 {
				if strings.HasSuffix(os.Args[2], ".uf") || strings.HasSuffix(os.Args[2], ".uo") {
					lexFile := l.LexFile(os.Args[2])
					b.GenerateAssembly(lexFile)
					fmt.Println(lexFile)
				} else {
					fmt.Println("err001: Please provide a valid file. (.uo, .uf)")
					os.Exit(1)
				}
				// Error
			} else {
				fmt.Println("err002: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni build <file>")
			}
		// Simulating of the program.
		case "run":
			sim()
		case "version":
			fmt.Println("Unilang 0.0.1. Development version. Report bugs at https://github.com/5CYTH3/unilang/issues")
		default:
			fmt.Println("err003: The command " + "\"" + os.Args[1] + "\"" + " is not valid.")
			os.Exit(1)
		}
	} else {
		Usage()
	}
}
