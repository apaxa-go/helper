package is_generator

import (
	"github.com/apaxa-go/helper/unicodeh/internal/generator/ucdparser"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"unicode"
	"github.com/apaxa-go/helper/unicodeh/internal/generator/name"
	"github.com/apaxa-go/helper/stringsh"
	"strings"
	"bytes"
)

// Settings.
// Can be changed before calling Generate.
var (
	Step0Factor int = 4      // maximum number of ranges
	Step0Factor1 int = 4 	// maximum number of elements in range with stride>1 to unwrap
	Step1Factor int = 6    // maximum number of ranges
)

const (
	decideTrivialFalse =iota
	decideTrivialTrue  =iota
	decideUnwrap       =iota
	decideKeepRanges   = iota
	decideArray        =iota
)

func decide(v ucdparser.Value)int{
	if rangetableh.RawRangesCount(v.Ranges)==0{
		return decideTrivialFalse
	}
	if rangetableh.RuneCount(v.Ranges)==unicode.MaxRune+1{
		return decideTrivialTrue
	}
	if rangetableh.RangesCount(v.Ranges)<=Step0Factor{
		return decideUnwrap
	}
	if rangetableh.RawRangesCount(v.Ranges)<=Step1Factor{
		return decideKeepRanges
	}
	return decideArray
}

func makeName(prop ucdparser.Property, valI int)string{
	return "Is"+name.Make(prop.LongName,prop.Values[valI].LongName,false)
}

func makeComment(prop ucdparser.Property, valI int)string{
	pName:=`"`+prop.LongName+`"`
	vName:=`"`+prop.Values[valI].LongName+`"`
	res:="// "+makeName(prop,valI) +" reports whether the rune has unicode property "+pName+"="+vName+".\n"
	if len(prop.KnownAs)>1{
		res+="// Property "+pName+" known as "+strings.Join(stringsh.Surround(prop.KnownAs,`"`,`"`),", ")+".\n"
	}
	if len(prop.Values[valI].KnownAs)>1{
		res+="// Value "+vName+" known as "+strings.Join(stringsh.Surround(prop.Values[valI].KnownAs,`"`,`"`),", ")+".\n"
	}
	return res
}

func Generate(p *ucdparser.Properties)(r [][]byte, importUnicode []bool){
	for propI:=range *p {
		data:=bytes.NewBuffer(nil)
		suffixData:=bytes.NewBuffer(nil)
		unicodeRequired:=false
		for valI := range (*p)[propI].Values {
			funcName:=makeName((*p)[propI],valI)
			funcComment:=makeComment((*p)[propI],valI)

			switch decide((*p)[propI].Values[valI]) {
			case decideTrivialFalse:
				generateTrivial(funcName,funcComment,false,data)
			case decideTrivialTrue:
				generateTrivial(funcName,funcComment,true,data)
			case decideUnwrap:
				generateUnwrap(funcName,funcComment,(*p)[propI].Values[valI].Ranges,data)
			case decideKeepRanges:
				generateKeepRanges(funcName, funcComment,data)
				unicodeRequired=true
			case decideArray:
				generateArray(funcName,funcComment,(*p)[propI].Values[valI].Ranges,data,suffixData)
			default:
				panic("Unknown decision")
			}
		}
		suffixData.WriteTo(data)
		r=append(r,data.Bytes())
		importUnicode=append(importUnicode,unicodeRequired)
	}
	return
}
