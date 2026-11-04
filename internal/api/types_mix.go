package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

type Mix struct {
	ID           int32            `json:"id"`
	Name         string           `json:"name"`
	Category     string           `json:"category"`
	Overdrive    NamedAPIResource `json:"overdrive"`
	Description  string           `json:"description"`
	Effect       string           `json:"effect"`
	Combinations []MixCombination `json:"combinations"`
}

type MixCombination struct {
	FirstItem  NamedAPIResource `json:"first_item"`
	SecondItem NamedAPIResource `json:"second_item"`
}

func convertMixCombination(server *Server, mc seedtypes.MixCombination) MixCombination {
	return MixCombination{
		FirstItem:  nameToNamedAPIResource(server, E.Items, mc.FirstItem, nil),
		SecondItem: nameToNamedAPIResource(server, E.Items, mc.SecondItem, nil),
	}
}
