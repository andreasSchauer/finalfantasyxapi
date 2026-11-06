package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type StatusInfliction struct {
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	MiscAbilities      []apiresource.NamedAPIResource `json:"misc_abilities"`
	EnemyAbilities     []apiresource.NamedAPIResource `json:"enemy_abilities"`
}

func (s StatusInfliction) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.MiscAbilities) == 0 &&
		len(s.EnemyAbilities) == 0
}

type StatusRemoval struct {
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	EnemyAbilities     []apiresource.NamedAPIResource `json:"enemy_abilities"`
	StatusConditions   []apiresource.NamedAPIResource `json:"status_conditions"`
}

func (s StatusRemoval) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.EnemyAbilities) == 0 &&
		len(s.StatusConditions) == 0
}