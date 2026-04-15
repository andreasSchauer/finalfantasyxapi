package api

type Property struct {
	ID                  int32				`json:"id"`
	Name                string           	`json:"name"`
	Effect              string           	`json:"effect"`
	NullifyArmored      *string          	`json:"nullify_armored,omitempty"`
	RelatedStats        []NamedAPIResource  `json:"related_stats"`
	StatChanges         []StatChange     	`json:"stat_changes"`
	ModifierChanges     []ModifierChange 	`json:"modifier_changes"`
	AutoAbilities		[]NamedAPIResource	`json:"auto_abilities"`
	Monsters			[]NamedAPIResource	`json:"monsters"`
}