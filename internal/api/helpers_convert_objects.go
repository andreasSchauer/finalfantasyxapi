package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func convertObjPtr[Old, New any](srv *Server, item *Old, convFn func(*Server, Old) New) *New {
	if item == nil {
		return nil
	}

	new := convFn(srv, *item)
	return &new
}

func convertObjSlice[Old, New any](srv *Server, items []Old, convFn func(*Server, Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := convFn(srv, item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func convertObjSliceOrNil[Old, New any](srv *Server, items []Old, convFn func(*Server, Old) New) []New {
	slice := convertObjSlice(srv, items, convFn)
	return h.SliceOrNil(slice)
}
