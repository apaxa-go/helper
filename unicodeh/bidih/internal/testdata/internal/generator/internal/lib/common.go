package lib

import (
	"strings"
	"github.com/apaxa-go/helper/strconvh"
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
	"go/format"
	"io/ioutil"
)

const maxFileSize = 1024*1024
const numLen = 3
const pkgName = "testdata"

func saveFile(fileName string, packageName string, imports []string, data []byte){
	start:="package "+packageName+"\n\n"
	for _,i:=range imports{
		start+="import \""+i+"\"\n"
	}
	start+="\n"

	data=append([]byte(start),data...)

	data, err := format.Source(data)
	if err!=nil{
		panic("Syntax error in file "+fileName+": "+err.Error())
	}

	if err := ioutil.WriteFile(fileName, data, 0); err != nil {
		panic(err)
	}
}

func parseLevels(str string)string{
	str=strings.Replace(str," ",",",-1)
	str=strings.Replace(str,"x",strconvh.FormatUint8(mathh.MaxUint8),-1)
	return str
}

func parseOrders(str string)string{
	return strings.Replace(str," ",",",-1)
}

func parseClass(str string)bidi.Class{
	switch str {
	case "AL":
		return bidi.ArabicLetter
	case "AN":
		return bidi.ArabicNumber
	case "B":
		return bidi.ParagraphSeparator
	case "BN":
		return bidi.BoundaryNeutral
	case "CS":
		return bidi.CommonSeparator
	case "EN":
		return bidi.EuropeanNumber
	case "ES":
		return bidi.EuropeanSeparator
	case "ET":
		return bidi.EuropeanTerminator
	case "FSI":
		return bidi.FirstStrongIsolate
	case "L":
		return bidi.LeftToRight
	case "LRE":
		return bidi.LeftToRightEmbedding
	case "LRI":
		return bidi.LeftToRightIsolate
	case "LRO":
		return bidi.LeftToRightOverride
	case "NSM":
		return bidi.NonSpacingMark
	case "ON":
		return bidi.OtherNeutral
	case "PDF":
		return bidi.PopDirectionalFormat
	case "PDI":
		return bidi.PopDirectionalIsolate
	case "R":
		return bidi.RightToLeft
	case "RLE":
		return bidi.RightToLeftEmbedding
	case "RLI":
		return bidi.RightToLeftIsolate
	case "RLO":
		return bidi.RightToLeftOverride
	case "S":
		return bidi.SegmentSeparator
	case "WS":
		return bidi.WhiteSpace
	default:
		panic("unknown BIDI class")
	}
}

func parseClasses(str string)[]string{
	classeStrs:=strings.Split(str," ")
	r:=make([]string,len(classeStrs))
	for i,v:=range classeStrs{
		r[i]=strconvh.FormatUint(uint(parseClass(v)))
	}
	return r
}
