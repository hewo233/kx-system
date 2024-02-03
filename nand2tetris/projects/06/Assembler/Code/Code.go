package Code

import (
	"Assembler/Parser"
	"Assembler/SymbolTable"
	"strconv"
)

// trans Code to binary

func FirstPass() {
	i := 0
	for _, ins := range Parser.Instruction {
		if ins.Type == 2 {
			SymbolTable.SymbolTable[ins.Name] = i
		} else {
			i++
		}
	}
}

var Ans []string

func SecondPass() {
	for _, ins := range Parser.Instruction {
		if ins.Type == 0 {
			//A-instruction
			var tenx int // 10 Address
			if ins.Address != "" {
				tenx, _ = strconv.Atoi(ins.Address)

			} else {
				tenx = SymbolTable.SymbolTable[ins.Name]
			}
			twox := strconv.FormatInt(int64(tenx), 2)
			for len(twox) < 15 {
				twox = "0" + twox
			}
			Ans = append(Ans, "0"+twox)
		} else if ins.Type == 1 {
			//C-instruction
			Ans = append(Ans, "111"+AluMap[ins.Alu]+DestMap[ins.Destination]+JumpMap[ins.Jump])
		}
	}
}
