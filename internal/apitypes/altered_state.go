package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AlteredState struct {
	URL         string `json:"url"`
	Condition   string `json:"condition"`
	IsTemporary bool   `json:"is_temporary"`
	Alts        []Alt  `json:"-"`
	Alterations
}

type Alt struct {
	AlterationType   database.AlterationType        `json:"alteration_type"`
	Distance         *int32                         `json:"distance,omitempty"`
	Properties       []apiresource.NamedAPIResource `json:"properties,omitempty"`
	AutoAbilities    []apiresource.NamedAPIResource `json:"auto_abilities,omitempty"`
	BaseStats        []BaseStat                     `json:"base_stats,omitempty"`
	ElemResists      []ElementalResist              `json:"elem_resists,omitempty"`
	StatusImmunities []apiresource.NamedAPIResource `json:"status_immunities,omitempty"`
	StatusResists    []StatusResist                 `json:"status_resists,omitempty"`
	AddedStatus      *InflictedStatus               `json:"added_status_condition,omitempty"`
	RemovedStatus    *apiresource.NamedAPIResource  `json:"removed_status_condition,omitempty"`
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