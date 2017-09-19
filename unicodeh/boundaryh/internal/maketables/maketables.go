package main

import (
	"fmt"
	"github.com/apaxa-go/helper/unicodeh"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"go/format"
	"golang.org/x/text/unicode/rangetable"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	cr   rune = '\u000D'
	lf   rune = '\u000A'
	zwnj rune = '\u200c'
	zwj  rune = '\u200d'
)

func genSpacingMark() *unicode.RangeTable {
	r := genSuffix()
	r = rangetableh.DeleteRunes(r, rangetableh.Runes(unicodeh.GraphemeExtendYes)...)
	return r
}

func genSuffix() *unicode.RangeTable {
	/*
		Grapheme_Extend | SpacingMark:
		=
			Grapheme_Extend = Yes
		+
			Grapheme_Cluster_Break ≠ Extend, and
			General_Category = Spacing_Mark, or
			any of the following (which have General_Category = Other_Letter):
			U+0E33 ( ำ ) THAI CHARACTER SARA AM
			U+0EB3 ( ຳ ) LAO VOWEL SIGN AM

			Exceptions: The following (which have General_Category = Spacing_Mark and would otherwise be included) are specifically excluded:
			U+102B ( ါ ) MYANMAR VOWEL SIGN TALL AA
			U+102C ( ာ ) MYANMAR VOWEL SIGN AA
			U+1038 ( း ) MYANMAR SIGN VISARGA
			U+1062 ( ၢ ) MYANMAR VOWEL SIGN SGAW KAREN EU
			..U+1064 ( ၤ ) MYANMAR TONE MARK SGAW KAREN KE PHO
			U+1067 ( ၧ ) MYANMAR VOWEL SIGN WESTERN PWO KAREN EU
			..U+106D ( ၭ ) MYANMAR SIGN WESTERN PWO KAREN TONE-5
			U+1083 ( ႃ ) MYANMAR VOWEL SIGN SHAN AA
			U+1087 ( ႇ ) MYANMAR SIGN SHAN TONE-2
			..U+108C ( ႌ ) MYANMAR SIGN SHAN COUNCIL TONE-3
			U+108F ( ႏ ) MYANMAR SIGN RUMAI PALAUNG TONE-5
			U+109A ( ႚ ) MYANMAR SIGN KHAMTI TONE-1
			..U+109C ( ႜ ) MYANMAR VOWEL SIGN AITON A
			U+1A61 ( ᩡ ) TAI THAM VOWEL SIGN A
			U+1A63 ( ᩣ ) TAI THAM VOWEL SIGN AA
			U+1A64 ( ᩤ ) TAI THAM VOWEL SIGN TALL AA
			U+AA7B ( ꩻ ) MYANMAR SIGN PAO KAREN TONE
			U+AA7D ( ꩽ ) MYANMAR SIGN TAI LAING TONE-5
			U+11720 ( 𑜠 ) AHOM VOWEL SIGN A
			U+11721 ( 𑜡 ) AHOM VOWEL SIGN AA

	*/
	includeRunes := []rune{'\u0e33', '\u0eb3', '\u200d'} // 'u200d' is ZWJ and looked missing in regexp.
	excludeRunes := []rune{'\u102b', '\u102c', '\u1038', '\u1062', '\u1064', '\u1067', '\u106d', '\u1083', '\u1087', '\u108c', '\u108f', '\u109a', '\u109c', '\u1a61', '\u1a63', '\u1a64', '\uaa7b', '\uaa7d', '\U00011720', '\U00011721'}
	return rangetable.Merge(unicodeh.GraphemeClusterBreakExtend /*ExtendYes*/, rangetableh.DeleteRunes(unicodeh.GeneralCategorySpacingMark, excludeRunes...), rangetable.New(includeRunes...))
}

func genPrepend() *unicode.RangeTable {
	/*
		Indic_Syllabic_Category = Consonant_Preceding_Repha, or
		Indic_Syllabic_Category = Consonant_Prefixed, or
		Prepended_Concatenation_Mark = Yes
	*/
	return rangetable.Merge(unicodeh.IndicSyllabicCategoryConsonantPrecedingRepha, unicodeh.IndicSyllabicCategoryConsonantPrefixed, unicodeh.PrependedConcatenationMarkYes)
}

func genControl() *unicode.RangeTable {
	/*
		General_Category = Line_Separator, or
		General_Category = Paragraph_Separator, or
		General_Category = Control, or
		General_Category = Unassigned and Default_Ignorable_Code_Point, or	// TODO strange construction here, not sure that implement this correctly
		General_Category = Surrogate, or
		General_Category = Format
		and not U+000D CARRIAGE RETURN
		and not U+000A LINE FEED
		and not U+200C ZERO WIDTH NON-JOINER (ZWNJ)
		and not U+200D ZERO WIDTH JOINER (ZWJ)
	*/
	excludeRunes := []rune{ /*cr, lf,*/ zwnj, zwj} // Ignore cr & ld !!!
	return rangetable.Merge(
		rangetableh.DeleteRunes(
			rangetable.Merge(
				unicodeh.GeneralCategoryLineSeparator,
				unicodeh.GeneralCategoryParagraphSeparator,
				unicodeh.GeneralCategoryControl,
				rangetableh.Intersect(unicodeh.GeneralCategoryUnassigned, unicodeh.DefaultIgnorableCodePointYes),
				unicodeh.GeneralCategorySurrogate,
				unicodeh.GeneralCategoryFormat,
			),
			excludeRunes...))
}

const (
	comment   = "//"
	prefix    = "maketables"
	delim     = ":"
	generated = "generated-file"

	maxFileSizeForOverwrite = 1024 * 1024 // 1 MB
	targetFn                = "tables-gen.go"
	packageName             = "grapheme"
)

// TODO this function is similar to some in generator, so move them to some lib
func isOverwriteSafe(fn string) bool {
	if stat, err := os.Stat(fn); os.IsNotExist(err) {
		return true
	} else if !stat.Mode().IsRegular() {
		return false
	} else if stat.Size() > maxFileSizeForOverwrite {
		return false
	} else if stat.Size() == 0 {
		return true
	}

	tmp, err := ioutil.ReadFile(fn)
	if err != nil {
		return false
	}

	const lookFor = comment + prefix + delim + generated
	return strings.Index(string(tmp), lookFor) != -1
}

func main() {
	tables := map[string]*unicode.RangeTable{}
	tables["suffixTable"] = genSuffix()
	tables["prependTable"] = genPrepend()
	tables["controlTable"] = genControl()
	tables["spacingMarkTable"] = genSpacingMark()

	if !isOverwriteSafe(targetFn) {
		panic("Target file " + targetFn + " : it is not safe to overwrite it")
	}

	data := comment + prefix + delim + generated + "\n\n"
	data += "package " + packageName + "\n\n"
	data += "import \"unicode\"\n\n"
	for n, t := range tables {
		data += fmt.Sprintf("var %s = %#v\n\n", n, t)
	}

	// Format output
	if fData, err := format.Source([]byte(data)); err == nil {
		data = string(fData)
	} else {
		log.Print("Unable to format result source file: ", err)
	}

	if err := ioutil.WriteFile(targetFn, []byte(data), 0777); err != nil {
		panic("Unable to write " + targetFn + " : " + err.Error())
	}
}
