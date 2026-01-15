package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EquipmentTable struct {
	ID                  int32
	Type                string `json:"type"`
	Classification      string `json:"classification"`
	SpecificCharacterID *int32
	SpecificCharacter   *string         `json:"specific_character"`
	Version             *int32          `json:"version"`
	Priority            *int32          `json:"priority"`
	RequiredAbilities   []string        `json:"required_abilities"`
	AbilityPool1        []string        `json:"ability_pool_1"`
	AbilityPool2        []string        `json:"ability_pool_2"`
	Pool1Amt            *int32          `json:"pool_1_amt"`
	Pool2Amt            *int32          `json:"pool_2_amt"`
	EmptySlotsAmt       int32           `json:"empty_slots_amount"`
	EquipmentNames      []EquipmentName `json:"names"`
}

func (e EquipmentTable) ToHashFields() []any {
	return []any{
		e.Type,
		e.Classification,
		h.DerefOrNil(e.SpecificCharacterID),
		h.DerefOrNil(e.Version),
		h.DerefOrNil(e.Priority),
		h.DerefOrNil(e.Pool1Amt),
		h.DerefOrNil(e.Pool2Amt),
		e.EmptySlotsAmt,
	}
}

func (e EquipmentTable) ToKeyFields() []any {
	return []any{
		e.Type,
		e.Classification,
		h.DerefOrNil(e.SpecificCharacter),
		h.DerefOrNil(e.Version),
		h.DerefOrNil(e.Priority),
	}
}

func (e EquipmentTable) GetID() int32 {
	return e.ID
}

func (e EquipmentTable) Error() string {
	return fmt.Sprintf("equipment table with type: %s, classification: %s, specific character: %v, version: %v, priority: %v", e.Type, e.Classification, h.DerefOrNil(e.SpecificCharacter), h.DerefOrNil(e.Version), h.DerefOrNil(e.Priority))
}

type EquipmentName struct {
	ID            int32
	CharacterID   int32
	CharacterName string `json:"character"`
	Name          string `json:"name"`
}

func (e EquipmentName) ToHashFields() []any {
	return []any{
		e.CharacterID,
		e.Name,
	}
}

func (e EquipmentName) GetID() int32 {
	return e.ID
}

func (e EquipmentName) Error() string {
	return fmt.Sprintf("equipment name %s, character name: %s", e.Name, e.CharacterName)
}

type EquipmentTableNameClstlJunction struct {
	Junction
	CelestialWeaponID *int32
}

func (j EquipmentTableNameClstlJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		h.DerefOrNil(j.CelestialWeaponID),
	}
}

type EquipmentAutoAbilityJunction struct {
	Junction
	AbilityPool string
}

func (j EquipmentAutoAbilityJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.AbilityPool,
	}
}

func (l *Lookup) seedEquipment(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/equipment.json"

	var equipmentTables []EquipmentTable
	err := loadJSONFile(string(srcPath), &equipmentTables)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, table := range equipmentTables {
			if table.SpecificCharacter != nil {
				var err error
				table.SpecificCharacterID, err = assignFKPtr(table.SpecificCharacter, l.Characters)
				if err != nil {
					return h.NewErr(table.Error(), err)
				}
			}

			dbEquipmentTable, err := qtx.CreateEquipmentTable(context.Background(), database.CreateEquipmentTableParams{
				DataHash:            generateDataHash(table),
				Type:                database.EquipType(table.Type),
				Classification:      database.EquipClass(table.Classification),
				SpecificCharacterID: h.GetNullInt32(table.SpecificCharacterID),
				Version:             h.GetNullInt32(table.Version),
				Priority:            h.GetNullInt32(table.Priority),
				Pool1Amt:            h.GetNullInt32(table.Pool1Amt),
				Pool2Amt:            h.GetNullInt32(table.Pool2Amt),
				EmptySlotsAmt:       table.EmptySlotsAmt,
			})
			if err != nil {
				return h.NewErr(table.Error(), err, "couldn't create equipment table")
			}

			table.ID = dbEquipmentTable.ID
			key := CreateLookupKey(table)
			l.EquipmentTables[key] = table
			l.EquipmentTablesID[table.ID] = table
		}
		return nil
	})
}

func (l *Lookup) seedEquipmentRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/equipment.json"

	var equipmentTables []EquipmentTable
	err := loadJSONFile(string(srcPath), &equipmentTables)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonTable := range equipmentTables {
			key := CreateLookupKey(jsonTable)
			table, err := GetResource(key, l.EquipmentTables)
			if err != nil {
				return err
			}

			err = l.seedEquipmentAutoAbilities(qtx, table, table.RequiredAbilities, string(database.AutoAbilityPoolRequired))
			if err != nil {
				return h.NewErr(table.Error(), err)
			}

			err = l.seedEquipmentAutoAbilities(qtx, table, table.AbilityPool1, string(database.AutoAbilityPoolOne))
			if err != nil {
				return h.NewErr(table.Error(), err)
			}

			err = l.seedEquipmentAutoAbilities(qtx, table, table.AbilityPool2, string(database.AutoAbilityPoolTwo))
			if err != nil {
				return h.NewErr(table.Error(), err)
			}

			err = l.seedEquipmentNames(qtx, table)
			if err != nil {
				return h.NewErr(table.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedEquipmentAutoAbilities(qtx *database.Queries, table EquipmentTable, autoAbilities []string, abilityPool string) error {
	for _, autoAbility := range autoAbilities {
		var err error
		eaJunction := EquipmentAutoAbilityJunction{}

		eaJunction.Junction, err = createJunction(table, autoAbility, l.AutoAbilities)
		if err != nil {
			return err
		}

		eaJunction.AbilityPool = abilityPool

		err = qtx.CreateEquipmentTablesAbilityPoolJunction(context.Background(), database.CreateEquipmentTablesAbilityPoolJunctionParams{
			DataHash:         generateDataHash(eaJunction),
			EquipmentTableID: eaJunction.ParentID,
			AutoAbilityID:    eaJunction.ChildID,
			AbilityPool:      database.AutoAbilityPool(eaJunction.AbilityPool),
		})
		if err != nil {
			return h.NewErr(autoAbility, err, "couldn't junction auto ability")
		}
	}

	return nil
}

func (l *Lookup) seedEquipmentNames(qtx *database.Queries, table EquipmentTable) error {
	for _, equipmentName := range table.EquipmentNames {
		var err error

		etncJunction := EquipmentTableNameClstlJunction{}
		etncJunction.Junction, err = createJunctionSeed(qtx, table, equipmentName, l.seedEquipmentName)
		if err != nil {
			return err
		}

		if table.Classification == string(database.EquipClassCelestialWeapon) {
			etncJunction.CelestialWeaponID, err = assignFKPtr(&equipmentName.Name, l.CelestialWeapons)
			if err != nil {
				return h.NewErr(equipmentName.Error(), err)
			}
		}

		err = qtx.CreateEquipmentTablesNamesJunction(context.Background(), database.CreateEquipmentTablesNamesJunctionParams{
			DataHash:          generateDataHash(etncJunction),
			EquipmentTableID:  etncJunction.ParentID,
			EquipmentNameID:   etncJunction.ChildID,
			CelestialWeaponID: h.GetNullInt32(etncJunction.CelestialWeaponID),
		})
		if err != nil {
			return h.NewErr(equipmentName.Error(), err, "couldn't junction equipment name")
		}

	}

	return nil
}

func (l *Lookup) seedEquipmentName(qtx *database.Queries, equipmentName EquipmentName) (EquipmentName, error) {
	var err error

	equipmentName.CharacterID, err = assignFK(equipmentName.CharacterName, l.Characters)
	if err != nil {
		return EquipmentName{}, h.NewErr(equipmentName.Error(), err)
	}

	dbEquipmentName, err := qtx.CreateEquipmentName(context.Background(), database.CreateEquipmentNameParams{
		DataHash:    generateDataHash(equipmentName),
		CharacterID: equipmentName.CharacterID,
		Name:        equipmentName.Name,
	})
	if err != nil {
		return EquipmentName{}, h.NewErr(equipmentName.Error(), err, "couldn't create equipment name")
	}

	equipmentName.ID = dbEquipmentName.ID
	l.EquipmentNames[equipmentName.Name] = equipmentName
	l.EquipmentNamesID[equipmentName.ID] = equipmentName

	return equipmentName, nil
}
