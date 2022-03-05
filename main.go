package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	t "scythe.com/uni/tokens"
)

func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

// TEMPORARY
/*
func parse(file string) Operator {
	var stack []interface{}
	trimmed := strings.Split(file, " ")

	for _, i := range trimmed {
		switch i {
		case "+":
			plus()
		case "-":
			min()
		case "":
		}
	}
}
*/

// TEST
func compile(entry []t.Token) {
	var arr []int
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

func parse(file string) {
	/* f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	fmt.Println(trimmed)*/

	program := []t.Token{
		t.Push(3),
		t.Push(18),
		t.Plus(),
		t.Dump(),
	}

	compile(program)

}

func interpret() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("$uni-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println(text)
	}
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "compile":
			if len(os.Args) >= 3 {
				parse(os.Args[2])
			} else {
				fmt.Println("err: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni compile <file>")
			}
		case "test":
			interpret()
		}
	} else {
		fmt.Println(`
Usage: uni <command> [argument]

Commands:
	- test
	- compile [file]
	`)
	}
}
