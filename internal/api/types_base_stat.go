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

func (bs BaseStat) ToResAmount() ResourceAmount[NamedAPIResource] {
	return ResourceAmount[NamedAPIResource]{
		Resource: bs.Stat,
		Amount:   bs.Value,
	}
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


