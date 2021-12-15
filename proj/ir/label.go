package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Label struct {
	source string // The first source register of the instruction
}

func NewLabels(source string) *Label {
	return &Label{source}
}

func (instr *Label) GetTargets() []int {
	return nil
}
func (instr *Label) GetSources() []int {
	return nil
}
func (instr *Label) GetImmediate() *int {
	return nil
}
func (instr *Label) GetLabel() string {
	// does not work with labels so we can just return a default value
	return instr.source
}
func (instr *Label) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Label) String() string {

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s:", instr.source))

	return out.String()

}

func (instr *Label) TranslateToArm(*codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, fmt.Sprintf("%v:", instr.GetLabel()))
	return output
}
