package codegen

import (
	"os"
)

type Frame struct {
	FrameName   string
	Arguments   []Argument
	LocalVar    []LocalVar
	LocalOffset int //len(Arguments) - 1
	TempVar     []TempVar
	Return      string
}

func (frame *Frame) FindVarOffset(register int) int {

	for _, arg := range frame.Arguments {
		if arg.Register == register {
			return arg.Offset
		}
	}
	for _, local := range frame.LocalVar {
		if local.Register == register {
			return local.Offset
		}
	}
	for _, temp := range frame.TempVar {
		if temp.Register == register {
			return temp.Offset
		}
	}
	// Shouldn't reach -1
	return -1
}

func (frame *Frame) CalcOffsetHex() int {
	sum := len(frame.Arguments) + len(frame.LocalVar) + len(frame.TempVar)
	if sum%2 == 0 {
		return sum * 8
	} else {
		return (sum + 1) * 8
	}
}

func (frame *Frame) IsArgument(register int) bool {
	for _, arg := range frame.Arguments {
		if arg.Register == register {
			return true
		}
	}
	return false
}

func GenerateFile(instructions []string, filename string) {
	f, _ := os.Create(filename + ".s")
	f.WriteString("\t\t.arch armv8-a\n\t\t.text\n")
	for _, instr := range instructions {
		f.WriteString(instr)
		f.WriteString("\n")
	}
	f.WriteString("\n.PRINT_LN:\n\t\t.asciz \"%ld\\n\"\n\t\t.size .PRINT_LN,5\n")
	f.WriteString("\n.PRINT:\n\t\t.asciz \"%ld\"\n\t\t.size .PRINT,4\n")
	f.WriteString("\n.READ:\n\t\t.asciz \"%ld\"\n\t\t.size .READ,4\n")
	f.Close()
}

type Argument struct {
	ArgName  string
	Register int
	Offset   int
}

type LocalVar struct {
	VarName  string
	Register int
	Offset   int
}

type TempVar struct {
	Register int
	Offset   int
}

type Instruction struct {
}
