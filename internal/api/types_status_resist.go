package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func newStatusResist(res apiresource.Named, resistance int32) apitypes.StatusResist {
	return apitypes.StatusResist{
		StatusCondition: res,
		Resistance:      resistance,
	}
}

func convertStatusResist(sr seedtypes.StatusResist) apitypes.StatusResist {
	return apitypes.StatusResist{
		StatusCondition: handlerinput.NameToNamedAPIResource(E.StatusConditions, sr.StatusCondition, nil),
		Resistance:      sr.Resistance,
	}
}
