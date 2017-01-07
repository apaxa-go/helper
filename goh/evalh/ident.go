package evalh

func undefIdentError(ident string) *intError {
	return newIntError("undefined: "+ident)
}
