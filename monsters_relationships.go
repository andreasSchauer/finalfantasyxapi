package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterAbility struct {
	Ability  NamedAPIResource `json:"ability"`
	IsForced bool             `json:"is_forced"`
	IsUnused bool             `json:"is_unused"`
}

func (ma MonsterAbility) GetAPIResource() APIResource {
	return ma.Ability
}

type BribeChance struct {
	Gil    int32 `json:"gil"`
	Chance int32 `json:"chance"`
}

type AgilityParams struct {
	TickSpeed int32  `json:"tick_speed"`
	MinICV    *int32 `json:"min_icv"`
	MaxICV    *int32 `json:"max_icv"`
}

func (cfg *Config) getMonsterRelationships(r *http.Request, mon seeding.Monster) (Monster, error) {
	areas, err := getResourcesDB(cfg, r, cfg.e.areas, mon, cfg.db.GetMonsterAreaIDs)
	if err != nil {
		return Monster{}, err
	}

	formations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, mon, cfg.db.GetMonsterMonsterFormationIDs)
	if err != nil {
		return Monster{}, err
	}

	abilities, err := cfg.getMonsterAbilities(mon)
	if err != nil {
		return Monster{}, err
	}

	monster := Monster{
		Properties:       namesToNamedAPIResources(cfg, cfg.e.properties, mon.Properties),
		AutoAbilities:    namesToNamedAPIResources(cfg, cfg.e.autoAbilities, mon.AutoAbilities),
		RonsoRages:       namesToNamedAPIResources(cfg, cfg.e.ronsoRages, mon.RonsoRages),
		Locations:        areas,
		Formations:       formations,
		BaseStats:        namesToResourceAmounts(cfg, cfg.e.stats, mon.BaseStats, cfg.newBaseStat),
		ElemResists:      cfg.getMonsterElemResists(mon.ElemResists),
		StatusImmunities: namesToNamedAPIResources(cfg, cfg.e.statusConditions, mon.StatusImmunities),
		StatusResists:    namesToResourceAmounts(cfg, cfg.e.statusConditions, mon.StatusResists, cfg.newStatusResist),
		Abilities:        abilities,
		AlteredStates:    cfg.getMonsterAlteredStates(r, mon),
	}

	return monster, nil
}


func (cfg *Config) getMonsterElemResists(resists []seeding.ElementalResist) []ElementalResist {
	elemResists := cfg.namesToElemResists(resists)
	elemResistMap := getResourceMap(elemResists)

	for key := range cfg.l.ElementsID {
		_, ok := elemResistMap[key]
		if !ok {
			element := cfg.l.ElementsID[key]
			elemResistMap[element.ID] = cfg.newElemResist(element.Name, "neutral")
		}
	}

	return resourceMapToSlice(elemResistMap)
}


func (cfg *Config) getMonsterAbilities(mon seeding.Monster) ([]MonsterAbility, error) {
	monAbilities := []MonsterAbility{}

	for _, seedAbility := range mon.Abilities {
		abilityResource, err := cfg.createAbilityResource(seedAbility.Name, seedAbility.Version, database.AbilityType(seedAbility.AbilityType))
		if err != nil {
			return nil, err
		}

		monAbility := MonsterAbility{
			Ability:  abilityResource,
			IsForced: seedAbility.IsForced,
			IsUnused: seedAbility.IsUnused,
		}

		monAbilities = append(monAbilities, monAbility)
	}

	return monAbilities, nil
}


func (cfg *Config) getMonsterPoisonDamage(mon Monster) (*int32, error) {
	if mon.PoisonRate == nil {
		return nil, nil
	}

	hpStat := cfg.getBaseStat("hp", mon.BaseStats)

	poisonDamageFloat := float32(hpStat.Value) * *mon.PoisonRate
	poisonDamage := int32(poisonDamageFloat)

	return &poisonDamage, nil
}


func (cfg *Config) getMonsterAgilityParams(r *http.Request, mon Monster) (*AgilityParams, error) {
	agilityStat := cfg.getBaseStat("agility", mon.BaseStats)
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
func (cfg *Config) getMonsterBribeChances(mon Monster) ([]BribeChance, error) {
	bribe := nameToNamedAPIResource(cfg, cfg.e.statusConditions, "bribe", nil)
	if resourcesContain(mon.StatusImmunities, bribe) || mon.Items == nil || mon.Items.Bribe == nil {
		return nil, nil
	}

	hpStat := cfg.getBaseStat("hp", mon.BaseStats)
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
