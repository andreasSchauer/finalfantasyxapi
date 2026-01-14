package main

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type BaseStat struct {
	Stat  NamedAPIResource `json:"stat"`
	Value int32            `json:"value"`
}

func (bs BaseStat) GetAPIResource() APIResource {
	return bs.Stat
}

func (cfg *Config) newBaseStat(id, value int32, name string) BaseStat {
	return BaseStat{
		Stat:  cfg.newNamedAPIResourceSimple(cfg.e.stats.endpoint, id, name),
		Value: value,
	}
}

func (bs BaseStat) GetName() string {
	return bs.Stat.Name
}

func (bs BaseStat) GetVal() int32 {
	return bs.Value
}

type ElementalResist struct {
	Element  NamedAPIResource `json:"element"`
	Affinity NamedAPIResource `json:"affinity"`
}

func (er ElementalResist) GetAPIResource() APIResource {
	return er.Element
}

func (cfg *Config) newElemResist(elem_id, affinity_id int32, element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  cfg.newNamedAPIResourceSimple(cfg.e.elements.endpoint, elem_id, element),
		Affinity: cfg.newNamedAPIResourceSimple(cfg.e.affinities.endpoint, affinity_id, affinity),
	}
}

type StatusResist struct {
	StatusCondition NamedAPIResource `json:"status_condition"`
	Resistance      int32            `json:"resistance"`
}

func (sr StatusResist) GetAPIResource() APIResource {
	return sr.StatusCondition
}

func (cfg *Config) newStatusResist(id, resistance int32, status string) StatusResist {
	return StatusResist{
		StatusCondition: cfg.newNamedAPIResourceSimple(cfg.e.statusConditions.endpoint, id, status),
		Resistance:      resistance,
	}
}

func (sr StatusResist) GetName() string {
	return sr.StatusCondition.Name
}

func (sr StatusResist) GetVal() int32 {
	return sr.Resistance
}

type InflictedStatus struct {
	StatusCondition NamedAPIResource      `json:"status_condition"`
	Probability     int32                 `json:"probability,omitempty"`
	DurationType    database.DurationType `json:"duration_type,omitempty"`
	Amount          *int32                `json:"amount,omitempty"`
}

func (is InflictedStatus) GetAPIResource() APIResource {
	return is.StatusCondition
}

func (is InflictedStatus) IsZero() bool {
	return is.StatusCondition.Name == ""
}

func (cfg *Config) newInflictedStatus(id, probability int32, status string, amount *int32, durationType database.DurationType) InflictedStatus {
	return InflictedStatus{
		StatusCondition: cfg.newNamedAPIResourceSimple(cfg.e.statusConditions.endpoint, id, status),
		Probability:     probability,
		DurationType:    durationType,
		Amount:          amount,
	}
}
