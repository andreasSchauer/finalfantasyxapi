package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type AgilityTier struct {
	ID              int32				`json:"id"`
	FromAgility     int32            	`json:"from_agility"`
	ToAgility       int32            	`json:"to_agility"`
	TickSpeed       int32            	`json:"tick_speed"`
	MonMinICV    	*int32           	`json:"mon_min_icv"`
	MonMaxICV    	*int32           	`json:"mon_max_icv"`
	CharMaxICV  	*int32           	`json:"char_max_icv"`
	CharMinICVs 	[]AgilitySubtier 	`json:"char_min_icvs"`
}


type AgilitySubtier struct {
	FromAgility int32  `json:"from_agility"`
	ToAgility  	int32  `json:"to_agility"`
	ICV 		*int32 `json:"icv"`
}

func convertAgilitySubtier(cfg *Config, st seeding.AgilitySubtier) AgilitySubtier {
	return AgilitySubtier{
		FromAgility: 	st.MinAgility,
		ToAgility: 		st.MaxAgility,
		ICV: 			st.CharacterMinICV,
	}
}