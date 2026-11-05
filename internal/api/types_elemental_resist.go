package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type ElementalResist struct {
	Element  NamedAPIResource `json:"element"`
	Affinity string           `json:"affinity"`
}

func (er ElementalResist) GetAPIResource() APIResource {
	return er.Element
}

func convertElemResist(er seedtypes.ElementalResist) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(E.Elements, er.Element, nil),
		Affinity: er.Affinity,
	}
}

func newElemResist(element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(E.Elements, element, nil),
		Affinity: affinity,
	}
}

func namesToElemResists(resists []seedtypes.ElementalResist) []ElementalResist {
	elemResists := []ElementalResist{}

	for _, seedResist := range resists {
		elemResist := newElemResist(seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}
