package helpers

func PtrsEqual[T comparable](a, b *T, derefFn func(*T) T) bool {
	if a == nil && b == nil {
		return true
	}

	return derefFn(a) == derefFn(b)
}

func DerefStringPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func DerefInt32Ptr(i *int32) int32 {
	if i == nil {
		return 0
	}

	return *i
}