#!/bin/sh
set -xe

nasm -felf64 out.asm
ld -o out out.o