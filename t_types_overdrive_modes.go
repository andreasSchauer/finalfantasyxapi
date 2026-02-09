package main

type expOverdriveMode struct {
	testGeneral
	expUnique
	description   string
	effect        string
	modeType      int32
	fillRate      *float32
	actionsAmount map[string]int32
}

func (e expOverdriveMode) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdriveModes(test test, exp expOverdriveMode, got OverdriveMode) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "description", exp.description, got.Description)
	compare(test, "effect", exp.effect, got.Effect)
	compIdApiResource(test, "type", test.cfg.e.overdriveModeType.endpoint, exp.modeType, got.Type)
	compare(test, "fill rate", exp.fillRate, got.FillRate)
	checkResAmtsNameVals(test, "actions", exp.actionsAmount, got.Actions)
}