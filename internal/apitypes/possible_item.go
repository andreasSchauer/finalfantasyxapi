package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type PossibleItem struct {
	apiresource.ResourceAmount[apiresource.TypedAPIResource]
	Chance int32 `json:"chance"`
}

func (ps PossibleItem) GetAPIResource() apiresource.APIResource {
	return ps.Resource.GetAPIResource()
}