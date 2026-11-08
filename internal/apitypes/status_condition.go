package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type StatusCondition struct {
	ID                      int32               `json:"id"`
	Name                    string              `json:"name"`
	Category                string              `json:"category"`
	IsPermanent             bool                `json:"is_permanent"`
	Visualization           *string             `json:"visualization"`
	Effect                  string              `json:"effect"`
	NullifyArmored          *string             `json:"nullify_armored,omitempty"`
	AddedElemResist         *ElementalResist    `json:"added_elem_resist,omitempty"`
	RelatedStats            []apiresource.Named `json:"related_stats"`
	RemovedStatusConditions []apiresource.Named `json:"removed_status_conditions"`
	StatChanges             []StatChange        `json:"stat_changes"`
	ModifierChanges         []ModifierChange    `json:"modifier_changes"`
	AutoAbilities           []apiresource.Named `json:"auto_abilities"`
	InflictedBy             *StatusInfliction   `json:"inflicted_by"`
	RemovedBy               *StatusRemoval      `json:"removed_by"`
	MonstersResistance      []apiresource.Named `json:"monsters_resistance"`
}
