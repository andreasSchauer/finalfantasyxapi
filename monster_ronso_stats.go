package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



func (cfg *apiConfig) applyRonsoStats(r *http.Request, mon Monster) ([]BaseStat, error) {
	baseStats := mon.BaseStats

	if mon.Name != "biran ronso" && mon.Name != "yenke ronso" {
		return mon.BaseStats, nil
	}

	query := r.URL.Query().Get("kimahri-stats")
	
	if query == "" {
		return mon.BaseStats, nil
	}

	kimahriStats, err := cfg.getKimahriStats(query)
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


func (cfg *apiConfig) getKimahriStats(query string) (map[string]int32, error) {
	statMap := make(map[string]int32)
	
	statKeyValuePairs := strings.SplitSeq(query, ",")

	for pair := range statKeyValuePairs {
		parts := strings.Split(pair, "-")
		if len(parts) != 2 {
			return nil, newHTTPError(http.StatusBadRequest, "invalid input. usage: kimahri-stats={stat}-{value},{stat}-{value}", nil)
		}

		stat := parts[0]
		valueStr := parts[1]

		statLookup, err := seeding.GetResource(stat, cfg.l.Stats)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat: %s. stat doesn't exist", stat), err)
		}

		switch statLookup.ID {
		case 2, 4, 6, 8, 9, 10:
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid stat: %s. kimahri-stats query parameter only uses hp, strength, magic, agility", stat), nil)
		}

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, "stat value needs to be a positive integer", err)
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
			return newHTTPError(http.StatusBadRequest, "kimahri's HP can't be higher than 99999", nil)
		}
	case "strength", "magic", "agility":
		if val > maxStatVal {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("kimahri's %s can't be higher than 255", key), nil)
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


func getRonsoStats(mon Monster, kimahriStats map[string]int32) map[string]int32 {
	ronsoStats := make(map[string]int32)
	
	ronsoStats["hp"] = getRonsoHP(mon, kimahriStats)
	ronsoStats["strength"] = getRonsoStrength(mon, kimahriStats)
	ronsoStats["magic"] = getRonsoMagic(mon, kimahriStats)
	ronsoStats["agility"] = getRonsoAgility(mon, kimahriStats)

	return ronsoStats
}


func getRonsoHP(mon Monster, kimahriStats map[string]int32) int32 {
	kimahriStr := kimahriStats["strength"]
	kimahriMag := kimahriStats["magic"]

	v1 := float64(h.PowInt(kimahriStr, 3))
	v2 := float64(h.PowInt(kimahriMag, 3))
	v3 := (v1 + v2) / 2 * 16 / 15

	hpMod := ((int32(v3) / 32) + 30) * 586 / 730 + 1

	if mon.Name == "biran ronso" {
		return int32(hpMod) * 8
	}

	if mon.Name == "yenke ronso" {
		return int32(hpMod) * 6
	}

	return 0
}


func getRonsoStrength(mon Monster, kimahriStats map[string]int32) int32 {
	kimahriHP := kimahriStats["hp"]
	strengthVals := []int32{11, 12, 13, 15, 17, 19, 21, 22, 23, 24, 25, 27}

	powerMod := min((kimahriHP - 644) / 200, 11)

	strength := strengthVals[powerMod]

	if mon.Name == "yenke ronso" {
		strength /= 2
	}

	return strength
}


func getRonsoMagic(mon Monster, kimahriStats map[string]int32) int32 {
	kimahriHP := kimahriStats["hp"]
	magicVals := []int32{8, 8, 9, 10, 12, 14, 16, 17, 19, 20, 21, 22}

	powerMod := min((kimahriHP - 644) / 200, 11)

	magic := magicVals[powerMod]

	if mon.Name == "biran ronso" {
		magic /= 2
	}

	return magic
}

func getRonsoAgility(mon Monster, kimahriStats map[string]int32) int32 {
	var agility int32
	kimahriAgility := kimahriStats["agility"]

	if mon.Name == "biran ronso" {
		agility = max(kimahriAgility - 4, 1)
	}

	if mon.Name == "yenke ronso" {
		agility = max(kimahriAgility - 6, 1)
	}

	return agility
}