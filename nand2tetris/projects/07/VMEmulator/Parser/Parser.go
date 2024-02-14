package Parser

import (
	"fmt"
	"strconv"
	"strings"
)

var Instruction []Model

func Remove(s string) string {
	// remove space and //
	s = strings.TrimSpace(s)
	index := strings.Index(s, "//")
	if index != -1 {
		s = s[:index]
	}
	s = strings.TrimSpace(s)

	return s
}

func PushOrPop(words []string, x *Model) {
	if words[0] == "push" {
		x.InstructionType = 0
	} else if words[0] == "pop" {
		x.InstructionType = 1
	}
	x.AddressType = AddressTypeMap[words[1]]
	newNum, err := strconv.Atoi(words[2])
	if err != nil {
		fmt.Println(err)
	}
	x.Num = newNum
}

func Single(words []string, x *Model) {
	if words[0] == "add" {
		x.InstructionType = 2
	} else if words[0] == "sub" {
		x.InstructionType = 3
	} else if words[0] == "neg" {
		x.InstructionType = 4
	} else if words[0] == "eq" {
		x.InstructionType = 5
	} else if words[0] == "gt" {
		x.InstructionType = 6
	} else if words[0] == "lt" {
		x.InstructionType = 7
	} else if words[0] == "and" {
		x.InstructionType = 8
	} else if words[0] == "or" {
		x.InstructionType = 9
	} else if words[0] == "not" {
		x.InstructionType = 10
	}
}

func GotoS(words []string, x *Model) {
	if words[0] == "label" {
		x.InstructionType = 11
		x.Label = words[1]
	} else if words[0] == "goto" {
		x.InstructionType = 12
		x.JumpTo = words[1]
	} else if words[0] == "if-goto" {
		x.InstructionType = 13
		x.JumpTo = words[1]
	}
}

func Parse(s string) {

	s = Remove(s)
	if s == "" {
		return
	}

	words := strings.Split(s, " ")

	var x Model

	if len(words) == 1 && words[0] != "return" {
		Single(words, &x)
	} else {
		if words[0] == "push" || words[0] == "pop" {
			PushOrPop(words, &x)

		} else if words[0] == "label" || words[0] == "goto" || words[0] == "if-goto" {
			GotoS(words, &x)

		} else {
			if words[0] == "function" {
				x.InstructionType = 14
				x.Label = words[1]

				newNum, err := strconv.Atoi(words[2])
				if err != nil {
					fmt.Println(err)
				}

				x.Num = newNum

			} else if words[0] == "call" {
				x.InstructionType = 15
				x.Label = words[1]

				newNum, err := strconv.Atoi(words[2])
				if err != nil {
					fmt.Println(err)
				}

				x.Num = newNum

			} else if words[0] == "return" {
				x.InstructionType = 16
			}
		}
	}
	Instruction = append(Instruction, x)
}
