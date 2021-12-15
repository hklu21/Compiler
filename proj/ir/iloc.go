package ir

import (
	"fmt"
	"proj/codegen"
)

type OperandTy int

const (
	REGISTER OperandTy = iota
	IMMEDIATE
)

type Instruction interface {
	GetTargets() []int // Get the registers targeted by this instruction

	GetSources() []int // Get the source registers for this instruction

	GetImmediate() *int // Get the immediate value (i.e., constant) of this instruction

	GetLabel() string // Get the label for this instruction

	SetLabel(newLabel string) //Set the label for this instruction

	String() string // Return a string representation of this instruction

	TranslateToArm(frame *codegen.Frame) []string
}

type FuncFrag struct {
	Label string         // Function name
	Body  []Instruction  // Function body of ILOC instructions
	Frame *codegen.Frame // Activation Records (i.e., stack frame) for this function
}

func (ff *FuncFrag) GenerateCode() []string {
	output := make([]string, 0)
	// Generate Prologue Here
	output = append(output, "")
	output = append(output, fmt.Sprintf("\t\t.type %s,%%function", ff.Label))
	output = append(output, fmt.Sprintf("\t\t.global %s", ff.Label))
	output = append(output, "\t\t.p2align 2")
	output = append(output, fmt.Sprintf("%s:", ff.Label))
	output = append(output, "\t\tsub sp,sp,#16\n\t\tstp x29,x30, [sp]\n\t\tmov x29,sp")
	output = append(output, fmt.Sprintf("\t\tsub sp,sp, #%d", ff.Frame.CalcOffsetHex()))
	for _, instr := range ff.Body[1:] {
		output = append(output, instr.TranslateToArm(ff.Frame)...)
	}
	// Generate Epilogue Here
	if ff.Frame.Return == "Epsilon" {
		output = append(output, fmt.Sprintf("\t\tadd sp,sp,#%d", ff.Frame.CalcOffsetHex()))
		output = append(output, "\t\tldp x29,x30, [sp]\n\t\tadd sp,sp, 16")
		output = append(output, "\t\tret")
	}
	output = append(output, fmt.Sprintf("\t\t.size %s, (. - %s)", ff.Label, ff.Label))
	return output
}
