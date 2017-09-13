package other_generator

import (
	"os"
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"strings"
	"strconv"
	"bytes"
)

// variable names
const (
	varBidiPairedBrackets = "bidiPairedBrackets"
	varBidiMirroringGlyph ="bidiMirroringGlyph"
)

// data sources
const(
 	srcBidiPairedBrackets = "BidiBrackets.txt"
	srcBidiMirroringGlyph ="BidiMirroring.txt"
)

func GenerateOther(srcDir string)[]byte{
	data :=bytes.NewBuffer(nil)
	generateBidiBracketsPublic(data)
	generateBidiMirroringGlyphPublic(data)
	generateBidiBrackets(srcDir, data)
	generateBidiMirroringGlyph(srcDir,data)
	return data.Bytes()
}

func generateBidiBracketsPublic(data *bytes.Buffer){
	data.WriteString(`// Unicode property "Bidi_Paired_Bracket" (known as "bpb", "Bidi_Paired_Bracket").`+"\n")
	data.WriteString(`// Kind of property: "Miscellaneous" (rune).`+"\n")
	data.WriteString(`// Based on file "`+ srcBidiPairedBrackets +`".`+"\n")
	data.WriteString("var "+strings.Title(varBidiPairedBrackets)+"="+ varBidiPairedBrackets+"\n\n")
}

func generateBidiBrackets(srcDir string, data *bytes.Buffer){
	src, err := os.Open(srcDir + string(os.PathSeparator) + srcBidiPairedBrackets)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	p := ucd.New(src)
	// Parse src
	data.WriteString("var "+ varBidiPairedBrackets +" = map[rune]rune{")
	for p.Next() {
		bracket0:=p.Rune(0)
		bracket1:=p.Rune(1)
		data.WriteString("0x"+strconv.FormatInt(int64(bracket0),16)+": "+"0x"+strconv.FormatInt(int64(bracket1),16)+",")	// "0x1234: 0x1236,"
	}
	// replace last unneeded "," with ")"
	data.Truncate(data.Len()-1)
	data.WriteString("}\n\n")
}

func generateBidiMirroringGlyphPublic(data *bytes.Buffer){
	data.WriteString(`// Unicode property "Bidi_Mirroring_Glyph" (known as "bmg", "Bidi_Mirroring_Glyph").`+"\n")
	data.WriteString(`// Kind of property: "Miscellaneous" (rune).`+"\n")
	data.WriteString(`// Based on file "`+srcBidiMirroringGlyph+`".`+"\n")
	data.WriteString("var "+strings.Title(varBidiMirroringGlyph)+"="+ varBidiMirroringGlyph+"\n\n")
}

func generateBidiMirroringGlyph(srcDir string, data *bytes.Buffer){
	src, err := os.Open(srcDir + string(os.PathSeparator) +srcBidiMirroringGlyph)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	p := ucd.New(src)
	// Parse src

	data.WriteString("var "+ varBidiMirroringGlyph +" = map[rune]rune{")
	for p.Next() {
		glyph0:=p.Rune(0)
		glyph1:=p.Rune(1)
		data.WriteString("0x"+strconv.FormatInt(int64(glyph0),16)+": "+"0x"+strconv.FormatInt(int64(glyph1),16)+",")	// "0x1234: 0x1236,"
	}
	// replace last unneeded "," with ")"
	data.Truncate(data.Len()-1)
	data.WriteString("}\n\n")
}