package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type OverdriveMode struct {
	TestGeneral
	Unique
	description   string
	effect        string
	modeType      string
	fillRate      *float32
	actionsAmount map[string]int32
}

func (e OverdriveMode) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e OverdriveMode) Compare(test Test, got apitypes.OverdriveMode) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("description", e.description, got.Description)
	test.compare("effect", e.effect, got.Effect)
	test.compare("type", e.modeType, got.Type)
	test.compare("fill rate", e.fillRate, got.FillRate)
	test.checkResAmts("actions", e.actionsAmount, got.Actions)
}
