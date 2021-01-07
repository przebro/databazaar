package collection

import (
	"errors"
)

var (
	//ErrEmptyOrInvalidID - empty or invalid id
	ErrEmptyOrInvalidID = errors.New("empty or invalid id")

	//ErrNotSliceOfStructs - not slice of structs
	ErrNotSliceOfStructs = errors.New("collection is not a slice of structs")

	//ErrNotStruct - not a struct
	ErrNotStruct = errors.New("document is not a struct")
)
