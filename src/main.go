package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	p "scythe.com/uni/src/parser"
	t "scythe.com/uni/src/tokens"
)

// Pop and return last element of a list
func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func interpreter(entry []t.Tokens) {
	arr := make([]int, 0)
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			arr = append(arr, i.GetValue())
		} else if i.GetOp() == t.OP_PLUS {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, a+b)
		} else if i.GetOp() == t.OP_DUMP {
			a := pop(&arr)
			fmt.Println(a)
		} else if i.GetOp() == t.OP_MIN {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, a-b)
		}
	}
}

func GenerateAssembly(entry []t.Tokens) {
	f, _ := os.Create("out.asm")
	f.WriteString(`segment .text
global _start

_start:` + "\n")
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			f.WriteString(fmt.Sprintf("	;; -- pushing value %d --\n", i.GetValue()))
			f.WriteString(fmt.Sprintf("	push %d\n", i.GetValue()))
		} else if i.GetOp() == t.OP_PLUS {
			f.WriteString(`	;; -- adding 2 values -
	pop rax
	pop rbx
	add rax, rbx
	push rax`)
		} else if i.GetOp() == t.OP_DUMP {
			// Asm
		} else if i.GetOp() == t.OP_MIN {
			// Asm
		}
	}
	f.Close()
	o1, _ := exec.Command("nasm", "-f", "elf64", "out.asm").Output()
	o2, _ := exec.Command("ld", "-o", "out", "out.o").Output()
	fmt.Printf("%s", o1)
	fmt.Printf("%s", o2)
	// defer os.Remove("out.asm")
	defer os.Remove("out.o")
}

// Interpret the user input from ParseLine function
func sim() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("$uni-> ")
		reader.Scan()
		stack := p.ParseLine(reader.Text())
		interpreter(stack)
	}
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "build":
			if len(os.Args) >= 3 {
				if strings.HasSuffix(os.Args[2], ".uf") || strings.HasSuffix(os.Args[2], ".uo") {
					GenerateAssembly(p.ParseFile(os.Args[2]))
				} else {
					fmt.Println("err: Please provide a valid file. (.uo, .uf)")
					os.Exit(1)
				}
			} else {
				fmt.Println("err: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni build <file>")
			}
		case "run":
			sim()
		default:
			fmt.Println("err: The command specified is not valid.")
			os.Exit(1)
		}
	} else {
		fmt.Println(`
Usage: uni <command> [argument]

Commands:
	- run
	- build [file]
	`)
	}
}
