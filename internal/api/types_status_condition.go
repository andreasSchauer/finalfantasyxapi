package api

type StatusCondition struct {
	ID                  int32             	`json:"id"`
	Name                string              `json:"name"`
	AutoAbilities		[]NamedAPIResource	`json:"auto_abilities"`
	InflictedBy			StatusInteractions	`json:"inflicted_by"`
	RemovedBy			StatusInteractions	`json:"removed_by"`
	MonstersResist		[]NamedAPIResource	`json:"monsters_resist"`
	MonstersImmune		[]NamedAPIResource	`json:"monsters_immune"`
}

type StatusInteractions struct {
	PlayerAbilities			[]NamedAPIResource	`json:"player_abilities"`
	OverdriveAbilities		[]NamedAPIResource	`json:"overdrive_abilities"`
	ItemAbilities			[]NamedAPIResource	`json:"item_abilities"`
	UnspecifiedAbilities	[]NamedAPIResource	`json:"unspecified_abilities"`
	EnemyAbilities			[]NamedAPIResource	`json:"enemy_abilities"`
	StatusConditions		[]NamedAPIResource	`json:"status_conditions,omitempty"`
}