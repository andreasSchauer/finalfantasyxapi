package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func convertObjPtr[Old, New any](item *Old, convFn func(Old) New) *New {
	if item == nil {
		return nil
	}

	new := convFn(*item)
	return &new
}

func convertObjSlice[Old, New any](items []Old, convFn func(Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := convFn(item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func convertObjSliceOrNil[Old, New any](items []Old, convFn func(Old) New) []New {
	slice := convertObjSlice(items, convFn)
	return h.SliceOrNil(slice)
}
