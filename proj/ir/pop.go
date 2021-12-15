package ir

import (
	"bytes"
	"fmt"
	"proj/codegen"
	"strconv"
)

type Pop struct {
	sourceRegs []int
}

func NewPop(sources []int) *Pop {
	return &Pop{sources}
}

func (instr *Pop) GetTargets() []int {
	return nil
}
func (instr *Pop) GetSources() []int {
	sources := make([]int, 0)
	sources = append(sources, instr.sourceRegs...)
	return sources
}
func (instr *Pop) GetImmediate() *int {
	return nil
}
func (instr *Pop) GetLabel() string {
	return ""
}
func (instr *Pop) SetLabel(newLabel string) {
}
func (instr *Pop) String() string {

	var out bytes.Buffer
	sources := ""
	if len(instr.sourceRegs) == 0 {
		out.WriteString(fmt.Sprintf("\tpop {%s}", sources))
		return out.String()
	}
	for _, source := range instr.sourceRegs {
		sources = sources + "r" + strconv.Itoa(source) + ","
	}
	out.WriteString(fmt.Sprintf("\tpop {%s}", sources[:len(sources)-1]))
	return out.String()

}

func (instr *Pop) TranslateToArm(frame *codegen.Frame) []string {
	output := make([]string, 0)
	output = append(output, "\t\tmov x1,x0")
	output = append(output, fmt.Sprintf("\t\tstr x1,[x29,#%d]", frame.FindVarOffset(instr.sourceRegs[0])))
	var align int
	if len(instr.sourceRegs)%2 == 0 {
		align = 16 * len(instr.sourceRegs)
	} else {
		align = 16 * (len(instr.sourceRegs) + 1)
	}
	for _, arg := range frame.Arguments {
		output = append(output, fmt.Sprintf("\t\tldr x%d, [x29, #%d]", arg.Register, arg.Offset))
	}
	output = append(output, fmt.Sprintf("\t\tadd sp, sp, #%d", align))
	/*
				var align int
				if len(instr.sourceRegs)%2 == 0 {
					align = 16 * len(instr.sourceRegs)
				} else {
					align = 16 * (len(instr.sourceRegs) + 1)
				}

				i := 0 // from x0 to ...
				for _, reg := range instr.sourceRegs {
					output = append(output, fmt.Sprintf("mov x%d, x%d", i+len(instr.sourceRegs), i))
					output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]", i+len(instr.sourceRegs), frame.FindVarOffset(reg)))
					i += 1
				}

		i := 0 // from x0 to ...
		for _, reg := range instr.sourceRegs {
			output = append(output, fmt.Sprintf("\t\tmov x%d, x%d", i+len(instr.sourceRegs), i))
			output = append(output, fmt.Sprintf("\t\tldr x%d,[x29,#%d]", i+len(instr.sourceRegs), frame.FindVarOffset(reg)))
			i += 1
		}

			output = append(output, fmt.Sprintf("\t\tsub sp, sp, #%d", align))
	*/
	return output
}
