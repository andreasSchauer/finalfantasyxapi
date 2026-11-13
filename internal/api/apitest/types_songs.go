package apitest

type Song struct {
	TestGeneral
	Unique
	composer          *string
	arranger          *string
	durationInSeconds int32
	canLoop           bool
	backgroundMusic   []TestBackgroundMusic
	cues              []TestCue
	bossFights        []int32
	fmvs              []int32
}

func (e Song) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestBackgroundMusic struct {
	index                  int
	replacesEncounterMusic bool
	areas                  []int32
}

func (t TestBackgroundMusic) GetIndex() int {
	return t.index
}

type TestCue struct {
	index                  int
	triggerArea            *int32
	includedAreas          []int32
	replacesEncounterMusic bool
	replacesBGMusic        *string
}

func (t TestCue) GetIndex() int {
	return t.index
}
