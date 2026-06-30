package api

import (

	"slices"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EnumApiResourceList struct {
	ListParams
	Results []EnumVal `json:"results"`
}

func (l EnumApiResourceList) getListParams() ListParams {
	return l.ListParams
}

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


// Searches an EnumAPIResource based on its value or an idString (mostly from queries)
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
	resources := []EnumVal{}

	for _, resource := range lookup {
		resources = append(resources, resource)
	}

	slices.SortStableFunc(resources, h.SortOnId)

	return resources
}

func enumSliceToMap(enumTypes []EnumVal) map[string]EnumVal {
	typeMap := make(map[string]EnumVal)

	for i, enumType := range enumTypes {
		typeMap[enumType.Name] = EnumVal{
			ID:          int32(i + 1),
			Name:        enumType.Name,
			Description: enumType.Description,
		}
	}

	return typeMap
}
