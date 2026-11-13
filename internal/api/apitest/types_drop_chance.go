package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

type DropChanceResponse struct {
	TestGeneral
	servicetypes.DropChanceResponse
}

func (e DropChanceResponse) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
