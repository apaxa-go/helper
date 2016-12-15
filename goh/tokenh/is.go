package tokenh

import "go/token"

func IsComparison(t token.Token) bool {
	return t == token.EQL || t == token.NEQ || t == token.LSS || t == token.LEQ || t == token.GTR || t == token.GEQ
}

func IsShift(t token.Token) bool {
	return t == token.SHL || t == token.SHR
}
