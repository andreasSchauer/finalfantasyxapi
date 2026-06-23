package api

import (
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AlteredState struct {
	URL         string `json:"url"`
	Condition   string `json:"condition"`
	IsTemporary bool   `json:"is_temporary"`
	Alts        []Alt  `json:"-"`
	Alterations
}

func getAlteredStatesAlterations(states []AlteredState) []AlteredState {
	for i := range states {
		states[i].Alterations = getAlterations(states[i].Alts)
	}

	return states
}

type Alt struct {
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

func (c *Alt) IsZero() bool {
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
		q.Set(string(qpnAlteredState), strconv.Itoa(i+1))

		alteredState := AlteredState{
			URL:         createResourceURLQuery(cfg, cfg.e.monsters.endpoint, mon.ID, q),
			Condition:   altState.Condition,
			IsTemporary: altState.IsTemporary,
			Alts:        getAlts(cfg, altState),
		}

		alteredStates = append(alteredStates, alteredState)
	}

	return alteredStates
}

func getAlts(cfg *Config, as seeding.AlteredState) []Alt {
	alts := []Alt{}

	for _, change := range as.Changes {
		altStateChange := convertAlt(cfg, change)
		altStateChange.AlterationType = database.AlterationType(change.AlterationType)
		altStateChange.Distance = change.Distance
		alts = append(alts, altStateChange)
	}

	return alts
}

func convertAlt(cfg *Config, asc seeding.Alt) Alt {
	return Alt{
		Properties:       h.SliceOrNil(namesToNamedAPIResources(cfg, cfg.e.properties, asc.Properties)),
		AutoAbilities:    h.SliceOrNil(namesToNamedAPIResources(cfg, cfg.e.autoAbilities, asc.AutoAbilities)),
		BaseStats:        h.SliceOrNil(toResAmtType(cfg, cfg.e.stats, asc.BaseStats, newBaseStat)),
		ElemResists:      h.SliceOrNil(namesToElemResists(cfg, asc.ElemResists)),
		StatusImmunities: h.SliceOrNil(namesToNamedAPIResources(cfg, cfg.e.statusConditions, asc.StatusImmunities)),
		AddedStatus:      convertObjPtr(cfg, asc.AddedStatus, convertInflictedStatus),
	}
}
