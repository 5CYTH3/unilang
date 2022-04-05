package build

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	t "scythe.com/uni/src/tokens"
	"scythe.com/uni/src/util"
)

func GenerateAssembly(entry []t.Tokens) {
	f, _ := os.Create("out.asm")
	f.WriteString(`segment .text
dump:
	mov r9, rax
	sub rsp, 40
	mov BYTE [rsp+31], 10
	lea rcx, [rsp+30]

.L2:
	mov rax, rdi
	lea r8, [rsp+32]
	mul r9
	mov rax, rdi
	sub r8, rcx
	shr rdx, 3
	lea rsi, [rdx+rdx*4]
	add rsi, rsi
	sub rax, rsi
	add eax, 48
	mov BYTE [rcx], al
	mov rax, rdi
	mov rdi, rdx
	mov rdx, rcx
	sub rcx, 1
	cmp rax, 9
	ja  .L2
	lea rax, [rsp+32]
	mov edi, 1
	sub rdx, rax
	xor eax, eax
	lea rsi, [rsp+32+rdx]
	mov rdx, r8
	mov rax, 1
	syscall
	add rsp, 40
	ret

global _start` + "\n")
	f.WriteString("_start:\n")
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			f.WriteString(fmt.Sprintf("	;; -- pushing value %d --\n", i.GetValue()))
			f.WriteString(fmt.Sprintf("	push %d\n", i.GetValue()))
		} else if i.GetOp() == t.OP_PLUS {
			f.WriteString(`	;; -- adding 2 values --
	pop rax
	pop rbx
	add rax, rbx
	push rax
	ret` + "\n")
		} else if i.GetOp() == t.OP_DIV {
			f.WriteString(`;; -- Divison not supported for now --`)
		} else if i.GetOp() == t.OP_MIN {
			f.WriteString(`	;; -- substracting 2 values --
	pop rax
	pop rbx
	sub rbx, rax
	push rbx
	`)
		} else if i.GetOp() == t.OP_MUL {
			f.WriteString(` ;; -- multiplication is not supported --`)
		} else if i.GetOp() == t.OP_DUMP {
			f.WriteString(`pop rdi
	call dump` + "\n")
		}
	}
	f.WriteString(`	mov rax, 60
	mov rdi, 0
	syscall`)
	f.Close()
	ExecuteCommand("nasm", "-f", "elf64", "out.asm")
	ExecuteCommand("ld", "-o", "out.exe", "out.o")
	// defer os.Remove("out.asm")
	// defer os.Remove("out.o")
}

func ExecuteCommand(cmdName string, cmdArg ...string) ([]byte, error) {
	out, err := exec.Command(cmdName, cmdArg...).Output()
	fmt.Printf("%s", out)
	if err != nil {
		log.Fatal(err)
	}
	return out, err
}

func Simulate(entry []t.Tokens) {
	arr := make([]int, 0)
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			arr = append(arr, int(i.GetValue()))
		} else if i.GetOp() == t.OP_PLUS {
			a := util.Pop(&arr)
			b := util.Pop(&arr)
			arr = append(arr, a+b)
		} else if i.GetOp() == t.OP_DUMP {
			a := util.Pop(&arr)
			fmt.Println(a)
		} else if i.GetOp() == t.OP_MIN {
			a := util.Pop(&arr)
			b := util.Pop(&arr)
			arr = append(arr, a-b)
		} else if i.GetOp() == t.OP_MUL {
			a := util.Pop(&arr)
			b := util.Pop(&arr)
			arr = append(arr, a*b)
		} else if i.GetOp() == t.OP_DIV {
			a := util.Pop(&arr)
			b := util.Pop(&arr)
			arr = append(arr, b/a)
		} else {
			fmt.Printf("Invalid operator")
		}
	}
	fmt.Println(arr)
}
