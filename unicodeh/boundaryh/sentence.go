package boundaryh

import (
	"github.com/apaxa-go/helper/unicodeh"
)

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

// TODO
func getSClass(r rune) sClass {
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

func (c sClass) isSkip() bool    { return c == sClassExtend || c == sClassFormat } // SB5
func (c sClass) isSB8() bool { // SB8
	return c != sClassOLetter && c != sClassUpper && c != sClassLower && c != sClassSep && c != sClassATerm && c != sClassSTerm
}

// Returns sClass for first rune and pos where to continue.
// It skips runes as described in SB5.
// It is also map <CR>, <LF> and <CRLF> to Sep class.
// <CRLF> parsed at once, so no special cases required in other places - just check for Sep class instead of ParaSep macro.
func getSClassSkip(runes []rune) (c sClass, i int) {
	l := len(runes)
	if l == 0 {
		c = sClassOther
		return
	}

	if runes[0] == crRune { // SB4
		if l > 1 && runes[1] == lfRune { // SB3
			return sClassSep, 2
		}
		return sClassSep, 1
	}

	if runes[0] == lfRune { // SB4
		return sClassSep, 1
	}

	c = getSClass(runes[0])

	if c == sClassSep { // SB4
		return sClassSep, 1
	}

	for i = 1; i < len(runes); i++ { // SB5
		if !getSClass(runes[i]).isSkip() {
			break
		}
	}

	return
}

func parseAfterSATerm(runes []rune) (closeSp bool, nextClass sClass, i1, i2 int) {
	l := len(runes)
	for i1 < l {
		if c, deltaI := getSClassSkip(runes[i1:]); c == sClassClose {
			closeSp = true
			i1 += deltaI
		} else {
			break
		}
	}
	for i1 < l {
		if c, deltaI := getSClassSkip(runes[i1:]); c == sClassSp {
			closeSp = true
			i1 += deltaI
		} else {
			break
		}
	}
	nextClass, deltaI := getSClassSkip(runes[i1:])
	i2 = i1 + deltaI
	return
}

// lastClass used as prevClass by parent function. It used only if stop==false and is undefined if stop==true.
//
//  Break?  | <prevClass> ATerm Close* Sp* <lastClass>        | Ruler
//   yes    |             ATerm Close* Sp* ParaSep            | SB11
//   no     |             ATerm -      -   Numeric            | SB6
//   no     | Upper       ATerm -      -   Upper              | SB7
//   no     | Lower       ATerm -      -   Upper              | SB7
//   no     |             ATerm Close* Sp* Lower              | SB8
//   no     |             ATerm Close* Sp* isSB8+       Lower | SB8
//   no     |             ATerm Close* Sp* SContinue          | SB8a
//   no     |             ATerm Close* Sp* ATerm              | SB8a
//   no     |             ATerm Close* Sp* STerm              | SB8a
//   yes    |             ATerm Close* Sp*                    | SB11
func decideAtATerm(prevClass sClass, runes []rune) (stop bool, i int, lastClass sClass) {
	closeSp, lastClass, i1, i := parseAfterSATerm(runes)
	switch {
	case lastClass == sClassSep: // SB11
		stop = true
	case !closeSp && lastClass == sClassNumeric: // SB6
	case (prevClass == sClassUpper || prevClass == sClassLower) && !closeSp && lastClass == sClassUpper: // SB7
	case lastClass == sClassSContinue || lastClass == sClassATerm || lastClass == sClassSTerm: // SB8a
	case lastClass == sClassLower: // SB8 (part 1)
	case lastClass.isSB8(): // SB8 (part 2)
		for i < len(runes) {
			var deltaI int
			lastClass, deltaI = getSClassSkip(runes[i:])
			i += deltaI
			if lastClass == sClassLower {
				return
			}
			if !lastClass.isSB8() {
				break
			}
		}
		stop = true
		i = i1
	default: // SB11
		stop = true
		i = i1
	}
	return
}

// lastClass used as prevClass by parent function. It used only if stop==false and is undefined if stop==true.
//
//  Break?  | STerm Close* Sp* some-type? | Ruler
//   yes    | STerm Close* Sp* ParaSep    | SB11
//   no     | STerm Close* Sp* SContinue  | SB8a
//   no     | STerm Close* Sp* ATerm      | SB8a
//   no     | STerm Close* Sp* STerm      | SB8a
//   yes    | STerm Close* Sp*            | SB11
func decideAtSTerm(runes []rune) (stop bool, i int, lastClass sClass) {
	_, lastClass, i1, i := parseAfterSATerm(runes)
	switch lastClass {
	case sClassSep: // SB11
		stop = true
	case sClassSContinue, sClassATerm, sClassSTerm: // SB8a
	default:
		stop = true // SB11
		i = i1
	}
	return
}

func FirstSentence(runes []rune) int {
	l := len(runes)
	if l <= 1 {
		return l
	}

	prevClass := sClassOther
	// Decide should we break sentence after i-th rune
	for i := 0; i < l; {
		curClass, deltaI := getSClassSkip(runes[i:])
		i += deltaI

		if i >= l { // end of runes
			return l
		}

		switch curClass {
		case sClassSep: // SB4
			return i
		case sClassATerm:
			stop, deltaI, newPrevClass := decideAtATerm(prevClass, runes[i:])
			i += deltaI
			if stop {
				return i
			}
			prevClass = newPrevClass
		case sClassSTerm:
			stop, deltaI, newPrevClass := decideAtSTerm(runes[i:])
			i += deltaI
			if stop {
				return i
			}
			prevClass = newPrevClass
		default:
			prevClass = curClass
		}
	}
	return l
}

func Sentences(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		length := FirstSentence(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}
