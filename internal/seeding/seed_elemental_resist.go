package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ElementalResist struct {
	ID         int32
	ElementID  int32
	AffinityID int32
	Element    string `json:"name"`
	Affinity   string `json:"affinity"`
}

func (er ElementalResist) ToHashFields() []any {
	return []any{
		er.ElementID,
		er.AffinityID,
	}
}

func (er ElementalResist) GetID() int32 {
	return er.ID
}

func (er ElementalResist) Error() string {
	return fmt.Sprintf("elemental resist with element: %s, affinity: %s", er.Element, er.Affinity)
}

func (l *Lookup) seedElementalResist(qtx *database.Queries, elemResist ElementalResist) (ElementalResist, error) {
	var err error

	elemResist.ElementID, err = assignFK(elemResist.Element, l.getElement)
	if err != nil {
		return ElementalResist{}, getErr(elemResist.Error(), err)
	}

	elemResist.AffinityID, err = assignFK(elemResist.Affinity, l.getAffinity)
	if err != nil {
		return ElementalResist{}, getErr(elemResist.Error(), err)
	}

	dbElemAffinity, err := qtx.CreateElementalResist(context.Background(), database.CreateElementalResistParams{
		DataHash:   generateDataHash(elemResist),
		ElementID:  elemResist.ElementID,
		AffinityID: elemResist.AffinityID,
	})
	if err != nil {
		return ElementalResist{}, getErr(elemResist.Error(), err, "couldn't create elemental resist")
	}

	elemResist.ID = dbElemAffinity.ID

	return elemResist, nil
}
