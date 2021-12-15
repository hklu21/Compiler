package scanner

import (
	"proj/token"
	"testing"
)

type ExpectedResult struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func VerifyTest(t *testing.T, tests []ExpectedResult, scanner *Scanner) {

	for i, tt := range tests {
		tok := scanner.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("FAILED[%d] - incorrect token.\nexpected=%v\ngot=%v\n",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("FAILED[%d] - incorrect token literal.\nexpected=%v\ngot=%v\n",
				i, tt.expectedLiteral, tok.Literal)
		}

	}
}

func Test1(t *testing.T) {
	// Test all of the recognisable string inputs
	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `package id five import fmt type struct
	int bool var func Scan Print Println if else
	for return true false nil
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.PACKAGE, "package"},
		{token.ID, "id"},
		{token.ID, "five"},
		{token.IMPORT, "import"},
		{token.FMT, "fmt"},
		{token.TYPE, "type"},
		{token.STRUCT, "struct"},
		{token.INT, "int"},
		{token.BOOL, "bool"},
		{token.VAR, "var"},
		{token.FUNC, "func"},
		{token.SCAN, "Scan"},
		{token.PRINT, "Print"},
		{token.PRINTLN, "Println"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.FOR, "for"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.NIL, "nil"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Golite program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test2(t *testing.T) {
	// Test all of the recognisable string inputs
	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `;{}().+-/*",
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.SEMICOLON, ";"},
		{token.LEFTCURL, "{"},
		{token.RIGHTCURL, "}"},
		{token.LEFTBRAC, "("},
		{token.RIGHTBRAC, ")"},
		{token.DOT, "."},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.ASTRIX, "*"},
		{token.QUOTATION, "\""},
		{token.COMMA, ","},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Golite program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test3(t *testing.T) {
	// Test all of the recognisable string inputs
	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `!!====&&||&<><=>=|
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.NOT, "!"},
		{token.NOTEQUAL, "!="},
		{token.EQUAL, "=="},
		{token.ASSIGN, "="},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.AMP, "&"},
		{token.LESSTHAN, "<"},
		{token.MORETHAN, ">"},
		{token.LEQ, "<="},
		{token.GEQ, ">="},
		{token.ILLEGAL, "|"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Golite program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test4(t *testing.T) {
	// Test all of the recognisable string inputs
	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := ` package main;
				import "fmt";
				func main() {
					var a int;
					a = 3 + 4 + 5;
					var b *int;
					b = &a;
					*b = 3 * 5 + 3 / 5;
					if (a == *b) {
						fmt.Print(a);
					} else {
						fmt.Println(a);
					}
				}
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.PACKAGE, "package"}, {token.ID, "main"}, {token.SEMICOLON, ";"},
		{token.IMPORT, "import"}, {token.QUOTATION, "\""}, {token.FMT, "fmt"}, {token.QUOTATION, "\""}, {token.SEMICOLON, ";"},
		{token.FUNC, "func"}, {token.ID, "main"}, {token.LEFTBRAC, "("}, {token.RIGHTBRAC, ")"}, {token.LEFTCURL, "{"},
		{token.VAR, "var"}, {token.ID, "a"}, {token.INT, "int"}, {token.SEMICOLON, ";"},
		{token.ID, "a"}, {token.ASSIGN, "="}, {token.NUMBER, "3"}, {token.PLUS, "+"}, {token.NUMBER, "4"}, {token.PLUS, "+"}, {token.NUMBER, "5"}, {token.SEMICOLON, ";"},
		{token.VAR, "var"}, {token.ID, "b"}, {token.ASTRIX, "*"}, {token.INT, "int"}, {token.SEMICOLON, ";"},
		{token.ID, "b"}, {token.ASSIGN, "="}, {token.AMP, "&"}, {token.ID, "a"}, {token.SEMICOLON, ";"},
		{token.ASTRIX, "*"}, {token.ID, "b"}, {token.ASSIGN, "="}, {token.NUMBER, "3"}, {token.ASTRIX, "*"}, {token.NUMBER, "5"},
		{token.PLUS, "+"}, {token.NUMBER, "3"}, {token.DIVIDE, "/"}, {token.NUMBER, "5"}, {token.SEMICOLON, ";"},
		{token.IF, "if"}, {token.LEFTBRAC, "("}, {token.ID, "a"}, {token.EQUAL, "=="}, {token.ASTRIX, "*"}, {token.ID, "b"}, {token.RIGHTBRAC, ")"}, {token.LEFTCURL, "{"},
		{token.FMT, "fmt"}, {token.DOT, "."}, {token.PRINT, "Print"}, {token.LEFTBRAC, "("}, {token.ID, "a"}, {token.RIGHTBRAC, ")"}, {token.SEMICOLON, ";"},
		{token.RIGHTCURL, "}"}, {token.ELSE, "else"}, {token.LEFTCURL, "{"},
		{token.FMT, "fmt"}, {token.DOT, "."}, {token.PRINTLN, "Println"}, {token.LEFTBRAC, "("}, {token.ID, "a"}, {token.RIGHTBRAC, ")"}, {token.SEMICOLON, ";"},
		{token.RIGHTCURL, "}"},
		{token.RIGHTCURL, "}"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Golite program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test5(t *testing.T) {
	// Test all of the recognisable string inputs
	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `fmt.Print(Hello World);
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.FMT, "fmt"},
		{token.DOT, "."},
		{token.PRINT, "Print"},
		{token.LEFTBRAC, "("},
		{token.ID, "Hello"},
		{token.ID, "World"},
		{token.RIGHTBRAC, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test8(t *testing.T) {

	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `let five = 5;
Print five; 
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.ID, "let"},
		{token.ID, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.PRINT, "Print"},
		{token.ID, "five"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test9(t *testing.T) {

	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `Print 03;
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.PRINT, "Print"},
		{token.NUMBER, "03"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test10(t *testing.T) {

	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `2Print;
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.NUMBER, "2"},
		{token.PRINT, "Print"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test11(t *testing.T) {

	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `123Num;
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.NUMBER, "123"},
		{token.ID, "Num"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test12(t *testing.T) {

	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `a=0111;
  printb b;
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.ID, "a"},
		{token.ASSIGN, "="},
		{token.NUMBER, "0111"},
		{token.SEMICOLON, ";"},
		{token.ID, "printb"},
		{token.ID, "b"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}

func Test7(t *testing.T) {

	// This is a raw string in Go (aka its a multiline string). This will be easy to
	input1 := `let a = 0 * 2 + 3 /4-5;
	Let b = */+-;
  printb a;
`
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.ID, "let"},
		{token.ID, "a"},
		{token.ASSIGN, "="},
		{token.NUMBER, "0"},
		{token.ASTRIX, "*"},
		{token.NUMBER, "2"},
		{token.PLUS, "+"},
		{token.NUMBER, "3"},
		{token.DIVIDE, "/"},
		{token.NUMBER, "4"},
		{token.MINUS, "-"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.ID, "Let"},
		{token.ID, "b"},
		{token.ASSIGN, "="},
		{token.ASTRIX, "*"},
		{token.DIVIDE, "/"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.SEMICOLON, ";"},
		{token.ID, "printb"},
		{token.ID, "a"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	// Define  a new scanner for some Cal program
	scanner := New(input1)

	// Verify that the scanner produces the tokens in the order that you expect.
	VerifyTest(t, expected, scanner)
}
