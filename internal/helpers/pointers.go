package helpers

func GetIntPtr(i int) *int {
	return &i
}

func GetInt32Ptr(i int32) *int32 {
	return &i
}

func GetFloat32Ptr(f float32) *float32 {
	return &f
}

func GetStrPtr(s string) *string {
	return &s
}

func GetStructPtr[T any](obj T) *T {
	return &obj
}

func StringPtrToString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}