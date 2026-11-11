package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func StatusResist(sr seedtypes.StatusResist) apitypes.StatusResist {
	return apitypes.StatusResist{
		StatusCondition: handlerinput.NameToNamedAPIResource(apilookups.E.StatusConditions, sr.StatusCondition, nil),
		Resistance:      sr.Resistance,
	}
}
