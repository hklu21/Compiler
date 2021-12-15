package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Cmp struct {
	sourceReg1 int // The first source register of the instruction
	sourceReg2 int // The second source register of the instruction
}

func NewCmp(sourceReg1 int, sourceReg2 int) *Cmp {
	return &Cmp{sourceReg1, sourceReg2}
}

func (instr *Cmp) GetTargets() []int {
	return nil
}
func (instr *Cmp) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg1)
	sources = append(sources, instr.sourceReg2)
	return sources
}
func (instr *Cmp) GetImmediate() *int {
	return nil
}
func (instr *Cmp) GetLabel() string {
	return "cmp"
}
func (instr *Cmp) SetLabel(newLabel string) {
}
func (instr *Cmp) String() string {
	var out bytes.Buffer
	sourceReg1 := fmt.Sprintf("r%v", instr.sourceReg1)
	sourceReg2 := fmt.Sprintf("r%v", instr.sourceReg2)
	out.WriteString(fmt.Sprintf("\tcmp %s,%s", sourceReg1, sourceReg2))
	return out.String()
}

func (instr *Cmp) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	// If soruce reg is not argument
	if !frame.IsArgument(instr.sourceReg1) {
		// Load srouce reg into local offset
		output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
			frame.LocalOffset, frame.FindVarOffset(instr.sourceReg1)))
		// If source Reg 2 is not argument
		if !frame.IsArgument(instr.sourceReg2) {
			// Load source reg 2 into local offset + 1
			output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
				frame.LocalOffset+1, frame.FindVarOffset(instr.sourceReg2)))
			// and print
			output = append(output, fmt.Sprintf("\t\tcmp x%d,x%d",
				frame.LocalOffset, frame.LocalOffset+1))
		} else {
			output = append(output, fmt.Sprintf("\t\tcmp x%d,x%d",
				instr.sourceReg1, instr.sourceReg2))
		}
	} else {
		// Source Reg 1 is an argument
		// Check if source reg 2 is an argument
		if !frame.IsArgument(instr.sourceReg2) {
			output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
				frame.LocalOffset+1, frame.FindVarOffset(instr.sourceReg2)))
			output = append(output, fmt.Sprintf("\t\tcmp x%d,x%d",
				instr.sourceReg1, frame.LocalOffset+1))
		} else {
			output = append(output, fmt.Sprintf("\t\tcmp x%d,x%d",
				instr.sourceReg1, instr.sourceReg2))
		}
	}
	return output
}
