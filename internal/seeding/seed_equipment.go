package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EquipmentTable struct {
	ID                      int32
	Type                    string `json:"type"`
	Classification          string `json:"classification"`
	SpecificCharacterID     *int32
	SpecificCharacter       *string         `json:"specific_character"`
	Version                 *int32          `json:"version"`
	Priority                *int32          `json:"priority"`
	RequiredAutoAbilities   []string        `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool   `json:"selectable_auto_abilities"`
	RequiredSlots           *int32           `json:"required_slots"`
	EquipmentNames          []EquipmentName `json:"names"`
}

func (e EquipmentTable) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Type,
		e.Classification,
		h.DerefOrNil(e.SpecificCharacterID),
		h.DerefOrNil(e.Version),
		h.DerefOrNil(e.Priority),
		h.DerefOrNil(e.RequiredSlots),
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
	return fmt.Sprintf("equipment table with type: %s, classification: %s, specific character: %v, version: %v, priority: %v", e.Type, e.Classification, h.PtrToString(e.SpecificCharacter), h.PtrToString(e.Version), h.PtrToString(e.Priority))
}

func (e EquipmentTable) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: e.ID,
	}
}

type AbilityPool struct {
	ID               int32
	EquipmentTableID int32
	PoolIdx          int32
	AutoAbilities    []string `json:"auto_abilities"`
	ReqAmount        int32    `json:"req_amount"`
}

func (p AbilityPool) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.EquipmentTableID,
		p.PoolIdx,
		p.ReqAmount,
	}
}

func (p AbilityPool) GetID() int32 {
	return p.ID
}

func (p AbilityPool) Error() string {
	return fmt.Sprintf("ability pool with equipment table id: %d, req amount: %d", p.EquipmentTableID, p.ReqAmount)
}

type EquipmentName struct {
	ID            int32
	CharacterID   int32
	CharacterName string `json:"character"`
	Name          string `json:"name"`
}

func (e EquipmentName) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
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

func (e EquipmentName) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   e.ID,
		Name: e.Name,
	}
}

type EquipmentTableNameClstlJunction struct {
	Junction
	CelestialWeaponID *int32
}

func (j EquipmentTableNameClstlJunction) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", j),
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
		fmt.Sprintf("%T", j),
		j.ParentID,
		j.ChildID,
		j.AbilityPool,
	}
}

func (l *Lookup) seedEquipment(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/equipment.json"

	var equipmentTables []EquipmentTable
	err := loadJSONFile(string(srcPath), &equipmentTables)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, table := range equipmentTables {
			var err error
			table.SpecificCharacterID, err = assignFKPtr(table.SpecificCharacter, l.Characters)
			if err != nil {
				return h.NewErr(table.Error(), err)
			}

			dbEquipmentTable, err := qtx.CreateEquipmentTable(context.Background(), database.CreateEquipmentTableParams{
				DataHash:            generateDataHash(table),
				Type:                database.EquipType(table.Type),
				Classification:      database.EquipClass(table.Classification),
				SpecificCharacterID: h.GetNullInt32(table.SpecificCharacterID),
				Version:             h.GetNullInt32(table.Version),
				Priority:            h.GetNullInt32(table.Priority),
				RequiredSlots:       table.RequiredSlots,
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
	const srcPath = "data/equipment.json"

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

			err = l.seedEquipmentReqAutoAbilities(qtx, table)
			if err != nil {
				return h.NewErr(table.Error(), err)
			}

			err = l.seedEquipmentTableAbilityPools(qtx, table)
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

func (l *Lookup) seedEquipmentReqAutoAbilities(qtx *database.Queries, table EquipmentTable) error {
	for _, autoAbility := range table.RequiredAutoAbilities {
		junction, err := createJunction(table, autoAbility, l.AutoAbilities)
		if err != nil {
			return err
		}

		err = qtx.CreateEquipmentTablesRequiredAutoAbilitiesJunction(context.Background(), database.CreateEquipmentTablesRequiredAutoAbilitiesJunctionParams{
			DataHash:         generateDataHash(junction),
			EquipmentTableID: junction.ParentID,
			AutoAbilityID:    junction.ChildID,
		})
		if err != nil {
			return h.NewErr(autoAbility, err, "couldn't junction auto ability")
		}
	}

	return nil
}

func (l *Lookup) seedEquipmentTableAbilityPools(qtx *database.Queries, table EquipmentTable) error {
	for i, abilityPool := range table.SelectableAutoAbilities {
		abilityPool.EquipmentTableID = table.ID
		abilityPool.PoolIdx = int32(i) + 1

		dbAbilityPool, err := qtx.CreateAbilityPool(context.Background(), database.CreateAbilityPoolParams{
			DataHash:         generateDataHash(abilityPool),
			EquipmentTableID: abilityPool.EquipmentTableID,
			PoolIdx:          abilityPool.PoolIdx,
			ReqAmount:        abilityPool.ReqAmount,
		})
		if err != nil {
			return h.NewErr(abilityPool.Error(), err, "couldn't create ability pool")
		}
		abilityPool.ID = dbAbilityPool.ID

		err = l.seedAbilityPoolAutoAbilities(qtx, abilityPool)
		if err != nil {
			return h.NewErr(abilityPool.Error(), err)
		}
	}

	return nil
}

func (l *Lookup) seedAbilityPoolAutoAbilities(qtx *database.Queries, abilityPool AbilityPool) error {
	for _, autoAbility := range abilityPool.AutoAbilities {
		junction, err := createJunction(abilityPool, autoAbility, l.AutoAbilities)
		if err != nil {
			return err
		}

		err = qtx.CreateAbilityPoolsAutoAbilitiesJunction(context.Background(), database.CreateAbilityPoolsAutoAbilitiesJunctionParams{
			DataHash:      generateDataHash(junction),
			AbilityPoolID: junction.ParentID,
			AutoAbilityID: junction.ChildID,
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

	equipmentNameLookup, ok := l.EquipmentNames[equipmentName.Name]
	if ok {
		return equipmentNameLookup, nil
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
