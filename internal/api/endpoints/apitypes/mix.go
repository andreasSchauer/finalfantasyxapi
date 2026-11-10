package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type Mix struct {
	ID           int32             `json:"id"`
	Name         string            `json:"name"`
	Category     string            `json:"category"`
	Overdrive    apiresource.Named `json:"overdrive"`
	Description  string            `json:"description"`
	Effect       string            `json:"effect"`
	Combinations []MixCombination  `json:"combinations"`
}

type MixCombination struct {
	FirstItem  apiresource.Named `json:"first_item"`
	SecondItem apiresource.Named `json:"second_item"`
}
