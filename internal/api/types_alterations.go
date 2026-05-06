package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
	Properties       []NamedAPIResource `json:"properties,omitempty"`
	AutoAbilities    []NamedAPIResource `json:"auto_abilities,omitempty"`
	StatusImmunities []NamedAPIResource `json:"status_immunities,omitempty"`
	StatusResists    []StatusResist     `json:"status_resists,omitempty"`
	Status           *InflictedStatus   `json:"added_status_condition,omitempty"`
}

type AltStateLoss struct {
	Properties       []NamedAPIResource `json:"properties,omitempty"`
	AutoAbilities    []NamedAPIResource `json:"auto_abilities,omitempty"`
	StatusImmunities []NamedAPIResource `json:"status_immunities,omitempty"`
	Status           *NamedAPIResource  `json:"removed_status_condition,omitempty"`
}

func getAlterations(alts []Alt) Alterations {
	alterations := Alterations{}

	for _, alt := range alts {
		switch alt.AlterationType {
		case database.AlterationTypeChange:
			change := AltStateChange{
				Distance:    alt.Distance,
				BaseStats:   alt.BaseStats,
				ElemResists: alt.ElemResists,
			}
			alterations.Change = &change

		case database.AlterationTypeGain:
			gain := AltStateGain{
				Properties:       alt.Properties,
				AutoAbilities:    alt.AutoAbilities,
				StatusImmunities: alt.StatusImmunities,
				StatusResists:    alt.StatusResists,
				Status:           alt.AddedStatus,
			}
			alterations.Gain = &gain

		case database.AlterationTypeLoss:
			loss := AltStateLoss{
				Properties:       alt.Properties,
				AutoAbilities:    alt.AutoAbilities,
				StatusImmunities: alt.StatusImmunities,
				Status:           alt.RemovedStatus,
			}
			alterations.Loss = &loss
		}
	}

	return alterations
}
