package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Movle struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	sourcety  OperandTy // The type for the source (REGISTER, IMMEDIATE)
}

func NewMovle(target int, sourceReg int, sourcety OperandTy) *Movle {
	return &Movle{target, sourceReg, sourcety}
}

func (instr *Movle) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Movle) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *Movle) GetImmediate() *int {

	// has two forms for the second operand: register, and immediate (constant)
	// make sure to check for that.
	if instr.sourcety == IMMEDIATE {
		return &instr.sourceReg
	}
	// Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Movle) GetLabel() string {
	// does not work with labels so we can just return a default value
	return ""
}
func (instr *Movle) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Movle) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	var prefix string

	if instr.sourcety == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	sourceReg := fmt.Sprintf("%v%v", prefix, instr.sourceReg)

	out.WriteString(fmt.Sprintf("\tmovle %s,%s", targetReg, sourceReg))

	return out.String()

}

/*b.ge skipMov_L2
  mov x1,#1
  str x1,[x29,#-16]
skipMov_L2:*/
func (instr *Movle) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	label := NewLabelWithPre("skipMov")
	output = append(output, fmt.Sprintf("\t\tb.le %s", label))
	if instr.sourcety == IMMEDIATE {
		output = append(output, fmt.Sprintf("\t\tmov x%d,#%d",
			frame.LocalOffset, *instr.GetImmediate()))
	} else {
		output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
			frame.LocalOffset+1, frame.FindVarOffset(instr.sourceReg)))
		output = append(output, fmt.Sprintf("\t\tmov x%d,x%d",
			frame.LocalOffset, frame.LocalOffset+1))
	}
	output = append(output, fmt.Sprintf("\t\tstr x%d, [x29,#%d]",
		frame.LocalOffset, frame.FindVarOffset(instr.target)))
	output = append(output, fmt.Sprintf("%s:", label))
	return output
}
