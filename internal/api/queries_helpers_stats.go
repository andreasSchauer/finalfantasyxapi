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

func (q QueryParam) ParseStats(r *http.Request, allowedStatIDs []int32) (map[string]int32, error) {
	query, err := q.CheckEmpty(r)
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