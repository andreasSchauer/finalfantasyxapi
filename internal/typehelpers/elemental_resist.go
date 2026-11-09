package typehelpers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func NewElemResist(element, affinity string) apitypes.ElementalResist {
	return apitypes.ElementalResist{
		Element:  handlerinput.NameToNamedAPIResource(apilookups.E.Elements, element, nil),
		Affinity: affinity,
	}
}

func NamesToElemResists(resists []seedtypes.ElementalResist) []apitypes.ElementalResist {
	elemResists := []apitypes.ElementalResist{}

	for _, seedResist := range resists {
		elemResist := NewElemResist(seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}
