package Parser

var AddressTypeMap = map[string]int{
	// 1 is local, 2 is argument, 3 is this, 4 is that, 5 is temp, 6 is static, 7 is pointer, 8 is constant
	"local":    1,
	"argument": 2,
	"this":     3,
	"that":     4,
	"temp":     5,
	"static":   6,
	"pointer":  7,
	"constant": 8,
}
