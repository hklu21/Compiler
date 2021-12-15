package sa

import (
	"flag"
	"fmt"
	"proj/ast"
	st "proj/symbolTable"
)

func reportErrors(errors []string) bool {
	if len(errors) == 0 {
		return false
	} else {
		for _, err := range errors {
			out := flag.CommandLine.Output()
			fmt.Fprintf(out, "semantic error: %v\n", err)
		}
		return true
	}
}

func PerformSA(program *ast.Program) *st.SymbolTable {
	globalST := st.New(nil)
	errors := make([]string, 0)

	errors = program.PerformSABuild(errors, globalST)
	//fmt.Printf("%v\n", globalST)
	// report any errors
	if !reportErrors(errors) {
		errors := make([]string, 0)
		errors = program.TypeCheck(errors, globalST)
		if !reportErrors(errors) {
			return globalST
		}
	}
	return nil
}

/*
type Stack []map[string]Attribute

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str map[string]Attribute) {
	newStr := []map[string]Attribute{str}
	*s = append(newStr, *s...) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (map[string]Attribute, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		element := (*s)[0] // Index into the slice and obtain the element.
		*s = (*s)[1:]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Print() {
	fmt.Println("Printing Stack:")
	if !s.IsEmpty() {
		for _, stack := range *s {
			fmt.Println(stack)
		}
	}
}

// semantic Analysis for Variable, Struct, Funtion scopes
func SemanticAnalysis(as *ast.Program, errors []string, st map[string]Attribute) bool {
	if as.Pack.Id != "main" {
		errors = append(errors, "package is not main")
	}
	var stack Stack
	stack.Push(st)
	for _, function := range as.Funcs {
		funcHash := st["func"+function.Id]
		fmt.Println(funcHash)
		stack.Push(funcHash.(FunctionEntry).HTable)
		stack.Print()
	}

	if len(errors) > 0 {
		for _, err := range errors {
			parseError(err)
		}
		return false
	} else {
		return true
	}
}
*/
