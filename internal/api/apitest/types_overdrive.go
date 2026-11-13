package apitest

type Overdrive struct {
	TestGeneral
	Unique
	rank               *int32
	countdownInSec     *int32
	user               int32
	overdriveCommand   *int32
	overdriveAbilities []int32
}

func (e Overdrive) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
