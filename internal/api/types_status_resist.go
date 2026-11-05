package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type StatusResist struct {
	StatusCondition apiresource.NamedAPIResource `json:"status_condition"`
	Resistance      int32                        `json:"resistance"`
}

func (sr StatusResist) GetAPIResource() apiresource.APIResource {
	return sr.StatusCondition
}

func convertStatusResist(sr seedtypes.StatusResist) StatusResist {
	return StatusResist{
		StatusCondition: nameToNamedAPIResource(E.StatusConditions, sr.StatusCondition, nil),
		Resistance:      sr.Resistance,
	}
}

func newStatusResist(res apiresource.NamedAPIResource, resistance int32) StatusResist {
	return StatusResist{
		StatusCondition: res,
		Resistance:      resistance,
	}
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
