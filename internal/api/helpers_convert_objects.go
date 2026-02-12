package api

import (
	"context"
	"fmt"
	"net/http"
)

func convertObjPtr[Old, New any](cfg *Config, item *Old, converter func(*Config, Old) New) *New {
	if item == nil {
		return nil
	}

	new := converter(cfg, *item)
	return &new
}

func convertObjSlice[Old, New any](cfg *Config, items []Old, converter func(*Config, Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := converter(cfg, item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func convertObjSliceNullable[Old, New any](cfg *Config, items []Old, converter func(*Config, Old) New) []New {
	slice := convertObjSlice(cfg, items, converter)

	if len(slice) == 0 {
		return nil
	}

	return slice
}

func dbQueryToSlice[T any](cfg *Config, r *http.Request, rtParent, rtChild string, id int32, dbQuery func(context.Context, int32) ([]int32, error), converter func(*Config, int32) T) ([]T, error) {
	IDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss of %s with id '%d'", rtChild, rtParent, id), err)
	}

	slice := convertObjSlice(cfg, IDs, converter)
	return slice, nil
}
