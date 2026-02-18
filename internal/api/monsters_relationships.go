package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



func getMonsterRelationships(cfg *Config, r *http.Request, mon seeding.Monster) (Monster, error) {
	areas, err := getResourcesDB(cfg, r, cfg.e.areas, mon, cfg.db.GetMonsterAreaIDs)
	if err != nil {
		return Monster{}, err
	}

	formations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, mon, cfg.db.GetMonsterMonsterFormationIDs)
	if err != nil {
		return Monster{}, err
	}

	monster := Monster{
		Areas:            areas,
		Formations:       formations,
	}

	return monster, nil
}


func completeMonsterResponse(cfg *Config, r *http.Request, mon Monster) (Monster, error) {
	mon, err := applyAlteredState(cfg, r, mon, "altered_state")
	if err != nil {
		return Monster{}, err
	}

	mon.BaseStats, err = applyAeonStatsMonsters(cfg, r, mon, "aeon_stats")
	if err != nil {
		return Monster{}, err
	}

	mon.BaseStats, err = applyRonsoStats(cfg, r, mon, "kimahri_stats")
	if err != nil {
		return Monster{}, err
	}

	mon.ElemResists, err = applyOmnisElements(cfg, r, mon, "omnis_elements")
	if err != nil {
		return Monster{}, err
	}

	mon.BribeChances, err = getMonsterBribeChances(cfg, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.PoisonDamage, err = getMonsterPoisonDamage(cfg, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.AgilityParameters, err = getMonsterAgilityParams(cfg, r, mon)
	if err != nil {
		return Monster{}, err
	}

	return mon, nil
}


func getMonsterElemResists(cfg *Config, resists []seeding.ElementalResist) []ElementalResist {
	elemResists := namesToElemResists(cfg, resists)
	elemResistMap := getResourceMap(elemResists)

	for key := range cfg.l.ElementsID {
		_, ok := elemResistMap[key]
		if !ok {
			element := cfg.l.ElementsID[key]
			elemResistMap[element.ID] = newElemResist(cfg, element.Name, "neutral")
		}
	}

	return resourceMapToSlice(elemResistMap)
}


func getMonsterPoisonDamage(cfg *Config, mon Monster) (*int32, error) {
	if mon.PoisonRate == nil {
		return nil, nil
	}

	hpStat := getBaseStat(cfg, "hp", mon.BaseStats)

	poisonDamageFloat := float32(hpStat.Value) * *mon.PoisonRate
	poisonDamage := int32(poisonDamageFloat)

	return &poisonDamage, nil
}


func getMonsterAgilityParams(cfg *Config, r *http.Request, mon Monster) (*AgilityParams, error) {
	agilityStat := getBaseStat(cfg, "agility", mon.BaseStats)
	agility := agilityStat.Value
	if agility == 0 {
		return nil, nil
	}

	dbAgilityTier, err := cfg.db.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't extract agility parameters from %s.", mon.Error()), err)
	}

	agilityParams := AgilityParams{
		TickSpeed: dbAgilityTier.TickSpeed,
		MinICV:    h.NullInt32ToPtr(dbAgilityTier.MonsterMinIcv),
		MaxICV:    h.NullInt32ToPtr(dbAgilityTier.MonsterMaxIcv),
	}

	fs := nameToNamedAPIResource(cfg, cfg.e.autoAbilities, "first strike", nil)
	if resourcesContain(mon.AutoAbilities, fs) {
		var fsICV int32 = -1
		agilityParams.MinICV = &fsICV
		agilityParams.MaxICV = &fsICV
	}

	return &agilityParams, nil
}


// HP x10 = 25%, HP x15 = 50%, HP x20 = 75%, HP x25 = 100%
func getMonsterBribeChances(cfg *Config, mon Monster) ([]BribeChance, error) {
	bribe := nameToNamedAPIResource(cfg, cfg.e.statusConditions, "bribe", nil)
	if resourcesContain(mon.StatusImmunities, bribe) || mon.Items == nil || mon.Items.Bribe == nil {
		return nil, nil
	}

	hpStat := getBaseStat(cfg, "hp", mon.BaseStats)
	hp := hpStat.Value

	bribeChances := []BribeChance{}
	var multiplier int32 = 10
	var chance int32 = 25

	for multiplier <= 25 {
		bribeChance := BribeChance{
			Gil:    hp * multiplier,
			Chance: chance,
		}
		bribeChances = append(bribeChances, bribeChance)
		multiplier += 5
		chance += 25
	}

	return bribeChances, nil
}