package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type BaseStat struct {
	Stat  NamedAPIResource `json:"stat"`
	Value int32            `json:"value"`
}

func (bs BaseStat) getAPIResource() IsAPIResource {
	return bs.Stat
}

func (cfg *Config) newBaseStat(id, value int32, name string) BaseStat {
	return BaseStat{
		Stat:  cfg.newNamedAPIResourceSimple("stats", id, name),
		Value: value,
	}
}

type ElementalResist struct {
	Element  NamedAPIResource `json:"element"`
	Affinity NamedAPIResource `json:"affinity"`
}

func (er ElementalResist) getAPIResource() IsAPIResource {
	return er.Element
}

func (cfg *Config) newElemResist(elem_id, affinity_id int32, element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  cfg.newNamedAPIResourceSimple("elements", elem_id, element),
		Affinity: cfg.newNamedAPIResourceSimple("affinities", affinity_id, affinity),
	}
}

type StatusResist struct {
	StatusCondition NamedAPIResource `json:"status_condition"`
	Resistance      int32            `json:"resistance"`
}

func (sr StatusResist) getAPIResource() IsAPIResource {
	return sr.StatusCondition
}

func (cfg *Config) newStatusResist(id, resistance int32, status string) StatusResist {
	return StatusResist{
		StatusCondition: cfg.newNamedAPIResourceSimple("status-conditions", id, status),
		Resistance:      resistance,
	}
}

type InflictedStatus struct {
	StatusCondition NamedAPIResource      `json:"status_condition"`
	Probability     int32                 `json:"probability,omitempty"`
	DurationType    database.DurationType `json:"duration_type,omitempty"`
	Amount          *int32                `json:"amount,omitempty"`
}

func (is InflictedStatus) getAPIResource() IsAPIResource {
	return is.StatusCondition
}

func (is InflictedStatus) IsZero() bool {
	return is.StatusCondition.Name == ""
}

func (cfg *Config) newInflictedStatus(id, probability int32, status string, amount *int32, durationType database.DurationType) InflictedStatus {
	return InflictedStatus{
		StatusCondition: cfg.newNamedAPIResourceSimple("status-conditions", id, status),
		Probability:     probability,
		DurationType:    durationType,
		Amount:          amount,
	}
}
