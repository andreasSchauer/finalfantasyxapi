package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop4SeedStatusResists(qtx *database.Queries, ctx context.Context) error {
	resists, err := l.extractStatusResists()
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractStatusResists() ([]StatusResist, error) {
	resists := []StatusResist{}
	var err error

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		for j := range mon.StatusResists {
			resist := &mon.StatusResists[j]
			resist.StatusConditionID, err = assignFK(resist.StatusCondition, l.StatusConditions)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}
	}

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		for j := range autoAbility.AddedStatusResists {
			resist := &autoAbility.AddedStatusResists[j]
			resist.StatusConditionID, err = assignFK(resist.StatusCondition, l.StatusConditions)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}
	}

	return dedupeRows(resists, l.Hashes), nil
}
