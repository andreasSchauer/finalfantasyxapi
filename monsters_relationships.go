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

func (cfg *Config) getMonsterRelationships(r *http.Request, dbMon database.Monster) (Monster, error) {
	properties, err := cfg.getMonsterProperties(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	autoAbilities, err := cfg.getMonsterAutoAbilities(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	ronsoRages, err := cfg.getMonsterRonsoRages(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	locations, err := cfg.getMonsterLocations(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	formations, err := cfg.getMonsterMonsterFormations(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	baseStats, err := cfg.getMonsterBaseStats(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	elemResists, err := cfg.getMonsterElemResists(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	statusResists, err := cfg.getMonsterStatusResists(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	immunities, err := cfg.getMonsterImmunities(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	alteredStates, err := cfg.getMonsterAlteredStates(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	abilities, err := cfg.getMonsterAbilities(r, dbMon)
	if err != nil {
		return Monster{}, err
	}

	monster := Monster{
		Properties:       properties,
		AutoAbilities:    autoAbilities,
		RonsoRages:       ronsoRages,
		Locations:        locations,
		Formations:       formations,
		BaseStats:        baseStats,
		ElemResists:      elemResists,
		StatusImmunities: immunities,
		StatusResists:    statusResists,
		AlteredStates:    alteredStates,
		Abilities:        abilities,
	}

	return monster, nil
}

func (ma MonsterAbility) GetAPIResource() IsAPIResource {
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

func (cfg *Config) getMonsterProperties(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbProperties, err := cfg.db.GetMonsterProperties(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get properties of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	properties := createNamedAPIResourcesSimple(cfg, dbProperties, cfg.e.properties.endpoint, func(prop database.GetMonsterPropertiesRow) (int32, string) {
		return prop.PropertyID, prop.Property
	})

	return properties, nil
}

func (cfg *Config) getMonsterAutoAbilities(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbAutoAbilities, err := cfg.db.GetMonsterAutoAbilities(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get auto abilities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	autoAbilities := createNamedAPIResourcesSimple(cfg, dbAutoAbilities, cfg.e.autoAbilities.endpoint, func(autoAbility database.GetMonsterAutoAbilitiesRow) (int32, string) {
		return autoAbility.AutoAbilityID, autoAbility.AutoAbility
	})

	return autoAbilities, nil
}

func (cfg *Config) getMonsterRonsoRages(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	const ronsoRageOffset int32 = 35
	dbRages, err := cfg.db.GetMonsterRonsoRages(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get ronso rages of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	rages := createNamedAPIResourcesSimple(cfg, dbRages, cfg.e.ronsoRages.endpoint, func(rage database.GetMonsterRonsoRagesRow) (int32, string) {
		return rage.RonsoRageID - ronsoRageOffset, rage.RonsoRage
	})

	return rages, nil
}

func (cfg *Config) getMonsterLocations(r *http.Request, mon database.Monster) ([]LocationAPIResource, error) {
	dbLocations, err := cfg.db.GetMonsterLocations(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get locations of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	locations := createLocationBasedAPIResources(cfg, dbLocations, func(loc database.GetMonsterLocationsRow) (string, string, string, *int32) {
		return loc.Location, loc.Sublocation, loc.Area, h.NullInt32ToPtr(loc.Version)
	})

	return locations, nil
}

func (cfg *Config) getMonsterMonsterFormations(r *http.Request, mon database.Monster) ([]UnnamedAPIResource, error) {
	dbFormations, err := cfg.db.GetMonsterMonsterFormations(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get formations of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	formations := createUnnamedAPIResources(cfg, dbFormations, cfg.e.monsterFormations.endpoint, func(formation database.GetMonsterMonsterFormationsRow) int32 {
		return formation.ID
	})

	return formations, nil
}

func (cfg *Config) getMonsterBaseStats(r *http.Request, mon database.Monster) ([]BaseStat, error) {
	dbBaseStats, err := cfg.db.GetMonsterBaseStats(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get base stats of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	baseStats := []BaseStat{}

	for _, dbStat := range dbBaseStats {
		baseStat := cfg.newBaseStat(dbStat.StatID, dbStat.Value, dbStat.Stat)

		baseStats = append(baseStats, baseStat)
	}

	return baseStats, nil
}

func (cfg *Config) getMonsterElemResists(r *http.Request, mon database.Monster) ([]ElementalResist, error) {
	dbElemResists, err := cfg.db.GetMonsterElemResists(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get elemental resists of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	elemResists := []ElementalResist{}

	for _, dbResist := range dbElemResists {
		elemResist := cfg.newElemResist(dbResist.ElementID, dbResist.AffinityID, dbResist.Element, dbResist.Affinity)

		elemResists = append(elemResists, elemResist)
	}

	return elemResists, nil
}

func (cfg *Config) changeElemResist(element NamedAPIResource, newAffinityName string) (ElementalResist, error) {
	newAffinity, err := seeding.GetResource(newAffinityName, cfg.l.Affinities)
	if err != nil {
		return ElementalResist{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get affinity %s", newAffinityName), err)
	}

	newResist := cfg.newElemResist(
		element.ID,
		newAffinity.ID,
		element.Name,
		newAffinity.Name,
	)

	return newResist, nil
}

func (cfg *Config) getMonsterStatusResists(r *http.Request, mon database.Monster) ([]StatusResist, error) {
	dbStatusResists, err := cfg.db.GetMonsterStatusResists(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get status resists of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	statusResists := []StatusResist{}

	for _, dbResist := range dbStatusResists {
		statusResist := cfg.newStatusResist(dbResist.StatusID, anyToInt32(dbResist.Resistance), dbResist.Status)

		statusResists = append(statusResists, statusResist)
	}

	return statusResists, nil
}

func (cfg *Config) getMonsterImmunities(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbImmunities, err := cfg.db.GetMonsterImmunities(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get status immunities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	immunities := createNamedAPIResourcesSimple(cfg, dbImmunities, cfg.e.statusConditions.endpoint, func(immunity database.GetMonsterImmunitiesRow) (int32, string) {
		return immunity.StatusID, immunity.Status
	})

	return immunities, nil
}

func (cfg *Config) getMonsterAbilities(r *http.Request, mon database.Monster) ([]MonsterAbility, error) {
	dbMonAbilities, err := cfg.db.GetMonsterAbilities(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get abilities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	monAbilities := []MonsterAbility{}

	for _, dbAbility := range dbMonAbilities {
		ref := seeding.AbilityReference{
			Name:        dbAbility.Ability,
			Version:     h.NullInt32ToPtr(dbAbility.Version),
			AbilityType: string(dbAbility.AbilityType),
		}

		abilityResource, err := cfg.getAbilityResource(ref, h.NullStringToPtr(dbAbility.Specification))
		if err != nil {
			return nil, err
		}

		monAbility := MonsterAbility{
			Ability:  abilityResource,
			IsForced: dbAbility.IsForced,
			IsUnused: dbAbility.IsUnused,
		}

		monAbilities = append(monAbilities, monAbility)
	}

	return monAbilities, nil
}


// can be used for various other functions related to abilities
func (cfg *Config) getAbilityResource(ref seeding.AbilityReference, specification *string) (NamedAPIResource, error) {
	i, err := cfg.getAbilityInput(database.AbilityType(ref.AbilityType))
	if err != nil {
		return NamedAPIResource{}, err
	}

	abilityLookup, err := seeding.GetResource(ref, i.ObjLookup())
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	abilityResource := cfg.newNamedAPIResource(i.Endpoint(), abilityLookup.GetID(), ref.Name, ref.Version, specification)

	return abilityResource, nil
}


// can be used for various other functions related to abilities
func (cfg *Config) getAbilityInput(abilityType database.AbilityType) (handlerView, error) {
	var i handlerView
	var err error

	switch abilityType {
	case database.AbilityTypePlayerAbility:
		i = cfg.e.playerAbilities

	case database.AbilityTypeEnemyAbility:
		i = cfg.e.enemyAbilities

	case database.AbilityTypeOverdriveAbility:
		i = cfg.e.overdriveAbilities

	case database.AbilityTypeItemAbility:
		i = cfg.e.itemAbilities

	case database.AbilityTypeTriggerCommand:
		i = cfg.e.triggerCommands

	default:
		err = newHTTPError(http.StatusInternalServerError, fmt.Sprintf("ability of type %s does not exist", abilityType), nil)
	}

	if err != nil {
		return nil, err
	}

	return i, nil
}


func (cfg *Config) getMonsterStat(mon Monster, stat string) (int32, error) {
	statLookup, err := seeding.GetResource(stat, cfg.l.Stats)
	if err != nil {
		return 0, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	statMap := getResourceMap(mon.BaseStats)
	key := cfg.createResourceURL(cfg.e.stats.endpoint, statLookup.ID)

	return statMap[key].Value, nil
}

func (cfg *Config) getMonsterPoisonDamage(mon Monster) (*int32, error) {
	if mon.PoisonRate == nil {
		return nil, nil
	}

	hp, err := cfg.getMonsterStat(mon, "hp")
	if err != nil {
		return nil, err
	}

	poisonDamageFloat := float32(hp) * *mon.PoisonRate
	poisonDamage := int32(poisonDamageFloat)

	return &poisonDamage, nil
}

func (cfg *Config) getMonsterAgilityVals(r *http.Request, mon Monster) (*AgilityParams, error) {
	agility, err := cfg.getMonsterStat(mon, "agility")
	if err != nil {
		return nil, err
	}

	if agility == 0 {
		return nil, nil
	}

	dbAgilityTier, err := cfg.db.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't extract agility parameters from Monster %s, Version %d", mon.Name, h.DerefOrNil(mon.Version)), err)
	}

	agilityParams := AgilityParams{
		TickSpeed: dbAgilityTier.TickSpeed,
		MinICV:    h.NullInt32ToPtr(dbAgilityTier.MonsterMinIcv),
		MaxICV:    h.NullInt32ToPtr(dbAgilityTier.MonsterMaxIcv),
	}

	return &agilityParams, nil
}

// HP x10 = 25%, HP x15 = 50%, HP x20 = 75%, HP x25 = 100%
func (cfg *Config) getMonsterBribeChances(mon Monster) ([]BribeChance, error) {
	bribeLookup, err := seeding.GetResource("bribe", cfg.l.StatusConditions)
	if err != nil {
		return nil, err
	}
	bribe := cfg.newNamedAPIResourceSimple("status-conditions", bribeLookup.ID, bribeLookup.Name)
	if resourcesContain(mon.StatusImmunities, bribe) || mon.Items == nil || mon.Items.Bribe == nil {
		return nil, nil
	}

	hp, err := cfg.getMonsterStat(mon, "hp")
	if err != nil {
		return nil, err
	}

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
