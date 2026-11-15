package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type MonsterFormation struct {
	TestGeneral
	IDOnly
	Category        string
	IsForcedAmbush  bool
	CanEscape       bool
	BossMusic       *int32
	Ap              int32
	ApOverkill      int32
	Gil             int32
	Monsters        map[string]int32
	Areas           []int32
	TriggerCommands []TestFormationTC
}

func (e MonsterFormation) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e MonsterFormation) Compare(test Test, got apitypes.MonsterFormation) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compare("category", e.Category, got.Category)
	test.compare("is forced ambush", e.IsForcedAmbush, got.IsForcedAmbush)
	test.compare("can escape", e.CanEscape, got.CanEscape)
	test.compare("ap", e.Ap, got.AP)
	test.compare("ap overkill", e.ApOverkill, got.ApOverkill)
	test.compare("gil", e.Gil, got.Gil)
	test.compIDAPIResourcePtrs("boss song", apilookups.E.Songs.Endpoint, e.BossMusic, got.BossMusic)
	test.checkResAmts("monsters", e.Monsters, got.Monsters)
	test.compTestStructSlices("trigger commands", e.TriggerCommands, got.TriggerCommands)
	test.CheckResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
}

type TestFormationTC struct {
	Ability int32
	Users   []int32
}

func (e TestFormationTC) Compare(test Test, fn string, got apitypes.FormationTriggerCommand) {
	test.T.Helper()
	tcEndpoint := apilookups.E.TriggerCommands.Endpoint
	charClassesEndpoint := apilookups.E.CharacterClasses.Endpoint

	test.compIDAPIResource(childFn(fn, "tc ability"), tcEndpoint, e.Ability, got.Ability)
	test.CheckResIDsInSlice(childFn(fn, "tc users"), charClassesEndpoint, e.Users, got.Users)
}
