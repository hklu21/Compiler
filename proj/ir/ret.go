package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
	"strconv"
)

type Ret struct {
	source string
}

func NewRet(source string) *Ret {
	return &Ret{source}
}

func (instr *Ret) GetTargets() []int {
	return nil
}
func (instr *Ret) GetSources() []int {
	return nil
}
func (instr *Ret) GetImmediate() *int {
	return nil
}
func (instr *Ret) GetLabel() string {
	return ""
}
func (instr *Ret) SetLabel(newLabel string) {
}
func (instr *Ret) String() string {

	var out bytes.Buffer
	if instr.source == "void" {
		out.WriteString(fmt.Sprintf("\tret %s", instr.source))
	} else {
		out.WriteString(fmt.Sprintf("\tret r%s", instr.source))
	}
	return out.String()

}

func (instr *Ret) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	if instr.source != "void" {
		i, _ := strconv.Atoi(instr.source)
		output = append(output, fmt.Sprintf("\t\tldr x1,[x29,#%d]", frame.FindVarOffset(i)))
		output = append(output, "\t\tmov x0,x1")
	}
	output = append(output, fmt.Sprintf("\t\tadd sp,sp,#%d", frame.CalcOffsetHex()))
	output = append(output, "\t\tldp x29,x30,[sp]")
	output = append(output, "\t\tadd sp,sp, 16")
	output = append(output, "\t\tret")

	return output
}
