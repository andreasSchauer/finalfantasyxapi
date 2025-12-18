package main

func filter[T any](items []T, filterFunc func(T) bool) []T {
	var newSlice []T

	for _, item := range items {
		if filterFunc(item) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}
