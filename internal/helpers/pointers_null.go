package helpers

type Zeroable interface {
	IsZero() bool
}

// Use derefOrNil for every field that is a pointer
func DerefOrNil[T any](ptr *T) any {
	if ptr == nil {
		return nil
	}
	return *ptr
}

func ObjPtrOrNil[T Zeroable](v T) *T {
	if v.IsZero() {
		return nil
	}
	return &v
}

func SliceOrNil[T any](s []T) []T {
	if len(s) == 0 {
		return nil
	}
	return s
}

func PtrIsNotNil[T any](ptr *T) bool {
	return ptr != nil
}
