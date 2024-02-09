package Code

import (
	"VMEmulator/Parser"
	"strconv"
)

var Ans []string

func Push() {
	Ans = append(Ans, "@SP")
	Ans = append(Ans, "A=M")
	Ans = append(Ans, "M=D")
	Ans = append(Ans, "@SP")
	Ans = append(Ans, "M=M+1")
}

func Pop() {
	Ans = append(Ans, "@R13")
	Ans = append(Ans, "M=D")
	Ans = append(Ans, "@SP")
	Ans = append(Ans, "AM=M-1")
	Ans = append(Ans, "D=M")
	Ans = append(Ans, "@R13")
	Ans = append(Ans, "A=M")
	Ans = append(Ans, "M=D")
}

func AddOrSub() {
	Ans = append(Ans, "@SP")
	Ans = append(Ans, "AM=M-1")
	Ans = append(Ans, "D=M")
	Ans = append(Ans, "A=A-1")
}

func BigOrSmall() {
	Ans = append(Ans, "@SP")
	Ans = append(Ans, "AM=M-1")
	Ans = append(Ans, "D=M")
	Ans = append(Ans, "A=A-1")
	Ans = append(Ans, "D=M-D")
	Ans = append(Ans, "M=-1")
}

func AndOrOr() {
	Ans = append(Ans, "@SP")
	Ans = append(Ans, "AM=M-1")
	Ans = append(Ans, "D=M")
	Ans = append(Ans, "A=A-1")
}

func PushAll(ins Parser.Model, FileName string) {
	if ins.AddressType <= 4 {
		// Local, Argument, This, That
		SegType := SegMap[ins.AddressType]
		Ans = append(Ans, "@"+SegType)
		Ans = append(Ans, "D=M")
		NumString := "@" + strconv.Itoa(ins.Num)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "A=D+A")
		Ans = append(Ans, "D=M")
		Push()

	} else if ins.AddressType == 5 {
		//Temp
		NumString := "@" + strconv.Itoa(ins.Num+5)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "D=M")
		Push()
		Ans = append(Ans, "@SP")
		Ans = append(Ans, "AM=M-1")
		Ans = append(Ans, "D=M")
		Ans = append(Ans, "A=A-1")
		Ans = append(Ans, "D=M-D")
	} else if ins.AddressType == 6 {
		//Static
		NumString := "@" + FileName + "." + strconv.Itoa(ins.Num)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "D=M")
		Push()

	} else if ins.AddressType == 7 {
		if ins.Num == 0 {
			// THIS
			NumString := "@" + "THIS"
			Ans = append(Ans, NumString)
			Ans = append(Ans, "D=M")
			Push()

		} else if ins.Num == 1 {
			// THAT
			NumString := "@" + "THAT"
			Ans = append(Ans, NumString)
			Ans = append(Ans, "D=M")
			Push()

		}
	} else if ins.AddressType == 8 {
		NumString := "@" + strconv.Itoa(ins.Num)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "D=A")
		Push()

	}
}

func PopAll(ins Parser.Model, FileName string) {
	if ins.AddressType <= 4 {
		// Local, Argument, This, That
		NumString := "@" + strconv.Itoa(ins.Num)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "D=A")
		Ans = append(Ans, "@"+SegMap[ins.AddressType])
		Ans = append(Ans, "D=D+M")
		Pop()

	} else if ins.AddressType == 5 {
		// Temp
		NumString := "@" + strconv.Itoa(ins.Num+5)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "D=A")
		Pop()

	} else if ins.AddressType == 6 {
		// Static
		NumString := "@" + FileName + "." + strconv.Itoa(ins.Num)
		Ans = append(Ans, NumString)
		Ans = append(Ans, "D=A")
		Pop()

	} else if ins.AddressType == 7 {
		if ins.Num == 0 {
			// THIS
			NumString := "@" + "THIS"
			Ans = append(Ans, NumString)
			Ans = append(Ans, "D=A")
			Pop()

		} else if ins.Num == 1 {
			// THAT
			NumString := "@" + "THAT"
			Ans = append(Ans, NumString)
			Ans = append(Ans, "D=A")
			Pop()

		}
	}
}

func Deal(ins Parser.Model, FileName string, eqCount, gtCount, ltCount *int) {
	if ins.InstructionType == 0 {
		//Push
		PushAll(ins, FileName)
	} else if ins.InstructionType == 1 {
		// Pop
		PopAll(ins, FileName)
	} else if ins.InstructionType == 2 {
		// Plus
		AddOrSub()
		Ans = append(Ans, "M=M+D")

	} else if ins.InstructionType == 3 {
		// Sub
		AddOrSub()
		Ans = append(Ans, "M=M-D")

	} else if ins.InstructionType == 4 {
		// Neg
		Ans = append(Ans, "@SP")
		Ans = append(Ans, "A=M-1")
		Ans = append(Ans, "M=-M")
	} else if ins.InstructionType == 5 {
		// Eq
		BigOrSmall()
		Ans = append(Ans, "@EQ_TRUE_"+FileName+"_"+strconv.Itoa(*eqCount))
		Ans = append(Ans, "D;JEQ")
		Ans = append(Ans, "@SP")
		Ans = append(Ans, "A=M-1")
		Ans = append(Ans, "M=0")
		Ans = append(Ans, "(EQ_TRUE_"+FileName+"_"+strconv.Itoa(*eqCount)+")")
		*eqCount++
	} else if ins.InstructionType == 6 {
		// Gt
		BigOrSmall()
		Ans = append(Ans, "@GT_TRUE_"+FileName+"_"+strconv.Itoa(*gtCount))
		Ans = append(Ans, "D;JGT")
		Ans = append(Ans, "@SP")
		Ans = append(Ans, "A=M-1")
		Ans = append(Ans, "M=0")
		Ans = append(Ans, "(GT_TRUE_"+FileName+"_"+strconv.Itoa(*gtCount)+")")
		*gtCount++
	} else if ins.InstructionType == 7 {
		// Lt
		BigOrSmall()
		Ans = append(Ans, "@LT_TRUE_"+FileName+"_"+strconv.Itoa(*ltCount))
		Ans = append(Ans, "D;JLT")
		Ans = append(Ans, "@SP")
		Ans = append(Ans, "A=M-1")
		Ans = append(Ans, "M=0")
		Ans = append(Ans, "(LT_TRUE_"+FileName+"_"+strconv.Itoa(*ltCount)+")")
		*ltCount++
	} else if ins.InstructionType == 8 {
		//And
		AndOrOr()
		Ans = append(Ans, "M=M&D")
	} else if ins.InstructionType == 9 {
		//Or
		AndOrOr()
		Ans = append(Ans, "M=M|D")
	} else if ins.InstructionType == 10 {
		//Not
		Ans = append(Ans, "@SP")
		Ans = append(Ans, "A=M-1")
		Ans = append(Ans, "M=!M")
	}
}

func Pass(FileName string) {

	var eqCount, gtCount, ltCount = 0, 0, 0

	for _, ins := range Parser.Instruction {
		Deal(ins, FileName, &eqCount, &gtCount, &ltCount)
	}
}
