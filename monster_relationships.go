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


func (cfg *apiConfig) getMonsterRelationships(r *http.Request, mon database.Monster) (Monster, error) {
	properties, err := cfg.getMonsterProperties(r, mon)
	if err != nil {
		return Monster{}, err
	}

	autoAbilities, err := cfg.getMonsterAutoAbilities(r, mon)
	if err != nil {
		return Monster{}, err
	}

	ronsoRages, err := cfg.getMonsterRonsoRages(r, mon)
	if err != nil {
		return Monster{}, err
	}

	locations, err := cfg.getMonsterLocations(r, mon)
	if err != nil {
		return Monster{}, err
	}

	baseStats, err := cfg.getMonsterBaseStats(r, mon)
	if err != nil {
		return Monster{}, err
	}

	elemResists, err := cfg.getMonsterElemResists(r, mon)
	if err != nil {
		return Monster{}, err
	}

	statusResists, err := cfg.getMonsterStatusResists(r, mon)
	if err != nil {
		return Monster{}, err
	}

	immunities, err := cfg.getMonsterImmunities(r, mon)
	if err != nil {
		return Monster{}, err
	}

	alteredStates, err := cfg.getMonsterAlteredStates(r, mon)
	if err != nil {
		return Monster{}, err
	}

	abilities, err := cfg.getMonsterAbilities(r, mon)
	if err != nil {
		return Monster{}, err
	}

	return Monster{
		Properties:       properties,
		AutoAbilities:    autoAbilities,
		RonsoRages:       ronsoRages,
		Locations: 		  locations,
		BaseStats:        baseStats,
		ElemResists:      elemResists,
		StatusImmunities: immunities,
		StatusResists:    statusResists,
		AlteredStates:    alteredStates,
		Abilities:        abilities,
	}, nil
}


type BribeChance struct {
	Gil			int32	`json:"gil"`
	Chance		int32	`json:"chance"`
}

type AgilityParams struct {
	TickSpeed	int32	`json:"tick_speed"`
	MinICV		*int32	`json:"min_icv"`
	MaxICV		*int32	`json:"max_icv"`
}



func (cfg *apiConfig) getMonsterProperties(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbProperties, err := cfg.db.GetMonsterProperties(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get properties of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	properties := createNamedAPIResourcesSimple(cfg, dbProperties, "properties", func(prop database.GetMonsterPropertiesRow) (int32, string) {
		return h.NullInt32ToVal(prop.PropertyID), h.NullStringToVal(prop.Property)
	})

	return properties, nil
}

func (cfg *apiConfig) getMonsterAutoAbilities(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbAutoAbilities, err := cfg.db.GetMonsterAutoAbilities(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get auto abilities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	autoAbilities := createNamedAPIResourcesSimple(cfg, dbAutoAbilities, "auto-abilities", func(autoAbility database.GetMonsterAutoAbilitiesRow) (int32, string) {
		return h.NullInt32ToVal(autoAbility.AutoAbilityID), h.NullStringToVal(autoAbility.AutoAbility)
	})

	return autoAbilities, nil
}

func (cfg *apiConfig) getMonsterRonsoRages(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbRages, err := cfg.db.GetMonsterRonsoRages(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get ronso rages of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	rages := createNamedAPIResourcesSimple(cfg, dbRages, "overdrives", func(rage database.GetMonsterRonsoRagesRow) (int32, string) {
		return h.NullInt32ToVal(rage.RonsoRageID), h.NullStringToVal(rage.RonsoRage)
	})

	return rages, nil
}


func (cfg *apiConfig) getMonsterLocations(r *http.Request, mon database.Monster) ([]LocationAPIResource, error) {
	dbLocations, err := cfg.db.GetMonsterLocations(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get locations of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	locations := createLocationBasedAPIResources(cfg, dbLocations, func(loc database.GetMonsterLocationsRow)(string, string, string, *int32) {
		return loc.Location, h.NullStringToVal(loc.Sublocation), h.NullStringToVal(loc.Area), h.NullInt32ToPtr(loc.Version)
	})

	return locations, nil
}


func (cfg *apiConfig) getMonsterBaseStats(r *http.Request, mon database.Monster) ([]BaseStat, error) {
	dbBaseStats, err := cfg.db.GetMonsterBaseStats(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get base stats of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	baseStats := []BaseStat{}

	for _, dbStat := range dbBaseStats {
		baseStat := cfg.newBaseStat(dbStat.StatID.Int32, dbStat.Value.Int32, dbStat.Stat.String)

		baseStats = append(baseStats, baseStat)
	}

	return baseStats, nil
}

func (cfg *apiConfig) getMonsterElemResists(r *http.Request, mon database.Monster) ([]ElementalResist, error) {
	dbElemResists, err := cfg.db.GetMonsterElemResists(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get elemental resists of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	elemResists := []ElementalResist{}

	for _, dbResist := range dbElemResists {
		elemResist := cfg.newElemResist(dbResist.ElementID.Int32, dbResist.AffinityID.Int32, dbResist.Element.String, dbResist.Affinity.String)

		elemResists = append(elemResists, elemResist)
	}

	return elemResists, nil
}

func (cfg *apiConfig) getMonsterStatusResists(r *http.Request, mon database.Monster) ([]StatusResist, error) {
	dbStatusResists, err := cfg.db.GetMonsterStatusResists(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get status resists of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	statusResists := []StatusResist{}

	for _, dbResist := range dbStatusResists {
		statusResist := cfg.newStatusResist(dbResist.StatusID.Int32, anyToInt32(dbResist.Resistance), dbResist.Status.String)

		statusResists = append(statusResists, statusResist)
	}

	return statusResists, nil
}

func (cfg *apiConfig) getMonsterImmunities(r *http.Request, mon database.Monster) ([]NamedAPIResource, error) {
	dbImmunities, err := cfg.db.GetMonsterImmunities(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get status immunities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	immunities := createNamedAPIResourcesSimple(cfg, dbImmunities, "status-conditions", func(immunity database.GetMonsterImmunitiesRow) (int32, string) {
		return h.NullInt32ToVal(immunity.StatusID), h.NullStringToVal(immunity.Status)
	})

	return immunities, nil
}

func (cfg *apiConfig) getMonsterAbilities(r *http.Request, mon database.Monster) ([]MonsterAbility, error) {
	dbMonAbilities, err := cfg.db.GetMonsterAbilities(r.Context(), mon.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get abilities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	monAbilities := []MonsterAbility{}

	for _, dbAbility := range dbMonAbilities {
		id, endpoint, err := cfg.getAbilityID(dbAbility)
		if err != nil {
			return nil, err
		}

		abilityResource := cfg.newNamedAPIResource(endpoint, id, dbAbility.Ability.String, h.NullInt32ToPtr(dbAbility.Version), h.NullStringToPtr(dbAbility.Specification))

		monAbility := MonsterAbility{
			Ability:  abilityResource,
			IsForced: dbAbility.IsForced.Bool,
			IsUnused: dbAbility.IsUnused.Bool,
		}

		monAbilities = append(monAbilities, monAbility)
	}

	return monAbilities, nil
}

func (cfg *apiConfig) getAbilityID(ability database.GetMonsterAbilitiesRow) (int32, string, error) {
	ref := seeding.AbilityReference{
		Name:        ability.Ability.String,
		Version:     h.NullInt32ToPtr(ability.Version),
		AbilityType: string(ability.AbilityType.AbilityType),
	}

	switch ability.AbilityType.AbilityType {
	case database.AbilityTypePlayerAbility:
		abilityLookup, err := seeding.GetResource(ref, cfg.l.PlayerAbilities)
		if err != nil {
			return 0, "", newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}
		return abilityLookup.ID, "player-abilities", nil

	case database.AbilityTypeEnemyAbility:
		abilityLookup, err := seeding.GetResource(ref, cfg.l.EnemyAbilities)
		if err != nil {
			return 0, "", newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}
		return abilityLookup.ID, "enemy-abilities", nil

	case database.AbilityTypeOverdriveAbility:
		abilityLookup, err := seeding.GetResource(ref, cfg.l.OverdriveAbilities)
		if err != nil {
			return 0, "", newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}
		return abilityLookup.ID, "overdrive-abilities", nil

	case database.AbilityTypeItemAbility:
		abilityLookup, err := seeding.GetResource(ability.Ability.String, cfg.l.Items)
		if err != nil {
			return 0, "", newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}
		return abilityLookup.ID, "item-abilities", nil

	case database.AbilityTypeTriggerCommand:
		abilityLookup, err := seeding.GetResource(ref, cfg.l.TriggerCommands)
		if err != nil {
			return 0, "", newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}
		return abilityLookup.ID, "trigger-commands", nil

	default:
		return 0, "", newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find id for monster ability %s, version %d, type %s", ref.Name, h.DerefOrNil(ref.Version), ref.AbilityType), nil)
	}
}



func (cfg *apiConfig) getMonsterStat(mon Monster, stat string) (int32, error) {
	statLookup, err := seeding.GetResource(stat, cfg.l.Stats)
	if err != nil {
		return 0, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	statMap := getResourceMap(mon.BaseStats)
	key := cfg.createURL("stats", statLookup.ID)

	return statMap[key].Value, nil
}


func (cfg *apiConfig) getMonsterPoisonDamage(mon Monster) (*int32, error) {
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


func (cfg *apiConfig) getMonsterAgilityVals(r *http.Request, mon Monster) (AgilityParams, error) {
	agility, err := cfg.getMonsterStat(mon, "agility")
	if err != nil {
		return AgilityParams{}, err
	}

	dbAgilityTier, err := cfg.db.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return AgilityParams{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't extract agility parameters from Monster %s, Version %d", mon.Name, h.DerefOrNil(mon.Version)), err)
	}

	return AgilityParams{
		TickSpeed: dbAgilityTier.TickSpeed,
		MinICV: h.NullInt32ToPtr(dbAgilityTier.MonsterMinIcv),
		MaxICV: h.NullInt32ToPtr(dbAgilityTier.MonsterMaxIcv),
	}, nil
}



// HP x10 = 25%, HP x15 = 50%, HP x20 = 75%, HP x25 = 100%
func (cfg *apiConfig) getMonsterBribeChances(mon Monster) ([]BribeChance, error) {
	bribeLookup, err := seeding.GetResource("bribe", cfg.l.StatusConditions)
	if err != nil {
		return nil, err
	}
	bribe := cfg.newNamedAPIResourceSimple("status-conditions", bribeLookup.ID, bribeLookup.Name)

	if resourcesContain(mon.StatusImmunities, bribe) {
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
			Gil: hp * multiplier,
			Chance: chance,
		}
		bribeChances = append(bribeChances, bribeChance)
		multiplier += 5
		chance += 25
	}

	return bribeChances, nil
}