package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type StatusResist struct {
	StatusCondition apiresource.NamedAPIResource `json:"status_condition"`
	Resistance      int32                        `json:"resistance"`
}

func (sr StatusResist) GetAPIResource() apiresource.APIResource {
	return sr.StatusCondition
}

func (sr StatusResist) GetName() string {
	return sr.StatusCondition.Name
}

func (sr StatusResist) GetVersion() *int32 {
	return nil
}

func (sr StatusResist) GetVal() int32 {
	return sr.Resistance
}

func (sr StatusResist) ToResAmount() apiresource.ResourceAmount[apiresource.NamedAPIResource] {
	return apiresource.ResourceAmount[apiresource.NamedAPIResource]{
		Resource: sr.StatusCondition,
		Amount:   sr.Resistance,
	}
}
