package main

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) applyRonsoStats(r *http.Request, mon Monster, queryName string) ([]BaseStat, error) {
	allowedStatIDs := []int32{1, 3, 5, 7}
	baseStats := mon.BaseStats
	queryParam := cfg.q.monsters[queryName]

	kimahri, _ := seeding.GetResource("kimahri", cfg.l.Characters)
	kimahriBS := namesToResourceAmounts(cfg, cfg.e.stats, kimahri.BaseStats, cfg.newBaseStat)

	kimahriStatMap, err := cfg.parseStatQuery(r, queryParam, kimahriBS, allowedStatIDs)
	if errors.Is(err, errEmptyQuery) {
		return baseStats, nil
	}
	if err != nil {
		return nil, err
	}

	ronsoStats := getRonsoStats(mon, kimahriStatMap)

	return ronsoStats, nil
}

func getRonsoStats(mon Monster, kimahriStatMap map[string]int32) []BaseStat {
	ronsoStatMap := make(map[string]int32)
	ronsoStatMap["hp"] = getRonsoHP(mon, kimahriStatMap)
	ronsoStatMap["strength"] = getRonsoStrength(mon, kimahriStatMap)
	ronsoStatMap["magic"] = getRonsoMagic(mon, kimahriStatMap)
	ronsoStatMap["agility"] = getRonsoAgility(mon, kimahriStatMap)

	return replaceBaseStats(mon.BaseStats, ronsoStatMap)
}

func getRonsoHP(mon Monster, kimahriStatMap map[string]int32) int32 {
	kimahriStr := kimahriStatMap["strength"]
	kimahriMag := kimahriStatMap["magic"]

	v1 := float64(h.PowInt(kimahriStr, 3))
	v2 := float64(h.PowInt(kimahriMag, 3))
	v3 := (v1 + v2) / 2 * 16 / 15

	hpMod := ((int32(v3)/32)+30)*586/730 + 1

	if mon.Name == "biran ronso" {
		return int32(hpMod) * 8
	}

	if mon.Name == "yenke ronso" {
		return int32(hpMod) * 6
	}

	return 0
}

func getRonsoStrength(mon Monster, kimahriStatMap map[string]int32) int32 {
	kimahriHP := kimahriStatMap["hp"]
	strengthVals := []int32{11, 12, 13, 15, 17, 19, 21, 22, 23, 24, 25, 27}

	powerMod := min((kimahriHP-644)/200, 11)

	strength := strengthVals[powerMod]

	if mon.Name == "yenke ronso" {
		strength /= 2
	}

	return strength
}

func getRonsoMagic(mon Monster, kimahriStatMap map[string]int32) int32 {
	kimahriHP := kimahriStatMap["hp"]
	magicVals := []int32{8, 8, 9, 10, 12, 14, 16, 17, 19, 20, 21, 22}

	powerMod := min((kimahriHP-644)/200, 11)

	magic := magicVals[powerMod]

	if mon.Name == "biran ronso" {
		magic /= 2
	}

	return magic
}

func getRonsoAgility(mon Monster, kimahriStatMap map[string]int32) int32 {
	var agility int32
	kimahriAgility := kimahriStatMap["agility"]

	if mon.Name == "biran ronso" {
		agility = max(kimahriAgility-4, 1)
	}

	if mon.Name == "yenke ronso" {
		agility = max(kimahriAgility-6, 1)
	}

	return agility
}
