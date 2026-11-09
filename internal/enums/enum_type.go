// Package enums bla bla
package enums

import (
	"slices"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type EnumType[E, N any] struct {
	Name         ts.EnumName
	Lookup       map[string]EnumVal
	ConvFunc     func(string) E
	NullConvFunc func(*string) N
	GetNullEnum  func(*E) N
	Aliasses     map[string][]E
}

// CheckVal verifies an EnumVal based on its value or an idString (mostly from query params)
func (e EnumType[E, N]) CheckVal(key string) (EnumVal, error) {
	id, err := strconv.Atoi(key)
	if err == nil {
		if id > len(e.Lookup) || id <= 0 {
			return EnumVal{}, errs.ErrIDNotFound
		}

		for _, res := range e.Lookup {
			if int32(id) == res.ID {
				return res, nil
			}
		}
	}

	res, found := e.Lookup[key]
	if found {
		return res, nil
	}

	keyWithSpaces := h.GetNameWithSpaces(key, "-")
	res, found = e.Lookup[keyWithSpaces]
	if found {
		return res, nil
	}

	return EnumVal{}, errs.ErrNoResource
}

func (e EnumType[E, N]) Vals() []EnumVal {
	vals := []EnumVal{}

	for _, val := range e.Lookup {
		vals = append(vals, val)
	}

	slices.SortStableFunc(vals, h.SortOnID)

	return vals
}

func (e EnumType[E, N]) Strings() []string {
	enumVals := e.Vals()
	strings := []string{}

	for _, val := range enumVals {
		strings = append(strings, val.Name)
	}

	return strings
}