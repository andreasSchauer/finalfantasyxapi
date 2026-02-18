package api

import (
	"errors"
	"net/http"
)

func applyAeonStatsMonsters(cfg *Config, r *http.Request, mon Monster, queryName string) ([]BaseStat, error) {
	allowedStatIDs := []int32{1, 3, 4, 5, 6, 7, 9, 10}
	aeonBaseStats := mon.BaseStats
	queryParam := cfg.q.monsters[queryName]

	queryStatMap, err := parseStatQuery(cfg, r, queryParam, aeonBaseStats, allowedStatIDs)
	if errors.Is(err, errEmptyQuery) {
		return mon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	newBaseStats := replaceBaseStats(aeonBaseStats, queryStatMap)

	return newBaseStats, nil
}
