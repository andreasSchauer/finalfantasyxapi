package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type StatusResist struct {
	StatusCondition NamedAPIResource `json:"status_condition"`
	Resistance      int32            `json:"resistance"`
}

func (sr StatusResist) GetAPIResource() APIResource {
	return sr.StatusCondition
}

func convertStatusResist(cfg *Config, sr seeding.StatusResist) StatusResist {
	return StatusResist{
		StatusCondition: nameToNamedAPIResource(cfg, cfg.e.statusConditions, sr.StatusCondition, nil),
		Resistance:      sr.Resistance,
	}
}

func newStatusResist(res NamedAPIResource, resistance int32) StatusResist {
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

func (sr StatusResist) ToResAmount() ResourceAmount[NamedAPIResource] {
	return ResourceAmount[NamedAPIResource]{
		Resource: sr.StatusCondition,
		Amount:   sr.Resistance,
	}
}
