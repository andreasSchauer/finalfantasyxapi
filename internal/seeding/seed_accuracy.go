package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)



type Accuracy struct {
	ID					int32
	AccSource			string		`json:"acc_source"`
	HitChance			*int32		`json:"hit_chance"`
	AccModifier			*float32	`json:"acc_modifier"`
}

func (a Accuracy) ToHashFields() []any {
	return []any{
		a.AccSource,
		derefOrNil(a.HitChance),
		derefOrNil(a.AccModifier),
	}
}

func (a Accuracy) GetID() int32 {
	return a.ID
}



func (l *lookup) seedAccuracy(qtx *database.Queries, accuracy Accuracy) (Accuracy, error) {
	dbAccuracy, err := qtx.CreateAbilityAccuracy(context.Background(), database.CreateAbilityAccuracyParams{
		DataHash: 		generateDataHash(accuracy),
		AccSource: 		database.AccSourceType(accuracy.AccSource),
		HitChance: 		getNullInt32(accuracy.HitChance),
		AccModifier: 	getNullFloat64(accuracy.AccModifier),
	})
	if err != nil {
		return Accuracy{}, fmt.Errorf("couldn't create accuracy: %v", err)
	}

	accuracy.ID = dbAccuracy.ID

	return accuracy, nil
}