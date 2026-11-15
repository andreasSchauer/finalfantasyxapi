package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type OverdriveMode struct {
	TestGeneral
	Unique
	Description   string
	Effect        string
	ModeType      string
	FillRate      *float32
	ActionsAmount map[string]int32
}

func (e OverdriveMode) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e OverdriveMode) Compare(test Test, got apitypes.OverdriveMode) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("description", e.Description, got.Description)
	test.compare("effect", e.Effect, got.Effect)
	test.compare("type", e.ModeType, got.Type)
	test.compare("fill rate", e.FillRate, got.FillRate)
	test.checkResAmts("actions", e.ActionsAmount, got.Actions)
}
