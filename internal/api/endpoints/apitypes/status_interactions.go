package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type StatusInfliction struct {
	PlayerAbilities    []apiresource.Named `json:"player_abilities"`
	OverdriveAbilities []apiresource.Named `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.Named `json:"item_abilities"`
	MiscAbilities      []apiresource.Named `json:"misc_abilities"`
	EnemyAbilities     []apiresource.Named `json:"enemy_abilities"`
}

func (s StatusInfliction) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.MiscAbilities) == 0 &&
		len(s.EnemyAbilities) == 0
}

type StatusRemoval struct {
	PlayerAbilities    []apiresource.Named `json:"player_abilities"`
	OverdriveAbilities []apiresource.Named `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.Named `json:"item_abilities"`
	EnemyAbilities     []apiresource.Named `json:"enemy_abilities"`
	StatusConditions   []apiresource.Named `json:"status_conditions"`
}

func (s StatusRemoval) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.EnemyAbilities) == 0 &&
		len(s.StatusConditions) == 0
}
