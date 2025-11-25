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
		sr.StatusConditionID,
		sr.Resistance,
	}
}

func (sr StatusResist) GetID() int32 {
	return sr.ID
}

func (sr StatusResist) Error() string {
	return fmt.Sprintf("status resist with status %s, resistance %d", sr.StatusCondition, sr.Resistance)
}

func (l *Lookup) seedStatusResist(qtx *database.Queries, statusResist StatusResist) (StatusResist, error) {
	var err error

	statusResist.StatusConditionID, err = assignFK(statusResist.StatusCondition, l.getStatusCondition)
	if err != nil {
		return StatusResist{}, h.GetErr(statusResist.Error(), err)
	}

	dbStatusResist, err := qtx.CreateStatusResist(context.Background(), database.CreateStatusResistParams{
		DataHash:          generateDataHash(statusResist),
		StatusConditionID: statusResist.StatusConditionID,
		Resistance:        statusResist.Resistance,
	})
	if err != nil {
		return StatusResist{}, h.GetErr(statusResist.Error(), err, "couldn't create status resist")
	}

	statusResist.ID = dbStatusResist.ID

	return statusResist, nil
}
