package ctlh

/*
type GraphemeCluster struct {
	FromI, ToI int // indexes in rune slice; [FromI; ToI)
	FromP, ToP int // in pixels; [FromP; ToP)
}

type GraphemeClusters struct {
	Clusters []GraphemeCluster
	DisplayOrder []int
}

func (c GraphemeClusters)At(x int)(i int){
	l:=len(c.DisplayOrder)
	if l==0 || x<0 || x>=c.Clusters[c.DisplayOrder[l-1]].ToP {
		return -1
	}
	i=x*l/c.Clusters[c.DisplayOrder[l-1]].ToP // Init i as "x/<average cluster width>"
	if x<c.Clusters[c.DisplayOrder[i]].FromP {
		for i--; i>0; i--{
			if x>=c.Clusters[c.DisplayOrder[i]].FromP {
				return
			}
		}
	}else if x>=c.Clusters[c.DisplayOrder[i]].ToP {
		for i++; i<len(c.DisplayOrder); i++{
			if x<c.Clusters[c.DisplayOrder[i]].ToP {
				return
			}
		}
	}else{
		return
	}
	return -1
}

// order - new order of runes (after BIDI).
func (c *GraphemeClusters) computeOrder(order []int){
	c.DisplayOrder=make([]int,0,len(c.Clusters))
	for len(order)>0{
		
	}
}
*/