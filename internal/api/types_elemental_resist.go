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

func convertElemResist(server *Server, er seedtypes.ElementalResist) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(server, E.Elements, er.Element, nil),
		Affinity: er.Affinity,
	}
}

func newElemResist(server *Server, element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(server, E.Elements, element, nil),
		Affinity: affinity,
	}
}

func namesToElemResists(server *Server, resists []seedtypes.ElementalResist) []ElementalResist {
	elemResists := []ElementalResist{}

	for _, seedResist := range resists {
		elemResist := newElemResist(server, seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}
