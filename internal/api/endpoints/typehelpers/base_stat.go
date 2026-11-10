// Package typehelpers asdüijasdpüojp
package typehelpers

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func NewBaseStat(res apiresource.Named, value int32) apitypes.BaseStat {
	return apitypes.BaseStat{
		Stat:  res,
		Value: value,
	}
}

func GetBaseStat(stat string, baseStats []apitypes.BaseStat) apitypes.BaseStat {
	statLookup, _ := lookup.GetResource(stat, lookup.L.Stats)
	statMap := handlerinput.GetResourceMap(baseStats)

	return statMap[statLookup.ID]
}

func GetBaseStatVal(stat string, baseStats []apitypes.BaseStat) int32 {
	bs := GetBaseStat(stat, baseStats)
	return bs.Value
}

func ReplaceBaseStats(baseStats []apitypes.BaseStat, statMap map[string]int32, allowedStatIDs []int32) []apitypes.BaseStat {
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

func AddToBaseStats(baseStats []apitypes.BaseStat, statMap map[string]int32) []apitypes.BaseStat {
	for i, baseStat := range baseStats {
		toAdd, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value += toAdd
		}
	}

	return baseStats
}

func GetAeonBaseStats(aeon seedtypes.Aeon) []apitypes.BaseStat {
	baseStats := aeon.BaseStats.XVals[0].BaseStats
	return apilookups.E.Stats.ToResAmtType(baseStats, NewBaseStat)
}
