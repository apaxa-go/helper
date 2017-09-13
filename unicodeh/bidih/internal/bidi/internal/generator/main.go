package main

import (
	"bytes"
	"github.com/apaxa-go/helper/unicodeh"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
	"go/format"
	"io/ioutil"
	"strconv"
	"unicode"
)

func GetBidiClass(r rune) bidi.Class {
	if unicodeh.IsBidiClassArabicLetter(r) {
		return bidi.ArabicLetter
	}
	if unicodeh.IsBidiClassArabicNumber(r) {
		return bidi.ArabicNumber
	}
	if unicodeh.IsBidiClassParagraphSeparator(r) {
		return bidi.ParagraphSeparator
	}
	if unicodeh.IsBidiClassBoundaryNeutral(r) {
		return bidi.BoundaryNeutral
	}
	if unicodeh.IsBidiClassCommonSeparator(r) {
		return bidi.CommonSeparator
	}
	if unicodeh.IsBidiClassEuropeanNumber(r) {
		return bidi.EuropeanNumber
	}
	if unicodeh.IsBidiClassEuropeanSeparator(r) {
		return bidi.EuropeanSeparator
	}
	if unicodeh.IsBidiClassEuropeanTerminator(r) {
		return bidi.EuropeanTerminator
	}
	if unicodeh.IsBidiClassFirstStrongIsolate(r) {
		return bidi.FirstStrongIsolate
	}
	if unicodeh.IsBidiClassLeftToRight(r) {
		return bidi.LeftToRight
	}
	if unicodeh.IsBidiClassLeftToRightEmbedding(r) {
		return bidi.LeftToRightEmbedding
	}
	if unicodeh.IsBidiClassLeftToRightIsolate(r) {
		return bidi.LeftToRightIsolate
	}
	if unicodeh.IsBidiClassLeftToRightOverride(r) {
		return bidi.LeftToRightOverride
	}
	if unicodeh.IsBidiClassNonspacingMark(r) {
		return bidi.NonSpacingMark
	}
	if unicodeh.IsBidiClassOtherNeutral(r) {
		return bidi.OtherNeutral
	}
	if unicodeh.IsBidiClassPopDirectionalFormat(r) {
		return bidi.PopDirectionalFormat
	}
	if unicodeh.IsBidiClassPopDirectionalIsolate(r) {
		return bidi.PopDirectionalIsolate
	}
	if unicodeh.IsBidiClassRightToLeft(r) {
		return bidi.RightToLeft
	}
	if unicodeh.IsBidiClassRightToLeftEmbedding(r) {
		return bidi.RightToLeftEmbedding
	}
	if unicodeh.IsBidiClassRightToLeftIsolate(r) {
		return bidi.RightToLeftIsolate
	}
	if unicodeh.IsBidiClassRightToLeftOverride(r) {
		return bidi.RightToLeftOverride
	}
	if unicodeh.IsBidiClassSegmentSeparator(r) {
		return bidi.SegmentSeparator
	}
	if unicodeh.IsBidiClassWhiteSpace(r) {
		return bidi.WhiteSpace
	}
	panic("unknown BIDI class for rune")
}

// TODO validate closedBracket>openBracket
func generateClassData() (data *bytes.Buffer) {
	data = bytes.NewBuffer(nil)
	data.WriteString("var data = [...]Class{")
	for r := rune(0); r <= unicode.MaxRune; r++ {
		class := GetBidiClass(r)
		if class >= 2*2*2*2*2 {
			panic("Unable to save class in 5 bits")
		}
		bracketPair, found := unicodeh.BidiPairedBracket[r]
		if !found {
			bracketPair = r
		}
		bracketShift := bracketPair - r + 3
		if bracketShift < 0 || bracketShift > 7 {
			panic("Unable to save brackets in 3 bits: " + strconv.FormatInt(int64(r), 10) + "-" + strconv.FormatInt(int64(bracketPair), 10) + "+3=" + strconv.FormatInt(int64(bracketShift), 10))
		}
		class |= bidi.Class(bracketShift) << 5
		data.WriteString("0x" + strconv.FormatInt(int64(class), 16) + ",")
	}
	data.Truncate(data.Len() - 1) // remove unnecessary ","
	data.WriteString("}")
	return
}

func main() {
	const targetFN = "data-gen.go"
	data := []byte("package bidi\n\n" + generateClassData().String() + "\n")
	data, err := format.Source(data)
	if err != nil {
		panic("Syntax error: " + err.Error())
	}
	if err := ioutil.WriteFile(targetFN, data, 0); err != nil {
		panic(err)
	}
}
