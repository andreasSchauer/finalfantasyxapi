package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type CharacterClass struct {
	ID                  int32                          `json:"id"`
	Name                string                         `json:"name"`
	Category            string                         `json:"category"`
	Members             []apiresource.TypedAPIResource `json:"members"`
	DefaultAbilities    []apiresource.TypedAPIResource `json:"default_abilities"`
	LearnableAbilities  []apiresource.TypedAPIResource `json:"learnable_abilities"`
	DefaultOverdrives   []apiresource.NamedAPIResource `json:"default_overdrives"`
	LearnableOverdrives []apiresource.NamedAPIResource `json:"learnable_overdrives"`
	Submenus            []apiresource.NamedAPIResource `json:"submenus"`
}
