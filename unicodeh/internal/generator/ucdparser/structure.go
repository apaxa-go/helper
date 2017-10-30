package ucdparser

import (
	"github.com/apaxa-go/helper/stringsh"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"golang.org/x/text/unicode/rangetable"
	"log"
	"strings"
	"unicode"
)

type Value struct {
	ShortName string
	LongName  string
	Num       uint // for "ccc" only	// TODO uint8?
	KnownAs   []string
	Ranges    *unicode.RangeTable
}

type Property struct {
	Kind      string
	ShortName string
	LongName  string
	KnownAs   []string
	Values    []Value
	File      string
	Missing   int // Value index or -1 if not applicable (i.e. file describe all values)
}

type Properties []Property

// -1 if not found
func (p Properties) PropIndexByName(str string) int {
	nstr := normalize(str)
	for i := range p {
		for j := range p[i].KnownAs {
			if nstr == normalize(p[i].KnownAs[j]) {
				return i
			}
		}
	}
	return -1
}

func (p Properties) MustPropIndexByName(str string) int {
	if r := p.PropIndexByName(str); r != -1 {
		return r
	}
	panic("property " + str + " does not exists")
}

func (p Property) MustValueIndexByName(str string) int {
	nstr := normalize(str)
	for i := range p.Values {
		for j := range p.Values[i].KnownAs {
			if nstr == normalize(p.Values[i].KnownAs[j]) {
				return i
			}
		}
	}
	panic("value " + str + " does not exists for property " + p.LongName + " (file " + p.File + ")")
}

func normalize(str string) string {
	return strings.ToLower(stringsh.ReplaceMulti(str, []string{" ", "_", "-"}, []string{"", "", ""}))
}

func (p *Parser) addAdditionalValues() {
	for i := range p.AdditionalValues {
		pI := p.Properties.MustPropIndexByName(i)
		for _, aV := range p.AdditionalValues[i] {
			v := Value{ShortName: aV[0], LongName: aV[1], KnownAs: aV}
			p.Properties[pI].Values = append(p.Properties[pI].Values, v)
		}
	}
}

func (p *Properties) CleanEmpty() {
	for i := 0; i < len(*p); i++ {
		if len((*p)[i].Values) > 0 {
			continue
		}
		log.Println("Remove empty property " + (*p)[i].LongName)
		*p = append((*p)[:i], (*p)[i+1:]...)
		i--
	}
}

func (p *Parser) cleanDeprecated() {
	for _, propName := range p.DeprecatedProperties {
		if pI := p.Properties.PropIndexByName(propName); pI != -1 {
			p.Properties = append(p.Properties[:pI], p.Properties[pI+1:]...)
		}
	}
}

func (p *Parser) computePseudoValues() {
	for prop := range p.PseudoValues {
		pI := p.Properties.MustPropIndexByName(prop)
		for val := range p.PseudoValues[prop] {
			// calc rangetable
			rs := []*unicode.RangeTable{}
			for _, x := range p.PseudoValues[prop][val] {
				vI := p.Properties[pI].MustValueIndexByName(x)
				rs = append(rs, p.Properties[pI].Values[vI].Ranges)
			}

			// save rangetable
			vI := p.Properties[pI].MustValueIndexByName(val)
			if p.Properties[pI].Values[vI].Ranges != nil {
				panic("range must be nil")
			}
			p.Properties[pI].Values[vI].Ranges = rangetable.Merge(rs...)
		}
	}
}

func (p *Parser) computeMissingValues() {
	for _, prop := range p.Properties {
		if prop.Missing == -1 {
			continue
		}
		allOtherValues := make([]*unicode.RangeTable, 0, len(prop.Values)-1)
		for i, val := range prop.Values {
			if val.Ranges != nil && i != prop.Missing {
				allOtherValues = append(allOtherValues, val.Ranges)
			}
		}
		prop.Values[prop.Missing].Ranges = rangetableh.Invert(allOtherValues...)
	}
}
