package parser

import (
	ct "proj/token"
	"testing"
)

var validPackage = []ct.Token{
	{ct.PACKAGE, "package"},
	{ct.ID, "tester"},
	{ct.SEMICOLON, ";"},
}

var validImport = []ct.Token{
	{ct.IMPORT, "import"},
	{ct.QUOTATION, "\""},
	{ct.FMT, "fmt"},
	{ct.QUOTATION, "\""},
	{ct.SEMICOLON, ";"},
}

var eof = ct.Token{ct.EOF, "EOF"}

var validTypeDec1 = []ct.Token{
	{ct.TYPE, "type"},
	{ct.ID, "coordinate"},
	{ct.STRUCT, "struct"},
	{ct.LEFTCURL, "{"},
	{ct.ID, "x"},
	{ct.INT, "int"},
	{ct.SEMICOLON, ";"},
	{ct.ID, "y"},
	{ct.INT, "int"},
	{ct.SEMICOLON, ";"},
	{ct.RIGHTCURL, "}"},
	{ct.SEMICOLON, ";"},
}

var invalidTypeDec1 = []ct.Token{
	{ct.TYPE, "type"},
	{ct.ID, "coordinate"},
	{ct.STRUCT, "struct"},
	{ct.LEFTCURL, "{"},
	{ct.RIGHTCURL, "}"},
	{ct.SEMICOLON, ";"},
}

var validDeclaration1 = []ct.Token{
	{ct.VAR, "var"},
	{ct.ID, "distance"},
	{ct.COMMA, ","},
	{ct.ID, "mass"},
	{ct.INT, "int"},
	{ct.SEMICOLON, ";"},
}

var validDeclaration2 = []ct.Token{
	{ct.VAR, "var"},
	{ct.ID, "distance"},
	{ct.COMMA, ","},
	{ct.ID, "mass"},
	{ct.ASTRIX, "*"},
	{ct.ID, "coordinate"},
	{ct.SEMICOLON, ";"},
}

var validFunctionToIncludeStatements = []ct.Token{
	{ct.FUNC, "func"},
	{ct.ID, "calcDistance"},
	{ct.LEFTBRAC, "("},
	{ct.RIGHTBRAC, ")"},
	{ct.LEFTCURL, "{"},
}

var validFunction2 = []ct.Token{
	{ct.FUNC, "func"},
	{ct.ID, "calcDistance"},
	{ct.LEFTBRAC, "("},
	{ct.ID, "x"},
	{ct.INT, "int"},
	{ct.COMMA, ","},
	{ct.ID, "y"},
	{ct.INT, "int"},
	{ct.RIGHTBRAC, ")"},
	{ct.LEFTCURL, "{"},
	{ct.RIGHTCURL, "}"},
}

var validFunction3 = []ct.Token{
	{ct.FUNC, "func"},
	{ct.ID, "calcDistance"},
	{ct.LEFTBRAC, "("},
	{ct.ID, "isDist"},
	{ct.BOOL, "bool"},
	{ct.RIGHTBRAC, ")"},
	{ct.LEFTCURL, "{"},
	{ct.VAR, "var"},
	{ct.ID, "distance"},
	{ct.COMMA, ","},
	{ct.ID, "mass"},
	{ct.ASTRIX, "*"},
	{ct.ID, "coordinate"},
	{ct.SEMICOLON, ";"},
}

/*
func calDistance(isDist bool){
	fmt.Scan(&datatxt);
}
*/

var validLoop1 = []ct.Token{
	{ct.FOR, "for"},
	{ct.LEFTBRAC, "("},

	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.NUMBER, "2"},
	{ct.ASTRIX, "*"},
	{ct.NUMBER, "3"},
	{ct.RIGHTBRAC, ")"},
	{ct.LEFTCURL, "{"},
	{ct.RIGHTCURL, "}"},
	{ct.RIGHTCURL, "}"},
}

var validConditional1 = []ct.Token{
	{ct.IF, "if"},
	{ct.LEFTBRAC, "("},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.NUMBER, "2"},
	{ct.ASTRIX, "*"},
	{ct.NUMBER, "3"},
	{ct.MORETHAN, ">"},
	{ct.NUMBER, "4"},
	{ct.RIGHTBRAC, ")"},
	{ct.RIGHTBRAC, ")"},
	{ct.LEFTCURL, "{"},
	{ct.RIGHTCURL, "}"},
	{ct.ELSE, "else"},
	{ct.LEFTCURL, "{"},
	{ct.ID, "a"},
	{ct.ASSIGN, "="},
	{ct.NUMBER, "1"},
	{ct.SEMICOLON, ";"},
	{ct.RIGHTCURL, "}"},
	{ct.RIGHTCURL, "}"},
}

var validAssignment1 = []ct.Token{
	{ct.ID, "coordinate"},
	{ct.DOT, "."},
	{ct.ID, "x"},
	{ct.DOT, "."},
	{ct.ID, "x"},
	{ct.DOT, "."},
	{ct.ID, "y"},
	{ct.ASSIGN, "="},
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.NUMBER, "2"},
	{ct.ASTRIX, "*"},
	{ct.NUMBER, "3"},
	{ct.SEMICOLON, ";"},
	{ct.RIGHTCURL, "}"},
}

//coordinate(Expression, Expression)
//Should be: {coordinate(1+true*false-1/(1+2));}
//Got: {coordinate ((1 + (true * false)) - (1 / (1 + 2)));}
var validInvocation1 = []ct.Token{
	{ct.ID, "coordinate"},
	{ct.LEFTBRAC, "("},
	// valieExpression2
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.TRUE, "true"},
	{ct.ASTRIX, "*"},
	{ct.FALSE, "false"},
	{ct.MINUS, "-"},
	{ct.NUMBER, "1"},
	{ct.DIVIDE, "/"},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.NUMBER, "2"},
	{ct.RIGHTBRAC, ")"},
	{ct.RIGHTBRAC, ")"},
	{ct.SEMICOLON, ";"},
	{ct.RIGHTCURL, "}"},
}

//coordinate(Expression, Expression)
//Should be: {coordinate(1||2&&3, 1 + 2);}
//Got:coordinate((1||(2&&3)),
//{coordinate((1||(2&&3)),(1 + 2));}
//Which is correct
var validInvocation2 = []ct.Token{
	{ct.ID, "coordinate"},
	{ct.LEFTBRAC, "("},
	// validExpression1
	{ct.NUMBER, "1"},
	{ct.OR, "||"},
	{ct.NUMBER, "2"},
	{ct.AND, "&&"},
	{ct.NUMBER, "3"},
	{ct.COMMA, ","},
	// validExpression3
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.NUMBER, "2"},
	{ct.RIGHTBRAC, ")"},
	{ct.SEMICOLON, ";"},
	{ct.RIGHTCURL, "}"},
}

var validFunction4 = []ct.Token{
	{ct.FUNC, "func"},
	{ct.ID, "calcDistance"},
	{ct.LEFTBRAC, "("},
	{ct.ID, "isDist"},
	{ct.BOOL, "bool"},
	{ct.RIGHTBRAC, ")"},
	{ct.LEFTCURL, "{"},
	{ct.VAR, "var"}, {ct.ID, "distance"}, {ct.COMMA, ","},
	{ct.ID, "mass"}, {ct.INT, "int"}, {ct.SEMICOLON, ";"},
	{ct.LEFTCURL, "{"},
	{ct.FMT, "fmt"}, {ct.DOT, "."}, {ct.PRINT, "Print"},
	{ct.LEFTBRAC, "("}, {ct.ID, "datatxt"}, {ct.RIGHTBRAC, ")"},
	{ct.SEMICOLON, ";"},
	{ct.LEFTCURL, "{"},
	{ct.RIGHTCURL, "}"},
	{ct.RIGHTCURL, "}"},
	{ct.RIGHTCURL, "}"},
}

var validExpression1 = []ct.Token{
	{ct.NUMBER, "1"},
	{ct.OR, "||"},
	{ct.NUMBER, "2"},
	{ct.AND, "&&"},
	{ct.NUMBER, "3"},
}

var validExpression2 = []ct.Token{
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.TRUE, "true"},
	{ct.ASTRIX, "*"},
	{ct.FALSE, "false"},
	{ct.MINUS, "-"},
	{ct.NUMBER, "1"},
	{ct.DIVIDE, "/"},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.NUMBER, "2"},
	{ct.RIGHTBRAC, ")"},
}

var validExpression3 = []ct.Token{
	{ct.NUMBER, "1"},
	{ct.PLUS, "+"},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "1"},
	{ct.ASTRIX, "*"},
	{ct.NUMBER, "3"},
	{ct.MINUS, "-"},
	{ct.NUMBER, "1"},
	{ct.DIVIDE, "/"},
	{ct.NUMBER, "4"},
	{ct.RIGHTBRAC, ")"},
	{ct.MORETHAN, ">"},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "3"},
	{ct.MINUS, "-"},
	{ct.NUMBER, "1"},
	{ct.RIGHTBRAC, ")"},
	{ct.OR, "||"},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "3"},
	{ct.EQUAL, "=="},
	{ct.NUMBER, "1"},
	{ct.RIGHTBRAC, ")"},
	{ct.OR, "&&"},
	{ct.LEFTBRAC, "("},
	{ct.NOT, "!"},
	{ct.ID, "add"},
	{ct.LEFTBRAC, "("},
	{ct.NUMBER, "1"},
	{ct.COMMA, ","},
	{ct.NUMBER, "2"},
	{ct.MORETHAN, ">"},
	{ct.NUMBER, "1"},
	{ct.RIGHTBRAC, ")"},
	{ct.EQUAL, "=="},
	{ct.ID, "num"},
	{ct.DOT, "."},
	{ct.ID, "n1"},
	{ct.DOT, "."},
	{ct.ID, "n2"},
	{ct.RIGHTBRAC, ")"},
	{ct.OR, "||"},
	{ct.MINUS, "-"},
	{ct.ID, "n2"},
	{ct.LEFTBRAC, "("},
	{ct.RIGHTBRAC, ")"},
	{ct.LEQ, "<="},
	{ct.NIL, "nil"},
}

func Test1(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{}
	tokens = append(tokens, validPackage...)
	tokens = append(tokens, validImport...)
	tokens = append(tokens, validTypeDec1...)
	tokens = append(tokens, validDeclaration1...)
	tokens = append(tokens, validFunctionToIncludeStatements...)
	tokens = append(tokens, validAssignment1...)
	tokens = append(tokens, eof)

	// Define  a new scanner for some Cal program
	parser := New(tokens)
	ast := parser.Parse()
	if ast != nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", ast)
	}
}

func Test2(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{}
	tokens = append(tokens, validPackage...)
	tokens = append(tokens, validImport...)
	tokens = append(tokens, validTypeDec1...)
	tokens = append(tokens, validFunction3...)
	tokens = append(tokens, validConditional1...)
	tokens = append(tokens, eof)

	// Define  a new scanner for some Cal program
	parser := New(tokens)
	ast := parser.Parse()
	if ast != nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", ast)
	}
}

func Test3(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{}
	tokens = append(tokens, validPackage...)
	tokens = append(tokens, validImport...)
	tokens = append(tokens, validTypeDec1...)
	tokens = append(tokens, validDeclaration2...)
	tokens = append(tokens, validFunctionToIncludeStatements...)
	tokens = append(tokens, validConditional1...)
	tokens = append(tokens, eof)

	// Define  a new scanner for some Cal program
	parser := New(tokens)
	ast := parser.Parse()
	if ast != nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", ast)
	}

}

func Test4(t *testing.T) {

	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{}
	tokens = append(tokens, validPackage...)
	tokens = append(tokens, validImport...)
	tokens = append(tokens, validDeclaration2...)
	tokens = append(tokens, validFunctionToIncludeStatements...)
	tokens = append(tokens, validInvocation2...)
	tokens = append(tokens, eof)

	// Define  a new scanner for some Cal program
	parser := New(tokens)
	ast := parser.Parse()
	if ast != nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, nil, ast)
	}
}
