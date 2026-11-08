package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type ElementalResist struct {
	Element  apiresource.Named `json:"element"`
	Affinity string            `json:"affinity"`
}

func (er ElementalResist) GetAPIResource() apiresource.APIResource {
	return er.Element
}
