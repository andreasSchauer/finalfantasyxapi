package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
		l.EquipmentTables[Key(tables[i])] = tables[i]
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
