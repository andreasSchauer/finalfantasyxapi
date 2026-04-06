package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type Mix struct {
	ID                 	int32               `json:"id"`
	Name               	string              `json:"name"`
	Category           	NamedAPIResource    `json:"category"`
	Overdrive          	NamedAPIResource    `json:"overdrive"`
	Description        	string              `json:"description"`
	Effect             	string              `json:"effect"`
	Combinations		[]MixCombination	`json:"combinations"`
}

type MixCombination struct {
	FirstItem	NamedAPIResource `json:"first_item"`
	SecondItem	NamedAPIResource `json:"second_item"`
}

func convertMixCombination(cfg *Config, mc seeding.MixCombination) MixCombination {
	return MixCombination{
		FirstItem: 	nameToNamedAPIResource(cfg, cfg.e.items, mc.FirstItem, nil),
		SecondItem: nameToNamedAPIResource(cfg, cfg.e.items, mc.SecondItem, nil),
	}
}