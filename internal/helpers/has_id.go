package helpers

type HasID interface {
	GetID() int32
}

func SortOnId[T HasID](a, b T) int {
	if a.GetID() < b.GetID() {
		return -1
	}

	if a.GetID() > b.GetID() {
		return 1
	}

	return 0
}