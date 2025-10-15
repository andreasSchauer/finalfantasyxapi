package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type StatusResist struct {
	ID                int32
	StatusConditionID int32
	StatusCondition   string `json:"name"`
	Resistance        int32  `json:"resistance"`
}

func (sr StatusResist) ToHashFields() []any {
	return []any{
		sr.StatusConditionID,
		sr.Resistance,
	}
}

func (sr StatusResist) GetID() int32 {
	return sr.ID
}

func (l *lookup) seedStatusResist(qtx *database.Queries, statusResist StatusResist) (StatusResist, error) {
	var err error

	statusResist.StatusConditionID, err = assignFK(statusResist.StatusCondition, l.getStatusCondition)
	if err != nil {
		return StatusResist{}, err
	}

	dbStatusResist, err := qtx.CreateStatusResist(context.Background(), database.CreateStatusResistParams{
		DataHash:          generateDataHash(statusResist),
		StatusConditionID: statusResist.StatusConditionID,
		Resistance:        statusResist.Resistance,
	})
	if err != nil {
		return StatusResist{}, fmt.Errorf("couldn't create status resist: %s - %d: %v", statusResist.StatusCondition, statusResist.Resistance, err)
	}

	statusResist.ID = dbStatusResist.ID

	return statusResist, nil
}
