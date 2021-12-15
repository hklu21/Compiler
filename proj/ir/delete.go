package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Delete struct {
	source int
}

func NewDelete(source int) *Delete {
	return &Delete{source}
}

func (instr *Delete) GetTargets() []int {
	return nil
}
func (instr *Delete) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.source)
	return sources
}
func (instr *Delete) GetImmediate() *int {
	return nil
}
func (instr *Delete) GetLabel() string {
	return ""
}
func (instr *Delete) SetLabel(newLabel string) {
}
func (instr *Delete) String() string {

	var out bytes.Buffer
	out.WriteString("Delete:\n")
	source := fmt.Sprintf("%v", instr.source)

	out.WriteString(fmt.Sprintf("\tdelete r%s", source))

	return out.String()

}

func (instr *Delete) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
