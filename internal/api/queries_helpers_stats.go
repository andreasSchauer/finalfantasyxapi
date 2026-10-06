package api

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func parseStatQuery(cfg *Config, r *http.Request, queryParam QueryParam, allowedStatIDs []int32) (map[string]int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	statMap := make(map[string]int32)
	statKeyValuePairs := strings.SplitSeq(query, ",")

	for pair := range statKeyValuePairs {
		stat, value, err := parseStatPair(cfg, pair, queryParam, allowedStatIDs)
		if err != nil {
			return nil, err
		}

		stat = h.GetNameWithSpaces(stat, "_")
		statMap[stat] = int32(value)
	}

	return statMap, nil
}

func parseStatPair(cfg *Config, pair string, queryParam QueryParam, allowedStatIDs []int32) (string, int, error) {
	stat, valueStr, found := strings.Cut(pair, ":")
	if !found {
		return "", 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s' . usage: '%s'.", queryParam.Name, stat, queryParam.Usage), nil)
	}

	err := validateQueryStatName(cfg, stat, allowedStatIDs, queryParam)
	if err != nil {
		return "", 0, err
	}

	value, err := validateQueryStatVal(stat, valueStr, queryParam)
	if err != nil {
		return "", 0, err
	}

	return stat, value, nil
}

func validateQueryStatName(cfg *Config, stat string, allowedStatIDs []int32, queryParam QueryParam) error {
	parseResp, err := checkUniqueName(stat, cfg.l.Stats)
	if err != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat: '%s' in '%s'. stat doesn't exist. use '/api/stats' to see existing stats.", stat, queryParam.Name), err)
	}

	if !slices.Contains(allowedStatIDs, parseResp.ID) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat '%s' in '%s'. '%s' only uses %s.", stat, queryParam.Name, queryParam.Name, getAllowedStatString(cfg, allowedStatIDs)), nil)
	}

	return nil
}

func validateQueryStatVal(statName string, valStr string, queryParam QueryParam) (int, error) {
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat value '%s' in '%s'. stat value needs to be a positive integer.", valStr, queryParam.Name), err)
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
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s in '%s' can't be higher than %d.", statName, queryParam.Name, maxStatVal), nil)
	}

	return val, nil
}

func getAllowedStatString(cfg *Config, allowedStatIDs []int32) string {
	stats := []string{}

	for _, id := range allowedStatIDs {
		stat := cfg.l.StatsID[id]
		statFormatted := fmt.Sprintf("'%s'", stat.Name)
		stats = append(stats, statFormatted)
	}

	return strings.Join(stats, ", ")
}

// parses the given stats and generates a complete stat map where each given stat is compared to the entity's given stats, usually the default stats. the higher value will be preferred, since lower than the default isn't possible. if no value for a stat was given, the default stat will be used. the result of this function is usually used to calculate the stats of another entity, not to replace any stats of the given entity.
func getStatCalcMap(cfg *Config, r *http.Request, queryParam QueryParam, baseStats []BaseStat, allowedStatIDs []int32) (map[string]int32, error) {
	if allowedStatIDs == nil {
		allowedStatIDs = h.GetNumSlice(1, int32(len(cfg.l.Stats)))
	}

	statMap, err := parseStatQuery(cfg, r, queryParam, allowedStatIDs)
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