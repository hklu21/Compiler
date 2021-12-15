package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type Str struct {
	target    int    // The target register for the instruction
	sourceReg string // The first source register of the instruction
}

func NewStr(target int, sourceReg string) *Str {
	return &Str{target, sourceReg}
}

func (instr *Str) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Str) GetSources() []int {
	return nil
}
func (instr *Str) GetImmediate() *int {
	return nil
}
func (instr *Str) GetLabel() string {
	// does not work with labels so we can just return a default value
	return ""
}
func (instr *Str) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}
func (instr *Str) String() string {

	var out bytes.Buffer
	out.WriteString("strGlobal:\n")
	targetReg := fmt.Sprintf("r%v", instr.target)
	prefix := "@"

	sourceReg := prefix + instr.sourceReg

	out.WriteString(fmt.Sprintf("\tstr %s,%s", targetReg, sourceReg))

	return out.String()

}

func (instr *Str) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
