package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/asth"
	"github.com/apaxa-go/helper/goh/constanth"
	"go/ast"
	"reflect"
	"strings"
)

type (
	Identifiers          map[string]Value
	IdentifiersRegular   map[string]reflect.Value
	IdentifiersInterface map[string]interface{}
)

func (idents IdentifiersRegular) Identifiers() (r Identifiers) {
	r = make(Identifiers)
	for i := range idents {
		r[i] = MakeRegular(idents[i])
	}
	return
}
func (idents IdentifiersInterface) IdentifiersRegular() (r IdentifiersRegular) {
	r = make(IdentifiersRegular)
	for i := range idents {
		r[i] = reflect.ValueOf(idents[i])
	}
	return
}
func (idents IdentifiersInterface) Identifiers() (r Identifiers) {
	r = make(Identifiers)
	for i := range idents {
		r[i] = MakeRegular(reflect.ValueOf(idents[i]))
	}
	return
}

func (idents Identifiers) normalize() error {
	packages := make(map[string]Identifiers)

	// Extract idents with package specific
	for ident := range idents {
		parts := strings.Split(ident, ".")
		switch len(parts) {
		case 1:
			continue
		case 2:
			if parts[0] == "_" || !asth.IsValidIdent(parts[0]) || !asth.IsValidExportedIdent(parts[1]) {
				return errors.New("invalid identifier " + ident)
			}

			if _, ok := packages[parts[0]]; !ok {
				packages[parts[0]] = make(Identifiers)
			}
			packages[parts[0]][parts[1]] = idents[ident]
			delete(idents, ident)
		default:
			return errors.New("invalid identifier " + ident)
		}
	}

	// Add computed packages
	for pk, pv := range packages {
		// Check for unique package name
		if _, ok := idents[pk]; ok {
			return errors.New("something with package name already exists " + pk)
		}
		idents[pk] = packageVal(pv)
	}

	return nil
}

func funcTranslateArgs(fields *ast.FieldList, ellipsisAlowed bool, idents Identifiers) (r []reflect.Type, variadic bool, err error) {
	if fields == nil || len(fields.List) == 0 {
		return
	}
	r = make([]reflect.Type, len(fields.List))
	for i := range fields.List {
		// check for variadic
		if _, ellipsis := fields.List[i].Type.(*ast.Ellipsis); ellipsis {
			if !ellipsisAlowed || i != len(fields.List)-1 {
				return nil, false, errors.New("only last input argument can be variadic")
			}
			variadic = true
		}
		// calc type
		var v Value
		v, err = astExpr(fields.List[i].Type, idents)
		if err != nil {
			return nil, false, err
		}

		if v.Kind() != Type {
			return nil, false, errors.New("") // TODO error
		}
		r[i] = v.Type()
	}
	return
}

func Expr(e ast.Expr, idents Identifiers) (r Value, err error) {
	err = idents.normalize()
	if err != nil {
		return
	}
	return astExpr(e, idents)
}

func ExprRegular(e ast.Expr, idents IdentifiersRegular) (r reflect.Value, err error) {
	rV, err := astExpr(e, idents.Identifiers())
	if err != nil {
		return
	}

	switch rV.Kind() {
	case Regular:
		r = rV.Regular()
	case Untyped:
		var ok bool
		r, ok = constanth.DefaultType(rV.Untyped())
		if !ok {
			return r, errors.New("unable to represent untyped value in default type")
		}
	case Nil:
		r = reflect.ValueOf(nil) // TODO is it normal?
	default:
		return r, errors.New("Regular, Untyped or Nil required")
	}
	return
}

func ExprInterface(e ast.Expr, idents IdentifiersInterface) (r interface{}, err error) {
	rV, err := ExprRegular(e, idents.IdentifiersRegular())
	if err != nil {
		return
	}

	if !rV.CanInterface() {
		return r, errors.New("result value can not be converted into interface")
	}

	return rV.Interface(), nil
}
