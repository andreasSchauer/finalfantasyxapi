package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func ElemResist(er seedtypes.ElementalResist) apitypes.ElementalResist {
	return apitypes.ElementalResist{
		Element:  handlerinput.NameToNamedAPIResource(apilookups.E.Elements, er.Element, nil),
		Affinity: er.Affinity,
	}
}
