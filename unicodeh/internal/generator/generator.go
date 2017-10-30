package main

import (
	"bytes"
	"github.com/apaxa-go/helper/strconvh"
	"github.com/apaxa-go/helper/stringsh"
	"github.com/apaxa-go/helper/unicodeh/internal/generator/is-generator"
	"github.com/apaxa-go/helper/unicodeh/internal/generator/name"
	"github.com/apaxa-go/helper/unicodeh/internal/generator/other-generator"
	"github.com/apaxa-go/helper/unicodeh/internal/generator/ucdparser"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

const numLen = 3
const pkgName = "unicodeh"

func prepareTables(p *ucdparser.Parser) (r [][]byte) {
	for _, prop := range p.Properties {
		data := bytes.NewBuffer(nil)
		//
		// Public variables
		//
		knownAs := ""
		if len(prop.KnownAs) > 1 {
			knownAs = " (known as " + strings.Join(stringsh.Surround(prop.KnownAs, `"`, `"`), ", ") + ")"
		}
		data.WriteString(`// Unicode property "` + prop.LongName + `"` + knownAs + ".\n")
		data.WriteString(`// Kind of property: "` + prop.Kind + "\".\n")
		data.WriteString(`// Based on file "` + prop.File + "\".\n")
		data.WriteString("var (\n")
		for _, val := range prop.Values {
			publicName := name.Make(prop.LongName, val.LongName, false)
			privateName := name.Make(prop.LongName, val.LongName, true)
			knownAs := ""
			if len(val.KnownAs) > 1 {
				knownAs = " (known as " + strings.Join(stringsh.Surround(val.KnownAs, `"`, `"`), ", ") + ")"
			}
			data.WriteString(publicName + " = " + privateName + ` // Value "` + val.LongName + `"` + knownAs + ".\n")
		}
		data.WriteString(")\n\n")
		//
		// Private variables
		//
		data.WriteString("var (\n")
		for _, val := range prop.Values {
			data.WriteString(name.Make(prop.LongName, val.LongName, true))
			data.WriteString("=")
			data.WriteString(rangetableh.GoString(val.Ranges))
			data.WriteString("\n")
		}
		data.WriteString(")\n\n")
		//
		r = append(r, data.Bytes())
	}
	return
}

func prepareMeta(p *ucdparser.Parser) []byte {
	return []byte("// Version is the Unicode edition from which the tables are derived.\nconst Version=\"" + p.Version + "\"\n")
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

func main() {
	const usage = "Bad usage. Usage: \"generator path-to-ucd-directory\""
	if len(os.Args) != 2 {
		panic(usage)
	}

	srcDir := os.Args[1]

	parser := ucdparser.NewParser(srcDir)
	parser.Properties = ucdparser.ParseStructureUCD(srcDir)
	parser.Parse()

	// TODO delete old files

	saveFile("meta-gen.go", pkgName, nil, prepareMeta(parser))

	for i, data := range prepareTables(parser) {
		iStr := stringsh.PadLeft(strconvh.FormatInt(i), "0", numLen) // 000,001,002,...
		fileName := "tables-" + iStr + "-gen.go"
		saveFile(fileName, pkgName, []string{"unicode"}, data)
	}

	saveFile("other-gen.go", pkgName, nil, other_generator.GenerateOther(srcDir))

	{
		datas, importUnicode := is_generator.Generate(&parser.Properties)
		for i, data := range datas {
			iStr := stringsh.PadLeft(strconvh.FormatInt(i), "0", numLen) // 000,001,002,...
			fileName := "is-" + iStr + "-gen.go"
			var imports []string
			if importUnicode[i] {
				imports = []string{"unicode"}
			}
			saveFile(fileName, pkgName, imports, data)
		}
	}
}
