package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type BaseStat struct {
	Stat  apiresource.NamedAPIResource `json:"stat"`
	Value int32                        `json:"value"`
}

func (bs BaseStat) GetAPIResource() apiresource.APIResource {
	return bs.Stat
}

func newBaseStat(res apiresource.NamedAPIResource, value int32) BaseStat {
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

func (bs BaseStat) ToResAmount() apiresource.ResourceAmount[apiresource.NamedAPIResource] {
	return apiresource.ResourceAmount[apiresource.NamedAPIResource]{
		Resource: bs.Stat,
		Amount:   bs.Value,
	}
}

func getBaseStat(stat string, baseStats []BaseStat) BaseStat {
	statLookup, _ := lookup.GetResource(stat, lookup.L.Stats)
	statMap := getResourceMap(baseStats)

	return statMap[statLookup.ID]
}

func getBaseStatVal(stat string, baseStats []BaseStat) int32 {
	bs := getBaseStat(stat, baseStats)
	return bs.Value
}

func replaceBaseStats(baseStats []BaseStat, statMap map[string]int32, allowedStatIDs []int32) []BaseStat {
	for i, baseStat := range baseStats {
		if !slices.Contains(allowedStatIDs, baseStat.Stat.ID) {
			continue
		}

		newVal, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value = newVal
		}
	}

	return baseStats
}

func addToBaseStats(baseStats []BaseStat, statMap map[string]int32) []BaseStat {
	for i, baseStat := range baseStats {
		toAdd, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value += toAdd
		}
	}

	return baseStats
}
