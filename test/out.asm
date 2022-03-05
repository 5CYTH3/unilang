%define SYS_EXIT 60

segment .text
	global _start

_start:
	mov rax, SYS_EXIT
	mov rdi, 69
	syscall
	ret