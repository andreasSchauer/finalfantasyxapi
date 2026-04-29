package seeding

import (
	"context"
	"database/sql"
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
		fmt.Sprintf("%T", a),
		a.AccSource,
		h.DerefOrNil(a.HitChance),
		h.DerefOrNil(a.AccModifier),
	}
}

func (a Accuracy) GetID() int32 {
	return a.ID
}

func (a Accuracy) Error() string {
	return fmt.Sprintf("accuracy with source: %s, hit chance: %v, modifier: %v", a.AccSource, h.PtrToString(a.HitChance), h.PtrToString(a.AccModifier))
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


func (l *Lookup) loop1SeedAccuracies(qtx *database.Queries, ctx context.Context) error {
	accuracies := l.extractAccuracies()

	params := database.CreateAbilityAccuracyBulkParams{
		DataHash:    make([]string, len(accuracies)),
		AccSource:   make([]database.AccSourceType, len(accuracies)),
		HitChance:   make([]sql.NullInt32, len(accuracies)),
		AccModifier: make([]sql.NullFloat64, len(accuracies)),
	}

	for i, a := range accuracies {
		params.DataHash[i] = generateDataHash(a)
		params.AccSource[i] = database.AccSourceType(a.AccSource)
		params.HitChance[i] = h.GetNullInt32(a.HitChance)
		params.AccModifier[i] = h.GetNullFloat64(a.AccModifier)
	}

	dbRows, err := qtx.CreateAbilityAccuracyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability accuracies: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAccuracies() []Accuracy {
	accuracies := []Accuracy{}

	for _, ability := range l.json.enemyAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.items {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.overdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.playerAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.triggerCommands {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.unspecifiedAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, aeon := range l.json.aeons {
		if aeon.PhysAtkAccuracy != nil {
			accuracies = append(accuracies, *aeon.PhysAtkAccuracy)
		}
	}

	return dedupeRows(accuracies, l.Hashes)
}