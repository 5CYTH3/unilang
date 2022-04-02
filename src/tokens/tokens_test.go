package tokens

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	got := InfixToPostfix([]Tokens{Push(33), Plus(), Push(22), Dump()})
	want := []Tokens{Push(33), Push(22), Plus(), Dump()}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("err! got : " + fmt.Sprint(got) + "for expected : " + fmt.Sprint(want))
	}
}
