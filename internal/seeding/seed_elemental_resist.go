package seeding

import (
	"context"

	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ElementalResist struct {
	ID        int32
	ElementID int32
	Element   string `json:"name"`
	Affinity  string `json:"affinity"`
}

func (er ElementalResist) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", er),
		er.ElementID,
		er.Affinity,
	}
}

func (er ElementalResist) ToKeyFields() []any {
	return []any{
		er.ElementID,
		er.Affinity,
	}
}

func (er ElementalResist) GetID() int32 {
	return er.ID
}

func (er *ElementalResist) SetID(id int32) {
	er.ID = id
}

func (er ElementalResist) Error() string {
	return fmt.Sprintf("elemental resist with element: %s, affinity: %s", er.Element, er.Affinity)
}

func (l *Lookup) loop2SeedElementalResists(qtx *database.Queries, ctx context.Context) error {
	resists, err := l.extractElementalResists()
	if err != nil {
		return err
	}

	params := database.CreateElementalResistBulkParams{
		DataHash:  make([]string, len(resists)),
		ElementID: make([]int32, len(resists)),
		Affinity:  make([]database.ElementalAffinity, len(resists)),
	}

	for i, er := range resists {
		params.DataHash[i] = generateDataHash(er)
		params.ElementID[i] = er.ElementID
		params.Affinity[i] = database.ElementalAffinity(er.Affinity)
	}

	dbRows, err := qtx.CreateElementalResistBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elemental resists: %v", err)
	}

	for i, row := range dbRows {
		resists[i].ID = row.ID
		key := Key(resists[i])
		l.ElementalResists[key] = resists[i]
		l.ElementalResistsID[row.ID] = resists[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractElementalResists() ([]ElementalResist, error) {
	resists := []ElementalResist{}
	var err error

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		if autoAbility.AddedElemResist == nil {
			continue
		}

		autoAbility.AddedElemResist.ElementID, err = assignFK(autoAbility.AddedElemResist.Element, l.Elements)
		if err != nil {
			return nil, err
		}

		resists = append(resists, *autoAbility.AddedElemResist)
	}

	for i := range l.json.monsters {
		monster := &l.json.monsters[i]

		for j := range monster.ElemResists {
			resist := &monster.ElemResists[j]

			resist.ElementID, err = assignFK(resist.Element, l.Elements)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}

		for j := range monster.AlteredStates {
			stateResists, err := l.extractAltStateElemResists(&monster.AlteredStates[j])
			if err != nil {
				return nil, err
			}

			resists = append(resists, stateResists...)
		}
	}

	for i := range l.json.statusConditions {
		condition := l.json.statusConditions[i]

		if condition.AddedElemResist == nil {
			continue
		}

		condition.AddedElemResist.ElementID, err = assignFK(condition.AddedElemResist.Element, l.Elements)
		if err != nil {
			return nil, err
		}

		resists = append(resists, *condition.AddedElemResist)
	}

	return dedupeRows(resists, l.Hashes), nil
}

func (l *Lookup) extractAltStateElemResists(state *AlteredState) ([]ElementalResist, error) {
	elemResists := []ElementalResist{}
	var err error

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.ElemResists == nil {
			continue
		}

		for j := range change.ElemResists {
			resist := &change.ElemResists[j]

			resist.ElementID, err = assignFK(resist.Element, l.Elements)
			if err != nil {
				return nil, err
			}
			elemResists = append(elemResists, *resist)
		}
	}

	return elemResists, nil
}
