package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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

type MonsterFilter struct {
	RequiredArea              *string
	RequiredSpecies           *string
	CreationsUnlockedCategory *string
	UnderwaterOnly            bool
}

func (mf MonsterFilter) IsZero() bool {
	return mf.RequiredArea == nil &&
		mf.RequiredSpecies == nil &&
		mf.CreationsUnlockedCategory == nil &&
		!mf.UnderwaterOnly
}

func createMonsterFilter(creation seedtypes.ArenaCreation) MonsterFilter {
	return MonsterFilter{
		RequiredArea:              creation.RequiredArea,
		RequiredSpecies:           creation.RequiredSpecies,
		CreationsUnlockedCategory: creation.CreationsUnlockedCategory,
		UnderwaterOnly:            creation.UnderwaterOnly,
	}
}
