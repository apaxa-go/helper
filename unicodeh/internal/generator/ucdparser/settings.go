package ucdparser

import "os"

var DefaultParseDetails = []ParseDetails{
	{"auxiliary" + string(os.PathSeparator) + "GraphemeBreakProperty.txt", -1, "Grapheme_Cluster_Break", 1, "", 0, "Other"},
	{"auxiliary" + string(os.PathSeparator) + "SentenceBreakProperty.txt", -1, "Sentence_Break", 1, "", 0, "Other"},
	{"auxiliary" + string(os.PathSeparator) + "WordBreakProperty.txt", -1, "Word_Break", 1, "", 0, "Other"},

	{"extracted" + string(os.PathSeparator) + "DerivedBidiClass.txt", -1, "Bidi_Class", 1, "", 0, "Left_To_Right"}, // TODO Missing is more complicated?
	{"extracted" + string(os.PathSeparator) + "DerivedBinaryProperties.txt", 1, "", -1, "Y", 0, "N"},
	{"extracted" + string(os.PathSeparator) + "DerivedCombiningClass.txt", -1, "Canonical_Combining_Class", 1, "", 0, "Not_Reordered"},
	{"extracted" + string(os.PathSeparator) + "DerivedDecompositionType.txt", -1, "Decomposition_Type", 1, "", 0, "None"},
	{"extracted" + string(os.PathSeparator) + "DerivedEastAsianWidth.txt", -1, "East_Asian_Width", 1, "", 0, "Neutral"},
	{"extracted" + string(os.PathSeparator) + "DerivedGeneralCategory.txt", -1, "General_Category", 1, "", 0, "Unassigned"},
	{"extracted" + string(os.PathSeparator) + "DerivedJoiningGroup.txt", -1, "Joining_Group", 1, "", 0, "No_Joining_Group"},
	{"extracted" + string(os.PathSeparator) + "DerivedJoiningType.txt", -1, "Joining_Type", 1, "", 0, "Non_Joining"},
	{"extracted" + string(os.PathSeparator) + "DerivedLineBreak.txt", -1, "Line_Break", 1, "", 0, "Unknown"},
	{"extracted" + string(os.PathSeparator) + "DerivedNumericType.txt", -1, "Numeric_Type", 1, "", 0, "None"},

	{"BidiBrackets.txt", -1, "Bidi_Paired_Bracket_Type", 2, "", 0, "n"},
	{"CompositionExclusions.txt", -1, "Composition_Exclusion", -1, "Y", 0, "N"},
	{"HangulSyllableType.txt", -1, "Hangul_Syllable_Type", 1, "", 0, "Not_Applicable"},
	{"Scripts.txt", -1, "Script", 1, "", 0, "Unknown"},
	{"Blocks.txt", -1, "Block", 1, "", 0, "No_Block"},
	{"DerivedAge.txt", -1, "Age", 1, "", 0, "Unassigned"},
	{"IndicPositionalCategory.txt", -1, "Indic_Positional_Category", 1, "", 0, "NA"},
	{"IndicSyllabicCategory.txt", -1, "Indic_Syllabic_Category", 1, "", 0, "Other"},
	{"Jamo.txt", -1, "Jamo_Short_Name", 1, "", 0, "none"},
	{"DerivedCoreProperties.txt", 1, "", -1, "Y", 0, "Y"},
	{"PropList.txt", 1, "", -1, "Y", 0, "Y"},
	{"VerticalOrientation.txt", -1, "Vertical_Orientation", 1, "", 0, "R"}, // TODO Missing is more complicated?
	{"DerivedNormalizationProps.txt", 1, "NFC_QC", 2, "", 1, "Yes"},
	{"DerivedNormalizationProps.txt", 1, "NFD_QC", 2, "", 1, "Yes"},
	{"DerivedNormalizationProps.txt", 1, "NFKC_QC", 2, "", 1, "Yes"},
	{"DerivedNormalizationProps.txt", 1, "NFKD_QC", 2, "", 1, "Yes"},
	{"DerivedNormalizationProps.txt", 1, "Full_Composition_Exclusion", -1, "Y", 1, "No"},
	{"DerivedNormalizationProps.txt", 1, "Changes_When_NFKC_Casefolded", -1, "Y", 1, "No"},
}

var DefaultDeprecatedProperies = []string{
	"FC_NFKC_Closure",
	"Expands_On_NFD",
	"Expands_On_NFC",
	"Expands_On_NFKD",
	"Expands_On_NFKC",
}

// map[<Property name in any form>][<index of additional value, does not matter>][<0-short name, 1-long name, 2..-other known names>]
var DefaultAdditionalValues = map[string][][]string{
	"Jamo_Short_Name": {{"", "None", "<none>"}},
}

// map[<Property name in any form>][<Pseudo value name in any form>]=slice of unioning values names
// From PropertyValueAliases.txt (look for "|")
var DefaultPseudoValues = map[string]map[string][]string{
	"General_Category": {
		"C":  {"Cc", "Cf", "Cn", "Co", "Cs"},
		"L":  {"Ll", "Lm", "Lo", "Lt", "Lu"},
		"LC": {"Ll", "Lt", "Lu"},
		"M":  {"Mc", "Me", "Mn"},
		"N":  {"Nd", "Nl", "No"},
		"P":  {"Pc", "Pd", "Pe", "Pf", "Pi", "Po", "Ps"},
		"S":  {"Sc", "Sk", "Sm", "So"},
		"Z":  {"Zl", "Zp", "Zs"},
	},
}

// Really empty values.
// List it here to avoid it in console log.
// map[<Property name in any form>]=slice of values names
var DefaultReallyEmptyValues = map[string][]string{
	"Script":                    {"Katakana_Or_Hiragana"},
	"Canonical_Combining_Class": {"CCC133", "Attached_Below_Left"},
}
