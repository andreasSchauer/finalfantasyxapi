package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"

type AlBhedResponse struct {
	TestGeneral
	servicetypes.AlBhedResponse
}

func (e AlBhedResponse) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
