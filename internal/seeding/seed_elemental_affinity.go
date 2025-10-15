package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ElementalAffinity struct {
	ID         int32
	ElementID  int32
	AffinityID int32
	Element    string `json:"name"`
	Affinity   string `json:"affinity"`
}

func (ea ElementalAffinity) ToHashFields() []any {
	return []any{
		ea.ElementID,
		ea.AffinityID,
	}
}

func (ea ElementalAffinity) GetID() int32 {
	return ea.ID
}

func (l *lookup) seedElementalAffinity(qtx *database.Queries, elemAffinity ElementalAffinity) (ElementalAffinity, error) {
	var err error

	elemAffinity.ElementID, err = assignFK(elemAffinity.Element, l.getElement)
	if err != nil {
		return ElementalAffinity{}, err
	}

	elemAffinity.AffinityID, err = assignFK(elemAffinity.Affinity, l.getAffinity)
	if err != nil {
		return ElementalAffinity{}, err
	}

	dbElemAffinity, err := qtx.CreateElementalAffinity(context.Background(), database.CreateElementalAffinityParams{
		DataHash:   generateDataHash(elemAffinity),
		ElementID:  elemAffinity.ElementID,
		AffinityID: elemAffinity.AffinityID,
	})
	if err != nil {
		return ElementalAffinity{}, fmt.Errorf("couldn't create Elemental Affinity: %v", err)
	}

	elemAffinity.ID = dbElemAffinity.ID

	return elemAffinity, nil
}
