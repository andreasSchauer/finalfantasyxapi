package typehelpers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func NewStatusResist(res apiresource.Named, resistance int32) apitypes.StatusResist {
	return apitypes.StatusResist{
		StatusCondition: res,
		Resistance:      resistance,
	}
}
