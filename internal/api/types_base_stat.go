package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func newBaseStat(res apiresource.NamedAPIResource, value int32) apitypes.BaseStat {
	return apitypes.BaseStat{
		Stat:  res,
		Value: value,
	}
}

func getBaseStat(stat string, baseStats []apitypes.BaseStat) apitypes.BaseStat {
	statLookup, _ := lookup.GetResource(stat, lookup.L.Stats)
	statMap := GetResourceMap(baseStats)

	return statMap[statLookup.ID]
}

func getBaseStatVal(stat string, baseStats []apitypes.BaseStat) int32 {
	bs := getBaseStat(stat, baseStats)
	return bs.Value
}

func replaceBaseStats(baseStats []apitypes.BaseStat, statMap map[string]int32, allowedStatIDs []int32) []apitypes.BaseStat {
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

func addToBaseStats(baseStats []apitypes.BaseStat, statMap map[string]int32) []apitypes.BaseStat {
	for i, baseStat := range baseStats {
		toAdd, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value += toAdd
		}
	}

	return baseStats
}
