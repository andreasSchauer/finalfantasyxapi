package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func convertObjPtr[Old, New any](server *Server, item *Old, convFn func(*Server, Old) New) *New {
	if item == nil {
		return nil
	}

	new := convFn(server, *item)
	return &new
}

func convertObjSlice[Old, New any](server *Server, items []Old, convFn func(*Server, Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := convFn(server, item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func convertObjSliceOrNil[Old, New any](server *Server, items []Old, convFn func(*Server, Old) New) []New {
	slice := convertObjSlice(server, items, convFn)
	return h.SliceOrNil(slice)
}
