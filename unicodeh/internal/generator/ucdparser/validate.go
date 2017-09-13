package ucdparser

import (
	"log"
	"unicode"
)

func (p *Parser)validate(){
	for propI,prop:=range p.Properties{
		OutTo:for valI, val :=range prop.Values{
			if val.Ranges==nil {//&& valI !=prop.Missing{
				p.Properties[propI].Values[valI].Ranges=&unicode.RangeTable{}

				for eProp:=range p.ReallyEmptyValues{
					if propI==p.Properties.MustPropIndexByName(eProp){
						for _,eVal:=range p.ReallyEmptyValues[eProp]{
							if valI==prop.MustValueIndexByName(eVal){
								continue OutTo
							}
						}
					}
				}

				log.Println(prop.LongName+"."+ val.LongName+" has no range")
			}
		}
	}
}
