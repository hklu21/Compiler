package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type B struct {
	source string // The first source register of the instruction
}

func NewB(source string) *B {
	return &B{source}
}

func (instr *B) GetTargets() []int {
	return nil
}
func (instr *B) GetSources() []int {
	return nil
}
func (instr *B) GetImmediate() *int {
	return nil
}
func (instr *B) GetLabel() string {
	// does not work with labels so we can just return a default value
	return instr.source
}
func (instr *B) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *B) String() string {

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("\tb %s", instr.source))

	return out.String()

}

func (instr *B) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("\t\tb %v", instr.GetLabel()))
	return output
}
