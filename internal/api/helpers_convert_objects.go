package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func convertObjPtr[Old, New any](cfg *Config, item *Old, convFn func(*Config, Old) New) *New {
	if item == nil {
		return nil
	}

	new := convFn(cfg, *item)
	return &new
}

func convertObjSlice[Old, New any](cfg *Config, items []Old, convFn func(*Config, Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := convFn(cfg, item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func convertObjSliceOrNil[Old, New any](cfg *Config, items []Old, convFn func(*Config, Old) New) []New {
	slice := convertObjSlice(cfg, items, convFn)
	return h.SliceOrNil(slice)
}
