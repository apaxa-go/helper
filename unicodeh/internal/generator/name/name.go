package name

import (
	"github.com/apaxa-go/helper/stringsh"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Make(propName, valName string, private bool) string {
	propName = stringsh.ReplaceMulti(propName, []string{"_", "-"}, []string{"", ""})
	valName = stringsh.ReplaceMulti(valName, []string{"_", "-"}, []string{"", ""})
	if private {
		// make first letter to lowercase
		runee, l := utf8.DecodeRuneInString(propName)
		if runee == utf8.RuneError {
			panic("Invalid property name: " + propName)
		}
		runee = unicode.ToLower(runee)
		propName = string(runee) + propName[l:]
	} else {
		propName = strings.Title(propName)
	}
	valName = strings.Title(valName)
	return propName + valName
}
