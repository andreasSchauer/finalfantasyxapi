package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type StatusResist struct {
	StatusConditionID	int32
	StatusCondition		string	`json:"name"`
	Resistance			int32	`json:"resistance"`
}

func (sr StatusResist) ToHashFields() []any {
	return []any{
		sr.StatusConditionID,
		sr.Resistance,
	}
}


func (l *lookup) seedStatusResist(qtx *database.Queries, statusResist StatusResist) (database.StatusResist, error) {
	condition, err := l.getStatusCondition(statusResist.StatusCondition)
	if err != nil {
		return database.StatusResist{}, err
	}

	statusResist.StatusConditionID = condition.ID

	dbStatusResist, err := qtx.CreateStatusResist(context.Background(), database.CreateStatusResistParams{
		DataHash: generateDataHash(statusResist),
		StatusConditionID: statusResist.StatusConditionID,
		Resistance: statusResist.Resistance,
	})
	if err != nil {
		return database.StatusResist{}, fmt.Errorf("couldn't create status resist: %s - %d: %v", statusResist.StatusCondition, statusResist.Resistance, err)
	}

	return dbStatusResist, nil
}