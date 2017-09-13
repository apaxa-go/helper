package ctl

import (
	"github.com/apaxa-go/helper/unicodeh/bidih"
	"github.com/apaxa-go/helper/unicodeh/grapheme"
)

type smth struct {}

func initGraphemeClusters(runes []rune)(clusters GraphemeClusters){
	clusters.Clusters=make([]GraphemeCluster,0,len(runes))	// TODO not memory effective
	for from:=0; from<len(runes);{
		length :=grapheme.FirstBoundaryInRunes(runes[from:])
		clusters.Clusters=append(clusters.Clusters,GraphemeCluster{FromI:from, ToI:from+ length})
		from+= length
	}
	return
}

func LayoutLineStringAdvanced(runes []rune, paragraphDirection bidih.EmbeddingLevel)(img smth, clusters GraphemeClusters){
	clusters=initGraphemeClusters(runes)
	order:=bidih.ParseParagraphExtended(runes, paragraphDirection, nil)
}
