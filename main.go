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

func push(x interface{}) Operator {
	return OP_PUSH
}

func plus() Operator {
	return OP_PLUS
}

func dump() {

}

func compile(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	t_file := string(f)
	fmt.Println(t_file)

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
