package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
)

// parses the given stats and generates a complete stat map where each given stat is compared to the entity's given stats, usually the default stats. the higher value will be preferred, since lower than the default isn't possible. if no value for a stat was given, the default stat will be used. the result of this function is usually used to calculate the stats of another entity, not to replace any stats of the given entity.
func getStatCalcMap(r *http.Request, queryParam queryparam.QueryParam, baseStats []apitypes.BaseStat, allowedStatIDs []int32) (map[string]int32, error) {
	if allowedStatIDs == nil {
		allowedStatIDs = h.GetNumSlice(1, int32(len(lookup.L.Stats)))
	}

	statMap, err := queryParam.ParseStats(r, allowedStatIDs)
	if err != nil {
		return nil, err
	}
	statMap = fillStatMapWithDefaults(statMap, baseStats)

	return statMap, nil
}

func fillStatMapWithDefaults(queryStatMap map[string]int32, baseStats []apitypes.BaseStat) map[string]int32 {
	statMap := make(map[string]int32)
	for _, baseStat := range baseStats {
		statName := baseStat.GetName()
		statMap[statName] = getStatValOrDefault(queryStatMap, statName, baseStat)
	}

	return statMap
}

func getStatValOrDefault(queryStatMap map[string]int32, statName string, baseStat apitypes.BaseStat) int32 {
	newVal, ok := queryStatMap[statName]
	if ok {
		return max(newVal, baseStat.Value)
	}
	return baseStat.Value
}
