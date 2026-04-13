package api

type Element struct {
	ID                  int32             	`json:"id"`
	Name                string              `json:"name"`
	OppositeElement		*NamedAPIResource	`json:"opposite_element"`
	StatusProtection	*NamedAPIResource	`json:"status_protection"`
	AutoAbilities		[]NamedAPIResource	`json:"auto_abilities"`
	PlayerAbilities		[]NamedAPIResource	`json:"player_abilities"`
	OverdriveAbilities	[]NamedAPIResource	`json:"overdrive_abilities"`
	ItemAbilities		[]NamedAPIResource	`json:"item_abilities"`
	EnemyAbilities		[]NamedAPIResource	`json:"enemy_abilities"`
	MonstersWeak		[]NamedAPIResource	`json:"monsters_weak"`
	MonstersHalved		[]NamedAPIResource	`json:"monsters_halved"`
	MonstersImmune		[]NamedAPIResource	`json:"monsters_immune"`
	MonstersAbsorb		[]NamedAPIResource	`json:"monsters_absorb"`
}