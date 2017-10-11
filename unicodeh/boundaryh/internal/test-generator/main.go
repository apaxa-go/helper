package main

import (
	"bytes"
	"fmt"
	"github.com/apaxa-go/helper/unicodeh/internal/ucd"
	"go/format"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	breakStr   = "÷"
	noBreakStr = "×"
)

/*type ClusterBoundary struct {
	From, To int
}*/

/*func (b ClusterBoundary) ToSource() string {
	return "{" + strconvh.FormatInt(b.From) + "," + strconvh.FormatInt(b.To) + "}"
}*/

/*func GraphemesToSource(gs []ClusterBoundary) string {
	r := "[]Boundary{"
	for _, g := range gs {
		r += g.ToSource() + ","
	}
	return r[:len(r)-1] + "}"
}*/

func RunesToSource(runes []rune) string {
	r := "[]rune{"
	for _, char := range runes {
		r += "0x" + strconv.FormatInt(int64(char), 16) + ","
	}
	return r[:len(r)-1] + "}"
}

func isBreak(str string)bool{
	switch str {
	case breakStr:
		return true
	case noBreakStr:
		return false
	default:
		panic("Invalid runes separator "+str)
	}
}

func parseLine(strs []string) (runes []rune, breaks []int) {
	if isBreak(strs[0]){
		breaks=[]int{0}
	}

	runeCount := (len(strs) - 1) / 2
	runes = make([]rune, runeCount)
	for runeI := 0; runeI < runeCount; runeI++ {
		//
		// Rune
		//
		runeStr := strs[1+runeI*2]
		var err error
		tmp, err := strconv.ParseInt(runeStr, 16, 32)
		if err != nil {
			panic(err.Error())
		}
		runes[runeI] = rune(tmp)

		//
		// Break
		//
		breakStr:=strs[2+runeI*2]
		if isBreak(breakStr){
			breaks=append(breaks,runeI+1)
		}
	}
	return
}

func saveFile(fileName string, packageName string, imports []string, data []byte) {
	start := "package " + packageName + "\n\n"
	for _, i := range imports {
		start += "import \"" + i + "\"\n"
	}
	start += "\n"

	data = append([]byte(start), data...)

	data, err := format.Source(data)
	if err != nil {
		panic("Syntax error in file " + fileName + ": " + err.Error())
	}

	if err := ioutil.WriteFile(fileName, data, 0); err != nil {
		panic(err)
	}
}

func parseTests(srcDir, testFilename, codeSuffix string) []byte {
	srcFile := srcDir + string(os.PathSeparator) + "auxiliary" + string(os.PathSeparator) + testFilename
	src, err := os.Open(srcFile)
	defer src.Close()
	if err != nil {
		panic(err)
	}

	data := bytes.NewBuffer(nil)
	data.WriteString("var ucd" + codeSuffix + "Tests = []ucd" + codeSuffix + "Test{\n")

	p := ucd.New(src)
	for p.Next() {
		runes, breaks := parseLine(p.Strings(0))
		data.WriteString(fmt.Sprintf("{%v, %#v},\n", RunesToSource(runes), breaks))
	}

	data.WriteString("}\n")
	return data.Bytes()
}

func main() {
	//
	// Init
	//
	const usage = "Bad usage. Usage: \"maketables path-to-ucd-directory\"" // TODO
	if len(os.Args) != 2 {
		panic(usage)
	}
	srcDir := os.Args[1]

	//
	// Define settings
	//
	type test struct {
		targetFilename string
		codeSuffix     string
	}
	tests := map[string]test{
		"GraphemeBreakTest.txt": {"grapheme-cluster-ucd-data-gen_test.go", "GraphemeCluster"},
		"WordBreakTest.txt":     {"word-ucd-data-gen_test.go", "Word"},
		"SentenceBreakTest.txt": {"sentence-ucd-data-gen_test.go", "Sentence"},
		"LineBreakTest.txt":     {"line-ucd-data-gen_test.go", "Line"},
	}

	//
	// Execute
	//
	for testFilename, test := range tests {
		data := parseTests(srcDir, testFilename, test.codeSuffix)
		saveFile(test.targetFilename, "boundaryh", nil, data)
	}
}
