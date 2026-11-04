package api

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func (q QueryParam) parseStats(r *http.Request, allowedStatIDs []int32) (map[string]int32, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return nil, err
	}

	statMap := make(map[string]int32)
	statKeyValuePairs := strings.SplitSeq(query, ",")

	for pair := range statKeyValuePairs {
		stat, value, err := q.parseStatPair(pair, allowedStatIDs)
		if err != nil {
			return nil, err
		}

		stat = h.GetNameWithSpaces(stat, "_")
		statMap[stat] = int32(value)
	}

	return statMap, nil
}

func (q QueryParam) parseStatPair(pair string, allowedStatIDs []int32) (string, int, error) {
	stat, valueStr, found := strings.Cut(pair, ":")
	if !found {
		return "", 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s' . usage: '%s'", q.Name, stat, q.Usage), nil)
	}

	err := q.validateStatName(stat, allowedStatIDs)
	if err != nil {
		return "", 0, err
	}

	value, err := q.validateStatVal(stat, valueStr)
	if err != nil {
		return "", 0, err
	}

	return stat, value, nil
}

func (q QueryParam) validateStatName(stat string, allowedStatIDs []int32) error {
	parseResp, err := checkUniqueName(stat, lookup.L.Stats)
	if err != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat: '%s' in '%s'. stat doesn't exist. use '/api/stats' to see existing stats", stat, q.Name), err)
	}

	if !slices.Contains(allowedStatIDs, parseResp.ID) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat '%s' in '%s'. '%s' only uses %s", stat, q.Name, q.Name, getAllowedStatString(allowedStatIDs)), nil)
	}

	return nil
}

func (q QueryParam) validateStatVal(statName string, valStr string) (int, error) {
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat value '%s' in '%s'. stat value needs to be a positive integer", valStr, q.Name), err)
	}

	var maxStatVal int

	switch statName {
	case "hp":
		maxStatVal = 99999
	case "mp":
		maxStatVal = 9999
	default:
		maxStatVal = 255
	}

	if val > maxStatVal {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s in '%s' can't be higher than %d", statName, q.Name, maxStatVal), nil)
	}

	return val, nil
}

func getAllowedStatString(allowedStatIDs []int32) string {
	stats := []string{}

	for _, id := range allowedStatIDs {
		stat := lookup.L.StatsID[id]
		statFormatted := fmt.Sprintf("'%s'", stat.Name)
		stats = append(stats, statFormatted)
	}

	return strings.Join(stats, ", ")
}

// parses the given stats and generates a complete stat map where each given stat is compared to the entity's given stats, usually the default stats. the higher value will be preferred, since lower than the default isn't possible. if no value for a stat was given, the default stat will be used. the result of this function is usually used to calculate the stats of another entity, not to replace any stats of the given entity.
func getStatCalcMap(r *http.Request, queryParam QueryParam, baseStats []BaseStat, allowedStatIDs []int32) (map[string]int32, error) {
	if allowedStatIDs == nil {
		allowedStatIDs = h.GetNumSlice(1, int32(len(lookup.L.Stats)))
	}

	statMap, err := queryParam.parseStats(r, allowedStatIDs)
	if err != nil {
		return nil, err
	}
	statMap = fillStatMapWithDefaults(statMap, baseStats)

	return statMap, nil
}

func fillStatMapWithDefaults(queryStatMap map[string]int32, baseStats []BaseStat) map[string]int32 {
	statMap := make(map[string]int32)
	for _, baseStat := range baseStats {
		statName := baseStat.GetName()
		statMap[statName] = getStatValOrDefault(queryStatMap, statName, baseStat)
	}

	return statMap
}

func getStatValOrDefault(queryStatMap map[string]int32, statName string, baseStat BaseStat) int32 {
	newVal, ok := queryStatMap[statName]
	if ok {
		return max(newVal, baseStat.Value)
	}
	return baseStat.Value
}
