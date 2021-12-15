package ast

import (
	"proj/ir"
	st "proj/symbolTable"
)

func (ast *Program) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	if ast.Pack.Id != "main" {
		return append(errors, "Package is not declared as main")
	}
	structs := make([]st.VarEntry, 0)
	for _, types := range ast.Types {
		errors = types.PerformSABuild(errors, symTable)
		structs = append(structs, st.VarEntry{Type: types.Id})
	}
	for _, decl := range ast.Decl {
		errors = decl.PerformSABuild(errors, symTable)
	}
	funcEntry := st.FunctionEntry{Params: 1, Parameters: structs, Return: "pointer"}
	symTable.St["funcnew"] = funcEntry
	funcEntry = st.FunctionEntry{Params: 1, Parameters: structs, Return: "epsilon"}
	symTable.St["funcdelete"] = funcEntry
	for _, funcs := range ast.Funcs {
		errors = funcs.PerformSABuild(errors, symTable)
	}

	// Check that function main exists and that it takes no arguments
	if funcmain, found := symTable.St["funcmain"]; !found {
		errors = append(errors, "function main is undeclared in the main package")
	} else if funcEntry, _ := funcmain.(st.FunctionEntry); funcEntry.Params != 0 {
		errors = append(errors, "function main cannot have any input parameters")
	}
	return errors
}

func (ast *Types) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	// Check if struct is already defined in symbol Table
	if st.IsInSymTable(symTable, "struct"+ast.Id) {
		return append(errors, "Struct "+ast.Id+" is already defined")
	}
	structEntry := st.StructEntry{Variables: make(map[string]st.Attribute)}

	// TODO: Determine if use pointers or not for symbol Table entries
	// Insert first to make sure it can include pointers of itself
	symTable.St["struct"+ast.Id] = structEntry
	for _, fields := range ast.Fields {
		if fields.Value == "bool" || fields.Value == "int" {
			if _, found := structEntry.Variables[fields.Id]; !found {
				structEntry.Variables[fields.Id] = st.VarEntry{Usable: true, Type: fields.Value}
			} else {
				errors = append(errors, "Varaible "+fields.Id+" is already defined within struct "+ast.Id)
			}
		} else {
			// Check if a struct pointer has been previously defined
			if _, found := symTable.St["struct"+fields.Value]; found {
				structEntry.Variables[fields.Id] = st.VarEntry{Usable: false, Type: fields.Value}
			} else {
				errors = append(errors, "Struct Pointer of type "+fields.Value+" not defined before "+ast.Id)
			}
		}
	}
	symTable.St["struct"+ast.Id] = structEntry
	return errors
}

func (ast *Declaration) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// DONE
	for _, ids := range ast.Id {
		// If it is a primative type
		if ast.Value == "bool" || ast.Value == "int" {
			if _, found := symTable.St["var"+ids]; !found {
				symTable.St["var"+ids] = st.VarEntry{Usable: true, Type: ast.Value}
			} else {
				errors = append(errors, "Global Varaible "+ids+" already defined")
			}
		} else {
			if _, found := symTable.St["var"+ids]; found {
				errors = append(errors, "Global Varaible "+ids+" already defined")
			}
			if _, found := symTable.St["struct"+ast.Value]; !found {
				errors = append(errors, "Global Variable "+ids+" is of undefined struct type "+ast.Value)
			}
			// Insert variable and structs into table regardless of error status
			// WRITE IN ESSAY: Above Decision
			symTable.St["var"+ids] = st.VarEntry{Usable: false, Type: ast.Value}
		}
	}
	return errors
}

func (ast *Function) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// Check if function is already defined
	if st.IsInSymTable(symTable, "func"+ast.Id) {
		errors = append(errors, "Function "+ast.Id+" is already defined")
	}
	symT := st.New(symTable)
	symT.Return = ast.ReturnT
	funcEntry := st.FunctionEntry{SymTable: symT, Params: len(ast.Param), Parameters: make([]st.VarEntry, 0), Return: ast.ReturnT}
	// Check Parameters
	for _, param := range ast.Param {
		// Check if struct pointer is defined
		if !st.IsInSymTable(symTable, "struct"+param.Value) {
			errors = append(errors, "Param "+param.Id+" in function "+ast.Id+" is of undefined Struct type "+param.Value)
		}
		// Check if param has previously been defined
		if st.IsInMySymTable(funcEntry.SymTable, "var"+param.Id) {
			errors = append(errors, "Param "+param.Id+" is already defined in function "+ast.Id)
		}
		// Insert param into table regardless of error
		varEntry := st.VarEntry{Usable: true, Type: param.Value, Register: ir.NewRegister()}
		funcEntry.Parameters = append(funcEntry.Parameters, varEntry)
		funcEntry.SymTable.St["var"+param.Id] = varEntry
	}
	// Check Varaiable Declarations
	for _, dec := range ast.Dec {
		for _, id := range dec.Id {
			if !st.IsInSymTable(symTable, "struct"+dec.Value) {
				errors = append(errors, "Variable "+id+" in function "+ast.Id+" is of undefined Struct type "+dec.Value)
			} else if st.IsInMySymTable(funcEntry.SymTable, "var"+id) {
				errors = append(errors, "Variable "+id+" is already defined in function "+ast.Id)
			} else {
				if dec.Value == "bool" || dec.Value == "int" {
					funcEntry.SymTable.St["var"+id] = st.VarEntry{Usable: true, Type: dec.Value, Register: ir.NewRegister()}
				} else {
					funcEntry.SymTable.St["var"+id] = st.VarEntry{Usable: false, Type: dec.Value, Register: ir.NewRegister()}
				}
			}
		}
	}
	symTable.St["func"+ast.Id] = funcEntry
	return errors
}
