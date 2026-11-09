package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertStatusResist(sr seedtypes.StatusResist) apitypes.StatusResist {
	return apitypes.StatusResist{
		StatusCondition: handlerinput.NameToNamedAPIResource(apilookups.E.StatusConditions, sr.StatusCondition, nil),
		Resistance:      sr.Resistance,
	}
}
