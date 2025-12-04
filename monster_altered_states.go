package main

import (
	"fmt"
	"net/http"

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
	AlterationType    database.AlterationType `json:"alteration_type"`
	Distance          *int32                  `json:"distance,omitempty"`
	Properties        []NamedAPIResource      `json:"properties,omitempty"`
	AutoAbilities     []NamedAPIResource      `json:"auto_abilities,omitempty"`
	BaseStats         []BaseStat              `json:"base_stats,omitempty"`
	ElemResists       []ElementalResist       `json:"elem_resists,omitempty"`
	StatusImmunities  []NamedAPIResource      `json:"status_immunities,omitempty"`
	StatusResistances []StatusResist          `json:"status_resistances,omitempty"`
	AddedStatusses    []InflictedStatus       `json:"added_status_conditions,omitempty"`
	RemovedStatus     *NamedAPIResource       `json:"removed_status_condition,omitempty"`
}



func (cfg *apiConfig) getMonsterAlteredStates(r *http.Request, mon database.Monster) ([]AlteredState, error) {
	dbAltStates, err := cfg.db.GetMonsterAlteredStates(r.Context(), mon.ID)
	if err != nil {
		return []AlteredState{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get altered states of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	alteredStates := []AlteredState{}

	for i, dbAltState := range dbAltStates {
		altStateChanges, err := cfg.getAltStateChanges(r, mon, dbAltState)
		if err != nil {
			return []AlteredState{}, err
		}

		alteredState := AlteredState{
			URL:         fmt.Sprintf("http://%s/api/monsters/%d?altered-state=%d", cfg.host, mon.ID, i+1),
			Condition:   dbAltState.Condition,
			IsTemporary: dbAltState.IsTemporary,
			Changes:     altStateChanges,
		}

		alteredStates = append(alteredStates, alteredState)
	}

	return alteredStates, nil
}

func (cfg *apiConfig) getAltStateChanges(r *http.Request, mon database.Monster, as database.AlteredState) ([]AltStateChange, error) {
	altStateChanges := []AltStateChange{}

	dbAltStateChanges, err := cfg.db.GetAltStateChanges(r.Context(), as.ID)
	if err != nil {
		return []AltStateChange{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get altered states relationships of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	for _, dbChange := range dbAltStateChanges {
		altStateChange, err := cfg.getAltStateChangeRelationships(r, mon, dbChange)
		if err != nil {
			return []AltStateChange{}, nil
		}

		altStateChange.AlterationType = dbChange.AlterationType
		altStateChange.Distance = anyToInt32Ptr(dbChange.Distance)

		altStateChanges = append(altStateChanges, altStateChange)
	}

	return altStateChanges, nil
}

func (cfg *apiConfig) getAltStateChangeRelationships(r *http.Request, mon database.Monster, asc database.AltStateChange) (AltStateChange, error) {
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

	addedStatusses, err := cfg.getAltStateStatusses(r, mon, asc)
	if err != nil {
		return AltStateChange{}, err
	}

	return AltStateChange{
		Properties:       properties,
		AutoAbilities:    autoAbilities,
		BaseStats:        baseStats,
		ElemResists:      elemResists,
		StatusImmunities: immunities,
		AddedStatusses:   addedStatusses,
	}, nil
}

func (cfg *apiConfig) getAltStateProperties(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]NamedAPIResource, error) {
	dbProperties, err := cfg.db.GetAltStateProperties(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state properties of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	properties := createNamedAPIResourcesSimple(cfg, dbProperties, "properties", func(prop database.GetAltStatePropertiesRow) (int32, string) {
		return h.NullInt32ToVal(prop.PropertyID), h.NullStringToVal(prop.Property)
	})

	if len(properties) == 0 {
		return nil, nil
	}

	return properties, nil
}

func (cfg *apiConfig) getAltStateAutoAbilities(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]NamedAPIResource, error) {
	dbAutoAbilities, err := cfg.db.GetAltStateAutoAbilities(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state auto abilities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	autoAbilities := createNamedAPIResourcesSimple(cfg, dbAutoAbilities, "auto-abilities", func(autoAbility database.GetAltStateAutoAbilitiesRow) (int32, string) {
		return h.NullInt32ToVal(autoAbility.AutoAbilityID), h.NullStringToVal(autoAbility.AutoAbility)
	})

	if len(autoAbilities) == 0 {
		return nil, nil
	}

	return autoAbilities, nil
}

func (cfg *apiConfig) getAltStateBaseStats(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]BaseStat, error) {
	dbBaseStats, err := cfg.db.GetAltStateBaseStats(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state base stats of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	var baseStats []BaseStat

	for _, dbStat := range dbBaseStats {
		baseStat := cfg.newBaseStat(dbStat.StatID.Int32, dbStat.Value.Int32, dbStat.Stat.String)

		baseStats = append(baseStats, baseStat)
	}

	return baseStats, nil
}

func (cfg *apiConfig) getAltStateElemResists(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]ElementalResist, error) {
	dbElemResists, err := cfg.db.GetAltStateElemResists(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state elemental resists of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	var elemResists []ElementalResist

	for _, dbResist := range dbElemResists {
		elemResist := cfg.newElemResist(dbResist.ElementID.Int32, dbResist.AffinityID.Int32, dbResist.Element.String, dbResist.Affinity.String)

		elemResists = append(elemResists, elemResist)
	}

	return elemResists, nil
}

func (cfg *apiConfig) getAltStateImmunities(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]NamedAPIResource, error) {
	dbImmunities, err := cfg.db.GetAltStateImmunities(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state status immunities of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	immunities := createNamedAPIResourcesSimple(cfg, dbImmunities, "status-conditions", func(immunity database.GetAltStateImmunitiesRow) (int32, string) {
		return h.NullInt32ToVal(immunity.StatusID), h.NullStringToVal(immunity.Status)
	})

	if len(immunities) == 0 {
		return nil, nil
	}

	return immunities, nil
}

func (cfg *apiConfig) getAltStateStatusses(r *http.Request, mon database.Monster, asc database.AltStateChange) ([]InflictedStatus, error) {
	dbAddedStatusses, err := cfg.db.GetAltStateStatusses(r.Context(), asc.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get alt state added statusses of Monster %s, Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	var addedStatusses []InflictedStatus

	for _, dbStatus := range dbAddedStatusses {
		addedStatus := cfg.newInflictedStatus(dbStatus.StatusID.Int32, anyToInt32(dbStatus.Probability), dbStatus.Status.String, h.NullInt32ToPtr(dbStatus.Amount), dbStatus.DurationType.DurationType)

		addedStatusses = append(addedStatusses, addedStatus)
	}

	return addedStatusses, nil
}
