package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Div struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	operand   int       // The operand either register or constant
	opty      OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewDiv(target int, sourceReg int, operand int, opty OperandTy) *Div {
	return &Div{target, sourceReg, operand, opty}
}

func (instr *Div) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Div) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *Div) GetImmediate() *int {

	// Div instruction has two forms for the second operand: register, and immediate (constant)
	// make sure to check for that.
	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	// Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Div) GetLabel() string {
	// Mul does not work with labels so we can just return a default value
	return ""
}
func (instr *Div) SetLabel(newLabel string) {
	// Div does not work with labels can we can skip implementing this method.
}
func (instr *Div) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.sourceReg)

	var prefix string

	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2 := fmt.Sprintf("%v%v", prefix, instr.operand)

	out.WriteString(fmt.Sprintf("\tdiv %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *Div) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	if !frame.IsArgument(instr.sourceReg) {
		output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
			frame.LocalOffset, frame.FindVarOffset(instr.sourceReg)))
		if instr.opty == IMMEDIATE {
			output = append(output, fmt.Sprintf("\t\tmov x%d,#%d",
				frame.LocalOffset+1, *instr.GetImmediate()))
		} else {
			if !frame.IsArgument(instr.operand) {
				output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
					frame.LocalOffset+1, frame.FindVarOffset(instr.operand)))
			}
		}
		output = append(output, fmt.Sprintf("\t\tdiv x%d,x%d, x%d",
			frame.LocalOffset+2, frame.LocalOffset, frame.LocalOffset+1))
	} else {
		if instr.opty == IMMEDIATE {
			output = append(output, fmt.Sprintf("\t\tmov x%d,#%d",
				frame.LocalOffset+2, *instr.GetImmediate()))
		} else {
			if !frame.IsArgument(instr.operand) {
				output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
					frame.LocalOffset+2, frame.FindVarOffset(instr.operand)))
			}
		}
		output = append(output, fmt.Sprintf("\t\tdiv x%d, x%d, x%d",
			frame.LocalOffset+2, instr.sourceReg, frame.LocalOffset+1))
	}
	// This is done because the target should always be a temporary register
	// in our implementation
	output = append(output, fmt.Sprintf("\t\tstr x%d,[x29,#%d]",
		frame.LocalOffset+2, frame.FindVarOffset(instr.target)))
	return output
}
