package is_generator

import (
	"unicode"
	"github.com/apaxa-go/helper/unicodeh/rangetableh"
	"strconv"
	"strings"
	"github.com/apaxa-go/helper/mathh"
	"bytes"
)

func generateTrivial(name,comment string, result bool, data *bytes.Buffer){
	data.WriteString(comment+"func "+name+"(r rune)bool{return ")
	if result{
		data.WriteString("true")
	}else{
		data.WriteString("false")
	}
	data.WriteString("}\n\n")
}

func formatUint16(i uint16)string{return "0x"+strconv.FormatUint(uint64(i),16)}
func formatUint32(i uint32)string{return "0x"+strconv.FormatUint(uint64(i),16)}
func formatInt64(i int64)string{return "0x"+strconv.FormatInt(i,16)}

func generateUnwrap(name,comment string, rt *unicode.RangeTable, data *bytes.Buffer){
	data.WriteString(comment+"func "+name+"(r rune)bool{ return ")
	conditions :=[]string{}
	for _,v:=range rt.R16{
		if rangetableh.RuneCount16(v)<=2 || (v.Stride>1 && rangetableh.RuneCount16(v)<=Step0Factor1){	// Direct in-code equivalence
			for j:=int64(v.Lo); j<=int64(v.Hi); j+=int64(v.Stride){
				conditions =append(conditions,"r=="+formatInt64(j))
			}
		}else{
			loHex:=formatUint16(v.Lo)
			hiHex:=formatUint16(v.Hi)
			comp:="r>="+loHex +" && r<="+hiHex
			if v.Stride==1{
				comp="("+comp+")"
			}else{
				strideHex:=formatUint16(v.Stride)
				comp="("+comp+" && (r-"+loHex+")%"+strideHex+"==0)"
			}
			conditions =append(conditions,comp)
		}
	}
	// TODO join if possible last 16 and first 32
	for _,v:=range rt.R32{
		if rangetableh.RuneCount32(v)<=2 || (v.Stride>1 && rangetableh.RuneCount32(v)<=Step0Factor1){	// Direct in-code equivalence
			for j:=int64(v.Lo); j<=int64(v.Hi); j+=int64(v.Stride){
				conditions =append(conditions,"r=="+formatInt64(j))
			}
		}else{
			loHex:=formatUint32(v.Lo)
			hiHex:=formatUint32(v.Hi)
			comp:="r>="+loHex +" && r<="+hiHex
			if v.Stride==1{
				comp="("+comp+")"
			}else{
				strideHex:=formatUint32(v.Stride)
				comp="("+comp+" && (r-"+loHex+")%"+strideHex+"==0)"
			}
			conditions =append(conditions,comp)
		}
	}
	data.WriteString(strings.Join(conditions," || "))
	data.WriteString("}\n\n")
}

func generateKeepRanges(name,comment string,data *bytes.Buffer){
	data.WriteString(comment+"func "+name+"(r rune)bool{ return unicode.Is("+name[2:]+",r) }\n\n")	// skip "Is" in name for rangetable
}

func generateArray(name,comment string, rt *unicode.RangeTable, data, suffixData *bytes.Buffer){
	//
	// compute dataArray, its position and values outside the dataArray
	//
	underLoRes:=false
	overHiRes:=false
	lo :=rangetableh.Lo(rt)
	hi:=rangetableh.Hi(rt)
	if lo==0{
		underLoRes=true
		for ; lo<=hi && unicode.Is(rt,lo); lo++{}
	}
	if hi==unicode.MaxRune{
		overHiRes=true
		for ; hi>=lo && unicode.Is(rt,hi); hi--{}
	}

	l:=hi- lo +1
	l=mathh.DivideCeilInt32(l,mathh.Uint64Bits)
	dataArray :=make([]uint64,l)
	for iter :=rangetableh.Start(rt); !iter.End(); iter.Next(){
		v:=iter.Value()
		if v<lo{
			continue
		}
		if v>hi{
			break
		}
		bitI:=int(v-lo)
		byteI:=bitI/mathh.Uint64Bits
		bitI=bitI%mathh.Uint64Bits
		dataArray[byteI]= dataArray[byteI] | (1<<uint(bitI))
	}

	//
	// Now we have lo, hi and dataArray
	//

	// Suffix data (array)
	suffixData.WriteString("var dataArray"+name+" = [...]uint64{")
	for i,v:=range dataArray {
		suffixData.WriteString("0x"+strconv.FormatUint(v,16))
		if i!=len(dataArray)-1{
			suffixData.WriteString(",")
		}
	}
	suffixData.WriteString("}\n\n")

	// Primary data (func)
	data.WriteString(comment+"func "+name+"(r rune)bool{\n")
	data.WriteString("const lo=0x"+strconv.FormatInt(int64(lo),16)+"\n")
	data.WriteString("const hi=0x"+strconv.FormatInt(int64(hi),16)+"\n")
	if underLoRes==overHiRes{
		data.WriteString("if r<lo||r>hi { return "+strconv.FormatBool(underLoRes)+" }\n")
	}else{
		data.WriteString("if r<lo { return "+strconv.FormatBool(underLoRes)+" }\n")
		data.WriteString("if r>hi { return "+strconv.FormatBool(overHiRes)+" }\n")
	}
	data.WriteString("var i=(r-lo)/64\n")
	data.WriteString("var shift=(r-lo)%64\n")
	data.WriteString("return dataArray"+name+"[i]&(1<<uint(shift))!=0\n")
	data.WriteString("}\n\n")
}
