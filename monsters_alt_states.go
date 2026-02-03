package main

import (
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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
	return c.Distance == nil &&
		c.Properties == nil &&
		c.AutoAbilities == nil &&
		c.BaseStats == nil &&
		c.ElemResists == nil &&
		c.StatusImmunities == nil &&
		c.StatusResists == nil &&
		c.AddedStatus == nil &&
		c.RemovedStatus == nil
}

func getMonsterAlteredStates(cfg *Config, r *http.Request, mon seeding.Monster) []AlteredState {
	alteredStates := []AlteredState{}

	for i, altState := range mon.AlteredStates {
		q := r.URL.Query()
		q.Set("altered_state", strconv.Itoa(i+1))

		alteredState := AlteredState{
			URL:         createResourceURLQuery(cfg, cfg.e.monsters.endpoint, mon.ID, q),
			Condition:   altState.Condition,
			IsTemporary: altState.IsTemporary,
			Changes:     getAltStateChanges(cfg, altState),
		}

		alteredStates = append(alteredStates, alteredState)
	}

	return alteredStates
}

func getAltStateChanges(cfg *Config, as seeding.AlteredState) []AltStateChange {
	altStateChanges := []AltStateChange{}

	for _, change := range as.Changes {
		altStateChange := convertAltStateChange(cfg, change)
		altStateChange.AlterationType = database.AlterationType(change.AlterationType)
		altStateChange.Distance = change.Distance
		altStateChanges = append(altStateChanges, altStateChange)
	}

	return altStateChanges
}

func convertAltStateChange(cfg *Config, asc seeding.AltStateChange) AltStateChange {
	var change AltStateChange

	if asc.Properties != nil {
		properties := namesToNamedAPIResources(cfg, cfg.e.properties, *asc.Properties)
		change.Properties = h.SliceOrNil(properties)
	}

	if asc.AutoAbilities != nil {
		autoAbilities := namesToNamedAPIResources(cfg, cfg.e.autoAbilities, *asc.AutoAbilities)
		change.AutoAbilities = h.SliceOrNil(autoAbilities)
	}

	if asc.BaseStats != nil {
		baseStats := namesToResourceAmounts(cfg, cfg.e.stats, *asc.BaseStats, newBaseStat)
		change.BaseStats = h.SliceOrNil(baseStats)
	}

	if asc.ElemResists != nil {
		elemResists := namesToElemResists(cfg, *asc.ElemResists)
		change.ElemResists = h.SliceOrNil(elemResists)
	}

	if asc.StatusImmunities != nil {
		immunities := namesToNamedAPIResources(cfg, cfg.e.statusConditions, *asc.StatusImmunities)
		change.StatusImmunities = h.SliceOrNil(immunities)
	}

	if asc.AddedStatus != nil {
		addedStatus := newInflictedStatus(cfg, *asc.AddedStatus)
		change.AddedStatus = &addedStatus
	}

	return change
}
