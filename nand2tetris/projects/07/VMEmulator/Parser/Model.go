package Parser

type Model struct {
	InstructionType int // 0 is push ,1 is pop , 2 is plus , 3 is sub
	AddressType     int // 1 is local, 2 is argument, 3 is this, 4 is that, 5 is temp, 6 is static, 7 is pointer, 8 is constant
	Num             int
}
