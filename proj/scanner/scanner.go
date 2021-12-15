package scanner

import (
	"proj/token"
	"strings"
)

type Scanner struct {
	input string // Stores the input string passed in when creating the scanner
	index int    // Stores the index of where the indexs have been read
}

func New(input string) *Scanner {
	scanner := &Scanner{input, 0}
	return scanner
}

func (l *Scanner) NextToken() token.Token {
	output := token.Token{}
	accept := "" // Identifies the type of accept state
	lexeme := "" // Stores the lexeme of the current token
	var char byte
	var state int
	goto s0
s0:
	// Check if the end of the file has been reached
	if l.index == len(l.input) {
		accept = "EOF"
		lexeme = "EOF " // the last char will be cut in sout
		goto sout
	}
	char = l.input[l.index]

	// Read the next character and go to different states based on the char read
	lexeme += string(char)
	l.index += 1
	if isAlpha(char) {
		goto s1
	} else if isInt(char) {
		goto s2
	} else if isSingleSpecialChar(char) {
		goto s3
	} else if isFirstSpecialChar(char) {
		goto s4
	} else if isIgnore(char) {
		// Ignore any whitespace, \t tab and \n newline characters
		lexeme = lexeme[:len(lexeme)-1]
		goto s0
	} else {
		// Go to out if there is an illegal input that is not any of the above cases
		lexeme += " "
		goto sout
	}

s1: // State for identifier
	char = l.input[l.index]
	lexeme += string(char)
	l.index += 1
	accept = "ID"
	if isAlpha(char) || isInt(char) {
		goto s1
	} else {
		goto sout
	}

s2: // State for reading integers
	char = l.input[l.index]
	lexeme += string(char)
	l.index += 1
	accept = "NUMBER"
	if isInt(char) {
		goto s2
	} else {
		goto sout
	}

s3: // State for reading unary special character (which are all single character)
	char = l.input[l.index]
	lexeme += string(char)
	l.index += 1
	accept = "SINGLESPEC"
	goto sout

s4: // State for reading multiple special characters
	char = l.input[l.index]
	l.index += 1
	// Check if the next character can form a longer special character with the
	// current lexeme
	state = isSecondSpecialChar(lexeme, char)
	lexeme += string(char)
	if state == 0 { // illegal
		goto sout
	} else if state == 1 { // single character
		accept = "COMBSPECIAL"
		goto sout
	} else if state == 2 { // multiple (two) characters
		char = l.input[l.index]
		lexeme += string(char)
		l.index += 1
		accept = "COMBSPECIAL"
		goto sout
	} else if state == 3 {
		lexeme = ""
		goto s5
	}
	goto sout
s5:
	for char != '\n' {
		l.index += 1
		if l.index == len(l.input) {
			goto s0
		}
		char = l.input[l.index]
	}
	goto s0
sout:
	if accept != "" {
		lexeme = lexeme[:len(lexeme)-1]
		l.index--
		output.Literal = lexeme
		// Depending on the accept state, define the output type
		switch accept {
		case "ID": // case for IDs or key words
			if lexeme == "package" {
				output.Type = token.PACKAGE
			} else if lexeme == "import" {
				output.Type = token.IMPORT
			} else if lexeme == "fmt" {
				output.Type = token.FMT
			} else if lexeme == "type" {
				output.Type = token.TYPE
			} else if lexeme == "struct" {
				output.Type = token.STRUCT
			} else if lexeme == "int" {
				output.Type = token.INT
			} else if lexeme == "bool" {
				output.Type = token.BOOL
			} else if lexeme == "var" {
				output.Type = token.VAR
			} else if lexeme == "func" {
				output.Type = token.FUNC
			} else if lexeme == "Scan" {
				output.Type = token.SCAN
			} else if lexeme == "Print" {
				output.Type = token.PRINT
			} else if lexeme == "Println" {
				output.Type = token.PRINTLN
			} else if lexeme == "if" {
				output.Type = token.IF
			} else if lexeme == "else" {
				output.Type = token.ELSE
			} else if lexeme == "for" {
				output.Type = token.FOR
			} else if lexeme == "return" {
				output.Type = token.RETURN
			} else if lexeme == "true" {
				output.Type = token.TRUE
			} else if lexeme == "false" {
				output.Type = token.FALSE
			} else if lexeme == "nil" {
				output.Type = token.NIL
			} else { // not a key word
				output.Type = token.ID
			}
		case "NUMBER":
			output.Type = token.NUMBER
		case "SINGLESPEC":
			if lexeme == "{" {
				output.Type = token.LEFTCURL
			} else if lexeme == "}" {
				output.Type = token.RIGHTCURL
			} else if lexeme == ";" {
				output.Type = token.SEMICOLON
			} else if lexeme == "\"" {
				output.Type = token.QUOTATION
			} else if lexeme == "(" {
				output.Type = token.LEFTBRAC
			} else if lexeme == ")" {
				output.Type = token.RIGHTBRAC
			} else if lexeme == "*" {
				output.Type = token.ASTRIX
			} else if lexeme == "+" {
				output.Type = token.PLUS
			} else if lexeme == "-" {
				output.Type = token.MINUS
			} else if lexeme == "," {
				output.Type = token.COMMA
			} else if lexeme == "." {
				output.Type = token.DOT
			} else {
				output.Type = token.ILLEGAL
			}
		case "COMBSPECIAL":
			if lexeme == "/" {
				output.Type = token.DIVIDE
			} else if lexeme == "!" {
				output.Type = token.NOT
			} else if lexeme == "&" {
				output.Type = token.AMP
			} else if lexeme == "<" {
				output.Type = token.LESSTHAN
			} else if lexeme == "<=" {
				output.Type = token.LEQ
			} else if lexeme == ">" {
				output.Type = token.MORETHAN
			} else if lexeme == ">=" {
				output.Type = token.GEQ
			} else if lexeme == "==" {
				output.Type = token.EQUAL
			} else if lexeme == "!=" {
				output.Type = token.NOTEQUAL
			} else if lexeme == "&&" {
				output.Type = token.AND
			} else if lexeme == "||" {
				output.Type = token.OR
			} else if lexeme == "=" {
				output.Type = token.ASSIGN
			} else {
				output.Type = token.ILLEGAL
			}
		case "EOF":
			output.Type = token.EOF
			output.Literal = "EOF"
		}
	} else {
		lexeme = lexeme[:len(lexeme)-1]
		output.Type = token.ILLEGAL
		output.Literal = lexeme
	}
	return output
}

// Check if the input char is a character from the English alphabet
func isAlpha(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') {
		return false
	} else {
		return true
	}
}

// Check if the input char is a character that has special token type in Cal
func isFirstSpecialChar(b byte) bool {
	specialChar := "<>|&=!/"
	if strings.ContainsAny(specialChar, string(b)) {
		return true
	} else {
		return false
	}
}

func isSecondSpecialChar(s string, b byte) int {
	if s == "!" || s == "<" || s == ">" || s == "=" {
		if b == '=' { // a single "=" is legal
			return 2
		} else {
			return 1
		}
	} else if s == "&" {
		if b == '&' { // a single "&" is legal
			return 2
		} else {
			return 1
		}
	} else if s == "|" { // a single "|" is illegal
		if b == '|' {
			return 2
		} else {
			return 0
		}
	} else if s == "/" { // a single "/" is legal
		if b == '/' {
			return 3 // return 3 for a special case to ignore til new line
		} else {
			return 1
		}
	}
	return 0
}

func isSingleSpecialChar(b byte) bool {
	specialChar := ";{}().+-*ε,\""
	if strings.ContainsAny(specialChar, string(b)) {
		return true
	} else {
		return false
	}
}

/*/ Check if the input char is a number between 1 and 9
func isFirstInt(b byte) bool {
	if b < '1' || b > '9' {
		return false
	} else {
		return true
	}
}*/

// Check if the input char is a number between 0 and 9
func isInt(b byte) bool {
	if b < '0' || b > '9' {
		return false
	} else {
		return true
	}
}

// Check if the input char is a space, newline or tab that should be ignored
func isIgnore(b byte) bool {
	if b == '\n' || b == '\r' || b == '\t' || b == ' ' {
		return true
	} else {
		return false
	}
}
