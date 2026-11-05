package api

type CritChanceResponse struct {
	CritChance float64 `json:"crit_chance"`
}

func handleCritChance(_ CritChanceParams) (CritChanceResponse, error) {
	return CritChanceResponse{}, nil
}
