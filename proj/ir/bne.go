package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Bne struct {
	source string // The first source register of the instruction
}

func NewBne(source string) *Bne {
	return &Bne{source}
}

func (instr *Bne) GetTargets() []int {
	return nil
}
func (instr *Bne) GetSources() []int {
	return nil
}
func (instr *Bne) GetImmediate() *int {
	return nil
}
func (instr *Bne) GetLabel() string {
	// does not work with labels so we can just return a default value
	return instr.source
}
func (instr *Bne) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Bne) String() string {

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("\tbne %s", instr.source))

	return out.String()

}

func (instr *Bne) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("\t\tb.eq %v", instr.GetLabel()))
	return output
}
