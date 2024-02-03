package Parser

// Model
//@a
//M=A+1;JMP
type Model struct {
	Type        int    // 0: A-instruction, 1: C-instruction, 2: Define instruction
	Address     string // A
	Destination string // C
	Alu         string // C
	Jump        string // C
	Name        string // Define instruction or A instruction
}

var Instruction []Model
