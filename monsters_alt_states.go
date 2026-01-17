package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AlteredState struct {
	URL         string           `json:"url"`
	Condition   string           `json:"condition"`
	IsTemporary bool             `json:"is_temporary"`
	Changes     []AltStateChange `json:"changes"`
}

type AltStateChange struct {
	AlterationType   database.AlterationType `json:"alteration_type"`
	Distance         *int32                  `json:"distance,omitempty"`
	Properties       []NamedAPIResource      `json:"properties,omitempty"`
	AutoAbilities    []NamedAPIResource      `json:"auto_abilities,omitempty"`
	BaseStats        []BaseStat              `json:"base_stats,omitempty"`
	ElemResists      []ElementalResist       `json:"elem_resists,omitempty"`
	StatusImmunities []NamedAPIResource      `json:"status_immunities,omitempty"`
	StatusResists    []StatusResist          `json:"status_resists,omitempty"`
	AddedStatus      *InflictedStatus        `json:"added_status_condition,omitempty"`
	RemovedStatus    *NamedAPIResource       `json:"removed_status_condition,omitempty"`
}

func (c *AltStateChange) IsZero() bool {
	return c.Distance 		== nil &&
		c.Properties 		== nil &&
		c.AutoAbilities 	== nil &&
		c.BaseStats 		== nil &&
		c.ElemResists 		== nil &&
		c.StatusImmunities 	== nil &&
		c.StatusResists 	== nil &&
		c.AddedStatus 		== nil &&
		c.RemovedStatus 	== nil
}

func (cfg *Config) getMonsterAlteredStates(r *http.Request, mon database.Monster) ([]AlteredState, error) {
	dbAltStates, err := cfg.db.GetMonsterAlteredStates(r.Context(), mon.ID)
	if err != nil {
		return []AlteredState{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get altered states of %s.", getMonsterName(mon)), err)
	}

	alteredStates := []AlteredState{}

	for i, dbAltState := range dbAltStates {
		altStateChanges, err := cfg.getAltStateChanges(r, mon, dbAltState)
		if err != nil {
			return []AlteredState{}, err
		}
		q := r.URL.Query()
		q.Set("altered-state", strconv.Itoa(i+1))

		alteredState := AlteredState{
			URL:         cfg.createResourceURLQuery(cfg.e.monsters.endpoint, mon.ID, q),
			Condition:   dbAltState.Condition,
			IsTemporary: dbAltState.IsTemporary,
			Changes:     altStateChanges,
		}

		alteredStates = append(alteredStates, alteredState)
	}

	return alteredStates, nil
}

func (cfg *Config) getAltStateChanges(r *http.Request, mon database.Monster, as database.AlteredState) ([]AltStateChange, error) {
	altStateChanges := []AltStateChange{}

	dbAltStateChanges, err := cfg.db.GetAltStateChanges(r.Context(), as.ID)
	if err != nil {
		return []AltStateChange{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state relationships of %s.", getMonsterName(mon)), err)
	}

	for _, dbChange := range dbAltStateChanges {
		altStateChange, err := cfg.getAltStateChangeRelationships(r, mon, dbChange)
		if err != nil {
			return []AltStateChange{}, err
		}

		altStateChange.AlterationType = dbChange.AlterationType
		altStateChange.Distance = anyToInt32Ptr(dbChange.Distance)

		altStateChanges = append(altStateChanges, altStateChange)
	}

	return altStateChanges, nil
}

func (cfg *Config) getAltStateChangeRelationships(r *http.Request, mon database.Monster, asc database.AltStateChange) (AltStateChange, error) {
	properties, err := cfg.getAltStateProperties(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	autoAbilities, err := cfg.getAltStateAutoAbilities(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	baseStats, err := cfg.getAltStateBaseStats(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	elemResists, err := cfg.getAltStateElemResists(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	immunities, err := cfg.getAltStateImmunities(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	addedStatusses, err := cfg.getAltStateStatus(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	return AltStateChange{
		Properties:       properties,
		AutoAbilities:    autoAbilities,
		BaseStats:        baseStats,
		ElemResists:      elemResists,
		StatusImmunities: immunities,
		AddedStatus:      addedStatusses,
	}, nil
}

// => GetResourcesDbOrNil
func (cfg *Config) getAltStateProperties(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]NamedAPIResource, error) {
	dbProperties, err := cfg.db.GetAltStateProperties(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state properties of %s.", getMonsterName(mon)), err)
	}

	properties := createNamedAPIResourcesSimple(cfg, dbProperties, cfg.e.properties.endpoint, func(prop database.GetAltStatePropertiesRow) (int32, string) {
		return prop.PropertyID, prop.Property
	})

	if len(properties) == 0 {
		return nil, nil
	}

	return properties, nil
}

func (cfg *Config) getAltStateAutoAbilities(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]NamedAPIResource, error) {
	dbAutoAbilities, err := cfg.db.GetAltStateAutoAbilities(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state auto-abilities of %s.", getMonsterName(mon)), err)
	}

	autoAbilities := createNamedAPIResourcesSimple(cfg, dbAutoAbilities, cfg.e.autoAbilities.endpoint, func(autoAbility database.GetAltStateAutoAbilitiesRow) (int32, string) {
		return autoAbility.AutoAbilityID, autoAbility.AutoAbility
	})

	if len(autoAbilities) == 0 {
		return nil, nil
	}

	return autoAbilities, nil
}

func (cfg *Config) getAltStateBaseStats(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]BaseStat, error) {
	dbBaseStats, err := cfg.db.GetAltStateBaseStats(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state base stats of %s.", getMonsterName(mon)), err)
	}

	var baseStats []BaseStat

	for _, dbStat := range dbBaseStats {
		baseStat := cfg.newBaseStat(dbStat.StatID, dbStat.Value, dbStat.Stat)

		baseStats = append(baseStats, baseStat)
	}

	return baseStats, nil
}

func (cfg *Config) getAltStateElemResists(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]ElementalResist, error) {
	dbElemResists, err := cfg.db.GetAltStateElemResists(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state elemental resists of %s.", getMonsterName(mon)), err)
	}

	var elemResists []ElementalResist

	for _, dbResist := range dbElemResists {
		elemResist := cfg.newElemResist(dbResist.ElementID, dbResist.AffinityID, dbResist.Element, dbResist.Affinity)

		elemResists = append(elemResists, elemResist)
	}

	return elemResists, nil
}

func (cfg *Config) getAltStateImmunities(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]NamedAPIResource, error) {
	dbImmunities, err := cfg.db.GetAltStateImmunities(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state status immunities of %s.", getMonsterName(mon)), err)
	}

	immunities := createNamedAPIResourcesSimple(cfg, dbImmunities, cfg.e.statusConditions.endpoint, func(immunity database.GetAltStateImmunitiesRow) (int32, string) {
		return immunity.StatusID, immunity.Status
	})

	if len(immunities) == 0 {
		return nil, nil
	}

	return immunities, nil
}

func (cfg *Config) getAltStateStatus(r *http.Request, mon database.Monster, asc database.AltStateChange) (*InflictedStatus, error) {
	dbStatus, err := cfg.db.GetAltStateStatus(r.Context(), asc.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state added status of %s.", getMonsterName(mon)), err)
	}

	addedStatus := cfg.newInflictedStatus(dbStatus.StatusID, anyToInt32(dbStatus.Probability), dbStatus.Status, h.NullInt32ToPtr(dbStatus.Amount), dbStatus.DurationType)

	if addedStatus.IsZero() {
		return nil, nil
	}

	return &addedStatus, nil
}
