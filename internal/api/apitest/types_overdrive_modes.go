package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expOverdriveMode struct {
	testGeneral
	expUnique
	description   string
	effect        string
	modeType      string
	fillRate      *float32
	actionsAmount map[string]int32
}

func (e expOverdriveMode) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdriveModes(test test, exp expOverdriveMode, got apitypes.OverdriveMode) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("description", exp.description, got.Description)
	test.compare("effect", exp.effect, got.Effect)
	test.compare("type", exp.modeType, got.Type)
	test.compare("fill rate", exp.fillRate, got.FillRate)
	test.checkResAmts("actions", exp.actionsAmount, got.Actions)
}
