package build

import (
	"fmt"
	"os"
	"os/exec"

	t "scythe.com/uni/src/tokens"
)

func GenerateAssembly(entry []t.Tokens) {
	f, _ := os.Create("out.asm")
	f.WriteString(`segment .text
global _start

_start:` + "\n")
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			f.WriteString(fmt.Sprintf("	;; -- pushing value %d --\n", i.GetValue()))
			f.WriteString(fmt.Sprintf("	push rax, %d\n", i.GetValue()))
		} else if i.GetOp() == t.OP_PLUS {
			f.WriteString(`	;; -- adding 2 values --
	add rax, rbx
	pop     eax
	ret`)
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
