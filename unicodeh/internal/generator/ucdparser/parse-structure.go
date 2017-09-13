package ucdparser

import (
	"strings"
	"os"
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"regexp"
	"bufio"
)

func (parser *Parser)parsePropertyValueAliases(){
	const aliasesFile = "PropertyValueAliases.txt"
	const specialCaseCCC = "ccc"
	specialCaseCCCExists:=false	// check if case actual and work
	src, err := os.Open(parser.dir + string(os.PathSeparator) + aliasesFile)
	defer src.Close()
	if err != nil {
		panic(err)
	}
	p := ucd.New(src)

	// Parse src
	for p.Next() {
		prop:=p.String(0)
		v:=Value{}
		i :=1
		if prop==specialCaseCCC{
			specialCaseCCCExists=true
			v.Num =p.Uint(1)
			i++
			v.KnownAs=append(v.KnownAs,p.String(1))
		}
		for str:=p.String(i); len(str)!=0; str=p.String(i){	// len==0 is not good, but the only available solution
			v.KnownAs =append(v.KnownAs,str)
			i++
		}
		v.ShortName =v.KnownAs[0]
		if prop==specialCaseCCC{
			v.LongName =v.KnownAs[2]
		}else {
			v.LongName = v.KnownAs[1]
		}
		// remove duplicated known-as
		NextString:for i:=1; i<len(v.KnownAs); i++{
			for _,j:=range v.KnownAs[:i]{
				if v.KnownAs[i]==j{
					v.KnownAs=append(v.KnownAs[:i],v.KnownAs[i+1:]...)
					i--
					continue NextString
				}
			}
		}
		// save value
		i=parser.Properties.MustPropIndexByName(prop)
		parser.Properties[i].Values =append(parser.Properties[i].Values,v)
	}

	// additional check
	if !specialCaseCCCExists {
		panic("special case is not used")
	}
}

func (parser *Parser)parsePropertyAliases(){
	const aliasesFile = "PropertyAliases.txt"
	const kindREStr = `^#[[:space:]]*([[:alnum:]]+)[[:space:]]+Properties[[:space:]]*$`	// for catching row like this "# Catalog Properties"

	kindRE:=regexp.MustCompile(kindREStr)

	file, err := os.Open(parser.dir + string(os.PathSeparator) + aliasesFile)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	kind:=""

	for scanner.Scan() {
		str:=scanner.Text()
		if rer:=kindRE.FindStringSubmatch(str); rer!=nil{
			kind=rer[1]	// TODO normalize Kind?
			continue
		}
		// Trim comment
		if pos:=strings.Index(str,"#"); pos!=-1{
			str=str[:pos]
		}
		// Split estimated string to Values
		strs:=strings.Split(str,";")
		if len(strs)==1 {
			if len(strings.TrimSpace(strs[0]))!=0{
				panic("invalid string "+str)
			}
			continue
		}

		if kind==""{
			panic("Properties Kind is not set")
		}
		// Trim Properties names
		for i:=range strs{
			strs[i]=strings.TrimSpace(strs[i])
			if len(strs[i])==0{
				panic("empty property LongName")
			}
		}
		// Construct result
		parser.Properties =append(parser.Properties,Property{Kind:kind,ShortName:strs[0],LongName:strs[1],KnownAs:strs})
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}

