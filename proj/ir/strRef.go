package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type StrRef struct {
	target    int // The target register for the instruction
	sourceReg int // The first source register of the instruction
	feild     string
}

func NewStrRef(target int, sourceReg int, feild string) *LoadRef {
	return &LoadRef{target, sourceReg, feild}
}

func (instr *StrRef) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *StrRef) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *StrRef) GetImmediate() *int {
	return nil
}
func (instr *StrRef) GetLabel() string {
	return ""
}
func (instr *StrRef) SetLabel(newLabel string) {
}
func (instr *StrRef) String() string {

	var out bytes.Buffer
	out.WriteString("StrRef:\n")
	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.sourceReg)

	prefix := "@"
	operand2 := fmt.Sprintf("%v%v", prefix, instr.feild)

	out.WriteString(fmt.Sprintf("\tstrRef %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *StrRef) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
