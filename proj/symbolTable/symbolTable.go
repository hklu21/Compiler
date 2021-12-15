package symbolTable

import (
	"fmt"
)

func parseError(msg string) {
	fmt.Printf("semantic error: %s\n", msg)
}

type SymbolTable struct {
	Parent *SymbolTable
	St     map[string]Attribute
	Return string
}

type Attribute interface {
}

type VarEntry struct {
	Usable   bool
	Type     string
	Register int
}

type StructEntry struct {
	Variables map[string]Attribute
}

type FunctionEntry struct {
	Params     int
	Parameters []VarEntry
	SymTable   *SymbolTable
	Return     string
}

// TODO Update this to map
/*
func (st *SymbolTable) String() string {
	var out bytes.Buffer
	if len(st.GlobalVar) > 0 {
			out.WriteString(variable.String())
			out.WriteString("\n")
		}
	}
	if len(st.Structs) > 0 {
		for _, structure := range st.Structs {
			out.WriteString(structure.String())
			out.WriteString("\n")
		}
	}
	if len(st.Functions) > 0 {
		for _, function := range st.Functions {
			out.WriteString(function.String())
			out.WriteString("\n")
		}
	}
	return out.String()
}*/

/*type SymbolTable struct {
	GlobalVar []VarEntry
	Structs   []StructEntry
	Functions []FunctionEntry
	BadVar    []string
}*/
// TODO: Add check for redefinition error when producing Symbol Table

func New(parent *SymbolTable) *SymbolTable {
	HMap := make(map[string]Attribute)
	HMap["structint"] = nil
	HMap["structbool"] = nil
	if parent == nil {
		return &SymbolTable{Parent: nil, St: HMap}
	} else {
		return &SymbolTable{Parent: parent, St: HMap}
	}
}

func IsInSymTable(st *SymbolTable, input string) bool {
	_, found := st.St[input]
	if !found && st.Parent != nil {
		return IsInSymTable(st.Parent, input)
	}
	return found
}

func IsInMySymTable(st *SymbolTable, input string) bool {
	_, found := st.St[input]
	return found
}

func UsableVarSymTable(st *SymbolTable, input string) bool {
	// Assumes that the input can be found in the symbol table
	value, found := st.St[input]
	if !found && st.Parent != nil {
		return UsableVarSymTable(st.Parent, input)
	}
	return value.(VarEntry).Usable
}

func GetSymTableEntry(st *SymbolTable, input string) Attribute {
	value, found := st.St[input]
	if !found {
		if st.Parent != nil {
			return GetSymTableEntry(st.Parent, input)
		} else {
			return nil
		}
	}
	return value
}

func GetTypeFromST(st *SymbolTable, input string) string {
	value, found := st.St["var"+input]
	if !found {
		if st.Parent != nil {
			return GetTypeFromST(st.Parent, input)
		} else {
			if value, found = st.St["func"+input]; found {
				return value.(FunctionEntry).Return
			} else {
				return input
			}
		}
	} else {
		varEntry, _ := value.(VarEntry)
		return varEntry.Type
	}
}

func GetVarEntryFromST(st *SymbolTable, input string) Attribute {
	value, found := st.St["var"+input]
	if !found {
		if st.Parent != nil {
			return GetVarEntryFromST(st.Parent, input)
		} else {
			return nil
		}
	} else {
		return value
	}
}

func GetFuncRetFromST(st *SymbolTable, input string) string {
	value, found := st.St["func"+input]
	if !found {
		if st.Parent != nil {
			return GetTypeFromST(st.Parent, input)
		} else {
			return "nil"
		}
	} else {
		funcEntry, _ := value.(FunctionEntry)
		return funcEntry.Return
	}
}

func UpdateUsable(symTable *SymbolTable, input string, newState bool) {
	if value, found := symTable.St["var"+input]; found {
		newVal := value.(VarEntry)
		newVal.Usable = newState
		symTable.St["var"+input] = newVal
	} else {
		UpdateUsable(symTable.Parent, input, newState)
	}
}

func PrintST(symTable *SymbolTable) {
	for key, element := range symTable.St {
		fmt.Printf("key: %s element: %v\n", key, element)
		if len(key) >= 4 {
			if key[:4] == "func" && key != "funcnew" && key != "funcdelete" {
				newEle := element.(FunctionEntry).SymTable.St
				if len(newEle) != 0 {
					for key, element := range newEle {
						fmt.Printf("key: %s element: %v\n", key, element)
					}
				}
			}
		}

	}
}

/*
func New(errors []string, ast *ast.Program) map[string]Attribute {
	st := make(map[string]Attribute)
	//st := SymbolTable{}
	for _, structType := range ast.Types {
		if _, found := st[structType.Id]; found {
			parseError("Struct " + structType.Id + "is already defined")
		}
		st["struct"+structType.Id] = *ReadStruct(structType)
	}
	for _, varDec := range ast.Decl {
		for _, id := range varDec.Id {
			saveVar := ReadVar(varDec.Value, st)
			if saveVar == nil {
				parseError("Global variable " + id + " has non-defined type " + varDec.Value)
			}
			st["var"+id] = *saveVar
		}
	}
	for _, function := range ast.Funcs {
		saveFunc := ReadFunction(function, st)
		if saveFunc == nil {
			parseError("Func " + function.Id + " is already defined")
		}
		st["func"+function.Id] = *saveFunc
	}
	return st
}
*/
/*
func (ve *VarEntry) String() string {
	var out bytes.Buffer
	out.WriteString("Variable Name: ")
	out.WriteString(ve.Name)
	out.WriteString(" Variable Type: ")
	out.WriteString(ve.Type)
	return out.String()
}*/
/*
func ReadVar(varType string, st map[string]Attribute) *VarEntry {
	if varType == "bool" || varType == "int" {
		return &VarEntry{Usable: true, Type: varType}
	} else {
		if _, found := st["struct"+varType]; found {
			return &VarEntry{Usable: false, Type: varType}
		} else {
			return nil
		}
	}
}
*/

/*
func (se *StructEntry) String() string {
	var out bytes.Buffer
	out.WriteString("Struct Name: ")
	out.WriteString(se.Name)
	if len(se.Variables) > 0 {
		out.WriteString(" Contains the Following Variables: \n")
		for _, variable := range se.Variables {
			out.WriteString("  ")
			out.WriteString(variable.String())
			out.WriteString("\n")
		}
	}
	return out.String()
}*/
/*
func ReadStruct(ast ast.Types) *StructEntry {
	output := make(map[string]Attribute)
	for _, variable := range ast.Fields {
		saveVar := ReadVar(variable.Value, output)
		if saveVar == nil {
			parseError("Variable " + variable.Id + " is already defined within struct " + ast.Id)
		}
		output[variable.Id] = *saveVar
	}
	return &StructEntry{Variables: output}
}
*/
/*
func (fe *FunctionEntry) String() string {
	var out bytes.Buffer
	out.WriteString("Function Name: ")
	out.WriteString(fe.Name)
	out.WriteString(", Function Return Type: ")
	out.WriteString(fe.Return)
	if len(fe.Inputs) > 0 {
		out.WriteString(" Requires the Following Inputs: \n")
		for _, input := range fe.Inputs {
			out.WriteString("  ")
			out.WriteString(input.String())
			out.WriteString("\n")
		}
	}
	if len(fe.Variables) > 0 {
		out.WriteString("Within the Function, the following Variables are Defined: \n")
		for _, input := range fe.Variables {
			out.WriteString("  ")
			out.WriteString(input.String())
			out.WriteString("\n")
		}
	}
	return out.String()
}*/

/*
func ReadFunction(ast ast.Function, st map[string]Attribute) *FunctionEntry {
	if _, found := st[ast.Id]; found {
		return nil
	}
	output := FunctionEntry{Parent: &st, Return: ast.ReturnT}
	hashTable := make(map[string]Attribute)
	for _, variable := range ast.Param {
		saveParam := ReadVar(variable.Value, st)
		if saveParam == nil {
			parseError("Input " + variable.Id + " has non-defined type " + variable.Value + " within function " + ast.Id)
		}
		if _, found := hashTable[variable.Id]; found {
			parseError("Multiple inputs named " + variable.Id + " within function " + ast.Id)
		}
		saveParam.Usable = true
		hashTable[variable.Id] = *saveParam
	}
	for _, decl := range ast.Dec {
		for _, id := range decl.Id {
			saveVar := ReadVar(decl.Value, st)
			if saveVar == nil {
				parseError("Variable " + id + " has non-defined type " + decl.Value + " within function " + ast.Id)
			}
			if _, found := hashTable[id]; found {
				parseError(id + " redefined within function " + ast.Id)
			}
			hashTable[id] = *saveVar
		}
	}
	output.HTable = hashTable
	return &output
}
*/
