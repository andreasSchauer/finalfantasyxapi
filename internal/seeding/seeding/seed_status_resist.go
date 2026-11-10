package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop4SeedStatusResists(qtx *database.Queries, ctx context.Context) error {
	resists, err := extractStatusResists()
	if err != nil {
		return err
	}

	params := database.CreateStatusResistBulkParams{
		DataHash:          make([]string, len(resists)),
		StatusConditionID: make([]int32, len(resists)),
		Resistance:        make([]int32, len(resists)),
	}

	for i, r := range resists {
		params.DataHash[i] = generateDataHash(r)
		params.StatusConditionID[i] = r.StatusConditionID
		params.Resistance[i] = r.Resistance
	}

	dbRows, err := qtx.CreateStatusResistBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create status resists: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractStatusResists() ([]seedtypes.StatusResist, error) {
	resists := []seedtypes.StatusResist{}
	var err error

	for i := range rawJSON.monsters {
		mon := &rawJSON.monsters[i]

		for j := range mon.StatusResists {
			resist := &mon.StatusResists[j]
			resist.StatusConditionID, err = assignFK(resist.StatusCondition, lookup.L.StatusConditions)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}
	}

	for i := range rawJSON.autoAbilities {
		autoAbility := &rawJSON.autoAbilities[i]

		for j := range autoAbility.AddedStatusResists {
			resist := &autoAbility.AddedStatusResists[j]
			resist.StatusConditionID, err = assignFK(resist.StatusCondition, lookup.L.StatusConditions)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}
	}

	return dedupeRows(resists, hashes), nil
}
