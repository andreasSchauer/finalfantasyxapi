package api

import (
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
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

func getMonsterAlteredStates(server *Server, r *http.Request, mon seedtypes.Monster) []AlteredState {
	alteredStates := []AlteredState{}

	for i, altState := range mon.AlteredStates {
		var url string

		// if this function is triggered via a jsonParam
		if r != nil {
			q := r.URL.Query()
			q.Set(string(ts.QpnAlteredState), strconv.Itoa(i+1))
			url = CreateResourceURLQuery(server, E.Monsters.Endpoint, mon.ID, q)
		}

		alteredState := AlteredState{
			URL:         url,
			Condition:   altState.Condition,
			IsTemporary: altState.IsTemporary,
			Alts:        getAlts(server, altState),
		}

		alteredStates = append(alteredStates, alteredState)
	}

	return alteredStates
}

func getAlts(server *Server, as seedtypes.AlteredState) []Alt {
	alts := []Alt{}

	for _, change := range as.Changes {
		altStateChange := convertAlt(server, change)
		altStateChange.AlterationType = database.AlterationType(change.AlterationType)
		altStateChange.Distance = change.Distance
		alts = append(alts, altStateChange)
	}

	return alts
}

func convertAlt(server *Server, asc seedtypes.Alt) Alt {
	return Alt{
		Properties:       h.SliceOrNil(namesToNamedAPIResources(server, E.Properties, asc.Properties)),
		AutoAbilities:    h.SliceOrNil(namesToNamedAPIResources(server, E.AutoAbilities, asc.AutoAbilities)),
		BaseStats:        h.SliceOrNil(E.Stats.toResAmtType(server, asc.BaseStats, newBaseStat)),
		ElemResists:      h.SliceOrNil(namesToElemResists(server, asc.ElemResists)),
		StatusImmunities: h.SliceOrNil(namesToNamedAPIResources(server, E.StatusConditions, asc.StatusImmunities)),
		AddedStatus:      convertObjPtr(server, asc.AddedStatus, convertInflictedStatus),
	}
}
