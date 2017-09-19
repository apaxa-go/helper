package main

import (
	"os"
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"bytes"
	"go/format"
	"io/ioutil"
	"fmt"
	"strconv"
	"github.com/apaxa-go/helper/strconvh"
)

const(
	breakStr = "÷"
	noBreakStr = "×"
)

type ClusterBoundary struct {
	From,To int
}

func (b ClusterBoundary)ToSource()string{
	return "{"+strconvh.FormatInt(b.From)+","+strconvh.FormatInt(b.To)+"}"
}

func GraphemesToSource(gs []ClusterBoundary)string{
	r:="[]Boundary{"
	for _,g :=range gs{
		r+=g.ToSource()+","
	}
	return r[:len(r)-1]+"}"
}

func RunesToSource(runes []rune)string{
	r:="[]rune{"
	for _,char:=range runes{
		r+="0x"+strconv.FormatInt(int64(char),16)+","
	}
	return r[:len(r)-1]+"}"
}

func parseLine(strs []string)(runes []rune, clusters []ClusterBoundary){
	if strs[0]!=breakStr || strs[len(strs)-1]!=breakStr{
		panic("Invalid test line.")
	}

	//
	//	Compute string
	//
	runeLength:=(len(strs)-1)/2
	runes=make([]rune,runeLength)
	for runeI:=0; runeI<runeLength; runeI++{
		runeStr:=strs[1+runeI*2]
		var err error
		tmp,err:=strconv.ParseInt(runeStr,16,32)
		if err!=nil{
			panic(err.Error())
		}
		runes[runeI]=rune(tmp)
	}

	//
	// Compute grapheme clusters
	//
	from:=0
	for runeI:=0; runeI<runeLength; runeI++{ // Check if runeI is the last rune in cluster
		i:=1+runeI*2
		switch strs[1+i]{
		case breakStr:
			clusters=append(clusters, ClusterBoundary{from,runeI+1})
			from=runeI+1
		case noBreakStr:
		default:
			panic("Invalid test line.")
		}
	}

	return
}

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

func parseTests(srcDir,testFilename,codeSuffix string)[]byte{
	srcFile := srcDir + string(os.PathSeparator)+"auxiliary"+string(os.PathSeparator)+testFilename
	src, err := os.Open(srcFile)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	data:=bytes.NewBuffer(nil)
	data.WriteString("var ucd"+codeSuffix+"Tests = []ucd"+codeSuffix+"Test{\n")

	p := ucd.New(src)
	for p.Next() {
		runes,clusters:=parseLine(p.Strings(0))
		data.WriteString(fmt.Sprintf("{%v, %v},\n",RunesToSource(runes),GraphemesToSource(clusters)))
	}

	data.WriteString("}\n")
	return data.Bytes()
}

func main(){
	//
	// Init
	//
	const usage="Bad usage. Usage: \"maketables path-to-ucd-directory\"" // TODO
	if len(os.Args)!=2{
		panic(usage)
	}
	srcDir := os.Args[1]

	//
	// Define settings
	//
	type test struct {
		targetFilename string
		codeSuffix string
	}
	tests:=map[string]test{
		"GraphemeBreakTest.txt":{"grapheme-cluster-ucd-data-gen_test.go","GraphemeCluster"},
		"WordBreakTest.txt":{"word-ucd-data-gen_test.go","Word"},
		"SentenceBreakTest.txt":{"sentence-ucd-data-gen_test.go","Sentence"},
	}

	//
	// Execute
	//
	for testFilename,test:=range tests {
		data := parseTests(srcDir,testFilename,test.codeSuffix)
		saveFile(test.targetFilename, "boundaryh", nil, data)
	}
}