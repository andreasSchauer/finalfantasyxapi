package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func applyStatQuery(cfg *Config, r *http.Request, queryParam QueryParam, baseStats []BaseStat, allowedStatIDs []int32) ([]BaseStat, error) {
	queryStatMap, err := parseStatQuery(cfg, r, queryParam, allowedStatIDs)
	if queryIsEmpty(err) {
		return baseStats, nil
	}
	if err != nil {
		return nil, err
	}

	newBaseStats := replaceBaseStats(baseStats, queryStatMap, allowedStatIDs)

	return newBaseStats, nil
}

func applyAeonStatsMonsters(cfg *Config, r *http.Request, mon Monster) ([]BaseStat, error) {
	allowedStatIDs := []int32{1, 3, 4, 5, 6, 7, 9, 10}

	aeonLookup, _ := seeding.GetResource(mon.Name, cfg.l.Aeons)
	aeon := Aeon{
		ID: aeonLookup.ID,
		Name: aeonLookup.Name,
		BaseStats: getAeonBaseStats(cfg, aeonLookup),
	}

	aeon, err := applyAeonStats(cfg, r, aeon, nil)
	if err != nil {
		return nil, err
	}
	
	aeonStatMap := getResAmtTypeMap(aeon.BaseStats)
	newBaseStats := replaceBaseStats(mon.BaseStats, aeonStatMap, allowedStatIDs)

	return newBaseStats, nil
}

func applyStatsOverrideCharacters(cfg *Config, r *http.Request, char Character, queryName QueryParamName) ([]BaseStat, error) {
	queryParam := cfg.q.characters[queryName]

	return applyStatQuery(cfg, r, queryParam, char.BaseStats, nil)
}

func applyStatsOverrideAeons(cfg *Config, r *http.Request, aeon Aeon, queryName QueryParamName) ([]BaseStat, error) {
	queryParam := cfg.q.aeons[queryName]

	return applyStatQuery(cfg, r, queryParam, aeon.BaseStats, nil)
}
