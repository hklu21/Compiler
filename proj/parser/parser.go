package parser

import (
	"fmt"
	"proj/ast"
	ct "proj/token"
	"strconv"
)

/****** Parser Language Definition EBNF **********
Program = Package Import Types Declarations Functions 'eof'                               ;
Package = 'package' 'id' ';'                                                              ;
Import = 'import' '"' 'fmt' '"'  ';'                                                      ;
Types = {TypeDeclaration}                                                                 ;
TypeDeclaration = 'type' 'id' 'struct' '{' Fields '}' ';'                                 ;
Fields = Decl ';' {Decl ';'}                                                              ;
Decl = 'id' Type                                                                          ;
Type = 'int' | 'bool' | '*' 'id'                                                          ;
Declarations = {Declaration}                                                              ;
Declaration = 'var' Ids Type ';'                                                          ;
Ids = 'id' {',' 'id'}                                                                     ;
Functions = {Function}                                                                    ;
Function = 'func' 'id' Parameters ReturnType '{' Declarations Statements '}'              ;
Parameters = '(' [ Decl {',' Decl}] ')'                                                   ;
ReturnType = type | 'ε'                                                                   ;
Statements = {Statement}                                                                  ;
Statement = Block | Assignment | Print | Conditional | Loop | Return | Read | Invocation  ;
Block = '{' Statements '}'                                                                ;
Assignment = LValue '=' Expression ';'                                                    ;
Read = 'fmt' '.' 'Scan' '(' '&' 'id' ')' ';'                                              ;
Print = 'fmt' '.' 'Print' '(' 'id' ')' ';'                                                ;
Print = 'fmt' '.' 'Println' '(' 'id' ')' ';'                                              ;
Conditional = 'if' '(' Expression ')' Block ['else' Block]                                ;
Loop = 'for' '(' Expression ')' Block                                                     ;
Return = 'return' [Expression] ';'                                                        ;
Invocation = 'id' Arguments ';'                                                           ;
Arguments = '(' [Expression {',' Expression}] ')'                                         ;
LValue = 'id' {'.' id}                                                                    ;
Expression = BoolTerm {'||' BoolTerm}                                                     ;
BoolTerm = EqualTerm {'&&' EqualTerm}                                                     ;
EqualTerm =  RelationTerm {('=='| '!=') RelationTerm}                                     ;
RelationTerm = SimpleTerm {('>'| '<' | '<=' | '>=') SimpleTerm}                           ;
SimpleTerm = Term {('+'| '-') Term}                                                       ;
Term = UnaryTerm {('*'| '/') UnaryTerm}                                                   ;
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm                            ;
SelectorTerm = Factor {'.' 'id'}                                                          ;
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'true' | 'false' | 'nil'      ;
*/

type Parser struct {
	tokens    []ct.Token /* A slice of tokens inputed */
	currIndex int        /* The current index read by the Parser */
}

func New(tokens []ct.Token) *Parser {
	return &Parser{tokens, -1}
}

func (p *Parser) currToken() ct.Token {
	// Returns the current token using the saved currIndex of p
	return p.tokens[p.currIndex]
}

func (p *Parser) Parse() *ast.Program {
	/* Include the edge case of an empty file without EOF */
	if len(p.tokens) == 0 {
		return nil
	}
	p.nextToken()
	return Program(p)
}

func (p *Parser) nextToken() ct.Token {

	var token ct.Token
	/*if p.currIndex == len(p.tokens)-1 {
		token = p.tokens[p.currIndex]
	}*/
	if p.currIndex == len(p.tokens) {
		token = p.tokens[p.currIndex-1]
	} else {
		p.currIndex += 1
		token = p.tokens[p.currIndex]
	}
	return token
}

func parseError(msg string) {
	fmt.Printf("syntax error:%s\n", msg)
}

func (p *Parser) match(token ct.TokenType) (ct.Token, bool) {
	// Match the tokenType of the current index
	// returns true if it matches and false otherwise
	if token == p.currToken().Type {
		token := p.currToken()
		p.nextToken()
		return token, true
	}
	return ct.Token{ct.ILLEGAL, ""}, false
}

/* Below is the Recursive Decent Parser for the goLite Language
 * Created Using the Right Recursive EBNF Grammar for the goLite Language
 * See top of file for the complete grammar
 */

func Program(p *Parser) *ast.Program {
	pack := Package(p)
	if pack == nil {
		return nil
	}
	impt := Import(p)
	if impt == false {
		return nil
	}
	typeList := make([]ast.Types, 0)
	typeList = Types(typeList, p)
	if typeList == nil {
		return nil
	}

	declList := make([]ast.Declaration, 0)
	declList = Declaration(declList, p)
	if declList == nil {
		return nil
	}

	funcList := make([]ast.Function, 0)
	funcList = Functions(funcList, p)
	if funcList == nil {
		return nil
	}

	//expr := Expression(p)

	return &ast.Program{Pack: *pack, Types: typeList, Decl: declList, Funcs: funcList}
	//types := Types(p)
	//decl := Declaration(p)
	//funcs := Functions(p)
	//return ast.Program{*pack, *impt, types, declaration, functions}
}

func Package(p *Parser) *ast.Package {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements
	if prevID, match := p.match(ct.PACKAGE); match {
		if id, match := p.match(ct.ID); match {
			if prevID, match := p.match(ct.SEMICOLON); match {
				return &ast.Package{Id: id.Literal}
			} else {
				parseError("expected \"package\", got" + prevID.Literal)
				return nil
			}
		} else {
			parseError("expected identifier, got" + string(prevID.Type))
			return nil
		}
	} else {
		parseError("expected \"package\", got" + prevID.Literal)
		return nil
	}
}

func Import(p *Parser) bool {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements

	if prevID, match := p.match(ct.IMPORT); match {
	} else {
		parseError("expected \"import\", got" + prevID.Literal)
		return false
	}
	if prevID, match := p.match(ct.QUOTATION); match {
	} else {
		parseError("expected \", got" + prevID.Literal)
		return false
	}
	if id, match := p.match(ct.FMT); match {
	} else {
		parseError("expected an identifer, got" + string(id.Type))
		return false
	}
	if prevID, match := p.match(ct.QUOTATION); match {
	} else {
		parseError("expected \", got" + prevID.Literal)
		return false
	}
	if prevID, match := p.match(ct.SEMICOLON); match {
	} else {
		parseError("expected \"package\", got" + prevID.Literal)
		return false
	}
	return true
}

func Types(lst []ast.Types, p *Parser) []ast.Types {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements
	if prevID, match := p.match(ct.TYPE); match {
		if id, match := p.match(ct.ID); match {
			if prevID, match := p.match(ct.STRUCT); match {
				if prevID, match := p.match(ct.LEFTCURL); match {
					var fieldLst []ast.Field
					fields := Fields(fieldLst, p)
					if len(fields) == 0 {
						parseError("expected at least 1 declaration within struct")
						return nil
					}
					if prevID, match := p.match(ct.RIGHTCURL); match {
						if prevID, match := p.match(ct.SEMICOLON); match {
							return Types(append(lst, ast.Types{Id: id.Literal, Fields: fields}), p)
						} else {
							parseError("expected semicolon, got" + prevID.Literal)
							return nil
						}
					} else {
						parseError("expected Right Brace, got " + prevID.Literal)
					}
				} else {
					parseError("expected Left Brace, got " + prevID.Literal)
				}
			} else {
				parseError("expected \"struct\", got" + prevID.Literal)
				return nil
			}
		} else {
			parseError("expected identifier, got" + string(prevID.Type))
			return nil
		}
	}
	return lst
}

func Fields(lst []ast.Field, p *Parser) []ast.Field {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements
	if id, match := p.match(ct.ID); match {
		if _, match := p.match(ct.INT); match {
			if prevID, match := p.match(ct.SEMICOLON); match {
				return Fields(append(lst, ast.Field{Id: id.Literal, Value: "int"}), p)
			} else {
				parseError("expected semicolon, got" + prevID.Literal)
				return nil
			}
		} else if _, match := p.match(ct.BOOL); match {
			if prevID, match := p.match(ct.SEMICOLON); match {
				return Fields(append(lst, ast.Field{Id: id.Literal, Value: "bool"}), p)
			} else {
				parseError("expected semicolon, got" + prevID.Literal)
				return nil
			}
		} else if _, match := p.match(ct.ASTRIX); match {
			if ptrID, match := p.match(ct.ID); match {
				if prevID, match := p.match(ct.SEMICOLON); match {
					return Fields(append(lst, ast.Field{Id: id.Literal, Value: ptrID.Literal}), p)
				} else {
					parseError("expected semicolon, got" + prevID.Literal)
					return nil
				}
			} else {
				parseError("expected proper identifier name, got " + id.Literal)
				return nil
			}
		} else {
			parseError("expected a type declaration, got " + id.Literal)
			return nil
		}
	}
	return lst
}

func Declaration(lst []ast.Declaration, p *Parser) []ast.Declaration {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements
	if prevID, match := p.match(ct.VAR); match {
		if id, match := p.match(ct.ID); match {
			idSlice := []string{id.Literal}
			idSlice = Ids(idSlice, p)
			if idSlice == nil {
				return nil
			}
			if _, match := p.match(ct.INT); match {
				if prevID, match := p.match(ct.SEMICOLON); match {
					return Declaration(append(lst, ast.Declaration{Id: idSlice, Value: "int"}), p)
				} else {
					parseError("expected semicolon, got" + prevID.Literal)
					return nil
				}
			} else if _, match := p.match(ct.BOOL); match {
				if prevID, match := p.match(ct.SEMICOLON); match {
					return Declaration(append(lst, ast.Declaration{Id: idSlice, Value: "bool"}), p)
				} else {
					parseError("expected semicolon, got" + prevID.Literal)
					return nil
				}
			} else if _, match := p.match(ct.ASTRIX); match {
				if ptrID, match := p.match(ct.ID); match {
					if prevID, match := p.match(ct.SEMICOLON); match {
						return Declaration(append(lst, ast.Declaration{Id: idSlice, Value: ptrID.Literal}), p)
					} else {
						parseError("expected semicolon, got" + prevID.Literal)
						return nil
					}
				} else {
					parseError("expected an indentifier, got " + ptrID.Literal)
					return nil
				}
			}
		} else {
			parseError("expected identifier, got" + string(prevID.Type))
			return nil
		}
	}
	return lst
}

func Ids(idSlice []string, p *Parser) []string {
	if _, match := p.match(ct.COMMA); match {
		if id, match := p.match(ct.ID); match {
			return Ids(append(idSlice, id.Literal), p)
		} else {
			parseError("expected proper identifier name, got " + id.Literal)
			return nil
		}
	}
	return idSlice
}

func Functions(funcSlice []ast.Function, p *Parser) []ast.Function {
	if _, match := p.match(ct.FUNC); match {
		if id, match := p.match(ct.ID); match {
			if prevID, match := p.match(ct.LEFTBRAC); match {
				// Implement Parameters functions
				var paramSlice []ast.Parameter
				param := Parameters(paramSlice, p)
				if prevID, match := p.match(ct.RIGHTBRAC); match {
					// Implement Return Type function
					var returnType string
					if _, match := p.match(ct.INT); match {
						returnType = "int"
					} else if _, match := p.match(ct.BOOL); match {
						returnType = "bool"
					} else if _, match := p.match(ct.ASTRIX); match {
						if ptrID, match := p.match(ct.ID); match {
							returnType = ptrID.Literal
						} else {
							parseError("expected identifier, got" + string(prevID.Type))
							return nil
						}
					} else {
						returnType = "Epsilon"
					}
					if prevID, match := p.match(ct.LEFTCURL); match {
						// Implement Declaration function
						declList := make([]ast.Declaration, 0)
						declList = Declaration(declList, p)
						if declList == nil {
							return nil
						}
						// TTODO: Implement Statement functions
						stateList := make([]ast.Statement, 0)
						stateList = Statements(stateList, p)
						if stateList == nil {
							return nil
						}
						if prevID, match := p.match(ct.RIGHTCURL); match {
							// TODO: Create Return Type
							return Functions(append(funcSlice, ast.Function{Id: id.Literal, Param: param, ReturnT: returnType, Dec: declList, States: stateList}), p)
						} else {
							parseError("expected right brace, got " + string(prevID.Type))
							return nil
						}
					} else {
						parseError("expected left brace, got " + string(prevID.Type))
						return nil
					}
				} else {
					parseError("expected right parenthesis, got " + string(prevID.Type))
					return nil
				}
			} else {
				parseError("expected left bracket, got " + string(prevID.Type))
				return nil
			}
		} else {
			parseError("expected identifier, got " + string(id.Type))
			return nil
		}
	}
	return funcSlice
}

func Parameters(lst []ast.Parameter, p *Parser) []ast.Parameter {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements
	if id, match := p.match(ct.ID); match {
		if _, match := p.match(ct.INT); match {
			return ParamPrime(append(lst, ast.Parameter{Id: id.Literal, Value: "int"}), p)
		} else if _, match := p.match(ct.BOOL); match {
			return ParamPrime(append(lst, ast.Parameter{Id: id.Literal, Value: "bool"}), p)
		} else if _, match := p.match(ct.ASTRIX); match {
			if ptrID, match := p.match(ct.ID); match {
				return ParamPrime(append(lst, ast.Parameter{Id: id.Literal, Value: ptrID.Literal}), p)
			}
		} else {
			parseError("expected a type declaration, got " + id.Literal)
			return nil
		}
	}
	return lst
}

func ParamPrime(lst []ast.Parameter, p *Parser) []ast.Parameter {
	// Checks if the program is at the end of the file
	// If not, test for more valid statements
	if prevID, match := p.match(ct.COMMA); match {
		if id, match := p.match(ct.ID); match {
			if _, match := p.match(ct.INT); match {
				return ParamPrime(append(lst, ast.Parameter{Id: id.Literal, Value: "int"}), p)
			} else if _, match := p.match(ct.BOOL); match {
				return ParamPrime(append(lst, ast.Parameter{Id: id.Literal, Value: "bool"}), p)
			} else if _, match := p.match(ct.ASTRIX); match {
				if ptrID, match := p.match(ct.ID); match {
					return ParamPrime(append(lst, ast.Parameter{Id: id.Literal, Value: ptrID.Literal}), p)
				}
			} else {
				parseError("expected a type declaration, got " + id.Literal)
				return nil
			}
		} else {
			parseError("expected an identifier, got " + prevID.Literal)
			return nil
		}
	}
	return lst
}

func Statements(lst []ast.Statement, p *Parser) []ast.Statement {
	//fmt.Println(p.currToken().Literal)
	switch p.currToken().Type {
	case ct.LEFTCURL:
		return Statements(append(lst, Block(p)), p)
	// TODO Test this RightCurl
	case ct.RIGHTCURL:
		return lst
	case ct.ID:
		return Statements(append(lst, ID(p)), p)
	case ct.FOR:
		return Statements(append(lst, For(p)), p)
	case ct.IF:
		return Statements(append(lst, If(p)), p)
	case ct.RETURN:
		return Statements(append(lst, Return(p)), p)
	case ct.FMT:
		return Statements(append(lst, FMT(p)), p)
	case ct.EOF:
		return nil
	default:
		return lst
	}
}

func ID(p *Parser) ast.Statement {
	if saveID, match := p.match(ct.ID); match {
		switch p.currToken().Type {
		case ct.LEFTBRAC:
			return Invocation(saveID.Literal, p)
		case ct.DOT:
			return Assignment(saveID.Literal, p)
		case ct.ASSIGN:
			return Assignment(saveID.Literal, p)
		default:
			parseError("expected valid ID, got " + saveID.Literal)
			return nil
		}
	}
	return nil
}

func Assignment(saveID string, p *Parser) ast.Statement {
	lvalue := []string{saveID}
	if p.currToken().Type == ct.DOT {
		lvalue = LValue(lvalue, p)
	}
	if prevID, match := p.match(ct.ASSIGN); match {
		expr := Expression(p)
		if prevID, match := p.match(ct.SEMICOLON); match {
			return &ast.Assignment{Lvalue: lvalue, Exprs: expr}
		} else {
			parseError("expected semicolon, got " + prevID.Literal)
		}
	} else {
		parseError("expected =, got " + prevID.Literal)
		return nil
	}
	parseError("unexpected error in Assignment()")
	return nil
}

func LValue(lst []string, p *Parser) []string {
	if p.currToken().Type == ct.DOT {
		p.nextToken()
		if saveID, match := p.match(ct.ID); match {
			return LValue(append(lst, saveID.Literal), p)
		} else {
			parseError("expected valid ID after ., got " + saveID.Literal)
			return nil
		}
	}
	return lst
}

func Invocation(saveID string, p *Parser) ast.Statement {
	if _, match := p.match(ct.LEFTBRAC); match {
		if p.currToken().Type == ct.RIGHTBRAC {
			p.nextToken()
			return &ast.Invocation{Id: saveID, Args: make([]ast.Expression, 0)}
		} else {
			argList := make([]ast.Expression, 0)
			arguments := Arguments(argList, p)
			if prevID, match := p.match(ct.RIGHTBRAC); match {
				if prevID, match := p.match(ct.SEMICOLON); match {
					return &ast.Invocation{Id: saveID, Args: arguments}
				} else {
					parseError("expected semicolon, got " + prevID.Literal)
					return nil
				}
			} else {
				parseError("expected right parenthesis, got " + prevID.Literal)
				return nil
			}
		}
	}
	return nil
}

func Return(p *Parser) ast.Statement {
	if _, match := p.match(ct.RETURN); match {
		if p.currToken().Type == ct.SEMICOLON {
			p.nextToken()
			return &ast.Return{Exprs: nil}
		} else {
			expression := Expression(p)
			if expression == nil {
				parseError("expected a valid expression")
				return nil
			} else if prevID, match := p.match(ct.SEMICOLON); match {
				return &ast.Return{Exprs: expression}
			} else {
				parseError("expected a semicolon, got " + prevID.Literal)
				return nil
			}
		}
	}
	return nil
}

func For(p *Parser) ast.Statement {
	if _, match := p.match(ct.FOR); match {
		if prevID, match := p.match(ct.LEFTBRAC); match {
			expression := Expression(p)
			if expression == nil {
				parseError("expected a valid expression")
				return nil
			}
			if prevID, match := p.match(ct.RIGHTBRAC); match {
				forBlock := Block(p)
				if forBlock == nil {
					parseError("expected a valid block")
					return nil
				}

				return &ast.Loop{Expr: expression, ForBlock: forBlock}
			} else {
				parseError("expected right parenthesis, got " + prevID.Literal)
				return nil
			}
		} else {
			parseError("expected left parenthesis, got " + prevID.Literal)
			return nil
		}
	}
	return nil
}

func If(p *Parser) ast.Statement {
	if _, match := p.match(ct.IF); match {
		if prevID, match := p.match(ct.LEFTBRAC); match {
			expression := Expression(p)
			if expression == nil {
				parseError("expected a valid expression")
				return nil
			}
			if prevID, match := p.match(ct.RIGHTBRAC); match {
				ifBlock := Block(p)
				if ifBlock == nil {
					parseError("expected a valid block")
					return nil
				}
				//fmt.Println("Else block to be read, next token is " + p.currToken().Literal)
				if _, match := p.match(ct.ELSE); match {
					elseBlock := Block(p)
					if elseBlock == nil {
						parseError("expected a valid block")
						return nil
					}
					return &ast.Conditional{Expr: expression, IfBlock: ifBlock, HasElse: true, ElseBlock: elseBlock}
				} else {
					return &ast.Conditional{Expr: expression, IfBlock: ifBlock, HasElse: false}
				}
			} else {
				parseError("expected right parenthesis, got " + prevID.Literal)
				return nil
			}
		} else {
			parseError("expected left parenthesis, got " + prevID.Literal)
			return nil
		}
	}
	return nil
}

// TODO Update Block to fit with later implementation
// Remove right curl from Statements
func Block(p *Parser) ast.Statement {
	if prevID, match := p.match(ct.LEFTCURL); match {
		if p.currToken().Type == ct.RIGHTCURL {
			p.nextToken()
			return &ast.Block{Block: nil}
		}
		state := make([]ast.Statement, 0)
		state = Statements(state, p)
		if prevID, match := p.match(ct.RIGHTCURL); match {
			return &ast.Block{Block: state}
		} else {
			parseError("expected right brace, got " + prevID.Literal)
			return nil
		}
	} else {
		parseError("expected left brace, got " + prevID.Literal)
		return nil
	}
}

func FMT(p *Parser) ast.Statement {
	if _, match := p.match(ct.FMT); match {
		if prevID, match := p.match(ct.DOT); match {
			var funcName string
			funcName = p.currToken().Literal
			if funcName != "Scan" && funcName != "Print" && funcName != "Println" {
				parseError("expected a valid fmt function, got " + funcName)
				return nil
			}
			p.nextToken()
			if prevID, match := p.match(ct.LEFTBRAC); match {
				if funcName == "Scan" {
					if _, match := p.match(ct.AMP); match {
					} else {
						parseError("expected an ampersand, got " + funcName)
						return nil
					}
				}
				if id, match := p.match(ct.ID); match {
					if prevID, match := p.match(ct.RIGHTBRAC); match {
						if prevID, match := p.match(ct.SEMICOLON); match {
							return &ast.ReadPrint{FuncName: funcName, Id: id.Literal}
						} else {
							parseError("expected a semicolon, got " + prevID.Literal)
							return nil
						}
					} else {
						parseError("expected a right parenthesis, got " + prevID.Literal)
						return nil
					}
				} else {
					parseError("expected an identifier, got " + prevID.Literal)
					return nil
				}
			} else {
				parseError("expected a left parenthesis, got " + prevID.Literal)
				return nil
			}
		} else {
			parseError("expected a dot, got " + prevID.Literal)
			return nil
		}
	}
	return nil
}

func Expression(p *Parser) ast.Expression {
	left := BoolTerm(p)
	if left == nil {
		return nil
	}
	expr := BoolPrime(left, p)
	return expr
}

func BoolTerm(p *Parser) ast.Expression {
	left := EqualTerm(p)
	if left == nil {
		return nil
	}
	expr := BoolPrime(left, p)
	return expr
}

func BoolPrime(left ast.Expression, p *Parser) ast.Expression {
	if prevToken, match := p.match(ct.OR); match {
		right := EqualTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.AndOrExpr{Token: prevToken, Left: left, Right: right}
		return BoolPrime(expr, p)
	}
	return left
}

func EqualTerm(p *Parser) ast.Expression {
	left := RelationTerm(p)
	if left == nil {
		return nil
	}
	expr := EqualPrime(left, p)
	return expr
}

func EqualPrime(left ast.Expression, p *Parser) ast.Expression {
	if prevToken, match := p.match(ct.AND); match {
		right := RelationTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.AndOrExpr{Token: prevToken, Left: left, Right: right}
		return EqualPrime(expr, p)
	}
	return left
}

func RelationTerm(p *Parser) ast.Expression {
	left := SimpleTerm(p)
	if left == nil {
		return nil
	}
	expr := RelationPrime(left, p)
	return expr
}

func RelationPrime(left ast.Expression, p *Parser) ast.Expression {
	prevToken := p.currToken()
	switch prevToken.Type {
	case ct.EQUAL:
		p.nextToken()
		right := SimpleTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.CompareExpr{Token: prevToken, Comparator: ast.EQUAL, Left: left, Right: right}
		return RelationPrime(expr, p)
	case ct.NOTEQUAL:
		p.nextToken()
		right := SimpleTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.CompareExpr{Token: prevToken, Comparator: ast.NEQ, Left: left, Right: right}
		return RelationPrime(expr, p)
	default:
		return left
	}
}

func SimpleTerm(p *Parser) ast.Expression {
	left := Term(p)
	if left == nil {
		return nil
	}
	expr := SimplePrime(left, p)
	return expr
}

func SimplePrime(left ast.Expression, p *Parser) ast.Expression {
	prevToken := p.currToken()
	switch prevToken.Type {
	case ct.LESSTHAN:
		p.nextToken()
		right := Term(p)
		if right == nil {
			return nil
		}
		expr := &ast.CompareExpr{Token: prevToken, Comparator: ast.LESSER, Left: left, Right: right}
		return SimplePrime(expr, p)
	case ct.MORETHAN:
		p.nextToken()
		right := Term(p)
		if right == nil {
			return nil
		}
		expr := &ast.CompareExpr{Token: prevToken, Comparator: ast.GREATER, Left: left, Right: right}
		return SimplePrime(expr, p)
	case ct.LEQ:
		p.nextToken()
		right := Term(p)
		if right == nil {
			return nil
		}
		expr := &ast.CompareExpr{Token: prevToken, Comparator: ast.LEQ, Left: left, Right: right}
		return SimplePrime(expr, p)
	case ct.GEQ:
		p.nextToken()
		right := Term(p)
		if right == nil {
			return nil
		}
		expr := &ast.CompareExpr{Token: prevToken, Comparator: ast.GEQ, Left: left, Right: right}
		return SimplePrime(expr, p)
	default:
		return left
	}
}

func Term(p *Parser) ast.Expression {
	left := UnaryTerm(p)
	if left == nil {
		return nil
	}
	expr := TermPrime(left, p)
	return expr
}

func TermPrime(left ast.Expression, p *Parser) ast.Expression {
	prevToken := p.currToken()
	switch prevToken.Type {
	case ct.PLUS:
		p.nextToken()
		right := UnaryTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.BinOpExpr{Token: prevToken, Operator: ast.ADD, Left: left, Right: right}
		return TermPrime(expr, p)
	case ct.MINUS:
		p.nextToken()
		right := UnaryTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.BinOpExpr{Token: prevToken, Operator: ast.SUB, Left: left, Right: right}
		return TermPrime(expr, p)
	default:
		return left
	}
}

func UnaryTerm(p *Parser) ast.Expression {
	left := SelectorTerm(p)
	if left == nil {
		return nil
	}
	expr := UnaryPrime(left, p)
	return expr
}

func UnaryPrime(left ast.Expression, p *Parser) ast.Expression {
	prevToken := p.currToken()
	switch prevToken.Type {
	case ct.ASTRIX:
		p.nextToken()
		right := SelectorTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.BinOpExpr{Token: prevToken, Operator: ast.MULT, Left: left, Right: right}
		return UnaryPrime(expr, p)
	case ct.DIVIDE:
		p.nextToken()
		right := SelectorTerm(p)
		if right == nil {
			return nil
		}
		expr := &ast.BinOpExpr{Token: prevToken, Operator: ast.DIV, Left: left, Right: right}
		return UnaryPrime(expr, p)
	default:
		return left
	}
}

func SelectorTerm(p *Parser) ast.Expression {
	if p.currToken().Type == ct.NOT || p.currToken().Type == ct.MINUS {
		prevToken := p.currToken()
		p.nextToken()
		selector := SelectorPrime(p)
		if selector == nil {
			return nil
		}
		return &ast.Unary{SelectID: prevToken.Literal, Selector: selector}
	} else {
		selector := SelectorPrime(p)
		if selector == nil {
			return nil
		}
		return selector
	}
}

func SelectorPrime(p *Parser) ast.Expression {
	left := Factor(p)
	SelectID := ""
	if left == nil {
		return nil
	}

	for p.currToken().Type == ct.DOT {
		//fmt.Println(p.currToken().Literal)
		p.nextToken()
		if id, match := p.match(ct.ID); match {
			SelectID = SelectID + "." + id.Literal
		} else {
			parseError("expected proper identifier name, got " + id.Literal)
			return nil
		}
	}
	if SelectID != "" {
		return &ast.Selector{Factor: left, SelectID: SelectID}
	}
	return left
}

func Factor(p *Parser) ast.Expression {
	switch p.currToken().Type {
	case ct.LEFTBRAC:
		if _, match := p.match(ct.LEFTBRAC); match {
			expr := Expression(p)
			if expr == nil {
				return nil
			}
			if prevID, match := p.match(ct.RIGHTBRAC); match {
				return &ast.Factor{Exprs: expr}
			} else {
				parseError("expected a right parenthesis, got " + prevID.Literal)
				return nil
			}
		}
	case ct.ID:
		storeID := p.currToken()
		p.nextToken()
		if p.currToken().Type == ct.LEFTBRAC {
			p.nextToken()
			if p.currToken().Type == ct.RIGHTBRAC {
				p.nextToken()
				return &ast.ID{Id: storeID.Literal, Call: true, Args: make([]ast.Expression, 0)}
			} else {
				argList := make([]ast.Expression, 0)
				arguments := Arguments(argList, p)
				if prevID, match := p.match(ct.RIGHTBRAC); match {
					return &ast.ID{Id: storeID.Literal, Call: true, Args: arguments}
				} else {
					parseError("expected a right parenthesis, got " + prevID.Literal)
					return nil
				}
			}
		} else {
			return &ast.ID{Id: storeID.Literal, Call: false}
		}
	case ct.NUMBER:
		if prevID, match := p.match(ct.NUMBER); match {
			i, _ := strconv.Atoi(prevID.Literal)
			return &ast.Number{Value: i}
		}
	case ct.TRUE:
		if prevID, match := p.match(ct.TRUE); match {
			i, _ := strconv.ParseBool(prevID.Literal)
			return &ast.Boolean{Value: i}
		}
	case ct.FALSE:
		if prevID, match := p.match(ct.FALSE); match {
			i, _ := strconv.ParseBool(prevID.Literal)
			return &ast.Boolean{Value: i}
		}
	case ct.NIL:
		if _, match := p.match(ct.NIL); match {
			return &ast.Nil{}
		}
	default:
		parseError("expected a Factor Term")
		return nil
	}
	return nil
}

func Arguments(lst []ast.Expression, p *Parser) []ast.Expression {
	expr := Expression(p)
	right := ArgumentsPrime(append(lst, expr), p)
	return right
}

func ArgumentsPrime(lst []ast.Expression, p *Parser) []ast.Expression {
	if _, match := p.match(ct.COMMA); match {
		expr := Expression(p)
		if expr == nil {
			parseError("expected an expression (in ArgumentsPrime)")
			return nil
		}
		return ArgumentsPrime(append(lst, expr), p)
	}
	return lst
}
