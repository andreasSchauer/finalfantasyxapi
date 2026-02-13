package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type ArenaCreation struct {
	ID                        int32              `json:"id"`
	Name                      string             `json:"name"`
	Category                  string             `json:"category"`
	Monster                   NamedAPIResource   `json:"monster"`
	ParentSubquest            NamedAPIResource   `json:"parent_subquest"`
	Reward                    ItemAmount         `json:"reward"`
	RequiredCatchAmount       int32              `json:"required_catch_amount"`
	UnlockedCreationsCategory *string            `json:"unlocked_creations_category,omitempty"`
	RequiredMonsters          []NamedAPIResource `json:"required_monsters,omitempty"`
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
		mf.UnderwaterOnly == false
}

func createMonsterFilter(creation seeding.ArenaCreation) MonsterFilter {
	return MonsterFilter{
		RequiredArea:              creation.RequiredArea,
		RequiredSpecies:           creation.RequiredSpecies,
		CreationsUnlockedCategory: creation.CreationsUnlockedCategory,
		UnderwaterOnly:            creation.UnderwaterOnly,
	}
}