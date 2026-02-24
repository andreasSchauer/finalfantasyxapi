package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type BaseStat struct {
	Stat  NamedAPIResource `json:"stat"`
	Value int32            `json:"value"`
}

func (bs BaseStat) GetAPIResource() APIResource {
	return bs.Stat
}

func newBaseStat(res NamedAPIResource, value int32) BaseStat {
	return BaseStat{
		Stat:  res,
		Value: value,
	}
}

func (bs BaseStat) GetName() string {
	return bs.Stat.Name
}

func (bs BaseStat) GetVersion() *int32 {
	return nil
}

func (bs BaseStat) GetVal() int32 {
	return bs.Value
}

func getBaseStat(cfg *Config, stat string, baseStats []BaseStat) BaseStat {
	statLookup, _ := seeding.GetResource(stat, cfg.l.Stats)
	statMap := getResourceMap(baseStats)

	return statMap[statLookup.ID]
}

func replaceBaseStats(baseStats []BaseStat, statMap map[string]int32) []BaseStat {
	for i, baseStat := range baseStats {
		newVal, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value = newVal
		}
	}

	return baseStats
}




type ElementalResist struct {
	Element  NamedAPIResource `json:"element"`
	Affinity NamedAPIResource `json:"affinity"`
}

func (er ElementalResist) GetAPIResource() APIResource {
	return er.Element
}

func newElemResist(cfg *Config, element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(cfg, cfg.e.elements, element, nil),
		Affinity: nameToNamedAPIResource(cfg, cfg.e.affinities, affinity, nil),
	}
}

func namesToElemResists(cfg *Config, resists []seeding.ElementalResist) []ElementalResist {
	elemResists := []ElementalResist{}

	for _, seedResist := range resists {
		elemResist := newElemResist(cfg, seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}




type StatusResist struct {
	StatusCondition NamedAPIResource `json:"status_condition"`
	Resistance      int32            `json:"resistance"`
}

func (sr StatusResist) GetAPIResource() APIResource {
	return sr.StatusCondition
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

