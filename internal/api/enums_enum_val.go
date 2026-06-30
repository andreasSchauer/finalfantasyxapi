package api

import (
	"slices"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type EnumVal struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (e EnumVal) IsZero() bool {
	return e.ID == 0
}

func (e EnumVal) GetID() int32 {
	return e.ID
}


// Verifies an EnumVal based on its value or an idString (mostly from query params)
func CheckEnumVal[E, N any](key string, et EnumType[E, N]) (EnumVal, error) {
	id, err := strconv.Atoi(key)
	if err == nil {
		if id > len(et.lookup) || id <= 0 {
			return EnumVal{}, errIdNotFound
		}
		for _, res := range et.lookup {
			if int32(id) == res.ID {
				return res, nil
			}
		}
	}

	res, found := et.lookup[key]
	if found {
		return res, nil
	}

	keyWithSpaces := h.GetNameWithSpaces(key, "-")
	res, found = et.lookup[keyWithSpaces]
	if found {
		return res, nil
	}
	
	return EnumVal{}, errNoResource
}


func createEnumValSlice(lookup map[string]EnumVal) []EnumVal {
	vals := []EnumVal{}

	for _, val := range lookup {
		vals = append(vals, val)
	}

	slices.SortStableFunc(vals, h.SortOnId)

	return vals
}

func enumSliceToMap(enumVals []EnumVal) map[string]EnumVal {
	typeMap := make(map[string]EnumVal)

	for i, enumVal := range enumVals {
		typeMap[enumVal.Name] = EnumVal{
			ID:          int32(i + 1),
			Name:        enumVal.Name,
			Description: enumVal.Description,
		}
	}

	return typeMap
}

func getEnumValIDs(enumVals []EnumVal) []EnumVal {
	for i := range enumVals {
		enumVals[i].ID = int32(i + 1)
	}

	return enumVals
}