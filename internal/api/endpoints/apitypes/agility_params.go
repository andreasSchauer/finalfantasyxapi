package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type AgilityParams struct {
	AgilityTier apiresource.Unnamed `json:"agility_tier"`
	TickSpeed   int32               `json:"tick_speed"`
	MinICV      *int32              `json:"min_icv"`
	MaxICV      *int32              `json:"max_icv"`
}
