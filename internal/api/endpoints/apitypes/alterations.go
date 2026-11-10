package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type Alterations struct {
	Change *AltStateChange `json:"change,omitempty"`
	Gain   *AltStateGain   `json:"gain,omitempty"`
	Loss   *AltStateLoss   `json:"loss,omitempty"`
}

type AltStateChange struct {
	Distance    *int32            `json:"distance,omitempty"`
	BaseStats   []BaseStat        `json:"base_stats,omitempty"`
	ElemResists []ElementalResist `json:"elem_resists,omitempty"`
}

type AltStateGain struct {
	Properties       []apiresource.Named `json:"properties,omitempty"`
	AutoAbilities    []apiresource.Named `json:"auto_abilities,omitempty"`
	StatusImmunities []apiresource.Named `json:"status_immunities,omitempty"`
	StatusResists    []StatusResist      `json:"status_resists,omitempty"`
	Status           *InflictedStatus    `json:"status,omitempty"`
}

type AltStateLoss struct {
	Properties       []apiresource.Named `json:"properties,omitempty"`
	AutoAbilities    []apiresource.Named `json:"auto_abilities,omitempty"`
	StatusImmunities []apiresource.Named `json:"status_immunities,omitempty"`
	Status           *apiresource.Named  `json:"status,omitempty"`
}
