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

func convertElemResist(srv *Server, er seedtypes.ElementalResist) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(srv, E.Elements, er.Element, nil),
		Affinity: er.Affinity,
	}
}

func newElemResist(srv *Server, element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(srv, E.Elements, element, nil),
		Affinity: affinity,
	}
}

func namesToElemResists(srv *Server, resists []seedtypes.ElementalResist) []ElementalResist {
	elemResists := []ElementalResist{}

	for _, seedResist := range resists {
		elemResist := newElemResist(srv, seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}
