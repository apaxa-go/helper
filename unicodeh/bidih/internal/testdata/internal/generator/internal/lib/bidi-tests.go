package lib

import (
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"strings"
	"github.com/apaxa-go/helper/stringsh"
	"github.com/apaxa-go/helper/strconvh"
	"os"
	"bytes"
	"github.com/apaxa-go/helper/mathh"
)

func parseParagraphDirections(directions uint)[]string{
	r:=make([]string,0,3)
	if directions&2==2{
		r=append(r,"0") // LTR
		directions &=^uint(2)
	}
	if directions&4==4{
		r=append(r,"1") // RTL
		directions &=^uint(4)
	}
	if directions&1==1{
		r=append(r,strconvh.FormatUint(mathh.MaxUint8)) // AUTO
		directions &=^uint(1)
	}
	if directions!=0{
		panic("unknown directions")
	}
	if len(r)==0{
		panic("no direction defined")
	}
	return r
}

func parseBidiTests(srcDir string)(r [][]byte){
	const srcFile = "BidiTest.txt"
	const strLevels = "Levels:"
	const strReorder = "Reorder:"
	src, err := os.Open(srcDir + string(os.PathSeparator) +srcFile)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	curLevels:=""
	curOrders:=""
	groupOpened:=false
	data:=bytes.NewBuffer(nil)

	openGroupIfRequired:=func(){
		if !groupOpened {
			data.WriteString("{\n")
			data.WriteString("[]uint8{" + curLevels + "},\n")
			data.WriteString("[]int{" + curOrders + "},\n")
			data.WriteString("[]BidiTest{\n")
			groupOpened = true
		}
	}

	closeGroupIfRequired:=func(){
		if groupOpened{
			data.WriteString("},\n},\n")
			groupOpened=false
			if data.Len()>=maxFileSize{
				r=append(r,data.Bytes())
				data=bytes.NewBuffer(nil)
			}
		}
	}

	catchAtLines:=func(p *ucd.Parser){
		str:=p.String(0)
		if strings.HasPrefix(str,strLevels){
			curLevels=parseLevels(str[len(strLevels):])
			closeGroupIfRequired()
		}else if strings.HasPrefix(str,strReorder){
			curOrders=parseOrders(str[len(strReorder):])
			closeGroupIfRequired()
		}
	}

	p := ucd.New(src, ucd.Part(catchAtLines))

	for p.Next() {
		openGroupIfRequired()
		classes:=parseClasses(p.String(0))
		paragraphDirections:=parseParagraphDirections(p.Uint(1))
		data.WriteString("{")
		data.WriteString("[]bidi.Class{"+strings.Join(classes,",")+"},")
		data.WriteString("[]uint8{"+strings.Join(paragraphDirections,",")+"}")
		data.WriteString("},\n")
	}

	closeGroupIfRequired()
	if data.Len()>0{
		r=append(r,data.Bytes())
	}
	return
}

func GenerateBidiTests(srcDir string){
	r:= parseBidiTests(srcDir)
	for i:=range r{
		iStr:=stringsh.PadLeft(strconvh.FormatInt(i),"0",numLen)	// 000,001,002,...

		varName:="bidiTests"+iStr
		prefix:="func init(){ BidiTests = append(BidiTests,"+varName+"...) }\n\n"
		prefix+="var "+varName+" = []BidiTestGroup{\n"
		suffix:="}\n"
		data:=append(append([]byte(prefix),r[i]...),[]byte(suffix)...)

		fileName:="bidi-tests-"+iStr+"-gen.go"
		imports:=[]string{"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"}

		saveFile(fileName,pkgName,imports,data)
	}
}
