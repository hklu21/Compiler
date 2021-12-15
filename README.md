# Project: Compiler for Golite

## Introduction

This is a compiler for a language called Golite (a "tiny Go"). Go [here](https://classes.cs.uchicago.edu/archive/2021/fall/51300-1/assignments/overview/index.html) to see its grammar and semantics.

## Files

token/: defined token type
scanner/: implementation of the scanner
paser/: implementation the paser
sa/: implementation of the semantic analysis
ir/: implementation of AST to ILOC
codegen/: implementation of code generation
ast/: structure definition and useful methods within ast
symbolTable/: structure definition of symbol table for semantic analysis.
golite/: the main function of the project (to run the each phase of the compiler).

## How to run

cd into the golite directory, and use the following command:

```shell
go run golite.go -f1 -f2 -f3 input.golite
```

you will get an input.s which an assembly file ready to be compiled be clang.

f1, f2, f3 are flags you can use to expect some output on stdout.

-lex: print the token lists input file generated by scanner.

-ast: print the ast strcucture generated by paser if there is no syntactical error

-iloc: print the iloc intermediate representation if no semanstic error


## DFA of Golite

Below is the DFA of Golite, where the green ones are final states and the blue ones are start state or normal states. Besides, there is red nodes which is special final states. It represent the token type "ID" or key words (such as "import", "fmt" and etc). And the token type is deiceded exactly by the token string it owns. This can extremly simply the DFA and the work precedure of scanner while maintaining the scanner's functions.

<details>
<summary><b>Click to see the DFA</b></summary>

![DFA of Golite](https://i.ibb.co/rtfnsqP/ey-Jjb2-Rl-Ijoi-Z3-Jhc-Ggg-TFJcbi-Ag-ICAw-KCh-Td-GFyd-Ckp-IC0t-XCIn-ICcgf-CAn-XFx0-Jy-B8-ICdc-XG4n-X.jpg).

</details>

## Renference

[Mermaid Tutorials](https://mermaid-js.github.io/mermaid/#/Tutorials?id=live-editor-tutorials)
