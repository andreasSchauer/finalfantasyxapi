package typehelpers


func GetEqualICVs(val int32) (*int32, *int32) {
	minICV := val
	maxICV := val

	return &minICV, &maxICV
}