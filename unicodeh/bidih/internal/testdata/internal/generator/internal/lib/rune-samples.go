package lib

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"github.com/apaxa-go/helper/unicodeh"
	"fmt"
)

func GenerateRuneSamples(){
	runeSamples := make([]rune, bidi.Count)

	runeSamples[bidi.ArabicLetter] = rangetableh.Lo(unicodeh.BidiClassArabicLetter)
	runeSamples[bidi.ArabicNumber] = rangetableh.Lo(unicodeh.BidiClassArabicNumber)
	runeSamples[bidi.ParagraphSeparator] = rangetableh.Lo(unicodeh.BidiClassParagraphSeparator)
	runeSamples[bidi.BoundaryNeutral] = rangetableh.Lo(unicodeh.BidiClassBoundaryNeutral)
	runeSamples[bidi.CommonSeparator] = rangetableh.Lo(unicodeh.BidiClassCommonSeparator)
	runeSamples[bidi.EuropeanNumber] = rangetableh.Lo(unicodeh.BidiClassEuropeanNumber)
	runeSamples[bidi.EuropeanSeparator] = rangetableh.Lo(unicodeh.BidiClassEuropeanSeparator)
	runeSamples[bidi.EuropeanTerminator] = rangetableh.Lo(unicodeh.BidiClassEuropeanTerminator)
	runeSamples[bidi.FirstStrongIsolate] = rangetableh.Lo(unicodeh.BidiClassFirstStrongIsolate)
	runeSamples[bidi.LeftToRight] = rangetableh.Lo(unicodeh.BidiClassLeftToRight)
	runeSamples[bidi.LeftToRightEmbedding] = rangetableh.Lo(unicodeh.BidiClassLeftToRightEmbedding)
	runeSamples[bidi.LeftToRightIsolate] = rangetableh.Lo(unicodeh.BidiClassLeftToRightIsolate)
	runeSamples[bidi.LeftToRightOverride] = rangetableh.Lo(unicodeh.BidiClassLeftToRightOverride)
	runeSamples[bidi.NonSpacingMark] = rangetableh.Lo(unicodeh.BidiClassNonspacingMark)
	runeSamples[bidi.OtherNeutral] = rangetableh.Lo(unicodeh.BidiClassOtherNeutral)
	runeSamples[bidi.PopDirectionalFormat] = rangetableh.Lo(unicodeh.BidiClassPopDirectionalFormat)
	runeSamples[bidi.PopDirectionalIsolate] = rangetableh.Lo(unicodeh.BidiClassPopDirectionalIsolate)
	runeSamples[bidi.RightToLeft] = rangetableh.Lo(unicodeh.BidiClassRightToLeft)
	runeSamples[bidi.RightToLeftEmbedding] = rangetableh.Lo(unicodeh.BidiClassRightToLeftEmbedding)
	runeSamples[bidi.RightToLeftIsolate] = rangetableh.Lo(unicodeh.BidiClassRightToLeftIsolate)
	runeSamples[bidi.RightToLeftOverride] = rangetableh.Lo(unicodeh.BidiClassRightToLeftOverride)
	runeSamples[bidi.SegmentSeparator] = rangetableh.Lo(unicodeh.BidiClassSegmentSeparator)
	runeSamples[bidi.WhiteSpace] = rangetableh.Lo(unicodeh.BidiClassWhiteSpace)

	data:="// Sample runes for each BIDI class.\n"
	data+="// Generated.\n"
	data+=fmt.Sprintf("var runeSamples=%#v\n",runeSamples)

	saveFile("bidi-tests-runes-gen.go",pkgName,nil,[]byte(data))
}
