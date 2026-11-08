package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Property struct {
	ID             int32               `json:"id"`
	Name           string              `json:"name"`
	Effect         string              `json:"effect"`
	NullifyArmored *string             `json:"nullify_armored,omitempty"`
	RelatedStats   []apiresource.Named `json:"related_stats"`
	ModifierChange *ModifierChange     `json:"modifier_change"`
	AutoAbilities  []apiresource.Named `json:"auto_abilities"`
	Monsters       []apiresource.Named `json:"monsters"`
}
