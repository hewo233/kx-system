package Parser

type Model struct {
	// 0 is push ,1 is pop , 2 is plus , 3 is sub, 4 is neg , 5 is eq, 6 is gt, 7 is lt, 8 is and, 9 is or, 10 is not
	// 11 is label, 12 is goto, 13 is if-goto, 14 is function, 15 is call, 16 is return
	InstructionType int
	AddressType     int // 1 is local, 2 is argument, 3 is this, 4 is that, 5 is temp, 6 is static, 7 is pointer, 8 is constant
	Num             int
	Label           string
	JumpTo          string
}
