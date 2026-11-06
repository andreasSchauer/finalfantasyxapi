package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type AgilityParams struct {
	AgilityTier apiresource.UnnamedAPIResource `json:"agility_tier"`
	TickSpeed   int32                          `json:"tick_speed"`
	MinICV      *int32                         `json:"min_icv"`
	MaxICV      *int32                         `json:"max_icv"`
}
