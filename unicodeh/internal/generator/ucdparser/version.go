package ucdparser

import (
	"bufio"
	"os"
	"regexp"
)

var (
	ucdVerRegexp   = regexp.MustCompile(`^# [[:alpha:]]+-([[:digit:]]+(?:\.[[:digit:]]+)+)\.txt$`)
	emojiVerRegexp = regexp.MustCompile(`^# Version: ([[:digit:]]+(?:\.[[:digit:]]+)+)$`)
)

func UCDVer(fileName string) string {
	//
	// Open file
	//
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//
	// Read first line
	//
	fb := bufio.NewReader(f)
	line, notComplete, err := fb.ReadLine()
	if err != nil {
		panic(err)
	}
	if notComplete {
		return ""
	}
	//
	// Validate first line
	//
	tmp := ucdVerRegexp.FindSubmatch(line)
	if tmp == nil {
		return ""
	}
	return string(tmp[1])
}

func EmojiVer(fileName string) string {
	//
	// Open file
	//
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//
	//
	//
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		tmp := emojiVerRegexp.FindStringSubmatch(sc.Text())
		if tmp != nil {
			return tmp[1]
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return ""
}

func (p *Parser) getVersion() {
	if p.VersionFunc != nil {
		alreadyParsed := map[string]bool{}
		for i := range p.ParseDetails {
			fileName := p.dir + string(os.PathSeparator) + p.ParseDetails[i].File
			if _, ok := alreadyParsed[fileName]; ok {
				continue
			}
			alreadyParsed[fileName] = true
			version := p.VersionFunc(fileName)
			if p.Version == "" && version != "" {
				p.Version = version
			} else if p.Version != "" && version != "" && p.Version != version {
				panic("Different versions")
			}
		}
	}
	if p.Version == "" {
		panic("Undefined version")
	}
}
