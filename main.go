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
	OP_DUMP
)

func plus() [1]Operator {
	return [1]Operator{OP_PLUS}
}

func min() [1]Operator {
	return [1]Operator{OP_MIN}
}

func push(value int) [2]interface{} {
	return [2]interface{}{OP_PLUS, value}
}

func dump() [1]interface{} {
	return [1]interface{}{OP_DUMP}
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
func test(entry []interface{}) {
	var arr []interface{}
	for _, i := range entry {
		if i[0] == byte(OP_PUSH) {
			arr = append(arr, i[1])
		} else if i[0] == byte(OP_MIN) {
			a := arr[:len(arr)-1]
			b := arr[:len(arr)-1]
			arr = append(arr, a+b)
		} else if i[0] == byte(OP_DUMP) {
			a := arr[:len(arr)-1]
			fmt.Println(a)
		}
		fmt.Println(arr)
	}

}

func compile(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	program := []interface{}{
		push(5),
		push(4),
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
