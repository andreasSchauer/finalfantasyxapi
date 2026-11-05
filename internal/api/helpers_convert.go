package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func ConvertObjPtr[Old, New any](item *Old, convFn func(Old) New) *New {
	if item == nil {
		return nil
	}

	new := convFn(*item)
	return &new
}

func ConvertObjSlice[Old, New any](items []Old, convFn func(Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := convFn(item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func ConvertObjSliceOrNil[Old, New any](items []Old, convFn func(Old) New) []New {
	slice := ConvertObjSlice(items, convFn)
	return h.SliceOrNil(slice)
}
