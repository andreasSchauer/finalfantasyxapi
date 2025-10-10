package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type EquipmentAbilities struct {
	//id 			int32
	//dataHash		string
	Type                string `json:"type"`
	Classification      string `json:"classification"`
	SpecificCharacterID *int32
	SpecificCharacter   *string `json:"specific_character"`
	Version             *int32  `json:"version"`
	Priority            *int32  `json:"priority"`
	Pool1Amt            *int32  `json:"pool_1_amt"`
	Pool2Amt            *int32  `json:"pool_2_amt"`
	EmptySlotsAmt       int32   `json:"empty_slots_amount"`
}

func (e EquipmentAbilities) ToHashFields() []any {
	return []any{
		e.Type,
		e.Classification,
		derefOrNil(e.SpecificCharacterID),
		derefOrNil(e.Version),
		derefOrNil(e.Priority),
		derefOrNil(e.Pool1Amt),
		derefOrNil(e.Pool2Amt),
		e.EmptySlotsAmt,
	}
}

func (l *lookup) seedEquipment(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/equipment.json"

	var equipmentAbilities []EquipmentAbilities
	err := loadJSONFile(string(srcPath), &equipmentAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, entry := range equipmentAbilities {
			if entry.SpecificCharacter != nil {
				character, err := l.getCharacter(*entry.SpecificCharacter)
				if err != nil {
					return err
				}

				entry.SpecificCharacterID = &character.ID
			}

			err = qtx.CreateEquipmentAbility(context.Background(), database.CreateEquipmentAbilityParams{
				DataHash:            generateDataHash(entry),
				Type:                database.EquipType(entry.Type),
				Classification:      database.EquipClass(entry.Classification),
				SpecificCharacterID: getNullInt32(entry.SpecificCharacterID),
				Version:             getNullInt32(entry.Version),
				Priority:            getNullInt32(entry.Priority),
				Pool1Amt:            getNullInt32(entry.Pool1Amt),
				Pool2Amt:            getNullInt32(entry.Pool2Amt),
				EmptySlotsAmt:       entry.EmptySlotsAmt,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Equipment Ability. Type: %s, Class: %s, Char: %s, Version %d, Priority: %d: %v", entry.Type, entry.Classification, derefOrNil(entry.SpecificCharacter), derefOrNil(entry.Version), derefOrNil(entry.Priority), err)
			}
		}
		return nil
	})
}
