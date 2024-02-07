package Parser

import (
	"fmt"
	"strconv"
	"strings"
)

var Instruction []Model

func Remove(s string) string {
	// remove space and //
	index := strings.Index(s, "//")
	if index != -1 {
		s = s[:index]
	}
	return s
}

func Parse(s string) {

	s = Remove(s)
	if s == "" {
		return
	}

	words := strings.Split(s, " ")

	var x Model

	if len(words) == 1 {
		// Add or Sub
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
	} else {
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
	Instruction = append(Instruction, x)
}
