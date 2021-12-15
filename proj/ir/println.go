package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Println struct {
	target int // The target register for the instruction
}

func NewPrintln(target int) *Println {
	return &Println{target}
}

func (instr *Println) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Println) GetSources() []int {
	sources := make([]int, 0)
	return sources
}
func (instr *Println) GetImmediate() *int {
	return nil
}
func (instr *Println) GetLabel() string {
	return ""
}
func (instr *Println) SetLabel(newLabel string) {
}
func (instr *Println) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	out.WriteString(fmt.Sprintf("\tprintln %s", targetReg))
	return out.String()

}

func (instr *Println) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
		frame.LocalOffset+1, frame.FindVarOffset(instr.target)))
	if frame.IsArgument(instr.target) {
		output = append(output, fmt.Sprintf("\t\tmov x%d,x%d",
			frame.LocalOffset+1, instr.target))
	}
	output = append(output, fmt.Sprintf("\t\tadrp x%d, .PRINT_LN",
		frame.LocalOffset+2))
	output = append(output, fmt.Sprintf("\t\tadd x%d, x%d, :lo12:.PRINT_LN",
		frame.LocalOffset+2, frame.LocalOffset+2))
	output = append(output, fmt.Sprintf("\t\tmov x0, x%d",
		frame.LocalOffset+2))
	output = append(output, fmt.Sprintf("\t\tmov x1, x%d",
		frame.LocalOffset+1))
	output = append(output, fmt.Sprintf("\t\tbl printf"))
	return output
}
