package api

type PlayerUnit struct {
	ID                     int32              `json:"id"`
	Name                   string             `json:"name"`
	Type				   NamedAPIResource	  `json:"type"`	
	TypedUnit			   NamedAPIResource	  `json:"typed_unit"`
	Area                   AreaAPIResource    `json:"area"`
	CelestialWeapon        *NamedAPIResource  `json:"celestial_weapon"`
	CharacterClasses       []NamedAPIResource `json:"character_classes"`
}
