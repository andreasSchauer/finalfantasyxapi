package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertElemResist(er seedtypes.ElementalResist) apitypes.ElementalResist {
	return apitypes.ElementalResist{
		Element:  handlerinput.NameToNamedAPIResource(apilookups.E.Elements, er.Element, nil),
		Affinity: er.Affinity,
	}
}
