package rangetableh

import (
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/strconvh"
	"golang.org/x/text/unicode/rangetable"
	"strconv"
	"unicode"
)

const MinR32 = 1 << 16 // Minimal possible value for R32 (all smaller value should be in R16).

func FromRuneRange(lo, hi rune) *unicode.RangeTable {
	rt := unicode.RangeTable{}
	if lo < MinR32 {
		hi16 := mathh.Min2Int32(hi, MinR32)
		rt.R16 = []unicode.Range16{{uint16(lo), uint16(hi16), 1}}
		if hi16 <= unicode.MaxLatin1 {
			rt.LatinOffset = 1
		}
	}
	if hi >= MinR32 {
		lo32 := mathh.Max2Int32(lo, MinR32)
		rt.R32 = []unicode.Range32{{uint32(lo32), uint32(hi), 1}}
	}
	return &rt
}

func Copy(t *unicode.RangeTable) *unicode.RangeTable {
	var res unicode.RangeTable
	res.R16 = make([]unicode.Range16, len(t.R16))
	copy(res.R16, t.R16)
	res.R32 = make([]unicode.Range32, len(t.R32))
	copy(res.R32, t.R32)
	res.LatinOffset = t.LatinOffset
	return &res
}

// TODO move to "golang.org/x/text/unicode/rangetable"
// perform removing inplace (modify passed RangeTable)
func deleteRuneR16(t *unicode.RangeTable, r uint16) {
	lo := 0
	hi := len(t.R16)
	for lo < hi {
		m := lo + (hi-lo)/2
		range_ := &t.R16[m]
		if range_.Lo <= r && r <= range_.Hi { // here we found single possible range element
			if (r-range_.Lo)%range_.Stride == 0 { // check, if rune belong founded element

				if range_.Hi <= unicode.MaxLatin1 {
					t.LatinOffset--
				}
				insert := make([]unicode.Range16, 0, 2)
				if hi := r - range_.Stride; hi >= range_.Lo {
					prefix := *range_
					prefix.Hi = hi
					insert = append(insert, prefix)
					if prefix.Hi <= unicode.MaxLatin1 {
						t.LatinOffset++
					}
				}
				if lo := r + range_.Stride; lo <= range_.Hi {
					suffix := *range_
					suffix.Lo = lo
					insert = append(insert, suffix)
					if suffix.Hi <= unicode.MaxLatin1 {
						t.LatinOffset++
					}
				}
				t.R16 = append(t.R16[:m], append(insert, t.R16[m+1:]...)...) // replace t.R16[m] with insert

			}
			return
		}
		if r < range_.Lo {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return
}

// perform removing inplace (modify passed RangeTable)
func deleteRuneR32(t *unicode.RangeTable, r uint32) {
	lo := 0
	hi := len(t.R32)
	for lo < hi {
		m := lo + (hi-lo)/2
		range_ := &t.R32[m]
		if range_.Lo <= r && r <= range_.Hi { // here we found single possible range element
			if (r-range_.Lo)%range_.Stride == 0 { // check, if rune belong founded element

				insert := make([]unicode.Range32, 0, 2)
				if hi := r - range_.Stride; hi >= range_.Lo {
					prefix := *range_
					prefix.Hi = hi
					insert = append(insert, prefix)
				}
				if lo := r + range_.Stride; lo <= range_.Hi {
					suffix := *range_
					suffix.Lo = lo
					insert = append(insert, suffix)
				}
				t.R32 = append(t.R32[:m], append(insert, t.R32[m+1:]...)...) // replace t.R32[m] with insert

			}
			return
		}
		if r < range_.Lo {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return
}

// returns copy of past RangeTable
func DeleteRunes(t *unicode.RangeTable, r ...rune) *unicode.RangeTable {
	// TODO Here is a bug in implementation, now used bad way.
	/*
	res := Copy(t)

	for _, curR := range r {
		if l := len(res.R16); l > 0 && curR <= rune(res.R16[l-1].Hi) {
			deleteRuneR16(t, uint16(curR))
		} else if len(res.R32) > 0 && curR >= rune(res.R32[0].Lo) {
			deleteRuneR32(t, uint32(curR))
		}
	}

	return res
	*/
	runes0:=Runes(t)
	runes1:=make([]rune,0,len(runes0))
	Next:for _,r0:=range runes0{
		for _,rD:=range r{
			if r0==rD{
				continue Next
			}
		}
		runes1=append(runes1,r0)
	}
	return rangetable.New(runes1...)
}

func Runes(t *unicode.RangeTable)[]rune{
	res:=make([]rune,RuneCount(t))
	i:=0

	f:=func(r rune){
		res[i]=r
		i++
	}

	rangetable.Visit(t,f)
	return res
}

// ={lo,hi}-b0..-bi
func Sub(lo, hi rune, b ...*unicode.RangeTable) *unicode.RangeTable {
	runes := make([]rune, 0, hi-lo+1)
NextRune:
	for i := lo; i <= hi; i++ {
		for _, j := range b {
			if unicode.Is(j, i) {
				continue NextRune
			}
		}
		runes = append(runes, i)
	}
	return rangetable.New(runes...)
}

// =Sub(0, MaxRune, b)
func Invert(b ...*unicode.RangeTable) *unicode.RangeTable {
	return Sub(0, unicode.MaxRune, b...)
}

func Intersect(rts ...*unicode.RangeTable) *unicode.RangeTable {
	if len(rts) == 1 {
		return Copy(rts[0])
	}
	runes := []rune{}
NextRune:
	for iter := Start(rts[0]); !iter.End(); iter.Next() {
		v := iter.Value()
		for i := 1; i < len(rts); i++ {
			if !unicode.Is(rts[i], v) {
				continue NextRune
			}
		}
		runes = append(runes, v)
	}
	return rangetable.New(runes...)
}

func formatInt64(i int64) string {
	return "0x" + strconv.FormatInt(i, 16)
}

func formatUint64(i uint64) string {
	return "0x" + strconv.FormatUint(i, 16)
}

func GoString(r *unicode.RangeTable) (str string) {
	if r == nil {
		return "nil"
	}
	str = "&unicode.RangeTable{"
	if len(r.R16) > 0 {
		str += "[]unicode.Range16{"
		for i, v := range r.R16 {
			str += "{" + formatUint64(uint64(v.Lo)) + "," + formatUint64(uint64(v.Hi)) + "," + formatUint64(uint64(v.Stride)) + "}"
			if i < len(r.R16)-1 {
				str += ","
			}
		}
		str += "}"
	} else {
		str += "nil"
	}
	str += ","
	if len(r.R32) > 0 {
		str += "[]unicode.Range32{"
		for i, v := range r.R32 {
			str += "{" + formatUint64(uint64(v.Lo)) + "," + formatUint64(uint64(v.Hi)) + "," + formatUint64(uint64(v.Stride)) + "}"
			if i < len(r.R32)-1 {
				str += ","
			}
		}
		str += "}"
	} else {
		str += "nil"
	}
	str += "," + strconvh.FormatInt(r.LatinOffset) + "}"
	return
}

func RawRangesCount(rt *unicode.RangeTable) int {
	return len(rt.R16) + len(rt.R32)
}

func RuneCount16(range16 unicode.Range16) int {
	return (int(range16.Hi)-int(range16.Lo))/int(range16.Stride) + 1
}
func RuneCount32(range32 unicode.Range32) int {
	return (int(range32.Hi)-int(range32.Lo))/int(range32.Stride) + 1
}

func RuneCount(rt *unicode.RangeTable) (count int) {
	for _, r16 := range rt.R16 {
		count += RuneCount16(r16)
	}
	for _, r32 := range rt.R32 {
		count += RuneCount32(r32)
	}
	return
}

// stride = 0 => not joinable
func joinableR16R32(lo16 unicode.Range16, hi32 unicode.Range32) (stride uint32) {
	lo32 := unicode.Range32{uint32(lo16.Lo), uint32(lo16.Hi), uint32(lo16.Stride)}
	{
		cLo := RuneCount32(lo32)
		cHi := RuneCount32(hi32)
		if cLo == 1 && cHi == 1 {
			return hi32.Lo - lo32.Hi
		}
		if cLo == 1 {
			stride = hi32.Stride
		} else if cHi == 1 {
			stride = lo32.Stride
		} else if lo32.Stride == hi32.Stride {
			stride = lo32.Stride
		} else {
			return 0
		}
	}
	if lo32.Hi+stride != hi32.Lo {
		return 0
	}
	return stride
}

// TODO may be hide this general-useless function???
func RangesCount(rt *unicode.RangeTable) int {
	if len(rt.R16) == 0 {
		return len(rt.R32)
	}
	if len(rt.R32) == 0 {
		return len(rt.R16)
	}
	count := RawRangesCount(rt)
	if joinableR16R32(rt.R16[len(rt.R16)-1], rt.R32[0]) != 0 {
		count--
	}
	return count
}

// -1 if no runes
func Lo(rt *unicode.RangeTable) rune {
	if len(rt.R16) > 0 {
		return rune(rt.R16[0].Lo)
	}
	if len(rt.R32) > 0 {
		return rune(rt.R32[0].Lo)
	}
	return -1
}

// -1 if no runes
func Hi(rt *unicode.RangeTable) rune {
	if l := len(rt.R32); l > 0 {
		return rune(rt.R32[l-1].Hi)
	}
	if l := len(rt.R16); l > 0 {
		return rune(rt.R16[l-1].Hi)
	}
	return -1
}
