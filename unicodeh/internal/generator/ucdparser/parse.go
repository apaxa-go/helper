package ucdparser

import (
	"os"
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"golang.org/x/text/unicode/rangetable"
	"unicode"
	"fmt"
)

type ParseDetails struct {
	File     string
	PropertyColumn int
	PropertyName string
	ValueColumn int
	ValueName string
	RangeColumn int
	Missing string // only for PropertyName!=""
}

func (parser *Parser)addParseDetails(){
	for _,v:=range parser.ParseDetails{
		if v.PropertyName==""{continue}	// TODO still need fill this field, but in other way
		i:=parser.Properties.MustPropIndexByName(v.PropertyName)
		parser.Properties[i].File=v.File

		if v.Missing=="" {
			parser.Properties[i].Missing=-1
		}else {
			parser.Properties[i].Missing = parser.Properties[i].MustValueIndexByName(v.Missing)
		}
	}
}

type parseResult map[string]map[string]*unicode.RangeTable

func makeParseResult()parseResult{
	return make(map[string]map[string]*unicode.RangeTable)
}

// propertyColumn==-1 => using propertyName
// propertyColumn!=-1 && propertyName!="" => filtering
// propertyColumn==-1 && propertyName=="" => undefined behaviour
//
// valueColumn!=-1 xor valueName!="", else undefined behaviour
func (parser *Parser)parseFile(srcFile string, propertyColumn int, propertyName string, valueColumn int, valueName string, rangeColumn int) {
	src, err := os.Open(parser.dir + string(os.PathSeparator) +srcFile)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	p := ucd.New(src)
	res := makeParseResult()
	// Parse src
	for p.Next() {
		// TODO skip default for block?
		var ranget *unicode.RangeTable
		{
			lo, hi := p.Range(rangeColumn)
			ranget=rangetableh.FromRuneRange(lo, hi)
		}

		var prop string
		if propertyColumn!=-1{
			prop=p.String(propertyColumn)
			if propertyName!="" && propertyName!=prop {
				continue
			}
		}else {
			prop=propertyName
		}

		var val string
		if valueColumn!=-1{
			val=p.String(valueColumn)
		}else{
			val=valueName
		}

		if _,ok:=res[prop];ok{
			if _,ok:=res[prop][val]; ok{
				res[prop][val]=rangetable.Merge(res[prop][val],ranget)
			}else{
				res[prop][val]=ranget
			}
		}else{
			res[prop]=map[string]*unicode.RangeTable{val: ranget}
		}
	}

	// Hack: store source files for some properties
	if propertyName==""{
		for prop:=range res {
			propI := parser.Properties.PropIndexByName(prop)
			if propI==-1{
				panic("Property \""+prop+"\" is not defined. Defined properties: "+fmt.Sprintf("%v",parser.Properties))
			}
			parser.Properties[propI].File = srcFile
		}
	}

	parser.resultChan <- res
}

func (parser *Parser)applyParseResult(res parseResult){
	for i:=range res {
		propI:=parser.Properties.MustPropIndexByName(i)
		for j:=range res[i]{
			valJ:=parser.Properties[propI].MustValueIndexByName(j)
			if parser.Properties[propI].Values[valJ].Ranges==nil{
				parser.Properties[propI].Values[valJ].Ranges=res[i][j]
			}else {
				panic("here should be no panic! "+i+"."+j) // TODO !!!
				parser.Properties[propI].Values[valJ].Ranges = rangetable.Merge(parser.Properties[propI].Values[valJ].Ranges, res[i][j])
			}
		}
	}
}