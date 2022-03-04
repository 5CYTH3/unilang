package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Operator int64

const (
	OP_PLUS Operator = iota
	OP_MIN
	OP_PUSH
)

func plus() Operator {
	return OP_PLUS
}

func min() Operator {
	return OP_MIN
}

func push() Operator {
	return OP_PUSH
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
func test(file string) {
	trimmed := strings.Split(file, " ")
	var arr []Operator
	for _, i := range trimmed {
		switch i {
		case "+":
			arr = append(arr, plus())
		case "-":
			arr = append(arr, min())
		}
	}

	for _, j := range arr {
		switch j {
		case OP_PLUS:

		}
	}
	fmt.Println(arr)
}

func compile(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	t_file := string(f)
	test(t_file)

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
