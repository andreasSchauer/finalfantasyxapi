package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAeonRelationships(cfg *Config, r *http.Request, ae seeding.Aeon) (Aeon, error) {
	celestialWeapon, err := getResPtrDB(cfg, r, cfg.e.celestialWeapons, ae, cfg.db.GetAeonCelestialWeaponID)
	if err != nil {
		return Aeon{}, err
	}

	characterClasses, err := getResourcesDB(cfg, r, cfg.e.characterClasses, ae, cfg.db.GetAeonCharClassIDs)
	if err != nil {
		return Aeon{}, err
	}

	aeonCommands, err := getResourcesDB(cfg, r, cfg.e.aeonCommands, ae, cfg.db.GetAeonAeonCommandIDs)
	if err != nil {
		return Aeon{}, err
	}

	overdrives, err := getResourcesDB(cfg, r, cfg.e.overdrives, ae, cfg.db.GetAeonOverdriveIDs)
	if err != nil {
		return Aeon{}, err
	}

	defaultAbilities, err := getResourcesDB(cfg, r, cfg.e.playerAbilities, ae, cfg.db.GetAeonDefaultAbilityIDs)
	if err != nil {
		return Aeon{}, err
	}

	aeon := Aeon{
		CelestialWeapon:  celestialWeapon,
		CharacterClasses: characterClasses,
		AeonCommands:     aeonCommands,
		Overdrives:       overdrives,
		DefaultAbilities: defaultAbilities,
	}

	return aeon, nil
}


func applyAeonStats(cfg *Config, r *http.Request, aeon Aeon) (Aeon, error) {
	var err error

	aeon.BaseStats, err = applyAeonStatsBattles(cfg, r, aeon, "battles")
	if err != nil {
		return Aeon{}, err
	}

	aeon.BaseStats, err = applyYunaStats(cfg, r, aeon, "yuna_stats")
	if err != nil {
		return Aeon{}, err
	}

	return aeon, nil
}


func applyAeonStatsBattles(cfg *Config, r *http.Request, aeon Aeon, queryName string) ([]BaseStat, error) {
	queryParam := cfg.q.aeons[queryName]
	battles, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return aeon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	i := battles / 30 - 1

	if battles < 60 {
		i = 0
	}

	seedAeon, _ := seeding.GetResourceByID(aeon.ID, cfg.l.AeonsID)
	newBaseStats := seedAeon.BaseStats.XVals[i].BaseStats
	baseStats := namesToResourceAmounts(cfg, cfg.e.stats, newBaseStats, newBaseStat)

	return baseStats, nil
}


func applyYunaStats(cfg *Config, r *http.Request, aeon Aeon, queryName string) ([]BaseStat, error) {
	allowedStatIDs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	baseStats := aeon.BaseStats
	queryParam := cfg.q.aeons[queryName]
	
	yuna, _ := seeding.GetResource("yuna", cfg.l.Characters)
	yunaBS := namesToResourceAmounts(cfg, cfg.e.stats, yuna.BaseStats, newBaseStat)

	yunaStatMapInt, err := parseStatQuery(cfg, r, queryParam, yunaBS, allowedStatIDs)
	if errors.Is(err, errEmptyQuery) {
		return baseStats, nil
	}
	if err != nil {
		return nil, err
	}

	yunaStatMapInt["hp"] = min(yunaStatMapInt["hp"], 9999)
	yunaStatMapInt["mp"] = min(yunaStatMapInt["mp"], 999)

	aeonStats := calcAeonStats(cfg, aeon, yunaStatMapInt)

	return aeonStats, nil
}


func calcAeonStats(cfg *Config, aeon Aeon, yuna map[string]int32) []BaseStat {
	baseStats := aeon.BaseStats
	aVals, bVals := getAeonStatTables(cfg, aeon)
	yunaFloat := getResAmountFloatMap(yuna)
	yParameter := yParamCalc(yuna)

	for i, baseStat := range baseStats {
		yStat := yStatCalc(baseStat.GetName(), yParameter, aVals, bVals, yunaFloat)
		baseStat.Value = max(baseStat.Value, yStat)
		baseStats[i] = baseStat
	}

	return baseStats
}


func getAeonStatTables(cfg *Config, aeon Aeon) (map[string]float64, map[string]float64) {
	aeonLookup, _ := seeding.GetResourceByID(aeon.ID, cfg.l.AeonsID)
	
	aValsSlice := namesToResourceAmounts(cfg, cfg.e.stats, aeonLookup.BaseStats.AVals, newBaseStat)
	aValsInt := getResourceAmountMap(aValsSlice)
	aVals := getResAmountFloatMap(aValsInt)
	
	bValsSlice := namesToResourceAmounts(cfg, cfg.e.stats, aeonLookup.BaseStats.BVals, newBaseStat)
	bValsInt := getResourceAmountMap(bValsSlice)
	bVals := getResAmountFloatMap(bValsInt)

	return aVals, bVals
}


func getResAmountFloatMap(intMap map[string]int32) map[string]float64 {
	floatMap := make(map[string]float64)

	for key := range intMap {
		float := float64(intMap[key])
		floatMap[key] = float
	}
	
	return floatMap
}


func yParamCalc(yuna map[string]int32) float64 {
	return float64(yuna["hp"] / 100 + yuna["mp"] / 10 + yuna["strength"] + yuna["defense"] + yuna["magic"] + yuna["magic defense"] + yuna["agility"] + yuna["evasion"] + yuna["accuracy"])
}


func yStatCalc(stat string, yParameter float64, aVals, bVals, yuna map[string]float64) int32 {
	var yFloat float64
	var yStat int32

	switch stat {
	case "hp":
		yFloat = aVals[stat] * yParameter + bVals[stat] * (yuna[stat] / 100)
		yStat = int32(yFloat)
		return min(yStat, 99999)

	case "mp":
		yFloat = aVals[stat] * (yParameter / 10) + bVals[stat] * (yuna[stat] / 100)
		yStat = int32(yFloat)
		return min(yStat, 9999)

	case "luck":
		return int32(yuna[stat])

	default:
		yFloat := yParameter / aVals[stat] + bVals[stat] * (yuna[stat] / 10)
		yStat = int32(yFloat)
		return min(yStat, 255)
	}
}