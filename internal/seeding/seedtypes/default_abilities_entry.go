package seedtypes

type DefaultAbilitiesEntry struct {
	Name             string                       `json:"name"`
	DefaultAbilities []AbilityReference `json:"default_abilities"`
}