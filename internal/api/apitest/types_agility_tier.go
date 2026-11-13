package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AgilityTier struct {
	TestGeneral
	IDOnly
	fromAgility int32
	toAgility   int32
	tickSpeed   int32
	monMaxICV   *int32
	monMinICV   *int32
	charMaxICV  *int32
	charMinICVs []apitypes.AgilitySubtier
}

func (e AgilityTier) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
