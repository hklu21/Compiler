package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Bl struct {
	source string // The first source register of the instruction
}

func NewBl(source string) *Bl {
	return &Bl{source}
}

func (instr *Bl) GetTargets() []int {
	return nil
}
func (instr *Bl) GetSources() []int {
	return nil
}
func (instr *Bl) GetImmediate() *int {
	return nil
}
func (instr *Bl) GetLabel() string {
	// does not work with labels so we can just return a default value
	return instr.source
}
func (instr *Bl) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Bl) String() string {

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("\tbl %s", instr.source))

	return out.String()

}

func (instr *Bl) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("\t\tbl %v", instr.GetLabel()))
	return output
}
