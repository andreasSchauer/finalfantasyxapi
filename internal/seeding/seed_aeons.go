package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Aeon struct {
	//id 		int32
	//dataHash	string
	PlayerUnit
	UnitID				int32
	UnlockCondition     string   `json:"unlock_condition"`
	Category            *string  `json:"category"`
	IsOptional          bool     `json:"is_optional"`
	BattlesToRegenerate int32    `json:"num_battles_to_regenerate"`
	PhysAtkDmgConstant  *int32   `json:"phys_atk_damage_constant"`
	PhysAtkRange        *int32   `json:"phys_atk_range"`
	PhysAtkShatterRate  *int32   `json:"phys_atk_shatter_rate"`
	PhysAtkAccSource    *string  `json:"phys_atk_acc_source"`
	PhysAtkHitChance    *int32   `json:"phys_atk_hit_chance"`
	PhysAtkAccModifier  *float32 `json:"phys_atk_acc_modifier"`
}

func (a Aeon) ToHashFields() []any {
	return []any{
		a.UnitID,
		a.UnlockCondition,
		a.IsOptional,
		a.BattlesToRegenerate,
		derefOrNil(a.PhysAtkDmgConstant),
		derefOrNil(a.PhysAtkRange),
		derefOrNil(a.PhysAtkShatterRate),
		derefOrNil(a.PhysAtkAccSource),
		derefOrNil(a.PhysAtkHitChance),
		derefOrNil(a.PhysAtkAccModifier),
	}
}


type AeonLookup struct {
	Aeon
	ID		int32
}


func (l *lookup) seedAeons(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeons.json"

	var aeons []Aeon
	err := loadJSONFile(string(srcPath), &aeons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, aeon := range aeons {
			aeon.Type = database.UnitTypeAeon

			dbPlayerUnit, err := l.seedPlayerUnit(qtx, aeon.PlayerUnit)
			if err != nil {
				return err
			}

			aeon.UnitID = dbPlayerUnit.ID

			dbAeon, err := qtx.CreateAeon(context.Background(), database.CreateAeonParams{
				DataHash:              generateDataHash(aeon),
				UnitID:                aeon.UnitID,
				UnlockCondition:       aeon.UnlockCondition,
				IsOptional:            aeon.IsOptional,
				BattlesToRegenerate:   aeon.BattlesToRegenerate,
				PhysAtkDamageConstant: getNullInt32(aeon.PhysAtkDmgConstant),
				PhysAtkRange:          getNullInt32(aeon.PhysAtkRange),
				PhysAtkShatterRate:    getNullInt32(aeon.PhysAtkShatterRate),
				PhysAtkAccSource:      nullAccuracySource(aeon.PhysAtkAccSource),
				PhysAtkHitChance:      getNullInt32(aeon.PhysAtkHitChance),
				PhysAtkAccModifier:    getNullFloat64(aeon.PhysAtkAccModifier),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Aeon: %s: %v", aeon.Name, err)
			}

			key := createLookupKey(aeon.PlayerUnit)
			l.aeons[key] = AeonLookup{
				Aeon: 	aeon,
				ID: 	dbAeon.ID,
			}

			err = l.seedCharacterClasses(qtx, aeon.PlayerUnit)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

