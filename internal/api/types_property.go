package api

type Property struct {
	ID                  int32				`json:"id"`
	Name                string           	`json:"name"`
	Effect              string           	`json:"effect"`
	RelatedStats        []NamedAPIResource  `json:"related_stats"`
	NullifyArmored      *string          	`json:"nullify_armored"`
	StatChanges         []StatChange     	`json:"stat_changes"`
	ModifierChanges     []ModifierChange 	`json:"modifier_changes"`
	AutoAbilities		[]NamedAPIResource	`json:"auto_abilities"`
	Monsters			[]NamedAPIResource	`json:"monsters"`
}