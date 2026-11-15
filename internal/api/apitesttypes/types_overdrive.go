package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Overdrive struct {
	TestGeneral
	Unique
	Rank               *int32
	CountdownInSec     *int32
	User               int32
	OverdriveCommand   *int32
	OverdriveAbilities []int32
}

func (e Overdrive) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Overdrive) Compare(test Test, got apitypes.Overdrive) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("rank", e.Rank, got.Rank)
	test.compare("countdown in sec", e.CountdownInSec, got.CountdownInSec)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, e.User, got.User)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, e.OverdriveCommand, got.OverdriveCommand)
	test.CheckResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
}
