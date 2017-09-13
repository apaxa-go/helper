package ucdparser

import (
	"log"
)

type Parser struct {
	dir                  string
	UnicodeVersion       string
	Properties           Properties
	AdditionalValues     map[string][][]string
	PseudoValues         map[string]map[string][]string
	DeprecatedProperties []string
	ReallyEmptyValues    map[string][]string
	ParseDetails         []ParseDetails
	SkipFiles            []string

	resultChan chan parseResult
}

func NewParser(srcDir string, version string) *Parser{
	var p Parser
	p.dir=srcDir
	p.UnicodeVersion =version
	p.AdditionalValues=DefaultAdditionalValues
	p.DeprecatedProperties=DefaultDeprecatedProperies
	p.ParseDetails=DefaultParseDetails
	p.PseudoValues=DefaultPseudoValues
	p.SkipFiles=DefaultSkipFiles
	p.ReallyEmptyValues=DefaultReallyEmptyValues
	return &p
}

func (parser *Parser)Parse(){
	parser.parsePropertyAliases()
	parser.parsePropertyValueAliases()
	parser.addAdditionalValues()
	parser.cleanDeprecated()
	parser.Properties.CleanEmpty()
	parser.addParseDetails()

	parser.resultChan=make(chan parseResult)
	defer close(parser.resultChan)

	for _, v := range parser.ParseDetails{
		go parser.parseFile(v.File, v.PropertyColumn,v.PropertyName,v.ValueColumn,v.ValueName,v.RangeColumn)
		// TODO remove files from list
	}
	for i,l := 0,len(parser.ParseDetails); i < l; i++ {
		log.Printf("waiting for %v/%v\n", i+1, l)
		parser.applyParseResult(<-parser.resultChan)
		log.Printf("merged %v/%v\n", i+1, l)
	}

	parser.computePseudoValues()
	parser.computeMissingValues()
	parser.validate()
}