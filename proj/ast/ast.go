package ast

import (
	"bytes"
	"fmt"
	"proj/ir"
	st "proj/symbolTable"
	"proj/token"
	"strconv"
	"strings"
)

// The base Node interface that all ast nodes have to access
type Node interface {
	TokenLiteral() string
	String() string
	TypeCheck(errors []string, symTable *st.SymbolTable) []string
	PrintAST(int)
}

type Statement interface {
	Node
	TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction
	CheckReturn(symTable *st.SymbolTable) bool
}

type Program struct {
	Pack  Package
	Types []Types
	Decl  []Declaration
	Funcs []Function
	//Expr Expression
}

func (pr *Program) TokenLiteral() string { return pr.String() }
func (pr *Program) String() string {
	var out bytes.Buffer
	out.WriteString(pr.Pack.String())

	out.WriteString("\n")
	for _, ar := range pr.Types {
		out.WriteString(ar.String())
		out.WriteString("\n")
	}
	for _, ar := range pr.Decl {
		out.WriteString(ar.String())
		out.WriteString("\n")
	}
	for _, ar := range pr.Funcs {
		out.WriteString(ar.String())
		out.WriteString("\n")
	}
	return out.String()
}
func (ast *Program) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Program")
	// TODO Consider if import "fmt"; need to be printed
	fmt.Println(out.String())
	ast.Pack.PrintAST(level + 1)
	for _, types := range ast.Types {
		types.PrintAST(level + 1)
	}
	// Contains Static Import Reporduction as follows
	// import "fmt";
	for _, decl := range ast.Decl {
		decl.PrintAST(level + 1)
	}
	for _, funcs := range ast.Funcs {
		funcs.PrintAST(level + 1)
	}
}

type Package struct {
	Node
	Id string
}

func (il *Package) TokenLiteral() string { return il.String() }
func (il *Package) String() string {
	var out bytes.Buffer
	out.WriteString("Package ")
	out.WriteString(il.Id)
	out.WriteString(";")
	out.WriteString("\n")
	out.WriteString("import ")
	out.WriteString("\"")
	out.WriteString("fmt")
	out.WriteString("\"")
	out.WriteString(";")
	return out.String()
}

func (ast *Package) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Package (Name: ")
	out.WriteString(ast.Id)
	out.WriteString(")")
	fmt.Println(out.String())
}

type Types struct {
	Node
	Id     string
	Fields []Field
}

func (il *Types) TokenLiteral() string { return il.String() }
func (il *Types) String() string {
	var out bytes.Buffer

	out.WriteString("type ")
	out.WriteString(il.Id)
	out.WriteString("struct")
	out.WriteString("{")
	for _, feild := range il.Fields {
		out.WriteString(feild.String())
	}
	out.WriteString("}")
	out.WriteString(";")

	return out.String()
}

func (ast *Types) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Types (Struct Name :")
	out.WriteString(ast.Id)
	out.WriteString(")")
	fmt.Println(out.String())
	for index, fields := range ast.Fields {
		fields.PrintAST(index, level+1)
	}
}

type Field struct {
	Node
	Id    string
	Value string
}

func (il *Field) TokenLiteral() string { return il.String() }
func (il *Field) String() string {
	var out bytes.Buffer

	out.WriteString(il.Id)
	out.WriteString(" ")
	out.WriteString(il.Value)
	out.WriteString(";")
	out.WriteString("\n")

	return out.String()
}
func (ast *Field) PrintAST(index int, level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Field (Struct Variable ")
	out.WriteString(strconv.Itoa(index))
	out.WriteString(") (Varaiable Name: ")
	out.WriteString(ast.Id)
	out.WriteString(", Variable Type: ")
	out.WriteString(ast.Value)
	out.WriteString(")")
	fmt.Println(out.String())
}

type Declaration struct {
	Node
	Id    []string
	Value string
}

func (il *Declaration) TokenLiteral() string { return il.String() }
func (il *Declaration) String() string {
	var out bytes.Buffer

	out.WriteString("var ")
	out.WriteString(il.Id[0])
	for _, id := range il.Id[1:] {
		out.WriteString(", ")
		out.WriteString(id)
	}
	out.WriteString(" ")
	out.WriteString(il.Value)
	out.WriteString("; ")

	return out.String()
}
func (ast *Declaration) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Declaration (Variable Name: ")
	for _, id := range ast.Id {
		out.WriteString(id)
		out.WriteString(",")
	}
	out.WriteString(" Variable Type: ")
	if ast.Value == "bool" || ast.Value == "int" {
		out.WriteString(ast.Value)
	} else {
		out.WriteString("*")
		out.WriteString(ast.Value)
	}
	out.WriteString(")")
	fmt.Println(out.String())
}

type Function struct {
	Node
	Id      string
	Param   []Parameter
	ReturnT string
	Dec     []Declaration
	States  []Statement
}

func (il *Function) TokenLiteral() string { return il.String() }
func (il *Function) String() string {
	var out bytes.Buffer

	out.WriteString("func ")
	out.WriteString(il.Id)
	out.WriteString("(")
	if len(il.Param) > 0 {
		out.WriteString(il.Param[0].String())
		if len(il.Param) > 1 {
			for _, Parameter := range il.Param[1:] {
				out.WriteString(", ")
				out.WriteString(Parameter.String())
			}
		}
	}
	out.WriteString(") ")
	out.WriteString(il.ReturnT)
	out.WriteString(" {")
	for _, Declaration := range il.Dec {
		out.WriteString(Declaration.String())
	}
	for _, statement := range il.States {
		out.WriteString(statement.String())
	}
	out.WriteString("}")

	return out.String()
}

func (ast *Function) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Function (Function Name: ")
	out.WriteString(ast.Id)
	out.WriteString(", Return Type: ")
	if ast.ReturnT == "bool" || ast.ReturnT == "int" || ast.ReturnT == "Epsilon" {
		out.WriteString(ast.ReturnT)
	} else {
		out.WriteString("*")
		out.WriteString(ast.ReturnT)
	}
	out.WriteString(")")
	fmt.Println(out.String())
	for index, param := range ast.Param {
		param.PrintAST(index, level+1)
	}
	for _, dec := range ast.Dec {
		dec.PrintAST(level + 1)
	}
	for _, states := range ast.States {
		states.PrintAST(level + 1)
	}
}

type Parameter struct {
	Node
	Id    string
	Value string
}

func (il *Parameter) TokenLiteral() string { return il.String() }
func (il *Parameter) String() string {
	var out bytes.Buffer

	out.WriteString(il.Id)
	out.WriteString(" ")
	out.WriteString(il.Value)

	return out.String()
}

func (ast *Parameter) PrintAST(input int, level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Param (Function Input ")
	out.WriteString(strconv.Itoa(input))
	out.WriteString(") (Parameter Name: ")
	out.WriteString(ast.Id)
	out.WriteString(", Parameter Type: ")
	if ast.Value != "bool" && ast.Value != "int" {
		out.WriteString("*")
	}
	out.WriteString(ast.Value)
	out.WriteString(")")
	fmt.Println(out.String())
}

type Block struct {
	Statement
	Block []Statement
}

func (il *Block) TokenLiteral() string { return il.String() }
func (il *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	for _, block := range il.Block {
		out.WriteString(block.String())
		out.WriteString(" ")
	}
	out.WriteString("}")

	return out.String()
}
func (ast *Block) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Block (Size: ")
	out.WriteString(strconv.Itoa(len(ast.Block)))
	out.WriteString(")")
	fmt.Println(out.String())
	for _, block := range ast.Block {
		block.PrintAST(level + 1)
	}
}

type Assignment struct {
	Statement
	Lvalue []string
	Exprs  Expression
}

func (as *Assignment) TokenLiteral() string { return as.String() }
func (as *Assignment) String() string {
	var out bytes.Buffer

	out.WriteString(as.Lvalue[0])
	if len(as.Lvalue) > 1 {
		for _, id := range as.Lvalue[1:] {
			out.WriteString(".")
			out.WriteString(id)
		}
	}
	out.WriteString("=")
	out.WriteString(as.Exprs.String())
	out.WriteString(";")

	return out.String()
}
func (ast *Assignment) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Assignment (LValue: ")
	out.WriteString(ast.Lvalue[0])
	if len(ast.Lvalue) > 1 {
		for _, lvalue := range ast.Lvalue[1:] {
			out.WriteString(".")
			out.WriteString(lvalue)
		}
	}
	out.WriteString(")")
	fmt.Println(out.String())
	ast.Exprs.PrintAST(level + 1)
}

type ReadPrint struct {
	Statement
	FuncName string
	Id       string
}

func (il *ReadPrint) TokenLiteral() string { return il.String() }
func (il *ReadPrint) String() string {
	var out bytes.Buffer

	out.WriteString("fmt.")
	out.WriteString(il.FuncName)
	out.WriteString("(")
	if il.FuncName == "Scan" {
		out.WriteString("&")
	}
	out.WriteString(il.Id)
	out.WriteString(");")

	return out.String()
}
func (ast *ReadPrint) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.ReadPrint (Function Type: ")
	out.WriteString(ast.FuncName)
	if ast.FuncName == "Scan" {
		out.WriteString(", Scan To Variable: &")
	} else {
		out.WriteString(", Print Variable: ")
	}
	out.WriteString(ast.Id)
	out.WriteString(")")
	fmt.Println(out.String())
}

type Conditional struct {
	Statement
	Expr      Expression
	IfBlock   Statement
	HasElse   bool
	ElseBlock Statement
}

func (cond *Conditional) TokenLiteral() string { return cond.String() }
func (cond *Conditional) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(cond.Expr.String())
	out.WriteString(") ")
	out.WriteString(cond.IfBlock.String())
	if cond.HasElse {
		out.WriteString(" else ")
		out.WriteString(cond.ElseBlock.String())
	}
	return out.String()
}
func (ast *Conditional) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Conditional (If Then")
	if ast.HasElse {
		out.WriteString(" Else")
	}
	out.WriteString(" Block)")
	fmt.Println(out.String())
	ast.Expr.PrintAST(level + 1)
	ast.IfBlock.PrintAST(level + 1)
	if ast.HasElse {
		ast.ElseBlock.PrintAST(level + 1)
	}
}

type Loop struct {
	Statement
	Expr     Expression
	ForBlock Statement
}

func (lp *Loop) TokenLiteral() string { return lp.String() }
func (lp *Loop) String() string {
	var out bytes.Buffer

	out.WriteString("for (")
	out.WriteString(lp.Expr.String())
	out.WriteString(")")
	out.WriteString(lp.ForBlock.String())

	return out.String()
}
func (ast *Loop) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Loop (For Loop)")
	fmt.Println(out.String())
	ast.Expr.PrintAST(level + 1)
	ast.ForBlock.PrintAST(level + 1)
}

type Return struct {
	Statement
	Exprs Expression
}

func (il *Return) TokenLiteral() string { return il.String() }
func (il *Return) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	out.WriteString(il.Exprs.String())
	out.WriteString(";")

	return out.String()
}
func (ast *Return) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	if ast.Exprs == nil {
		out.WriteString("*ast.Return (Return Value: Epsilon)")
		fmt.Println(out.String())
	} else {
		out.WriteString("*ast.Return (Return Value: Non-Empty)")
		fmt.Println(out.String())
		ast.Exprs.PrintAST(level + 1)
	}
}

type Invocation struct {
	Statement
	Id   string
	Args []Expression
}

func (inv *Invocation) TokenLiteral() string { return inv.String() }
func (inv *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.Id)
	out.WriteString("(")
	if len(inv.Args) > 0 {
		out.WriteString(inv.Args[0].String())
	}
	if len(inv.Args) > 1 {
		for _, arg := range inv.Args[1:] {
			out.WriteString(", ")
			out.WriteString(arg.String())
		}
	}

	out.WriteString(")")
	out.WriteString(";")

	return out.String()
}
func (ast *Invocation) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Invocation (Function Name: ")
	out.WriteString(ast.Id)
	out.WriteString(")")
	fmt.Println(out.String())
	for _, arg := range ast.Args {
		arg.PrintAST(level + 1)
	}
}

type Expression interface {
	Node
	GetType(symTable *st.SymbolTable) string
	TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction
}

/*
type Arguments struct {
	Expressions []Expression
}

func (arg *Arguments) TokenLiteral() string { return arg.String() }
func (arg *Arguments) String() string {
	var out bytes.Buffer
	if len(arg.Expressions) > 0 {
		out.WriteString(arg.Expressions[0].String())
		if len(arg.Expressions) > 1 {
			for _, ar := range arg.Expressions[1:] {
				out.WriteString(",")
				out.WriteString(ar.String())
			}
		}

	}
	return out.String()
}
func (ast *Arguments) PrintAST(level int) {
	var out bytes.Buffer
	if len(ast.Expressions) > 0 {
		out.WriteString(strings.Repeat("   ", level))
		out.WriteString("*ast.Arguments (Number of Arguments: ")
		out.WriteString(strconv.Itoa(len(ast.Expressions)))
		out.WriteString(")")
		fmt.Println(out.String())
		for _, expression := range ast.Expressions {
			expression.PrintAST(level + 1)
		}
	}
}
func (ast *Arguments) TypeCheck(errors []string, st *st.SymbolTable) []string {
	// TODO
	return make([]string, 0)
}
func (ast *Arguments) GetType(st *st.SymbolTable) string {
	// TODO
	return "implement"
}

type Expression interface {
	Node
	GetType(symTable *st.SymbolTable) string
	TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction
}
func (ast *Arguments) GetType(symTable *st.SymbolTable) string {
	var out bytes.Buffer
	out.WriteString(ast.Expressions[0].GetType(symTable))
	for _, arg := range ast.Expressions[1:] {
		out.WriteString(",")
		out.WriteString(arg.GetType(symTable))
	}
	return out.String()
}*/

type AndOrExpr struct {
	Token token.Token
	Left  Expression
	Right Expression
	// Left and Right must be boolean, output boolean
}

func (ao *AndOrExpr) TokenLiteral() string { return ao.Token.Literal }
func (ao *AndOrExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ao.Left.String())
	out.WriteString(ao.Token.Literal)
	out.WriteString(ao.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ast *AndOrExpr) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.AndOrExpression (Operator: ")
	out.WriteString(ast.TokenLiteral())
	out.WriteString(")")
	fmt.Println(out.String())
	ast.Left.PrintAST(level + 1)
	ast.Right.PrintAST(level + 1)
}

type CompareExpr struct {
	Token      token.Token
	Comparator Comparator
	Left       Expression
	Right      Expression
	// Output boolean
}

func (cmp *CompareExpr) TokenLiteral() string {
	return cmp.Token.Literal
}
func (cmp *CompareExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(cmp.Left.String())
	out.WriteString(cmp.Token.Literal)
	out.WriteString(cmp.Right.String())
	out.WriteString(")")

	return out.String()
}
func (ast *CompareExpr) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.CompareExpression (Operator: ")
	out.WriteString(ast.TokenLiteral())
	out.WriteString(")")
	fmt.Println(out.String())
	ast.Left.PrintAST(level + 1)
	ast.Right.PrintAST(level + 1)
}

type Unary struct {
	SelectID string
	Selector Expression
}

func (un *Unary) TokenLiteral() string {
	return un.String()
}
func (un *Unary) String() string {
	var out bytes.Buffer

	out.WriteString(un.SelectID)
	out.WriteString(un.Selector.String())

	return out.String()
}
func (ast *Unary) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Unary (Unary Operator: ")
	out.WriteString(ast.SelectID)
	out.WriteString(")")
	fmt.Println(out.String())
	ast.Selector.PrintAST(level + 1)
}

type Selector struct {
	Factor   Expression
	SelectID string
}

func (sl *Selector) TokenLiteral() string {
	return ""
}
func (sl *Selector) String() string {
	var out bytes.Buffer

	out.WriteString(sl.Factor.String())
	out.WriteString(sl.SelectID)

	return out.String()
}
func (ast *Selector) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Unary (Ids to Select: ")
	out.WriteString(ast.SelectID)
	out.WriteString(")")
	fmt.Println(out.String())
	ast.Factor.PrintAST(level + 1)
}

type Factor struct {
	Exprs Expression
}

func (fc *Factor) TokenLiteral() string {
	return ""
}
func (fc *Factor) String() string {
	var out bytes.Buffer

	out.WriteString(fc.Exprs.String())

	return out.String()
}
func (ast *Factor) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Factor")
	fmt.Println(out.String())
	ast.Exprs.PrintAST(level + 1)
}

type Number struct {
	Value int
}

func (num *Number) TokenLiteral() string {
	return ""
}
func (num *Number) String() string {
	var out bytes.Buffer
	out.WriteString(strconv.Itoa(num.Value))

	return out.String()
}
func (ast *Number) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Number (Integer Value: ")
	out.WriteString(strconv.Itoa(ast.Value))
	out.WriteString(")")
	fmt.Println(out.String())
}

type Boolean struct {
	Value bool
}

func (bl *Boolean) TokenLiteral() string { return bl.String() }
func (bl *Boolean) String() string {
	var out bytes.Buffer
	out.WriteString(strconv.FormatBool(bl.Value))

	return out.String()
}
func (ast *Boolean) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Boolean (Value: ")
	out.WriteString(strconv.FormatBool(ast.Value))
	out.WriteString(")")
	fmt.Println(out.String())
}

type Nil struct {
}

func (nl *Nil) TokenLiteral() string {
	return ""
}
func (nl *Nil) String() string {
	var out bytes.Buffer
	out.WriteString("nil")

	return out.String()
}
func (ast *Nil) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Nil")
	fmt.Println(out.String())
}

type ID struct {
	Id   string
	Call bool
	Args []Expression
}

func (id *ID) TokenLiteral() string {
	return id.Id
}
func (id *ID) String() string {
	var out bytes.Buffer
	out.WriteString(id.Id)
	if id.Call {
		out.WriteString("(")
		out.WriteString(id.Args[0].String())
		for _, id := range id.Args[1:] {
			out.WriteString(", ")
			out.WriteString(id.String())
		}
		out.WriteString(")")
	}

	return out.String()
}
func (ast *ID) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Id ")
	if ast.Call {
		out.WriteString("(Function Name: ")
	} else {
		out.WriteString("(Variable Name: ")
	}
	out.WriteString(ast.Id)
	out.WriteString(")")
	fmt.Println(out.String())
	if ast.Call {
		for _, arg := range ast.Args {
			arg.PrintAST(level + 1)
		}

	}
}

/****************************************** */
// All expression nodes implement this interface

// Now you will need to implement the methods defined in the Node and Expression interface for MultEXpr

// You would define these structs for all operators so on.
/****************************************** */

// 2nd way - Define a more generalized version of the binary operators and make a field for the type of operator
type Operator int

const (
	ADD Operator = iota
	MULT
	SUB
	DIV
)

func OpString(op Operator) string {
	switch op {
	case ADD:
		return "+"
	case MULT:
		return "*"
	case SUB:
		return "-"
	case DIV:
		return "/"
	}
	panic("Error: Could not determine operator")
}

type Comparator int

const (
	EQUAL Comparator = iota
	NEQ
	GREATER
	LESSER
	GEQ
	LEQ
)

func CompString(op Comparator) string {

	switch op {
	case EQUAL:
		return "=="
	case NEQ:
		return "!="
	case GREATER:
		return ">"
	case LESSER:
		return "<"
	case GEQ:
		return ">="
	case LEQ:
		return "<="
	}
	panic("Error: Could not determine operator")
}

/****** Expressions ***********/

type BinOpExpr struct {
	Token    token.Token //The token from the scanner
	Operator Operator    // The operator for the binary expression (+ or - or / ...)
	Right    Expression
	Left     Expression
}

func (binOp *BinOpExpr) TokenLiteral() string {
	return binOp.Token.Literal
}

func (binOp *BinOpExpr) String() string {

	// You could use th string concatenation operator (+) to join strings together
	// str := str1 + str2 + str3
	// However, each time the concatenation operator joins two strings together it allocates a new
	// intermediate string. The   bytes.Buffer and WriteString prevents the generation a new string like in + operator from two or more strings.
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(binOp.Left.String())
	out.WriteString(" " + OpString(binOp.Operator) + " ")
	out.WriteString(binOp.Right.String())
	out.WriteString(")")

	return out.String()
}
func (ast *BinOpExpr) PrintAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.BinOpExpression (Operator: ")
	out.WriteString(ast.TokenLiteral())
	out.WriteString(")")
	fmt.Println(out.String())
	ast.Left.PrintAST(level + 1)
	ast.Right.PrintAST(level + 1)
}
