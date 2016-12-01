// Package goutilh provides functions for generating Go code.
// It can be used in utilities for go:generate, such as including any file at compile time into project as []byte or duplicate existing Go code with some modification (some kind of generics/macros).
package goutilh

import (
	"fmt"
	"github.com/apaxa-go/helper/strconvh"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"strconv"
)

const astThreshold = 1024 // Maximum number of elements in slice for using Ast implementation

// WriteBytes convert given slice of bytes to GoLang source code representation and write it to given io.Writer.
// Result will be representation of array type if passed array is true. Otherwise representation will be of type slice.
// Example: for b=[]byte{1, 2, 3} WriteBytesStr will write "[]byte{1, 2, 3}" AS STRING if array is false and "[3]byte{1, 2, 3}" if array is true.
// This function may be useful with "//go:generate" directive to include some (binary) files into app at compile time.
// There are 2 different implementations of this function: using "go/ast" and generating representation "by hand".
// WriteBytes decides at runtime which implementation to use depending on len(b).
func WriteBytes(b []byte, w io.Writer, array bool) error {
	if len(b) > astThreshold {
		return WriteBytesStr(b, w, array)
	}
	return WriteBytesAst(b, w, array)
}

func writeBytesStr(b []byte, prefix string, w io.Writer) error {
	const format = "%d" // It is possible to use "%#x" for hex output, but Ast implementation does not support it and result file is smaller with "%d"

	// Prefix
	_, err := fmt.Fprint(w, prefix)
	if err != nil {
		return err
	}

	// Data
	if len(b) > 0 {
		// First byte
		_, err = fmt.Fprintf(w, format, b[0])
		if err != nil {
			return err
		}

		// All other bytes
		for _, v := range b[1:] {
			_, err = fmt.Fprintf(w, ", "+format, v)
			if err != nil {
				return err
			}
		}
	}

	// Suffix
	_, err = fmt.Fprint(w, "}")

	return err
}

// WriteBytesStr is "by hand" implementation of WriteBytes.
// It should works quicker even with a large number of elements in b.
func WriteBytesStr(b []byte, w io.Writer, array bool) error {
	var prefix string
	if array {
		prefix = "[" + strconv.Itoa(len(b)) + "]byte{"
	} else {
		prefix = "[]byte{"
	}
	return writeBytesStr(b, prefix, w)
}

func writeBytesAst(b []byte, w io.Writer, l ast.Expr) error {
	const baseType = "byte"
	compLit := &ast.CompositeLit{
		Type: &ast.ArrayType{
			Lbrack: 1, //token.NoPos,
			Len:    l, // nil => slice, else array
			Elt:    ast.NewIdent(baseType),
		},
		Lbrace: 1, //token.NoPos,
		Elts:   make([]ast.Expr, len(b)),
		Rbrace: 1, //token.NoPos,
	}

	for i, v := range b {
		compLit.Elts[i] = &ast.BasicLit{
			ValuePos: 1, //token.NoPos,
			Kind:     token.INT,
			Value:    strconvh.FormatUint8(v),
		}
	}

	//ast.Print(token.NewFileSet(), compLit)
	return printer.Fprint(w, token.NewFileSet(), compLit)
}

// WriteBytesAst is Ast (see go/ast) implementation of WriteBytes.
// It should generate proper code even if minor changes in language will happened, but its speed and resource consumption is poor.
func WriteBytesAst(b []byte, w io.Writer, array bool) error {
	var l ast.Expr	// nil
	if array {
		l = &ast.BasicLit{
			ValuePos: token.NoPos,
			Kind:     token.INT,
			Value:    strconvh.FormatInt(len(b)),
		}
	}
	return writeBytesAst(b, w, l)
}
