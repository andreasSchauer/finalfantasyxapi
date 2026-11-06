package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type ArenaCreation struct {
	ID                        int32                                                    `json:"id"`
	Name                      string                                                   `json:"name"`
	Category                  string                                                   `json:"category"`
	Monster                   apiresource.NamedAPIResource                             `json:"monster"`
	ParentSubquest            apiresource.QuestAPIResource                             `json:"parent_subquest"`
	Reward                    apiresource.ResourceAmount[apiresource.TypedAPIResource] `json:"reward"`
	RequiredCatchAmount       int32                                                    `json:"required_catch_amount"`
	UnlockedCreationsCategory *string                                                  `json:"unlocked_creations_category,omitempty"`
	RequiredMonsters          []apiresource.NamedAPIResource                           `json:"required_monsters,omitempty"`
}