package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) applyRonsoStats(r *http.Request, mon Monster) ([]BaseStat, error) {
	baseStats := mon.BaseStats
	queryParam := cfg.q.monsters["kimahri-stats"]

	kimahriStats, err := cfg.getKimahriStats(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return mon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}
	ronsoStats := getRonsoStats(mon, kimahriStats)

	for i, baseStat := range baseStats {
		ronsoStatVal, ok := ronsoStats[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value = ronsoStatVal
		}
	}

	return baseStats, nil
}


func (cfg *Config) getKimahriStats(r *http.Request, queryParam QueryType) (map[string]int32, error) {
	statMap := make(map[string]int32)
	query := r.URL.Query().Get(queryParam.Name)
	queryLower := strings.ToLower(query)

	if query == "" {
		return nil, errEmptyQuery
	}

	statKeyValuePairs := strings.SplitSeq(queryLower, ",")

	for pair := range statKeyValuePairs {
		parts := strings.Split(pair, "-")
		if len(parts) != 2 {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input. usage: '%s'.", queryParam.Usage), nil)
		}

		stat := parts[0]
		valueStr := parts[1]

		statLookup, err := seeding.GetResource(stat, cfg.l.Stats)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat: '%s' in '%s'. stat doesn't exist. use '/api/stats' to see existing stats.", stat, queryParam.Name), err)
		}

		switch statLookup.ID {
		case 2, 4, 6, 8, 9, 10:
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat '%s' in '%s'. '%s' only uses 'hp', 'strength', 'magic', 'agility'.", stat, queryParam.Name, queryParam.Name), nil)
		}

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat value '%s' in '%s'. stat value needs to be a positive integer.", valueStr, queryParam.Name), err)
		}

		err = validateKimahriStat(stat, value)
		if err != nil {
			return nil, err
		}

		statMap[stat] = int32(value)
	}

	statMap = getKimahriDefaultStats(statMap)

	return statMap, nil
}


func validateKimahriStat(key string, val int) error {
	maxHP := 99999
	maxStatVal := 255

	switch key {
	case "hp":
		if val > maxHP {
			return newHTTPError(http.StatusBadRequest, "kimahri's hp can't be higher than 99999.", nil)
		}
	case "strength", "magic", "agility":
		if val > maxStatVal {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("kimahri's %s can't be higher than 255.", key), nil)
		}
	}

	return nil
}


func getKimahriDefaultStats(kimahriStats map[string]int32) map[string]int32 {
	var hpDefault int32 = 644
	var strDefault int32 = 16
	var magDefault int32 = 17
	var agilDefault int32 = 6

	kimahriStats["hp"] = max(kimahriStats["hp"], hpDefault)
	kimahriStats["strength"] = max(kimahriStats["strength"], strDefault)
	kimahriStats["magic"] = max(kimahriStats["magic"], magDefault)
	kimahriStats["agility"] = max(kimahriStats["agility"], agilDefault)

	return kimahriStats
}
