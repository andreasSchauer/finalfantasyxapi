package helpers

import "math"

func PowInt(x, y int32) int32 {
	var i int32 = 1
	var result int32 = 1

	for i <= y {
		result *= x
		i++
	}

	return result
}

func FloatRound(num float64, digits int32) float64 {
	denominator := float64(PowInt(10, digits))

	return math.Round(num * denominator) / denominator
}

func FloatPtrRound(ptr *float64, digits int32) *float64 {
	if ptr == nil {
		return nil
	}

	rounded := FloatRound(*ptr, digits)
	return &rounded
}

func FloatLen[T any](arr []T) float64 {
	return float64(len(arr))
}

func Len32[T any](arr []T) int32 {
	return int32(len(arr))
}