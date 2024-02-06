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

func Pass(FileName string) {
	for _, ins := range Parser.Instruction {
		if ins.InstructionType == 0 {
			//Push
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
		} else if ins.InstructionType == 1 {
			// Pop
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
		} else if ins.InstructionType == 2 {
			// Plus
			AddOrSub()
			Ans = append(Ans, "M=M+D")

		} else if ins.InstructionType == 3 {
			// Sub
			AddOrSub()
			Ans = append(Ans, "M=M-D")

		}

	}
}
