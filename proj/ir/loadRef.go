package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type LoadRef struct {
	target    int // The target register for the instruction
	sourceReg int // The first source register of the instruction
	feild     string
}

func NewLoadRef(target int, sourceReg int, feild string) *LoadRef {
	return &LoadRef{target, sourceReg, feild}
}

func (instr *LoadRef) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *LoadRef) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *LoadRef) GetImmediate() *int {
	return nil
}
func (instr *LoadRef) GetLabel() string {
	return ""
}
func (instr *LoadRef) SetLabel(newLabel string) {
}
func (instr *LoadRef) String() string {

	var out bytes.Buffer
	out.WriteString("LoadRef:\n")
	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.sourceReg)

	prefix := "@"
	operand2 := fmt.Sprintf("%v%v", prefix, instr.feild)

	out.WriteString(fmt.Sprintf("\tloadRef %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *LoadRef) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
