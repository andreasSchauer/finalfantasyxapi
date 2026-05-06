package api

type Property struct {
	ID             int32              `json:"id"`
	Name           string             `json:"name"`
	Effect         string             `json:"effect"`
	NullifyArmored *string            `json:"nullify_armored,omitempty"`
	RelatedStats   []NamedAPIResource `json:"related_stats"`
	ModifierChange *ModifierChange    `json:"modifier_change"`
	AutoAbilities  []NamedAPIResource `json:"auto_abilities"`
	Monsters       []NamedAPIResource `json:"monsters"`
}
