package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Operator int64
type Token struct {
	op    Operator
	value int
}

const (
	OP_PLUS Operator = iota
	OP_MIN
	OP_PUSH
	OP_DUMP
)

func plus() Token {
	return Token{OP_PLUS, 0}
}

func min() Token {
	return Token{OP_MIN, 0}
}

func push(value int) Token {
	return Token{OP_PUSH, value}
}

func dump() Token {
	return Token{OP_DUMP, 0}
}

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
func test(entry []Token) {
	var arr []int
	for _, i := range entry {
		if i.op == OP_PUSH {
			arr = append(arr, i.value)
		} else if i.op == OP_PLUS {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, a+b)
		} else if i.op == OP_DUMP {
			a := pop(&arr)
			fmt.Println(a)
		}
	}

}

func compile(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	fmt.Println(trimmed)
	program := []Token{
		push(3),
		push(18),
		plus(),
		dump(),
	}

	test(program)

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
			compile(os.Args[2])
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
