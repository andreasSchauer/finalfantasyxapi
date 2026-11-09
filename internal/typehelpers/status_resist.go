package typehelpers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

func NewStatusResist(res apiresource.Named, resistance int32) apitypes.StatusResist {
	return apitypes.StatusResist{
		StatusCondition: res,
		Resistance:      resistance,
	}
}
