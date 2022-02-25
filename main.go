package main

import (
	"fmt"
	"os"
)

type Operator int64

const (
	OP_PLUS Operator = iota
	OP_MIN
	OP_PUSH
)

func push(x interface{}) Operator {
	return OP_PUSH
}

func plus() Operator {
	return OP_PLUS
}

func dump() {

}

func compile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}

func interpret() {
	var input string
	fmt.Println("Interpreter function called.")
	for {
		fmt.Printf("uni>>")
		fmt.Scanln(&input)
		fmt.Println(input)
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
