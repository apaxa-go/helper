package asth

import "go/ast"

const (
	SendDir ast.ChanDir = ast.SEND
	RecvDir             = ast.RECV
	BothDir             = SendDir | RecvDir
)
