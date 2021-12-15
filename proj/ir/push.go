package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
	"strconv"
)

type Push struct {
	sourceRegs []int
}

func NewPush(sources []int) *Push {
	return &Push{sources}
}

func (instr *Push) GetTargets() []int {
	return nil
}
func (instr *Push) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceRegs...)
	return sources
}
func (instr *Push) GetImmediate() *int {
	return nil
}
func (instr *Push) GetLabel() string {
	return ""
}
func (instr *Push) SetLabel(newLabel string) {
}
func (instr *Push) String() string {

	var out bytes.Buffer
	sources := ""
	for _, source := range instr.sourceRegs {
		sources = sources + "r" + strconv.Itoa(source) + ","
	}
	sources = sources[:len(sources)-1]
	out.WriteString(fmt.Sprintf("\tpush {%s}", sources))
	return out.String()

}

func (instr *Push) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	var align int
	if len(instr.sourceRegs)%2 == 0 {
		align = 16 * len(instr.sourceRegs)
	} else {
		align = 16 * (len(instr.sourceRegs) + 1)
	}
	for _, arg := range frame.Arguments {
		output = append(output, fmt.Sprintf("\t\tstr x%d, [x29, #%d]", arg.Register, arg.Offset))
	}
	output = append(output, fmt.Sprintf("\t\tsub sp, sp, #%d", align))
	i := 0 // from x0 to ...
	for _, reg := range instr.sourceRegs {
		output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]", i+len(instr.sourceRegs), frame.FindVarOffset(reg)))
		output = append(output, fmt.Sprintf("\t\tmov x%d, x%d", i, i+len(instr.sourceRegs)))
		i += 1
	}
	return output
}
