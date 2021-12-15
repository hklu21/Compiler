package main

import (
	"flag"
	"fmt"
	"os"
	"proj/codegen"
	"proj/ir"
	"proj/parser"
	"proj/sa"
	"proj/scanner"
	ct "proj/token"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Simple main program that takes the second input as the filename
	// Scans the file for legal tokens and prints them to the standard output
	lexFlag := flag.Bool("lex", false, "Print out the scanned Tokens from Scanner")
	astFlag := flag.Bool("ast", false, "Print out the ast Representation of the code")
	iLocFlag := flag.Bool("iloc", false, "Print out the generated iLoc code")
	flag.Parse()

	argLen := len(os.Args[1:])
	var arguments []string
	filename := "simple1.golite"

	if argLen > 1 {
		arguments = os.Args[:]
		filename = arguments[argLen]
	} else {
		fmt.Println("No input file was detected")
		return
	}

	dat, err := os.ReadFile(filename)
	check(err)

	input := string(dat) + " "
	scanner := scanner.New(string(input))
	tok := scanner.NextToken()
	tokens := []ct.Token{}
	tokens = append(tokens, tok)
	for tok.Literal != "EOF" {
		//fmt.Println("Token." + string(tok.Type) + ": " + tok.Literal)
		tok = scanner.NextToken()
		tokens = append(tokens, tok)
	}
	if *lexFlag {
		fmt.Printf("%v", tokens)
		fmt.Println("Token." + string(tok.Type) + ": " + tok.Literal)
	}
	parser := parser.New(tokens)
	ast := parser.Parse()
	if *astFlag {
		ast.PrintAST(0)
	}

	symbolTable := sa.PerformSA(ast)
	//fmt.Printf("%v\n", symbolTable)
	var frags []*ir.FuncFrag
	if symbolTable != nil {
		frags = ast.TranslateToILOC(*symbolTable)
	}
	if *iLocFlag {
		//fmt.Println("iLoc printing functionality to be implemented")
		for _, frag := range frags {
			for _, instruction := range frag.Body {
				fmt.Println(instruction.String())
			}
		}
	}
	output := make([]string, 0)
	for _, frag := range frags {
		output = append(output, frag.GenerateCode()...)
	}
	codegen.GenerateFile(output, filename[:len(filename)-7])
}
