# Unilang
The Uni language is a project for a language that uses different paradigms relatively to the file type.
For a file ending with .uo, your code will be Object Oriented. For a file ending with .uf, yout code will be functional.

The compiler that I am trying to build is actually made in Golang. It's a fast, readable and garbage-collected language.

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
- [ ] Basic array calculation
- [ ] Basic file parser
- [ ] Basic interpreter parsing system
- [ ] Interpreter
- [ ] AST
- [ ] Compiler
- [ ] Turing-complete