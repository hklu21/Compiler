package ast

import (
	st "proj/symbolTable"
)

func (ast *Block) CheckReturn(symTable *st.SymbolTable) bool {
	if len(ast.Block) == 0 {
		return false
	}
	return ast.Block[len(ast.Block)-1].CheckReturn(symTable)
}

func (ast *Assignment) CheckReturn(symTable *st.SymbolTable) bool {
	return false
}

func (ast *ReadPrint) CheckReturn(symTable *st.SymbolTable) bool {
	return false
}

func (ast *Conditional) CheckReturn(symTable *st.SymbolTable) bool {
	if !ast.HasElse {
		return false
	} else {
		return ast.IfBlock.CheckReturn(symTable) && ast.ElseBlock.CheckReturn(symTable)
	}
}

func (ast *Loop) CheckReturn(symTable *st.SymbolTable) bool {
	return ast.ForBlock.CheckReturn(symTable)
}

func (ast *Return) CheckReturn(symTable *st.SymbolTable) bool {
	return true
}

func (ast *Invocation) CheckReturn(symTable *st.SymbolTable) bool {
	return false
}
