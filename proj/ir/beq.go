package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Beq struct {
	source string // The first source register of the instruction
}

func NewBeq(source string) *Beq {
	return &Beq{source}
}

func (instr *Beq) GetTargets() []int {
	return nil
}
func (instr *Beq) GetSources() []int {
	return nil
}
func (instr *Beq) GetImmediate() *int {
	return nil
}
func (instr *Beq) GetLabel() string {
	// does not work with labels so we can just return a default value
	return instr.source
}
func (instr *Beq) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Beq) String() string {

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("\tbeq %s", instr.source))

	return out.String()

}

func (instr *Beq) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("\t\tb.eq %v", instr.GetLabel()))
	return output
}
