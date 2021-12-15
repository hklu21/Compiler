package ast

import (
	st "proj/symbolTable"
	"strings"
)

func (ast *AndOrExpr) GetType(symTable *st.SymbolTable) string {
	left := ast.Left.GetType(symTable)
	right := ast.Right.GetType(symTable)
	if left != "bool" || right != "bool" {
		leftInvalid := strings.Contains(left, "invalid operation")
		rightInvalid := strings.Contains(right, "invalid operation")
		if leftInvalid && rightInvalid {
			return left + "," + right
		} else if leftInvalid {
			return left
		} else if rightInvalid {
			return right
		}
		return "invalid operation: " + ast.String() + "(mismatched types " + left + " and " + right + ")"
	}
	return "bool"
	/*if ast.Left.GetType(symTable) != ast.Right.GetType(symTable) {
		return "Mismatched Type between " + ast.Left.String() + " and " + ast.Right.String()
	}
	return ast.Left.GetType(symTable)
	*/
}
func (ast *CompareExpr) GetType(symTable *st.SymbolTable) string {
	left := ast.Left.GetType(symTable)
	right := ast.Right.GetType(symTable)
	if left != right {
		if right == "nil" {
			if st.IsInSymTable(symTable, "struct"+left) {
				return "bool"
			} else {
				return "invalid operation: " + ast.String() + "(mismatched types " + left + " and " + right + ")"
			}
		}
		leftInvalid := strings.Contains(left, "invalid operation")
		rightInvalid := strings.Contains(right, "invalid operation")
		if leftInvalid && rightInvalid {
			return left + "," + right
		} else if leftInvalid {
			return left
		} else if rightInvalid {
			return right
		}
		return "invalid operation: " + ast.String() + "(mismatched types " + left + " and " + right + ")"
	}
	return "bool"
	/*if ast.Left.GetType(symTable) != ast.Right.GetType(symTable) {
		return "Mismatched Type between " + ast.Left.String() + " and " + ast.Right.String()
	}
	return ast.Left.GetType(symTable)
	*/
}

func (ast *BinOpExpr) GetType(symTable *st.SymbolTable) string {
	left := ast.Left.GetType(symTable)
	right := ast.Right.GetType(symTable)
	if left != "int" || right != "int" {
		leftInvalid := strings.Contains(left, "invalid operation")
		rightInvalid := strings.Contains(right, "invalid operation")
		if leftInvalid && rightInvalid {
			return left + "," + right
		} else if leftInvalid {
			return left
		} else if rightInvalid {
			return right
		}
		return "invalid operation: " + ast.String() + "(mismatched types " + left + " and " + right + ")"
	}
	return "int"
}

func (ast *Unary) GetType(symTable *st.SymbolTable) string {
	return ast.Selector.GetType(symTable)
}

func (ast *Selector) GetType(symTable *st.SymbolTable) string {
	if len(ast.SelectID) == 0 {
		return ast.Factor.GetType(symTable)
	}
	selectSlice := strings.Split(ast.SelectID[1:], ".")
	factorType := ast.Factor.GetType(symTable)
	sEntry := st.GetSymTableEntry(symTable, "struct"+factorType)

	structEntry, _ := sEntry.(st.StructEntry)
	entry, _ := structEntry.Variables[selectSlice[0]]
	varEntry, _ := entry.(st.VarEntry)
	if len(selectSlice) > 1 {
		// Recursively check function calls
		counter := 1
		for counter < len(selectSlice) {
			varType := varEntry.Type
			sEntry := st.GetSymTableEntry(symTable, "struct"+varType)
			structEntry, _ := sEntry.(st.StructEntry)
			entry, _ = structEntry.Variables[selectSlice[counter]]
			varEntry, _ = entry.(st.VarEntry)
			counter++
		}
	}
	return varEntry.Type
}
func (ast *Factor) GetType(symTable *st.SymbolTable) string {
	return ast.Exprs.GetType(symTable)
}
func (ast *Number) GetType(symTable *st.SymbolTable) string {
	return "int"
}
func (ast *Boolean) GetType(symTable *st.SymbolTable) string {
	return "bool"
}
func (ast *Nil) GetType(symTable *st.SymbolTable) string {
	return "nil"
}
func (ast *ID) GetType(symTable *st.SymbolTable) string {
	if !ast.Call {
		value := st.GetVarEntryFromST(symTable, ast.Id)
		varEntry := value.(st.VarEntry)
		return varEntry.Type
	} else {
		if ast.Id == "new" {
			return ast.Args[0].String()
		}
		return st.GetFuncRetFromST(symTable, ast.Id)
	}
}
