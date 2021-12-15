package ast

import (
	"fmt"
	st "proj/symbolTable"
	"strings"
)

func (ast *Program) TypeCheck(errors []string, st *st.SymbolTable) []string {
	for _, funcs := range ast.Funcs {
		errors = funcs.TypeCheck(errors, st)
	}
	return errors
}

func (ast *Package) TypeCheck(errors []string, st *st.SymbolTable) []string {
	return errors
}

func (ast *Types) TypeCheck(errors []string, st *st.SymbolTable) []string {
	return errors
}

func (ast *Field) TypeCheck(errors []string, st *st.SymbolTable) []string {
	return errors
}

func (ast *Declaration) TypeCheck(errors []string, st *st.SymbolTable) []string {
	return errors
}

func (ast *Function) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	funcEntry, _ := symTable.St["func"+ast.Id].(st.FunctionEntry)
	/* TODO: Determine if parameters within Function need to be gone through again
	for _, param := range ast.Param {
		errors = param.TypeCheck(errors, funcEntry.SymTable)
	}*/
	/* TODO: Determine if Variable Declaration within Function needs to be gone through again
	for _, dec := range ast.Dec {
		errors = dec.TypeCheck(errors, funcEntry.SymTable)
	}*/
	for _, states := range ast.States {
		errors = states.TypeCheck(errors, funcEntry.SymTable)
	}
	if ast.ReturnT != "Epsilon" && ast.Id != "main" {
		if !ast.States[len(ast.States)-1].CheckReturn(funcEntry.SymTable) {
			errors = append(errors, "missing return at end of function")
		}
	}
	return errors
}

func (ast *Parameter) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	/* Parameters were already checked during SABuild
	if ast.Value != "bool" && ast.Value != "int" {
		if !st.IsInSymTable(symTable, "struct"+ast.Value) {
			errors = append(errors, "function have undefined struct type "+ast.Value)
		}
	}*/
	return errors
}

func (ast *Block) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	for _, block := range ast.Block {
		errors = block.TypeCheck(errors, symTable)
	}
	return errors
}

func (ast *Assignment) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	// If Leftvalue is a single Id
	entry := st.GetSymTableEntry(symTable, "var"+ast.Lvalue[0])
	if entry == nil {
		errors = append(errors, "Variable "+ast.Lvalue[0]+" is not declared")
	}
	varEntry, _ := entry.(st.VarEntry)
	varType := varEntry.Type
	// If Leftvalue is yet to be initialized
	right := ast.Exprs.GetType(symTable)
	if !varEntry.Usable {
		if st.IsInSymTable(symTable, "struct"+right) {
			if varType != right {
				if strings.Contains(right, "invalid operation") {
					errors = append(errors, strings.Split(right, ",")...)
					return errors
				} else {
					errors = append(errors, "invalid assignment: "+ast.String()+" (mismatched types "+varType+" and "+right+")")
					return errors
				}
			} else {
				st.UpdateUsable(symTable, ast.Lvalue[0], true)
			}
		}
		// Type Check for new is defined in ID.GetType()
		return errors
	}

	// Leftvalue is a value in some struct pointer
	if len(ast.Lvalue) > 1 {
		// Recursively check function calls
		counter := 1
		var found bool
		for counter < len(ast.Lvalue) {
			varType := varEntry.Type
			sEntry := st.GetSymTableEntry(symTable, "struct"+varType)
			structEntry, _ := sEntry.(st.StructEntry)
			entry, found = structEntry.Variables[ast.Lvalue[counter]]
			if !found {
				errors = append(errors, "Variable "+ast.Lvalue[counter]+" is not found within struct "+varEntry.Type)
			}
			varEntry, _ = entry.(st.VarEntry)
			counter++
		}
	}
	errors = ast.Exprs.TypeCheck(errors, symTable)
	varType = varEntry.Type
	// Check right hand side

	if varType != right {
		if strings.Contains(right, "invalid operation") {
			errors = append(errors, strings.Split(right, ",")...)
		} else {
			if !(varType != "bool" && varType != "int" && right == "nil") {
				errors = append(errors, "invalid assignment: "+ast.String()+" (mismatched types "+varType+" and "+right+")")
			}

		}
	}
	/*
		value, found := symTable.St["var"+ast.Lvalue[0]]
		// If the variable has not been declared
		if !found {
			errors = append(errors, "Variable "+ast.Lvalue[0]+" is not declared")
		}
		varEntry, found := value.(st.VarEntry)
		if !varEntry.Usable {
			errors = append(errors, "Varaiable "+ast.Lvalue[0]+" is not initialized with new()")
		}
		//var finalType string
		for _, id := range ast.Lvalue[1:] {
			/*if structValue, found := symTable.Parent.St["struct"+varEntry.Type]; found {
				structEntry, _ := structValue.(st.StructEntry)
				if val, found := structEntry.Variables["var"+id]; found {
					finalType = val.(st.VarEntry).Type
				} else {
					errors = append(errors, "TODO Implement Invocations for "+id)
				}
			}*/
	return errors
}

func (ast *ReadPrint) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// DONE for now
	// Check if the variable is decalred
	if !st.IsInSymTable(symTable, "var"+ast.Id) {
		errors = append(errors, ast.FuncName+" function has undeclared variable "+ast.Id)
	} else {
		if ast.FuncName == "Println" || ast.FuncName == "Print" {
			if symTable.St["var"+ast.Id].(st.VarEntry).Type != "int" {
				errors = append(errors, ast.FuncName+" function uses non-int variable "+ast.Id)
			}
		}
		// TODO: Determine if Scan needs typecheck
	}
	return errors
}

func (ast *Conditional) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = ast.Expr.TypeCheck(errors, symTable)
	if ast.Expr.GetType(symTable) != "bool" {
		fmt.Println(ast.Expr.GetType(symTable))
		errors = append(errors, "If Expression does not evaluate to type bool")
	} else {
		errors = ast.IfBlock.TypeCheck(errors, symTable)
		if ast.HasElse {
			errors = ast.ElseBlock.TypeCheck(errors, symTable)
		}
	}
	return errors
}

func (ast *Loop) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = ast.Expr.TypeCheck(errors, symTable)
	if ast.Expr.GetType(symTable) != "bool" {
		errors = append(errors, "Expression in For Loop does not evaluate to type bool")
	} else {
		errors = ast.ForBlock.TypeCheck(errors, symTable)
	}
	return errors
}

func (ast *Return) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if ast.Exprs == nil {
		errors = append(errors, "cannot return nothing as type "+symTable.Return+" in return argument")
		return errors
	}
	retType := ast.Exprs.GetType(symTable)
	if retType != symTable.Return {
		if retType != "nil" && symTable.Return != "bool" && symTable.Return != "int" {
			errors = append(errors, "cannot use "+ast.Exprs.String()+" (type "+retType+") as type "+symTable.Return+" in return argument")
		}

	}
	return errors
}

func (ast *Invocation) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	entry := st.GetSymTableEntry(symTable, "func"+ast.Id)
	if entry == nil {
		errors = append(errors, "Invalid undefined function call to "+ast.Id)
		return errors
	}
	funcEntry, _ := entry.(st.FunctionEntry)
	if funcEntry.Params != len(ast.Args) {
		errors = append(errors, "Invalid function call to "+ast.Id+": unmatched number of arguments")
		return errors
	}
	// Include cases for new and delete
	if ast.Id == "delete" {
		argType := ast.Args[0].String()
		varEntry := st.GetVarEntryFromST(symTable, argType)
		if varEntry == nil {
			errors = append(errors, "invalid function call to delete: (variable "+argType+" not found)")
		} else {
			entry, _ := varEntry.(st.VarEntry)
			if entry.Usable {
				return errors
			} else {
				errors = append(errors, "invalid function call to delete: (variable "+argType+" not initialized)")
				return errors
			}
		}
	}
	// Checks Input Parameters
	for index, expr := range ast.Args {
		exprType := expr.GetType(symTable)
		funcType := funcEntry.Parameters[index].Type
		if exprType != funcType {
			errors = append(errors, "invalid function call to "+ast.Id+": (mismatched types"+exprType+" and "+funcType+")")
		}
	}
	// Check return types
	if funcEntry.Return == "int" || funcEntry.Return == "bool" {
		errors = append(errors, ast.String()+" evaluated but not used")
		return errors
	}
	return errors
	// If the new function is called
	/*
		exprs := ast.Args.(*Arguments)
		if ast.Id == "new" {
			if len(exprs.Expressions) != 1 {
				errors = append(errors, "new function call has too many or too few arguments")
				return errors
			} else {
				if exprs.Expressions[0].GetType(symTable) != "id" {
					errors = append(errors, "new function call can only take variable as argument")
				} else if {

				}
				return errors
			}
		}
		if ast.Id == "delete" {

		}
		if value, found := symTable.Parent.St["func"+ast.Id]; found {
			funcEntry, _ := value.(st.FunctionEntry)
			if funcEntry.Params != len(exprs.Expressions) {
				errors = append(errors, "Function declaration parameters does not match ")
			} else {
				for _, expr := range ast.Expressions {
					if expr.GetType(symTable) == st.IsInMySymTable() {

					}
				}
			}
		} else {
			errors = append(errors, "function "+ast.Id+" is not declaration")
		}*/
}

func (ast *AndOrExpr) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	err := ast.GetType(symTable)
	if err != "bool" {
		errors = append(errors, strings.Split(err, ",")...)
		return errors
	}
	errors = ast.Left.TypeCheck(errors, symTable)
	errors = ast.Right.TypeCheck(errors, symTable)
	return errors
}

func (ast *CompareExpr) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	err := ast.GetType(symTable)
	if err != "bool" && err != "int" {
		errors = append(errors, strings.Split(err, ",")...)
		return errors
	}
	errors = ast.Left.TypeCheck(errors, symTable)
	errors = ast.Right.TypeCheck(errors, symTable)
	return errors
}

func (ast *BinOpExpr) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	err := ast.GetType(symTable)
	if err != "int" {
		errors = append(errors, strings.Split(err, ",")...)
		return errors
	}
	errors = ast.Left.TypeCheck(errors, symTable)
	errors = ast.Right.TypeCheck(errors, symTable)
	return errors
}

func (ast *Unary) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if ast.SelectID == "!" {
		if ast.Selector.GetType(symTable) != "bool" {
			errors = append(errors, "Mismatched Type: Selector ! must have boolean expression")
		}
	} else if ast.SelectID == "-" {
		if ast.Selector.GetType(symTable) != "int" {
			errors = append(errors, "Mismatched Type: Selector - must have int expression")
		}
	}
	return ast.Selector.TypeCheck(errors, symTable)
}

func (ast *Selector) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	selectSlice := strings.Split(ast.SelectID[1:], ".")
	entry := st.GetSymTableEntry(symTable, "var"+ast.Factor.String())
	if entry == nil {
		errors = append(errors, "Variable "+ast.Factor.String()+" is not declared")
	} else if varEntry, _ := entry.(st.VarEntry); !varEntry.Usable {
		errors = append(errors, "Variable "+ast.Factor.String()+" is not initialized")
	}
	varEntry, _ := entry.(st.VarEntry)
	// Recursively check function calls
	counter := 0
	var found bool
	for counter < len(selectSlice) {
		varType := varEntry.Type
		sEntry := st.GetSymTableEntry(symTable, "struct"+varType)
		structEntry, _ := sEntry.(st.StructEntry)
		entry, found = structEntry.Variables[selectSlice[counter]]
		if !found {
			errors = append(errors, "Variable "+selectSlice[counter]+" is not found within struct "+varEntry.Type)
		}
		varEntry, _ = entry.(st.VarEntry)
		counter++
	}
	errors = ast.Factor.TypeCheck(errors, symTable)
	return errors
}

func (ast *Factor) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = ast.Exprs.TypeCheck(errors, symTable)
	return errors
}

func (ast *Number) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}

func (ast *Boolean) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}

func (ast *Nil) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}

func (ast *ID) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if !ast.Call {
		if !st.IsInSymTable(symTable, "var"+ast.Id) {
			errors = append(errors, "undefined variable "+ast.Id)
			return errors
		}
	} else {
		entry := st.GetSymTableEntry(symTable, "func"+ast.Id)
		if entry == nil {
			errors = append(errors, "Invalid function call to undefined function "+ast.Id)
			return errors
		}
		funcEntry, _ := entry.(st.FunctionEntry)
		if funcEntry.Params != len(ast.Args) {
			errors = append(errors, "Invalid function call to "+ast.Id+": unmatched number of arguments")
			return errors
		}
		// Check new function cases
		// delete will report error due to unmatched return type
		if ast.Id == "new" {
			isDefinedStruct := false
			for _, param := range funcEntry.Parameters {
				if st.IsInSymTable(symTable, "struct"+param.Type) {
					isDefinedStruct = true
				}
			}
			if !isDefinedStruct {
				errors = append(errors, "invalid function call to new: initializing undefined struct")
				return errors
			} else {
				return errors
			}
		}
		for index, expr := range ast.Args {
			exprType := expr.GetType(symTable)
			funcType := funcEntry.Parameters[index].Type
			if exprType != funcType {
				errors = append(errors, "invalid function call to "+ast.Id+": (mismatched types"+exprType+" and "+funcType+")")
			}
			errors = expr.TypeCheck(errors, symTable)
		}
		return errors
	}
	return errors
}
