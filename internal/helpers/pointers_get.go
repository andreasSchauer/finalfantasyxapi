package helpers

func GetIntPtr(i int) *int {
	return new(i)
}

func GetInt32Ptr(i int32) *int32 {
	return new(i)
}

func GetFloat32Ptr(f float32) *float32 {
	return new(f)
}

func GetFloat64Ptr(f float64) *float64 {
	return new(f)
}

func GetStrPtr(s string) *string {
	return new(s)
}

func GetBoolPtr(b bool) *bool {
	return new(b)
}

func GetStructPtr[T any](obj T) *T {
	return new(obj)
}