package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Element struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	OppositeElement    *apiresource.Named  `json:"opposite_element"`
	StatusProtection   *apiresource.Named  `json:"status_protection"`
	AutoAbilities      []apiresource.Named `json:"auto_abilities"`
	PlayerAbilities    []apiresource.Named `json:"player_abilities"`
	OverdriveAbilities []apiresource.Named `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.Named `json:"item_abilities"`
	EnemyAbilities     []apiresource.Named `json:"enemy_abilities"`
	MonstersWeak       []apiresource.Named `json:"monsters_weak"`
	MonstersHalved     []apiresource.Named `json:"monsters_halved"`
	MonstersImmune     []apiresource.Named `json:"monsters_immune"`
	MonstersAbsorb     []apiresource.Named `json:"monsters_absorb"`
}
