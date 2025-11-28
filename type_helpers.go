package main


func anyToFloat32Ptr(value any) *float32 {
	switch t := any(value).(type) {
	case float64:
		val := float32(t)
		return &val
	default:
		return nil
	}
}