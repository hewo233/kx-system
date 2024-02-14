package Code

func Init() {
	Ans = []string{}
	CallCount = 0
	SegMap = map[int]string{
		1: "LCL",
		2: "ARG",
		3: "THIS",
		4: "THAT",
	}
}
