package main

import (
	"github.com/apaxa-go/helper/unicodeh"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"golang.org/x/text/unicode/norm"
	"log"
	"unicode"
)

// TODO use this for assumption and remove

func normRune(r rune) rune {
	tmp := []rune(norm.NFC.String(string(r)))
	if len(tmp) != 1 {
		return -1
	}
	return tmp[0]
}

func main() {
	tmpNormPairs := make(map[rune]rune) // norm open => norm close
	for iter := rangetableh.Start(unicodeh.BidiPairedBracketTypeOpen); !iter.End(); iter.Next() {
		openBracket := iter.Value()
		closeBracket := unicodeh.BidiPairedBracket[openBracket]
		tmpNormPairs[normRune(openBracket)] = normRune(closeBracket)
	}

	tmpNormClose := make(map[rune]rune) // norm close => close
	for iter := rangetableh.Start(unicodeh.BidiPairedBracketTypeClose); !iter.End(); iter.Next() {
		tmpNormClose[normRune(iter.Value())] = iter.Value()
	}

	normOpen := make(map[rune]rune)  // maps open bracket rune to its norm pair (closed bracket)
	normClose := make(map[rune]rune) // maps close bracket rune to its norm (not pair)
	for r := rune(0); r <= unicode.MaxRune; r++ {
		normR := normRune(r)
		if pair, ok := tmpNormPairs[normR]; ok {
			normOpen[r] = pair
		}
		if _, ok := tmpNormClose[normR]; ok {
			normClose[r] = normR
		}
	}

	log.Println("Open => norm close: ", normOpen)
	log.Println("Close => norm close:", normClose)
}
