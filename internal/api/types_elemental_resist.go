package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ElementalResist struct {
	Element  NamedAPIResource `json:"element"`
	Affinity NamedAPIResource `json:"affinity"`
}

func (er ElementalResist) GetAPIResource() APIResource {
	return er.Element
}

func convertElemResist(cfg *Config, er seeding.ElementalResist) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(cfg, cfg.e.elements, er.Element, nil),
		Affinity: enumToNamedAPIResource(cfg, cfg.e.elementalAffinity.endpoint, er.Affinity, cfg.t.ElementalAffinity),
	}
}

func newElemResist(cfg *Config, element, affinity string) ElementalResist {
	return ElementalResist{
		Element:  nameToNamedAPIResource(cfg, cfg.e.elements, element, nil),
		Affinity: enumToNamedAPIResource(cfg, cfg.e.elementalAffinity.endpoint, affinity, cfg.t.ElementalAffinity),
	}
}

func namesToElemResists(cfg *Config, resists []seeding.ElementalResist) []ElementalResist {
	elemResists := []ElementalResist{}

	for _, seedResist := range resists {
		elemResist := newElemResist(cfg, seedResist.Element, seedResist.Affinity)
		elemResists = append(elemResists, elemResist)
	}

	return elemResists
}
