package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Not struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	sourcety  OperandTy // The type for the source (REGISTER, IMMEDIATE)
}

func NewNot(target int, sourceReg int, sourcety OperandTy) *Not {
	return &Not{target, sourceReg, sourcety}
}

func (instr *Not) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Not) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *Not) GetImmediate() *int {

	// has two forms for the second operand: register, and immediate (constant)
	// make sure to check for that.
	if instr.sourcety == IMMEDIATE {
		return &instr.sourceReg
	}
	// Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Not) GetLabel() string {
	// does not work with labels so we can just return a default value
	return ""
}
func (instr *Not) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Not) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	var prefix string

	if instr.sourcety == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	sourceReg := fmt.Sprintf("%v%v", prefix, instr.sourceReg)

	out.WriteString(fmt.Sprintf("\tnot %s,%s", targetReg, sourceReg))

	return out.String()

}

func (instr *Not) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	if instr.sourcety == IMMEDIATE {
		output = append(output, fmt.Sprintf("\t\tneg x%d,#%d",
			frame.LocalOffset, *instr.GetImmediate()))
	} else {
		output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
			frame.LocalOffset, frame.FindVarOffset(instr.sourceReg)))
		output = append(output, fmt.Sprintf("\t\tneg x%d,x%d",
			frame.LocalOffset, frame.LocalOffset))
	}
	output = append(output, fmt.Sprintf("\t\tstr x%d,[x29,#%d]",
		frame.LocalOffset, frame.FindVarOffset(instr.target)))
	return output
}
