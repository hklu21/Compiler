package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
)

type And struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	operand   int       // The operand either register or constant
	opty      OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewAnd(target int, sourceReg int, operand int, opty OperandTy) *And {
	return &And{target, sourceReg, operand, opty}
}

func (instr *And) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *And) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceReg)
	return sources
}
func (instr *And) GetImmediate() *int {

	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}
func (instr *And) GetLabel() string {
	return ""
}
func (instr *And) SetLabel(newLabel string) {
}
func (instr *And) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.sourceReg)

	var prefix string

	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2 := fmt.Sprintf("%v%v", prefix, instr.operand)

	out.WriteString(fmt.Sprintf("\tand %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *And) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
