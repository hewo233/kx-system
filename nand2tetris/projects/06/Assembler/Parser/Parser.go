package Parser

import "strings"

func Remove(s string) string {
	// remove space and //
	index := strings.Index(s, "//")
	if index != -1 {
		s = s[:index]
	}
	s = strings.Replace(s, " ", "", -1)
	return s
}

func Parse(s string) {
	// parse s
	s = Remove(s)

	if s == "" {
		return
	}

	var x Model

	if s[0] == '@' {
		// A-instruction
		x.Type = 0
		if s[1] >= '0' && s[1] <= '9' {
			x.Address = s[1:]
		} else {
			x.Name = s[1:]
		}
	} else if s[0] == '(' {
		// Define instruction
		x.Type = 2
		x.Name = s[1 : len(s)-1]
	} else {
		// C-instruction
		x.Type = 1

		indexEqual := strings.Index(s, "=")
		indexSemicolon := strings.Index(s, ";")
		if indexEqual != -1 {
			x.Destination = s[:indexEqual]
			if indexSemicolon != -1 {
				x.Alu = s[indexEqual+1 : indexSemicolon]
				x.Jump = s[indexSemicolon+1:]
			} else {
				x.Alu = s[indexEqual+1:]
			}
		} else {
			if indexSemicolon != -1 {
				x.Alu = s[:indexSemicolon]
				x.Jump = s[indexSemicolon+1:]
			} else {
				x.Alu = s
			}
		}
	}
	Instruction = append(Instruction, x)
}
