package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop2ElementalResists(qtx *database.Queries, ctx context.Context) error {
	resists, err := extractElementalResists()
	if err != nil {
		return err
	}

	params := database.CreateElementalResistBulkParams{
		DataHash:  make([]string, len(resists)),
		ElementID: make([]int32, len(resists)),
		Affinity:  make([]database.ElementalAffinity, len(resists)),
	}

	for i, er := range resists {
		params.DataHash[i] = sh.GenerateDataHash(er)
		params.ElementID[i] = er.ElementID
		params.Affinity[i] = database.ElementalAffinity(er.Affinity)
	}

	dbRows, err := qtx.CreateElementalResistBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elemental resists: %v", err)
	}

	for i, row := range dbRows {
		resists[i].ID = row.ID
		lookup.L.ElementalResists[seedtypes.Key(resists[i])] = resists[i]
		lookup.L.ElementalResistsID[row.ID] = resists[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractElementalResists() ([]seedtypes.ElementalResist, error) {
	resists := []seedtypes.ElementalResist{}
	var err error

	for i := range sh.RawJSON.AutoAbilities {
		autoAbility := &sh.RawJSON.AutoAbilities[i]

		if autoAbility.AddedElemResist == nil {
			continue
		}

		autoAbility.AddedElemResist.ElementID, err = sh.AssignFK(autoAbility.AddedElemResist.Element, lookup.L.Elements)
		if err != nil {
			return nil, err
		}

		resists = append(resists, *autoAbility.AddedElemResist)
	}

	for i := range sh.RawJSON.Monsters {
		monster := &sh.RawJSON.Monsters[i]

		for j := range monster.ElemResists {
			resist := &monster.ElemResists[j]

			resist.ElementID, err = sh.AssignFK(resist.Element, lookup.L.Elements)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}

		for j := range monster.AlteredStates {
			stateResists, err := extractAltStateElemResists(&monster.AlteredStates[j])
			if err != nil {
				return nil, err
			}

			resists = append(resists, stateResists...)
		}
	}

	for i := range sh.RawJSON.StatusConditions {
		condition := sh.RawJSON.StatusConditions[i]

		if condition.AddedElemResist == nil {
			continue
		}

		condition.AddedElemResist.ElementID, err = sh.AssignFK(condition.AddedElemResist.Element, lookup.L.Elements)
		if err != nil {
			return nil, err
		}

		resists = append(resists, *condition.AddedElemResist)
	}

	return sh.DedupeRows(resists, sh.Hashes), nil
}

func extractAltStateElemResists(state *seedtypes.AlteredState) ([]seedtypes.ElementalResist, error) {
	elemResists := []seedtypes.ElementalResist{}
	var err error

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.ElemResists == nil {
			continue
		}

		for j := range change.ElemResists {
			resist := &change.ElemResists[j]

			resist.ElementID, err = sh.AssignFK(resist.Element, lookup.L.Elements)
			if err != nil {
				return nil, err
			}
			elemResists = append(elemResists, *resist)
		}
	}

	return elemResists, nil
}
