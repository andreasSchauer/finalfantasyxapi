package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

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
