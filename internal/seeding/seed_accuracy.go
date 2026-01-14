package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Accuracy struct {
	ID          int32
	AccSource   string   `json:"acc_source"`
	HitChance   *int32   `json:"hit_chance"`
	AccModifier *float32 `json:"acc_modifier"`
}

func (a Accuracy) ToHashFields() []any {
	return []any{
		a.AccSource,
		h.DerefOrNil(a.HitChance),
		h.DerefOrNil(a.AccModifier),
	}
}

func (a Accuracy) GetID() int32 {
	return a.ID
}

func (a Accuracy) Error() string {
	return fmt.Sprintf("accuracy with source: %s, hit chance: %v, modifier: %v", a.AccSource, h.DerefOrNil(a.HitChance), h.DerefOrNil(a.AccModifier))
}

func (l *Lookup) seedAccuracy(qtx *database.Queries, accuracy Accuracy) (Accuracy, error) {
	dbAccuracy, err := qtx.CreateAbilityAccuracy(context.Background(), database.CreateAbilityAccuracyParams{
		DataHash:    generateDataHash(accuracy),
		AccSource:   database.AccSourceType(accuracy.AccSource),
		HitChance:   h.GetNullInt32(accuracy.HitChance),
		AccModifier: h.GetNullFloat64(accuracy.AccModifier),
	})
	if err != nil {
		return Accuracy{}, h.NewErr(accuracy.Error(), err, "couldn't create accuracy")
	}

	accuracy.ID = dbAccuracy.ID

	return accuracy, nil
}
