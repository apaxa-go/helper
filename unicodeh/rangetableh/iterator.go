package rangetableh

import "unicode"

type Iterator struct {
	rt *unicode.RangeTable
	rangeT uint8	// 0 - R16, 1 - R32, 2 - End
	rangeI int
	runeI int
}

func Start(rt *unicode.RangeTable)*Iterator{
	iter:=Iterator{rt:rt}
	if len(rt.R16)>0{
		iter.rangeT=0
	}else if len(rt.R32)>0{
		iter.rangeT=1
	}else{
		iter.rangeT=2
	}
	return &iter
}

func (iter *Iterator)Next(){
	if iter.rangeT==2{
		return
	}
	iter.runeI++

	if iter.rangeT==0{
		r16:=iter.rt.R16[iter.rangeI]
		if int(r16.Lo)+iter.runeI*int(r16.Stride)<=int(r16.Hi){
			return
		}
	}else{
		r32:=iter.rt.R32[iter.rangeI]
		if int(r32.Lo)+iter.runeI*int(r32.Stride)<=int(r32.Hi){
			return
		}
	}

	iter.runeI=0
	iter.rangeI++
	if iter.rangeT==0{
		if iter.rangeI<len(iter.rt.R16){
			return
		}
	}else{
		if iter.rangeI<len(iter.rt.R32){
			return
		}
	}

	iter.rangeI=0
	iter.rangeT++
	if iter.rangeT==1 && len(iter.rt.R32)==0{
		iter.rangeT++
	}
}

func (iter *Iterator)End()bool{return iter.rangeT==2}

// -1 if end
func (iter *Iterator)Value()rune{
	if iter.rangeT==0{
		r16:=iter.rt.R16[iter.rangeI]
		return rune(r16.Lo)+rune(iter.runeI)*rune(r16.Stride)
	}
	if iter.rangeT==1{
		r32:=iter.rt.R32[iter.rangeI]
		return rune(r32.Lo)+rune(iter.runeI)*rune(r32.Stride)
	}
	return -1
}