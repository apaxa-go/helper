package boundaryh

import "github.com/apaxa-go/helper/unicodeh"

type sClass uint8

const (
	sClassOther     sClass = iota
	sClassSep       sClass = iota
	sClassATerm     sClass = iota
	sClassUpper     sClass = iota
	sClassLower     sClass = iota
	sClassExtend    sClass = iota // TODO merge with Format?
	sClassFormat    sClass = iota // TODO merge with Extend?
	sClassClose     sClass = iota
	sClassSp        sClass = iota
	sClassOLetter   sClass = iota // TODO may be add to it other isSB8? It does not used standalone.
	sClassSTerm     sClass = iota
	sClassSContinue sClass = iota // TODO may be add to it ATerm & STerm? It does not used standalone.
	sClassNumeric   sClass = iota
)

func sGetClass(r rune) sClass {
	switch {
	default:
		return sClassOther
	case unicodeh.IsSentenceBreakSep(r):
		return sClassSep
	case unicodeh.IsSentenceBreakATerm(r):
		return sClassATerm
	case unicodeh.IsSentenceBreakUpper(r):
		return sClassUpper
	case unicodeh.IsSentenceBreakLower(r):
		return sClassLower
	case unicodeh.IsSentenceBreakExtend(r):
		return sClassExtend
	case unicodeh.IsSentenceBreakFormat(r):
		return sClassFormat
	case unicodeh.IsSentenceBreakClose(r):
		return sClassClose
	case unicodeh.IsSentenceBreakSp(r):
		return sClassSp
	case unicodeh.IsSentenceBreakOLetter(r):
		return sClassOLetter
	case unicodeh.IsSentenceBreakSTerm(r):
		return sClassSTerm
	case unicodeh.IsSentenceBreakSContinue(r):
		return sClassSContinue
	case unicodeh.IsSentenceBreakNumeric(r):
		return sClassNumeric
	}
}

func (c sClass) isSkip() bool { return c == sClassExtend || c == sClassFormat } // SB5
func (c sClass) isSB8() bool { // SB8
	return c != sClassOLetter && c != sClassUpper && c != sClassLower && c != sClassSep && c != sClassATerm && c != sClassSTerm
}
