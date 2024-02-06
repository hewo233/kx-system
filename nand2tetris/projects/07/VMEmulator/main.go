package main

import (
	"VMEmulator/Code"
	"VMEmulator/IO"
)

func main() {
	s := IO.ReadFile()
	Code.Pass(s)
	IO.WriteFile()
}
