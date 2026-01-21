package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
)


func parseStatQuery(cfg *Config, r *http.Request, queryParam QueryType, baseStats []BaseStat, allowedStatIDs []int32) (map[string]int32, error) {
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

		statMap[stat] = int32(value)
	}

	statMap = getDefaultStats(statMap, baseStats)

	return statMap, nil
}



func parseStatPair(cfg *Config, pair string, queryParam QueryType, allowedStatIDs []int32) (string, int, error) {
	parts := strings.Split(pair, "-")
	if len(parts) != 2 {
		return "", 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input. usage: '%s'.", queryParam.Usage), nil)
	}

	stat := parts[0]
	valueStr := parts[1]

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


func getDefaultStats(statMap map[string]int32, baseStats []BaseStat) map[string]int32 {
	for _, baseStat := range baseStats {
		statName := baseStat.GetName()
		_, ok := statMap[statName]
		if ok {
			statMap[statName] = max(statMap[statName], baseStat.Value)
		}
	}

	return statMap
}

func validateQueryStatName(cfg *Config, stat string, allowedStatIDs []int32, queryParam QueryType) error {
	parseResp, err := checkUniqueName(stat, cfg.l.Stats)
	if err != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat: '%s' in '%s'. stat doesn't exist. use '/api/stats' to see existing stats.", stat, queryParam.Name), err)
	}

	if !slices.Contains(allowedStatIDs, parseResp.ID) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat '%s' in '%s'. '%s' only uses %s.", stat, queryParam.Name, queryParam.Name, getAllowedStatString(cfg, allowedStatIDs)), nil)
	}

	return nil
}


func validateQueryStatVal(statName string, valStr string, queryParam QueryType) (int, error) {
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