package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type InflictedStatus struct {
	ID					*int32
	StatusConditionID	int32
	StatusCondition		string	`json:"name"`
	Probability			int32	`json:"probability"`
	DurationType		string	`json:"duration_type"`
	Amount				*int32	`json:"amount"`
}

func (is InflictedStatus) ToHashFields() []any {
	return []any{
		is.StatusConditionID,
		is.Probability,
		is.DurationType,
		derefOrNil(is.Amount),
	}
}

func (is InflictedStatus) GetID() *int32 {
	return is.ID
}


func (l *lookup) seedInflictedStatus(qtx *database.Queries, inflictedStatus InflictedStatus) (InflictedStatus, error) {
	condition, err := l.getStatusCondition(inflictedStatus.StatusCondition)
	if err != nil {
		return InflictedStatus{}, err
	}
	inflictedStatus.StatusConditionID = condition.ID

	dbInflictedStatus, err := qtx.CreateInflictedStatus(context.Background(), database.CreateInflictedStatusParams{
		DataHash: 			generateDataHash(inflictedStatus),
		StatusConditionID: 	inflictedStatus.StatusConditionID,
		Probability: 		inflictedStatus.Probability,
		DurationType: 		database.DurationType(inflictedStatus.DurationType),
		Amount: 			getNullInt32(inflictedStatus.Amount),
	})
	if err != nil {
		return InflictedStatus{}, fmt.Errorf("couldn't create Inflicted Status: %v", err)
	}
	inflictedStatus.ID = &dbInflictedStatus.ID

	return inflictedStatus, nil
}