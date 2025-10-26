package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type EquipmentTable struct {
	ID					int32
	Type                string 				`json:"type"`
	Classification      string 				`json:"classification"`
	SpecificCharacterID *int32
	SpecificCharacter   *string 			`json:"specific_character"`
	Version             *int32  			`json:"version"`
	Priority            *int32  			`json:"priority"`
	RequiredAbilities	[]string			`json:"required_abilities"`
	AbilityPool1		[]string			`json:"ability_pool_1"`
	AbilityPool2		[]string			`json:"ability_pool_2"`
	Pool1Amt            *int32  			`json:"pool_1_amt"`
	Pool2Amt            *int32  			`json:"pool_2_amt"`
	EmptySlotsAmt       int32   			`json:"empty_slots_amount"`
	EquipmentNames		[]EquipmentName		`json:"names"`
}

func (e EquipmentTable) ToHashFields() []any {
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

func (e EquipmentTable) ToKeyFields() []any {
	return []any{
		e.Type,
		e.Classification,
		derefOrNil(e.SpecificCharacter),
		derefOrNil(e.Version),
		derefOrNil(e.Priority),
	}
}

func (e EquipmentTable) GetID() int32 {
	return e.ID
}


type EquipmentName struct {
	ID				int32
	CharacterID		int32
	CharacterName	string		`json:"character"`
	Name			string		`json:"name"`
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


type EquipmentTableNameClstlJunction struct {
	Junction
	CelestialWeaponID	*int32
}

func (j EquipmentTableNameClstlJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		derefOrNil(j.CelestialWeaponID),
	}
}


type EquipmentAutoAbilityJunction struct {
	Junction
	AbilityPool		string
}

func (j EquipmentAutoAbilityJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.AbilityPool,
	}
}


func (l *lookup) seedEquipment(db *database.Queries, dbConn *sql.DB) error {
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
				table.SpecificCharacterID, err = assignFKPtr(table.SpecificCharacter, l.getCharacter)
				if err != nil {
					return err
				}
			}

			dbEquipmentTable ,err := qtx.CreateEquipmentTable(context.Background(), database.CreateEquipmentTableParams{
				DataHash:            generateDataHash(table),
				Type:                database.EquipType(table.Type),
				Classification:      database.EquipClass(table.Classification),
				SpecificCharacterID: getNullInt32(table.SpecificCharacterID),
				Version:             getNullInt32(table.Version),
				Priority:            getNullInt32(table.Priority),
				Pool1Amt:            getNullInt32(table.Pool1Amt),
				Pool2Amt:            getNullInt32(table.Pool2Amt),
				EmptySlotsAmt:       table.EmptySlotsAmt,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Equipment Ability. Type: %s, Class: %s, Char: %s, Version %d, Priority: %d: %v", table.Type, table.Classification, derefOrNil(table.SpecificCharacter), derefOrNil(table.Version), derefOrNil(table.Priority), err)
			}

			table.ID = dbEquipmentTable.ID
			key := createLookupKey(table)
			l.equipmentTables[key] = table
		}
		return nil
	})
}



func (l *lookup) createEquipmentRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/equipment.json"

	var equipmentTables []EquipmentTable
	err := loadJSONFile(string(srcPath), &equipmentTables)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonTable := range equipmentTables {
			key := createLookupKey(jsonTable)
			table, err := l.getEquipmentTable(key)
			if err != nil {
				return err
			}

			err = l.seedEquipmentAutoAbilities(qtx, table, table.RequiredAbilities, string(database.AutoAbilityPoolRequired))
			if err != nil {
				return err
			}

			err = l.seedEquipmentAutoAbilities(qtx, table, table.AbilityPool1, string(database.AutoAbilityPoolOne))
			if err != nil {
				return err
			}

			err = l.seedEquipmentAutoAbilities(qtx, table, table.AbilityPool2, string(database.AutoAbilityPoolTwo))
			if err != nil {
				return err
			}

			err = l.seedEquipmentNames(qtx, table)
			if err != nil {
				return err
			}
		}

		return nil
	})
}



func (l *lookup) seedEquipmentAutoAbilities(qtx *database.Queries, table EquipmentTable, autoAbilities []string, abilityPool string) error {
	for _, autoAbility := range autoAbilities {
		var err error
		eaJunction := EquipmentAutoAbilityJunction{}
		
		eaJunction.Junction, err = createJunction(table, autoAbility, l.getAutoAbility)
		if err != nil {
			return err
		}

		eaJunction.AbilityPool = abilityPool

		err = qtx.CreateEquipmentTablesAbilityPoolJunction(context.Background(), database.CreateEquipmentTablesAbilityPoolJunctionParams{
			DataHash: 			generateDataHash(eaJunction),
			EquipmentTableID: 	eaJunction.ParentID,
			AutoAbilityID: 		eaJunction.ChildID,
			AbilityPool: 		database.AutoAbilityPool(eaJunction.AbilityPool),
		})
		if err != nil {
			return fmt.Errorf("couldn't create %s auto abilities for equipment: %s: %v", abilityPool, createLookupKey(table), err)
		}
	}

	return nil
}


func (l *lookup) seedEquipmentNames(qtx *database.Queries, table EquipmentTable) error {
	for _, equipmentName := range table.EquipmentNames {
		var err error

		etncJunction := EquipmentTableNameClstlJunction{}
		etncJunction.Junction, err = createJunctionSeed(qtx, table, equipmentName, l.seedEquipmentName)
		if err != nil {
			return err
		}

		if table.Classification == string(database.EquipClassCelestialWeapon) {
			etncJunction.CelestialWeaponID, err = assignFKPtr(&equipmentName.Name, l.getCelestialWeapon)
			if err != nil {
				return err
			}
		}

		err = qtx.CreateEquipmentTablesNamesJunction(context.Background(), database.CreateEquipmentTablesNamesJunctionParams{
			DataHash: generateDataHash(etncJunction),
			EquipmentTableID: etncJunction.ParentID,
			EquipmentNameID: etncJunction.ChildID,
			CelestialWeaponID: getNullInt32(etncJunction.CelestialWeaponID),
		})
		if err != nil {
			return fmt.Errorf("couldn't create name %s for equipment table: %s: %v", equipmentName.Name, createLookupKey(table), err)
		}

	}
	
	return nil
}


func (l *lookup) seedEquipmentName(qtx *database.Queries, equipmentName EquipmentName) (EquipmentName, error) {
	var err error

	equipmentName.CharacterID, err = assignFK(equipmentName.CharacterName, l.getCharacter)
	if err != nil {
		return EquipmentName{}, err
	}

	dbEquipmentName, err := qtx.CreateEquipmentName(context.Background(), database.CreateEquipmentNameParams{
		DataHash: 		generateDataHash(equipmentName),
		CharacterID: 	equipmentName.CharacterID,
		Name: 			equipmentName.Name,
	})
	if err != nil {
		return EquipmentName{}, fmt.Errorf("couldn't create equipment name: %s: %v", equipmentName.Name, err)
	}

	equipmentName.ID = dbEquipmentName.ID
	l.equipmentNames[equipmentName.Name] = equipmentName

	return equipmentName, nil
}