package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type ArenaCreation struct {
	TestGeneral
	Unique
	category                  string
	monster                   int32
	parentSubquest            int32
	reward                    TestResAmount[apiresource.Typed]
	requiredCatchAmount       int32
	unlockedCreationsCategory *string
	requiredMonsters          []int32
}

func (e ArenaCreation) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
