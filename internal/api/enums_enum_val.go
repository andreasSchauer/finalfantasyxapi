package api

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