package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

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
	RequiredSlots           *int32          `json:"required_slots"`
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

func (p *AbilityPool) SetID(id int32) {
	p.ID = id
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

func (e *EquipmentName) SetID(id int32) {
	e.ID = id
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

type EquipmentTableNameJunction struct {
	StdJunction
	CelestialWeaponID *int32
}

func (j EquipmentTableNameJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		h.DerefOrNil(j.CelestialWeaponID),
	}
}

func (j EquipmentTableNameJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
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
			key := Key(table)
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
			key := Key(jsonTable)
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

		etncJunction := EquipmentTableNameJunction{}
		etncJunction.StdJunction, err = createJunctionSeed(qtx, table, equipmentName, l.seedEquipmentName)
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

func (l *Lookup) loop5SeedEquipmentTables(qtx *database.Queries, ctx context.Context) error {
	tables, err := l.extractEquipmentTables()
	if err != nil {
		return err
	}

	params := database.CreateEquipmentTableBulkParams{
		DataHash:            make([]string, len(tables)),
		Type:                make([]database.EquipType, len(tables)),
		Classification:      make([]database.EquipClass, len(tables)),
		SpecificCharacterID: make([]sql.NullInt32, len(tables)),
		Version:             make([]sql.NullInt32, len(tables)),
		Priority:            make([]sql.NullInt32, len(tables)),
		RequiredSlots:       make([]sql.NullInt32, len(tables)),
	}

	for i, t := range tables {
		params.DataHash[i] = generateDataHash(t)
		params.Type[i] = database.EquipType(t.Type)
		params.Classification[i] = database.EquipClass(t.Classification)
		params.SpecificCharacterID[i] = h.GetNullInt32(t.SpecificCharacterID)
		params.Version[i] = h.GetNullInt32(t.Version)
		params.Priority[i] = h.GetNullInt32(t.Priority)
		params.RequiredSlots[i] = h.GetNullInt32(t.RequiredSlots)
	}

	dbRows, err := qtx.CreateEquipmentTableBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment tables: %v", err)
	}

	for i, row := range dbRows {
		tables[i].ID = row.ID
		l.json.equipment[i].ID = row.ID
		key := Key(tables[i])
		l.EquipmentTables[key] = tables[i]
		l.EquipmentTablesID[row.ID] = tables[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEquipmentTables() ([]EquipmentTable, error) {
	tables := []EquipmentTable{}
	var err error

	for i := range l.json.equipment {
		table := &l.json.equipment[i]

		table.SpecificCharacterID, err = assignFKPtr(table.SpecificCharacter, l.Characters)
		if err != nil {
			return nil, err
		}

		tables = append(tables, *table)
	}

	return dedupeRows(tables, l.Hashes), nil
}

func (l *Lookup) completeEquipment() error {
	for i := range l.json.equipment {
		table := &l.json.equipment[i]
		err := assignIDs(l, table.SelectableAutoAbilities)
		if err != nil {
			return err
		}

		err = assignIDs(l, table.EquipmentNames)
		if err != nil {
			return err
		}

		l.EquipmentTables[Key(*table)] = *table
		l.EquipmentTablesID[table.ID] = *table
	}

	return nil
}

func (l *Lookup) getEquipmentReqAutoAbilities(et EquipmentTable) ([]AutoAbility, error) {
	return getResources(et.RequiredAutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncEquipmentReqAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment table + required auto-abilities"
	jParams, err := processJunctions(l, desc, l.json.equipment, l.getEquipmentReqAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulk(ctx, database.CreateEquipmentTablesRequiredAutoAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		EquipmentTableID: jParams.ParentIDs,
		AutoAbilityID:    jParams.ChildIDs,
	})
}

func (l *Lookup) loop5SeedEquipmentNames(qtx *database.Queries, ctx context.Context) error {
	names, err := l.extractEquipmentNames()
	if err != nil {
		return err
	}

	params := database.CreateEquipmentNameBulkParams{
		DataHash:    make([]string, len(names)),
		CharacterID: make([]int32, len(names)),
		Name:        make([]string, len(names)),
	}

	for i, n := range names {
		params.DataHash[i] = generateDataHash(n)
		params.CharacterID[i] = n.CharacterID
		params.Name[i] = n.Name
	}

	dbRows, err := qtx.CreateEquipmentNameBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment names: %v", err)
	}

	for i, row := range dbRows {
		names[i].ID = row.ID
		l.EquipmentNames[names[i].Name] = names[i]
		l.EquipmentNamesID[row.ID] = names[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEquipmentNames() ([]EquipmentName, error) {
	names := []EquipmentName{}
	var err error

	for i := range l.json.equipment {
		table := &l.json.equipment[i]

		for j := range table.EquipmentNames {
			name := &table.EquipmentNames[j]

			name.CharacterID, err = assignFK(name.CharacterName, l.Characters)
			if err != nil {
				return nil, err
			}

			names = append(names, *name)
		}
	}

	return dedupeRows(names, l.Hashes), nil
}

func(l *Lookup) seedJuncEquipmentTablesNames(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment tables + equipment names"
	params := database.CreateEquipmentTablesNamesJunctionBulkParams{
		DataHash: 			make([]string, 0),
		EquipmentTableID: 	make([]int32, 0),
		EquipmentNameID: 	make([]int32, 0),
		CelestialWeaponID: 	make([]sql.NullInt32, 0),
	}

	for _, table := range l.json.equipment {
		for _, name := range table.EquipmentNames {
			j := EquipmentTableNameJunction{}
			j.ParentID = table.ID
			j.ChildID = name.ID
			
			if table.Classification == string(database.EquipClassCelestialWeapon) {
				var err error
				j.CelestialWeaponID, err = assignFKPtr(&name.Name, l.CelestialWeapons)
				if err != nil {
					return err
				}
			}

			dataHash := generateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.EquipmentTableID = append(params.EquipmentTableID, table.ID)
			params.EquipmentNameID = append(params.EquipmentNameID, name.ID)
			params.CelestialWeaponID = append(params.CelestialWeaponID, h.GetNullInt32(j.CelestialWeaponID))
		}
	}

	return qtx.CreateEquipmentTablesNamesJunctionBulk(ctx, params)
}

func (l *Lookup) loop6SeedAbilityPools(qtx *database.Queries, ctx context.Context) error {
	pools := l.extractAbilityPools()

	params := database.CreateAbilityPoolBulkParams{
		DataHash:         make([]string, len(pools)),
		EquipmentTableID: make([]int32, len(pools)),
		PoolIdx:          make([]int32, len(pools)),
		ReqAmount:        make([]int32, len(pools)),
	}

	for i, p := range pools {
		params.DataHash[i] = generateDataHash(p)
		params.EquipmentTableID[i] = p.EquipmentTableID
		params.PoolIdx[i] = p.PoolIdx
		params.ReqAmount[i] = p.ReqAmount
	}

	dbRows, err := qtx.CreateAbilityPoolBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability pools: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilityPools() []AbilityPool {
	pools := []AbilityPool{}

	for i := range l.json.equipment {
		table := &l.json.equipment[i]

		for j := range table.SelectableAutoAbilities {
			pool := &table.SelectableAutoAbilities[j]
			pool.EquipmentTableID = table.ID
			pool.PoolIdx = int32(j) + 1
			pools = append(pools, *pool)
		}
	}

	return dedupeRows(pools, l.Hashes)
}

func (l *Lookup) getAbilityPools() []AbilityPool {
	pools := []AbilityPool{}

	for _, table := range l.json.equipment {
		pools = append(pools, table.SelectableAutoAbilities...)
	}

	return pools
}

func (l *Lookup) getAbilityPoolAutoAbilities(ap AbilityPool) ([]AutoAbility, error) {
	return getResources(ap.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncAbilityPoolsAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "ability pools + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getAbilityPools(), l.getAbilityPoolAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAbilityPoolsAutoAbilitiesJunctionBulk(ctx, database.CreateAbilityPoolsAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AbilityPoolID: jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}
