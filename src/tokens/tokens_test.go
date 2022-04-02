package tokens

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	var tests = [][][]Tokens{
		{
			InfixToPostfix([]Tokens{Push(33), Plus(), Push(22), Dump()}),
			[]Tokens{Push(33), Push(22), Plus(), Dump()},
		},
		{
			InfixToPostfix([]Tokens{Push(33), Plus(), Push(22), Min(), Push(20)}),
			[]Tokens{Push(33), Push(22), Plus(), Push(20), Min()},
		},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Testing two [][]Tokens")
		t.Run(testname, func(t *testing.T) {
			if reflect.DeepEqual(tt[0], tt[1]) == false {
				t.Errorf("got %v, want %v", tt[0], tt[1])
			}
		})
	}

}
