package main

import (
	"Assembler/Code"
	"Assembler/IO"
)

func main() {
	IO.ReadFile()
	Code.FirstPass()
	Code.SecondPass()
	IO.WriteFile()
}
