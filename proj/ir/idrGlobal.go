package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
	st "proj/symbolTable"
)

type Ldr struct {
	target    int    // The target register for the instruction
	sourceReg string // The first source register of the instruction
}

func NewLdr(target int, sourceReg string) *Ldr {
	return &Ldr{target, sourceReg}
}

func (instr *Ldr) GetTargets() []int {
	targets := make([]int, 0)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Ldr) GetSources() []int {
	return nil
}
func (instr *Ldr) GetImmediate() *int {
	return nil
}
func (instr *Ldr) GetLabel() string {
	// does not work with labels so we can just return a default value
	return ""
}
func (instr *Ldr) SetLabel(newLabel string) {
	//  does not work with labels can we can skip implementing this method.

}

func (instr *Ldr) TranslateToARM(symTable *st.SymbolTable) []codegen.Instruction {
	return make([]codegen.Instruction, 0)
}

func (instr *Ldr) String() string {

	var out bytes.Buffer

	targetReg := fmt.Sprintf("r%v", instr.target)
	prefix := "@"

	sourceReg := prefix + instr.sourceReg

	out.WriteString(fmt.Sprintf("\tldr %s,%s", targetReg, sourceReg))

	return out.String()

}

func (instr *Ldr) TranslateToArm(*codegen.Frame) []string {
	return make([]string, 0)
}
