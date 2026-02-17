package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"



type Accuracy struct {
	AccSource   string   `json:"acc_source"`
	HitChance   *int32   `json:"hit_chance"`
	AccModifier *float32 `json:"acc_modifier"`
}

func convertAccuracy(_ *Config, a seeding.Accuracy) Accuracy {
	return Accuracy{
		AccSource: 		a.AccSource,
		HitChance: 		a.HitChance,
		AccModifier: 	a.AccModifier,
	}
}