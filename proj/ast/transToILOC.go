package ast

import (
	"fmt"
	"proj/codegen"
	"proj/ir"
	st "proj/symbolTable"
	ct "proj/token"
	"strconv"
	"strings"
)

func (p *Program) TranslateToILOC(table st.SymbolTable) []*ir.FuncFrag {
	funcFrags := make([]*ir.FuncFrag, 0)
	for _, f := range p.Funcs {
		// st for a func
		st := table.St["func"+f.Id].(st.FunctionEntry).SymTable
		funcFrags = append(funcFrags, f.TranslateToILoc(*st))
	}
	return funcFrags
}

func (f *Function) TranslateToILoc(symTable st.SymbolTable) *ir.FuncFrag {
	FuncFrag := &ir.FuncFrag{Label: f.Id, Body: make([]ir.Instruction, 0)}
	FuncFrag.Body = append(FuncFrag.Body, ir.NewLabels(f.Id))
	startReg := ir.NewRegister() // start of temp regester
	fmt.Println(startReg)
	// function body
	for _, statement := range f.States {
		FuncFrag.Body = append(FuncFrag.Body, statement.TranslateToILoc(&symTable)...)
	}
	endReg := ir.NewRegister() // end of temp regester
	var frameArguments []codegen.Argument
	var frameLocalVar []codegen.LocalVar
	var frameTempVar []codegen.TempVar
	offset := 24
	for _, arg := range f.Param {
		frameArguments = append(frameArguments, codegen.Argument{ArgName: arg.Id, Register: symTable.St["var"+arg.Id].(st.VarEntry).Register, Offset: offset})
		offset += 8
	}
	offset = -8
	for _, decl := range f.Dec {
		for _, id := range decl.Id {
			frameLocalVar = append(frameLocalVar, codegen.LocalVar{VarName: id, Register: symTable.St["var"+id].(st.VarEntry).Register, Offset: offset})
			offset -= 8
		}
	}

	for reg := startReg + 1; reg < endReg; reg++ {
		frameTempVar = append(frameTempVar, codegen.TempVar{Register: reg, Offset: offset})
		offset -= 8
	}
	FuncFrag.Frame = &codegen.Frame{FrameName: f.Id, Arguments: frameArguments, LocalOffset: len(frameArguments), LocalVar: frameLocalVar, TempVar: frameTempVar, Return: f.ReturnT}
	fmt.Println(FuncFrag.Frame)

	return FuncFrag
}

func (s *Block) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instructions := make([]ir.Instruction, 0)
	for _, stat := range s.Block {
		instructions = append(instructions, stat.TranslateToILoc(symTable)...)
	}
	return instructions
}

func (s *ReadPrint) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instructions := make([]ir.Instruction, 0)
	var target int
	if _, found := symTable.St["var"+s.Id]; found {
		// read&print a local id
		target = symTable.St["var"+s.Id].(st.VarEntry).Register
	} else {
		// read&print a global id
		target = symTable.Parent.St["var"+s.Id].(st.VarEntry).Register
	}
	switch s.FuncName {
	case "Scan":
		instructions = append(instructions, ir.NewRead(target))
	case "Print":
		instructions = append(instructions, ir.NewPrint(target))
	case "Println":
		instructions = append(instructions, ir.NewPrintln(target))
	}
	return instructions
}

func (s *Assignment) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	instruction = append(instruction, s.Exprs.TranslateToILoc(symTable)...)
	var target int
	right := instruction[len(instruction)-1].GetTargets()[0] // final target of right value
	if len(s.Lvalue) > 1 {
		var source int
		var feild string
		target = ir.NewRegister()
		if _, found := symTable.St["var"+s.Lvalue[0]]; found {
			// local struct
			source = symTable.St["var"+s.Lvalue[0]].(st.VarEntry).Register
		} else {
			// global struct
			source = ir.NewRegister()
			instruction = append(instruction, ir.NewLdr(source, s.Lvalue[0]))
		}
		for i := 1; i < len(s.Lvalue)-1; i++ {
			feild = s.Lvalue[i]
			instruction = append(instruction, ir.NewLoadRef(target, source, feild))
			source = target
			target = ir.NewRegister()
		}
		feild = s.Lvalue[len(s.Lvalue)-1]
		instruction = append(instruction, ir.NewStrRef(right, source, feild))
	} else {
		if _, found := symTable.St["var"+s.Lvalue[0]]; found {
			// local
			target = symTable.St["var"+s.Lvalue[0]].(st.VarEntry).Register
			instruction = append(instruction, ir.NewMove(target, right, ir.REGISTER))
		} else {
			// global
			instruction = append(instruction, ir.NewStr(right, s.Lvalue[0]))
		}
	}
	return instruction
}

func (s *Invocation) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	var target int
	if s.Id == "new" {
		target = ir.NewRegister()
		instruction = append(instruction, ir.NewNew(target, s.Args[0].String()))
	} else if s.Id == "delete" {
		tmp := s.Args[0].TranslateToILoc(symTable)
		target = tmp[len(tmp)-1].GetTargets()[0]
		instruction = append(instruction, ir.NewDelete(target))
	} else {
		sourceRegs := make([]int, 0)
		for _, para := range s.Args {
			instruction = append(instruction, para.TranslateToILoc(symTable)...)
			sourceRegs = append(sourceRegs, instruction[len(instruction)-1].GetTargets()[0])
		}
		instruction = append(instruction, ir.NewPush(sourceRegs))

		instruction = append(instruction, ir.NewBl(s.Id))
		instruction = append(instruction, ir.NewPop([]int{ir.NewRegister()}))
		source := instruction[len(instruction)-1].GetSources()[0]
		instruction = append(instruction, ir.NewMove(ir.NewRegister(), source, ir.REGISTER))
	}
	return instruction
}

func (s *Conditional) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	var source1 int
	var source2 int
	if !s.HasElse {
		instruction = append(instruction, s.Expr.TranslateToILoc(symTable)...)
		source1 = instruction[len(instruction)-1].GetTargets()[0]
		source2 = ir.NewRegister()
		instruction = append(instruction, ir.NewMove(source2, 1, ir.IMMEDIATE))
		instruction = append(instruction, ir.NewCmp(source1, source2))
		label1 := ir.NewLabelWithPre("ifLabel")
		label2 := ir.NewLabelWithPre("done")
		instruction = append(instruction, ir.NewBne(label2))
		instruction = append(instruction, ir.NewLabels(label1))
		instruction = append(instruction, s.IfBlock.TranslateToILoc(symTable)...)
		instruction = append(instruction, ir.NewLabels(label2))
	} else {
		instruction = append(instruction, s.Expr.TranslateToILoc(symTable)...)
		source1 = instruction[len(instruction)-1].GetTargets()[0]
		source2 = ir.NewRegister()
		instruction = append(instruction, ir.NewMove(source2, 1, ir.IMMEDIATE))
		instruction = append(instruction, ir.NewCmp(source1, source2))
		label1 := ir.NewLabelWithPre("ifLabel")
		label2 := ir.NewLabelWithPre("elseLabel")
		instruction = append(instruction, ir.NewBne(label2))
		instruction = append(instruction, ir.NewLabels(label1))
		instruction = append(instruction, s.IfBlock.TranslateToILoc(symTable)...)
		label3 := ir.NewLabelWithPre("done")
		instruction = append(instruction, ir.NewB(label3))
		instruction = append(instruction, ir.NewLabels(label2))
		instruction = append(instruction, s.ElseBlock.TranslateToILoc(symTable)...)
		instruction = append(instruction, ir.NewLabels(label3))
	}
	return instruction
}

func (s *Loop) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	var source1 int
	var source2 int
	label1 := ir.NewLabelWithPre("loopBody")
	label2 := ir.NewLabelWithPre("testCond")
	instruction = append(instruction, ir.NewB(label2))

	instruction = append(instruction, ir.NewLabels(label1))
	instruction = append(instruction, s.ForBlock.TranslateToILoc(symTable)...)

	instruction = append(instruction, ir.NewLabels(label2))
	instruction = append(instruction, s.Expr.TranslateToILoc(symTable)...)
	source1 = instruction[len(instruction)-1].GetTargets()[0]
	source2 = ir.NewRegister()
	instruction = append(instruction, ir.NewMove(source2, 1, ir.IMMEDIATE))
	instruction = append(instruction, ir.NewCmp(source1, source2))
	instruction = append(instruction, ir.NewBeq(label1))

	return instruction
}

func (s *Return) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	if s.Exprs == nil {
		instruction = append(instruction, ir.NewRet("void"))
	} else {
		instruction = append(instruction, s.Exprs.TranslateToILoc(symTable)...)
		source := instruction[len(instruction)-1].GetTargets()[0]
		// push to stack
		instruction = append(instruction, ir.NewRet(strconv.Itoa(source)))
	}
	return instruction
}

func (s *AndOrExpr) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := s.Left.TranslateToILoc(symTable)
	instructionRight := s.Right.TranslateToILoc(symTable)
	source := instruction[len(instruction)-1].GetTargets()[0]
	instruction = append(instruction, instructionRight...)
	operator := instructionRight[len(instructionRight)-1].GetTargets()[0]
	target := ir.NewRegister()
	switch s.Token.Type {
	case ct.AND:
		instruction = append(instruction, ir.NewAnd(target, source, operator, ir.REGISTER))
	case ct.OR:
		instruction = append(instruction, ir.NewOr(target, source, operator, ir.REGISTER))
	}
	return instruction
}

func (s *CompareExpr) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := s.Left.TranslateToILoc(symTable)
	instructionRight := s.Right.TranslateToILoc(symTable)
	left := instruction[len(instruction)-1].GetTargets()[0]
	instruction = append(instruction, instructionRight...)
	right := instructionRight[len(instructionRight)-1].GetTargets()[0]
	target := ir.NewRegister()
	instruction = append(instruction, ir.NewMove(target, -1, ir.IMMEDIATE))
	instruction = append(instruction, ir.NewCmp(left, right))
	switch s.Comparator {
	case EQUAL:
		instruction = append(instruction, ir.NewMoveq(target, 1, ir.IMMEDIATE))
	case NEQ:
		instruction = append(instruction, ir.NewMovne(target, 1, ir.IMMEDIATE))
	case GREATER:
		instruction = append(instruction, ir.NewMovgt(target, 1, ir.IMMEDIATE))
	case LESSER:
		instruction = append(instruction, ir.NewMovlt(target, 1, ir.IMMEDIATE))
	case GEQ:
		instruction = append(instruction, ir.NewMovge(target, 1, ir.IMMEDIATE))
	case LEQ:
		instruction = append(instruction, ir.NewMovle(target, 1, ir.IMMEDIATE))
	}
	return instruction
}

func (s *BinOpExpr) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := s.Left.TranslateToILoc(symTable)
	instructionRight := s.Right.TranslateToILoc(symTable)
	source := instruction[len(instruction)-1].GetTargets()[0]
	instruction = append(instruction, instructionRight...)
	operator := instructionRight[len(instructionRight)-1].GetTargets()[0]
	target := ir.NewRegister()
	switch s.Operator {
	case ADD:
		instruction = append(instruction, ir.NewAdd(target, source, operator, ir.REGISTER))
	case SUB:
		instruction = append(instruction, ir.NewSub(target, source, operator, ir.REGISTER))
	case MULT:
		instruction = append(instruction, ir.NewMul(target, source, operator, ir.REGISTER))
	case DIV:
		instruction = append(instruction, ir.NewDiv(target, source, operator, ir.REGISTER))
	}
	return instruction
}

func (s *Unary) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := s.Selector.TranslateToILoc(symTable)
	target := instruction[len(instruction)-1].GetTargets()[0]
	if s.SelectID == "!" {
		tmp := ir.NewRegister()
		instruction = append(instruction, ir.NewMove(tmp, target, ir.REGISTER))
		instruction = append(instruction, ir.NewNot(tmp, tmp, ir.REGISTER))
	} else if s.SelectID == "-" {
		tmp := ir.NewRegister()
		instruction = append(instruction, ir.NewMove(tmp, target, ir.REGISTER))
		instruction = append(instruction, ir.NewMul(tmp, tmp, -1, ir.IMMEDIATE))
	}
	return instruction
}

func (s *Selector) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := s.Factor.TranslateToILoc(symTable)
	ids := strings.Split(s.SelectID[1:], ".")
	var target int
	var source int
	var feild string
	target = ir.NewRegister()
	source = symTable.St["var"+s.Factor.String()].(st.VarEntry).Register
	for i := 0; i < len(ids); i++ {
		feild = ids[i]
		instruction = append(instruction, ir.NewLoadRef(target, source, feild))
		source = target
		target = ir.NewRegister()
	}
	return instruction
}

func (s *Factor) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	instruction = append(instruction, s.Exprs.TranslateToILoc(symTable)...)
	return instruction
}

func (s *Number) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	target := ir.NewRegister()
	i, _ := strconv.Atoi(s.String())
	instruction = append(instruction, ir.NewMove(target, i, ir.IMMEDIATE))
	return instruction
}

func (s *Boolean) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	target := ir.NewRegister()
	var i int
	if s.String() == "true" {
		i = 1
	} else if s.String() == "false" {
		i = -1
	}
	instruction = append(instruction, ir.NewMove(target, i, ir.IMMEDIATE))
	return instruction
}

func (s *ID) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	if s.Call {
		invocation := Invocation{Id: s.Id, Args: s.Args}
		instr := invocation.TranslateToILoc(symTable)
		instruction = append(instruction, instr...)
		return instruction
	}
	var target int
	if _, found := symTable.St["var"+s.String()]; found {
		// local
		target = symTable.St["var"+s.String()].(st.VarEntry).Register
		instruction = append(instruction, ir.NewMove(target, target, ir.REGISTER))
	} else {
		// global
		target = ir.NewRegister()
		instruction = append(instruction, ir.NewLdr(target, s.String()))
	}
	return instruction
}

func (s *Nil) TranslateToILoc(symTable *st.SymbolTable) []ir.Instruction {
	instruction := make([]ir.Instruction, 0)
	target := ir.NewRegister()
	instruction = append(instruction, ir.NewMove(target, 0, ir.IMMEDIATE))
	return instruction
}
