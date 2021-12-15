package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type New struct {
	target int // The target register for the instruction
	source string
}

func NewNew(target int, source string) *New {
	return &New{target, source}
}

func (instr *New) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *New) GetSources() []int {
	return nil
}
func (instr *New) GetImmediate() *int {
	return nil
}
func (instr *New) GetLabel() string {
	return ""
}
func (instr *New) SetLabel(newLabel string) {
}
func (instr *New) String() string {

	var out bytes.Buffer
	out.WriteString("new:\n")
	targetReg := fmt.Sprintf("r%v", instr.target)

	source := fmt.Sprintf("%v", instr.source)

	out.WriteString(fmt.Sprintf("\tnew %s,%s", targetReg, source))

	return out.String()

}

func (instr *New) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
