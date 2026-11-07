package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func newElemResist(element, affinity string) apitypes.ElementalResist {
	return apitypes.ElementalResist{
		Element:  NameToNamedAPIResource(E.Elements, element, nil),
		Affinity: affinity,
	}
}

func convertElemResist(er seedtypes.ElementalResist) apitypes.ElementalResist {
	return apitypes.ElementalResist{
		Element:  NameToNamedAPIResource(E.Elements, er.Element, nil),
		Affinity: er.Affinity,
	}
}

func namesToElemResists(resists []seedtypes.ElementalResist) []apitypes.ElementalResist {
	elemResists := []apitypes.ElementalResist{}

	for _, seedResist := range resists {
		elemResist := newElemResist(seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}
