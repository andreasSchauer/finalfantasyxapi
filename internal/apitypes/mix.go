package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type Mix struct {
	ID           int32                        `json:"id"`
	Name         string                       `json:"name"`
	Category     string                       `json:"category"`
	Overdrive    apiresource.NamedAPIResource `json:"overdrive"`
	Description  string                       `json:"description"`
	Effect       string                       `json:"effect"`
	Combinations []MixCombination             `json:"combinations"`
}

type MixCombination struct {
	FirstItem  apiresource.NamedAPIResource `json:"first_item"`
	SecondItem apiresource.NamedAPIResource `json:"second_item"`
}