package lexer

import (
	"fmt"
	"testing"

	"reflect"

	"scythe.com/uni/src/tokens"
)

func TestLexString(t *testing.T) {
	got := LexString([]string{"38", "22", "+", "dmp"})
	want := []tokens.Tokens{tokens.Push(38), tokens.Push(22), tokens.Plus(), tokens.Dump()}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("err! got : " + fmt.Sprint(got) + "for expected : " + fmt.Sprint(want))
	}
}
