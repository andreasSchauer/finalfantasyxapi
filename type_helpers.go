package main


// these are used if sqlc exports domain values as interface{}

func anyToFloat32Ptr(value any) *float32 {
	switch t := any(value).(type) {
	case float64:
		val := float32(t)
		return &val
	default:
		return nil
	}
}


func anyToInt32(value any) int32 {
	switch t := any(value).(type) {
	case int32:
		return t
	default:
		return 0
	}
}


func anyToInt32Ptr(value any) *int32 {
	switch t := any(value).(type) {
	case int32:
		val := int32(t)
		return &val
	default:
		return nil
	}
}