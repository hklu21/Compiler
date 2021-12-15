package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Read struct {
	target int // The target register for the instruction
}

func NewRead(target int) *Read {
	return &Read{target}
}

func (instr *Read) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Read) GetSources() []int {
	sources := make([]int, 0)
	return sources
}
func (instr *Read) GetImmediate() *int {
	return nil
}
func (instr *Read) GetLabel() string {
	return ""
}
func (instr *Read) SetLabel(newLabel string) {
}
func (instr *Read) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	out.WriteString(fmt.Sprintf("\tread %s", targetReg))
	return out.String()

}

func (instr *Read) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]",
		frame.LocalOffset+1, frame.FindVarOffset(instr.target)))
	output = append(output, fmt.Sprintf("\t\tadrp x%d, .READ",
		frame.LocalOffset+2))
	output = append(output, fmt.Sprintf("\t\tadd x%d, x%d, :lo12:.READ",
		frame.LocalOffset+2, frame.LocalOffset+2))
	output = append(output, fmt.Sprintf("\t\tadd x1, x29,#%d",
		frame.FindVarOffset(instr.target)))
	output = append(output, fmt.Sprintf("\t\tmov x0, x%d",
		frame.LocalOffset+2))
	output = append(output, fmt.Sprintf("\t\tbl scanf"))
	return output
}
