package main


func convertObjPtr[Old, New any](cfg *Config, item *Old, constructor func(*Config, Old) New) *New {
	if item == nil {
		return nil
	}

	new := constructor(cfg, *item)
	return &new
}

func convertObjSlice[Old, New any](cfg *Config, items []Old, constructor func(*Config, Old) New) []New {
	newSlice := []New{}

	for _, item := range items {
		new := constructor(cfg, item)
		newSlice = append(newSlice, new)
	}

	return newSlice
}

func convertObjSliceNullable[Old, New any](cfg *Config, items []Old, constructor func(*Config, Old) New) []New {
	slice := convertObjSlice(cfg, items, constructor)

	if len(slice) == 0 {
		return nil
	}

	return slice
}