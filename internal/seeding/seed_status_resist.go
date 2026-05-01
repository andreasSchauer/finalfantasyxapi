package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type StatusResist struct {
	ID                int32
	StatusConditionID int32
	StatusCondition   string `json:"name"`
	Resistance        int32  `json:"resistance"`
}

func (sr StatusResist) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", sr),
		sr.StatusConditionID,
		sr.Resistance,
	}
}

func (sr StatusResist) GetID() int32 {
	return sr.ID
}

func (sr *StatusResist) SetID(id int32) {
	sr.ID = id
}

func (sr StatusResist) GetName() string {
	return sr.StatusCondition
}

func (sr StatusResist) GetVersion() *int32 {
	return nil
}

func (sr StatusResist) GetVal() int32 {
	return sr.Resistance
}

func (sr StatusResist) Error() string {
	return fmt.Sprintf("status resist with status %s, resistance %d", sr.StatusCondition, sr.Resistance)
}

func (l *Lookup) seedStatusResist(qtx *database.Queries, statusResist StatusResist) (StatusResist, error) {
	var err error

	statusResist.StatusConditionID, err = assignFK(statusResist.StatusCondition, l.StatusConditions)
	if err != nil {
		return StatusResist{}, h.NewErr(statusResist.Error(), err)
	}

	dbStatusResist, err := qtx.CreateStatusResist(context.Background(), database.CreateStatusResistParams{
		DataHash:          generateDataHash(statusResist),
		StatusConditionID: statusResist.StatusConditionID,
		Resistance:        statusResist.Resistance,
	})
	if err != nil {
		return StatusResist{}, h.NewErr(statusResist.Error(), err, "couldn't create status resist")
	}

	statusResist.ID = dbStatusResist.ID

	return statusResist, nil
}



func (l *Lookup) loop4SeedStatusResists(qtx *database.Queries, ctx context.Context) error {
	resists, err := l.extractStatusResists()
	if err != nil {
		return err
	}

	params := database.CreateStatusResistBulkParams{
		DataHash:   		make([]string, len(resists)),
		StatusConditionID: 	make([]int32, len(resists)),
		Resistance: 		make([]int32, len(resists)),
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