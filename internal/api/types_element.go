package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Element struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	OppositeElement    *apiresource.NamedAPIResource  `json:"opposite_element"`
	StatusProtection   *apiresource.NamedAPIResource  `json:"status_protection"`
	AutoAbilities      []apiresource.NamedAPIResource `json:"auto_abilities"`
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	EnemyAbilities     []apiresource.NamedAPIResource `json:"enemy_abilities"`
	MonstersWeak       []apiresource.NamedAPIResource `json:"monsters_weak"`
	MonstersHalved     []apiresource.NamedAPIResource `json:"monsters_halved"`
	MonstersImmune     []apiresource.NamedAPIResource `json:"monsters_immune"`
	MonstersAbsorb     []apiresource.NamedAPIResource `json:"monsters_absorb"`
}
