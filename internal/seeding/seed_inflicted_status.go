package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type InflictedStatus struct {
	ID                int32
	StatusConditionID int32
	StatusCondition   string `json:"name"`
	Probability       int32  `json:"probability"`
	DurationType      string `json:"duration_type"`
	Amount            *int32 `json:"amount"`
}

func (is InflictedStatus) ToHashFields() []any {
	return []any{
		is.StatusConditionID,
		is.Probability,
		is.DurationType,
		h.DerefOrNil(is.Amount),
	}
}

func (is InflictedStatus) GetID() int32 {
	return is.ID
}

func (is InflictedStatus) Error() string {
	return fmt.Sprintf("inflicted status with condition: %s, probability: %d, duration type: %s, amount: %v", is.StatusCondition, is.Probability, is.DurationType, is.Amount)
}

func (l *Lookup) seedInflictedStatus(qtx *database.Queries, inflictedStatus InflictedStatus) (InflictedStatus, error) {
	var err error

	inflictedStatus.StatusConditionID, err = assignFK(inflictedStatus.StatusCondition, l.statusConditions)
	if err != nil {
		return InflictedStatus{}, h.GetErr(inflictedStatus.Error(), err)
	}

	dbInflictedStatus, err := qtx.CreateInflictedStatus(context.Background(), database.CreateInflictedStatusParams{
		DataHash:          generateDataHash(inflictedStatus),
		StatusConditionID: inflictedStatus.StatusConditionID,
		Probability:       inflictedStatus.Probability,
		DurationType:      database.DurationType(inflictedStatus.DurationType),
		Amount:            h.GetNullInt32(inflictedStatus.Amount),
	})
	if err != nil {
		return InflictedStatus{}, h.GetErr(inflictedStatus.Error(), err, "couldn't create inflicted status")
	}
	inflictedStatus.ID = dbInflictedStatus.ID

	return inflictedStatus, nil
}
