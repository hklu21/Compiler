package token

type TokenType string

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Tokens that are strings
	PACKAGE = "package"
	ID      = "id"
	IMPORT  = "import"
	FMT     = "fmt"
	TYPE    = "type"
	STRUCT  = "struct"
	INT     = "int"
	BOOL    = "bool"
	VAR     = "var"
	FUNC    = "func"
	IF      = "if"
	ELSE    = "else"
	FOR     = "for"
	RETURN  = "return"

	// Functions in GoLite
	SCAN    = "Scan"
	PRINT   = "Print"
	PRINTLN = "Println"

	// Data in GoLite
	NUMBER = "number"
	TRUE   = "true"
	FALSE  = "false"
	NIL    = "nil"

	// Special Characters
	SEMICOLON = "SEMICOLON" // ；
	QUOTATION = "QUOTATION" // "
	LEFTCURL  = "LEFTCURL"  // {
	RIGHTCURL = "RIGHTCURL" // }
	LEFTBRAC  = "LEFTBRAC"  // (
	RIGHTBRAC = "RIGHTBRAC" // )
	ASSIGN    = "ASSIGN"    // =
	DOT       = "DOT"       // .
	AMP       = "AMP"       // &
	OR        = "OR"        // ||
	AND       = "AND"       // &&
	EQUAL     = "EQUAL"     // ==
	NOTEQUAL  = "NOTEQUAL"  // !=
	MORETHAN  = "MORETHAN"  // >
	LESSTHAN  = "LESSTHAN"  // <
	GEQ       = "GEQ"       // >=
	LEQ       = "LEQ"       // <=
	PLUS      = "PLUS"      // +
	MINUS     = "MINUS"     // -
	ASTRIX    = "ASTRIX"    // *
	DIVIDE    = "DIVIDE"    // /
	NOT       = "NOT"       // !
	COMMA     = "COMMA"     // ,
)

type Token struct {
	Type    TokenType
	Literal string
}
