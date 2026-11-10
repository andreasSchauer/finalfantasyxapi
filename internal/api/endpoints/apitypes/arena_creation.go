package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type ArenaCreation struct {
	ID                        int32                                         `json:"id"`
	Name                      string                                        `json:"name"`
	Category                  string                                        `json:"category"`
	Monster                   apiresource.Named                             `json:"monster"`
	ParentSubquest            apiresource.Quest                             `json:"parent_subquest"`
	Reward                    apiresource.ResourceAmount[apiresource.Typed] `json:"reward"`
	RequiredCatchAmount       int32                                         `json:"required_catch_amount"`
	UnlockedCreationsCategory *string                                       `json:"unlocked_creations_category,omitempty"`
	RequiredMonsters          []apiresource.Named                           `json:"required_monsters,omitempty"`
}
