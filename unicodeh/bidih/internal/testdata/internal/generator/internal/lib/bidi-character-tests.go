package lib

import (
	"os"
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"bytes"
	"github.com/apaxa-go/helper/mathh"
	"fmt"
	"github.com/apaxa-go/helper/strconvh"
	"github.com/apaxa-go/helper/stringsh"
)

func parseBidiCharacterTests(srcDir string)(r [][]byte){
	const srcFile = "BidiCharacterTest.txt"
	src, err := os.Open(srcDir + string(os.PathSeparator) +srcFile)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	p := ucd.New(src)

	data:=bytes.NewBuffer(nil)

	for p.Next() {
		//
		// Read test
		//
		// Input
		runes:=p.Runes(0)
		paragraphDirection:=p.Uint(1)
		// Output
		paragraphEmbeddingLevel:=p.Uint(2)
		embeddingLevels:=p.String(3)
		orders:=p.String(4)

		//
		// Prepare test data
		//
		if paragraphDirection==2{
			paragraphDirection=mathh.MaxUint8
		}
		embeddingLevels=parseLevels(embeddingLevels)
		orders=parseOrders(orders)
		//
		// Write test
		//
		data.WriteString("{")
		data.WriteString(fmt.Sprintf("%#v,",runes))
		data.WriteString(strconvh.FormatUint(paragraphDirection)+",")
		data.WriteString(strconvh.FormatUint(paragraphEmbeddingLevel)+",")
		data.WriteString("[]uint8{"+embeddingLevels+"},")
		data.WriteString("[]int{"+orders+"}")
		data.WriteString("},\n")

		//
		// Split to files if required
		//
		if data.Len()>=maxFileSize{
			r=append(r,data.Bytes())
			data=bytes.NewBuffer(nil)
		}
	}

	if data.Len()>0{
		r=append(r,data.Bytes())
	}

	return
}

func GenerateBidiCharacterTests(srcDir string){
	r:= parseBidiCharacterTests(srcDir)
	for i:=range r{
		iStr:=stringsh.PadLeft(strconvh.FormatInt(i),"0",numLen)	// 000,001,002,...

		varName:="bidiCharacterTests"+iStr
		prefix:="func init(){ BidiCharacterTests = append(BidiCharacterTests,"+varName+"...) }\n\n"
		prefix+="var "+varName+" = []BidiCharacterTest{\n"
		suffix:="}\n"
		data:=append(append([]byte(prefix),r[i]...),[]byte(suffix)...)

		fileName:="bidi-character-tests-"+iStr+"-gen.go"

		saveFile(fileName,pkgName,nil,data)
	}
}