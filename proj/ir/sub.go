package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Sub struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	operand   int       // The operand either register or constant
	opty      OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewSub(target int, sourceReg int, operand int, opty OperandTy) *Sub {
	return &Sub{target, sourceReg, operand, opty}
}

func (instr *Sub) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Sub) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *Sub) GetImmediate() *int {

	//Sub instruction has two forms for the second operand: register, and immediate (constant)
	//make sure to check for that.
	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	//Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Sub) GetLabel() string {
	// Sub does not work with labels so we can just return a default value
	return ""
}
func (instr *Sub) SetLabel(newLabel string) {
	// Sub does not work with labels can we can skip implementing this method.
}
func (instr *Sub) String() string {

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

	out.WriteString(fmt.Sprintf("\tsub %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *Sub) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	// If The source register is not an argument
	if !frame.IsArgument(instr.sourceReg) {
		// Load it onto LocalOffset
		output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
			frame.LocalOffset, frame.FindVarOffset(instr.sourceReg)))
		// Load Operand onto LocalOffset + 1
		if instr.opty == IMMEDIATE {
			output = append(output, fmt.Sprintf("\t\tmov x%d,#%d",
				frame.LocalOffset+1, *instr.GetImmediate()))
		} else {
			if !frame.IsArgument(instr.operand) {
				output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
					frame.LocalOffset+1, frame.FindVarOffset(instr.operand)))
			}
		}
		// Add to local offset + 2
		output = append(output, fmt.Sprintf("\t\tsub x%d,x%d, x%d",
			frame.LocalOffset+2, frame.LocalOffset, frame.LocalOffset+1))
	} else {
		// Source register is an argument
		if instr.opty == IMMEDIATE {
			output = append(output, fmt.Sprintf("\t\tmov x%d,#%d",
				frame.LocalOffset+1, *instr.GetImmediate()))
		} else {
			if !frame.IsArgument(instr.operand) {
				output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
					frame.LocalOffset+1, frame.FindVarOffset(instr.operand)))
			} else {
				output = append(output, fmt.Sprintf("\t\tmov x%d,x%d",
					frame.LocalOffset+1, instr.operand))
			}
		}
		// Save Value in LocalOffset + 2
		output = append(output, fmt.Sprintf("\t\tsub x%d, x%d, x%d",
			frame.LocalOffset+2, instr.sourceReg, frame.LocalOffset+1))
	}
	// This is done because the target should always be a temporary register
	// in our implementation
	output = append(output, fmt.Sprintf("\t\tstr x%d,[x29,#%d]",
		frame.LocalOffset+2, frame.FindVarOffset(instr.target)))
	return output
}
