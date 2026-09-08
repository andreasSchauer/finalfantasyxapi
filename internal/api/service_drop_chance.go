package api

type DropChanceResponse struct {
	URL            		string 					`json:"url"`
	TotalDropChances	TotalDropChances		`json:"total_drop_chances"`
	IsolatedDropChances IsolatedDropChances		`json:"isolated_drop_chances"`
}

func (r DropChanceResponse) GetURL() string {
	return r.URL
}

type TotalDropChances struct {
	WithFinBlow		float64		`json:"with_fin_blow"`
	WithoutFinBlow	float64		`json:"without_fin_blow"`
}

type IsolatedDropChances struct {
	EquipmentDrop				float64		`json:"equipment_drop"`
	EquipmentMatch				float64		`json:"equipment_match"`
	AnyChar						float64		`json:"any_char"`
	MatchingCharFinBlow			float64		`json:"matching_char_fin_blow"`
	MatchingCharWithoutFinBlow	float64		`json:"matching_char_without_fin_blow"`
}

func calcDropChance(cfg *Config, params DropChanceParams, url string) (DropChanceResponse, error) {
	response := DropChanceResponse{
		URL: url,
	}
	
	return response, nil
}