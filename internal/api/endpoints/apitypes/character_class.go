package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type CharacterClass struct {
	ID                  int32               `json:"id"`
	Name                string              `json:"name"`
	Category            string              `json:"category"`
	Members             []apiresource.Typed `json:"members"`
	DefaultAbilities    []apiresource.Typed `json:"default_abilities"`
	LearnableAbilities  []apiresource.Typed `json:"learnable_abilities"`
	DefaultOverdrives   []apiresource.Named `json:"default_overdrives"`
	LearnableOverdrives []apiresource.Named `json:"learnable_overdrives"`
	Submenus            []apiresource.Named `json:"submenus"`
}
