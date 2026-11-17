package helpers

func PtrsEqual[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}

	return DerefPtr(a) == DerefPtr(b)
}

func DerefPtr[T any](p *T) T {
	var zero T

	if p == nil {
		return zero
	}

	return *p
}

func CopyPtr[T any](p *T) *T {
	if p == nil {
		return nil
	}

	return new(*p)
}