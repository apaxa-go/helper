package asth

import (
	"unicode"
	"unicode/utf8"
)

// req:
// 	-1: s should be valid not exported identifier
//	 0: s should be valid identifier
// 	 1: s should be valid exported identifier
func validateIdent(s string, req int) bool {
	r, i := utf8.DecodeRuneInString(s)
	switch req {
	case -1:
		if !unicode.IsLower(r) && r != '_' {
			return false
		}
	case 0:
		if !unicode.IsLetter(r) && r != '_' {
			return false
		}
	case 1:
		if !unicode.IsUpper(r) {
			return false
		}
	default:
		panic("unknown requirements")
	}
	for _, r = range s[i:] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return true
}

func IsValidIdent(s string) bool            { return validateIdent(s, 0) }
func IsValidExportedIdent(s string) bool    { return validateIdent(s, 1) }
func IsValidNotExportedIdent(s string) bool { return validateIdent(s, -1) }
