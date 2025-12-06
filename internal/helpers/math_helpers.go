package helpers


func PowInt(x, y int32) int32 {
	var i int32 = 1
	var result int32 = 1

	for i <= y {
		result *= x
		i++
	}

	return result
}