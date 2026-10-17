package helpers

type ChanceContainer[T any] interface {
	Percent() T
	Round(int32) T

}

func SlicePercent[T ChanceContainer[T]](s []T) []T {
	new := make([]T, len(s))
	for i, item := range s {
		new[i] = item.Percent()
	}

	return new
}

func SliceRound[T ChanceContainer[T]](s []T, n int32) []T {
	new := make([]T, len(s))
	for i, item := range s {
		new[i] = item.Round(n)
	}

	return new
}