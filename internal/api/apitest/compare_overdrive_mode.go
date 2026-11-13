package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareOverdriveModes(test Test, exp OverdriveMode, got apitypes.OverdriveMode) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("description", exp.description, got.Description)
	test.compare("effect", exp.effect, got.Effect)
	test.compare("type", exp.modeType, got.Type)
	test.compare("fill rate", exp.fillRate, got.FillRate)
	test.checkResAmts("actions", exp.actionsAmount, got.Actions)
}
