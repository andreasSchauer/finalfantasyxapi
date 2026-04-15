package api

type StatusCondition struct {
	ID                  	int32             	`json:"id"`
	Name                	string              `json:"name"`
	Category				NamedAPIResource	`json:"category"`
	IsPermanent				bool			 	`json:"is_permanent"`
	Visualization           *string          	`json:"visualization"`
	Effect                  string           	`json:"effect"`
	NullifyArmored          *string          	`json:"nullify_armored,omitempty"`
	AddedElemResist         *ElementalResist 	`json:"added_elem_resist,omitempty"`
	CtbOnInfliction			*InflictedDelay	 	`json:"ctb_on_infliction,omitempty"`
	RelatedStats            []NamedAPIResource	`json:"related_stats"`
	RemovedStatusConditions []NamedAPIResource  `json:"removed_status_conditions"`
	StatChanges             []StatChange     	`json:"stat_changes"`
	ModifierChanges         []ModifierChange 	`json:"modifier_changes"`
	AutoAbilities			[]NamedAPIResource	`json:"auto_abilities"`
	InflictedBy				*StatusInteractions	`json:"inflicted_by"`
	RemovedBy				*StatusInteractions	`json:"removed_by"`
	MonstersResistance		[]NamedAPIResource	`json:"monsters_resistance"`
}

