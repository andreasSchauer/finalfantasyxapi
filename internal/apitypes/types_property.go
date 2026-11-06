package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Property struct {
	ID             int32                          `json:"id"`
	Name           string                         `json:"name"`
	Effect         string                         `json:"effect"`
	NullifyArmored *string                        `json:"nullify_armored,omitempty"`
	RelatedStats   []apiresource.NamedAPIResource `json:"related_stats"`
	ModifierChange *ModifierChange                `json:"modifier_change"`
	AutoAbilities  []apiresource.NamedAPIResource `json:"auto_abilities"`
	Monsters       []apiresource.NamedAPIResource `json:"monsters"`
}
