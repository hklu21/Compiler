package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Movgt struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	sourcety  OperandTy // The type for the source (REGISTER, IMMEDIATE)
}

func NewMovgt(target int, sourceReg int, sourcety OperandTy) *Movgt {
	return &Movgt{target, sourceReg, sourcety}
}

func (instr *Movgt) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Movgt) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *Movgt) GetImmediate() *int {

	// has two forms for the second operand: register, and immediate (constant)
	// make sure to check for that.
	if instr.sourcety == IMMEDIATE {
		return &instr.sourceReg
	}
	// Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Movgt) GetLabel() string {
	// does not work with labels so we can just return a default value
	return ""
}
func (instr *Movgt) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Movgt) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	var prefix string

	if instr.sourcety == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	sourceReg := fmt.Sprintf("%v%v", prefix, instr.sourceReg)

	out.WriteString(fmt.Sprintf("\tmovgt %s,%s", targetReg, sourceReg))

	return out.String()

}

func (instr *Movgt) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	if instr.sourceReg == instr.target {
		return output
	} else {
		if instr.sourcety == IMMEDIATE {
			output = append(output, fmt.Sprintf("\t\tmovgt x%d,#%d",
				frame.LocalOffset, *instr.GetImmediate()))
		} else {
			output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#-%d]",
				frame.LocalOffset+1, frame.FindVarOffset(instr.sourceReg)))
			output = append(output, fmt.Sprintf("\t\tmovgt x%d,x%d",
				frame.LocalOffset, frame.LocalOffset+1))
		}
		output = append(output, fmt.Sprintf("\t\tstrdgt x%d,[x29,#-%d]",
			frame.LocalOffset, frame.FindVarOffset(instr.target)))
		return output
	}
}
