package tokenh

import "go/token"

func IsComparisonCompare(t token.Token) bool {
	return t == token.EQL || t == token.NEQ
}

func IsComparisonOrder(t token.Token) bool {
	return t == token.LSS || t == token.LEQ || t == token.GTR || t == token.GEQ
}

func IsComparison(t token.Token) bool {
	return IsComparisonCompare(t) || IsComparisonOrder(t)
}

func IsShift(t token.Token) bool {
	return t == token.SHL || t == token.SHR
}
