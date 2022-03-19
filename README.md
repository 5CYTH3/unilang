# Unilang
<img src="https://img.shields.io/github/downloads/5CYTH3/unilang/total?style=flat-square"> <img src="https://img.shields.io/github/commit-activity/m/5CYTH3/unilang?style=flat-square"> <img src="https://img.shields.io/github/languages/count/5CYTH3/unilang?style=flat-square">

The Uni language is a project for a language that uses different paradigms relatively to the file type.
For a file ending with .uo, your code will be Object Oriented. For a file ending with .uf, yout code will be functional.

The compiler that I am trying to build is actually made in Golang. It's a fast, readable and garbage-collected language.

:warning: If you are a Go developer and interested in contributing to the project, go check the [CONTRIBUTE.md](CONTRIBUTE.md) file.

## Syntax

Here are few examples of the syntax:

main.uf
```
package main

import "Std"

func main(): void {
    hw: string = "Hello World";
    Std.print(string);
}
```

main.uo
```
package main

import "Std"

class Main {
    func main(): void {
        hw: string = "Hello World";
        Std.print(string);
    }
}
```

## Objectives

All the steps I want to achieve for this programming language:

##### Basic parsing
- [x] Basic array calculation
- [x] Basic file parser
- [x] Basic interpreter parsing system
##### Advanced compilation
- [ ] Advanced syntax
- [ ] Interpreter
- [ ] AST
- [ ] Compiler
##### Advanced programming concepts
- [ ] Turing-complete
