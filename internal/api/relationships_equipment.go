package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getEquipmentRelationships(cfg *Config, r *http.Request, equipment seeding.EquipmentName) (EquipmentName, error) {
	i := cfg.e.equipment
	var table seeding.EquipmentTable
	
	rel := EquipmentName{
		Character: nameToNamedAPIResource(cfg, cfg.e.characters, equipment.CharacterName, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := getRelAvailabilityParams(cfg, r, i, equipment.ID)
	if err != nil {
		return EquipmentName{}, err
	}

	g.Go(func() error{
		var err error
		rel.Treasures, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, equipment, availabilityParams, getEquipmentSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Shops, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, equipment, availabilityParams, getEquipmentSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.EquipmentTable, err = getEquipmentTableRef(cfg, r, ctx, i.queryLookup[qpnTable], equipment)
		if err != nil {
			return err
		}
		table, _ = seeding.GetResourceByID(rel.EquipmentTable.ID, cfg.l.EquipmentTablesID)
		rel.Type = table.Type
		rel.Classification = table.Classification
		rel.Priority = table.Priority
		rel.RequiredSlots = table.RequiredSlots
		rel.RequiredAutoAbilities = namesToNamedAPIResources(cfg, cfg.e.autoAbilities, table.RequiredAutoAbilities)
		rel.SelectableAutoAbilities = convertObjSlice(cfg, table.SelectableAutoAbilities, convertAbilityPool)
		return nil
	})

	err = g.Wait()
	if err != nil {
		return EquipmentName{}, err
	}

	if table.Classification == string(database.EquipClassCelestialWeapon) {
		char, _ := seeding.GetResource(equipment.CharacterName, cfg.l.Characters)
		cwRes, err := getResPtrDB(cfg, r.Context(), cfg.e.celestialWeapons, char, cfg.db.GetCharacterCelestialWeaponID)
		if err != nil {
			return EquipmentName{}, err
		}
		rel.CelestialWeapon = cwRes
	}

	return rel, nil
}

func getEquipmentTableRef(cfg *Config, r *http.Request, ctx context.Context, queryParam QueryParam, equipment seeding.EquipmentName) (UnnamedAPIResource, error) {
	tables, err := getResourcesDbItem(cfg, ctx, cfg.e.equipmentTables, equipment, cfg.db.GetEquipmentEquipmentTableIDs)
	if err != nil {
		return UnnamedAPIResource{}, err
	}

	tableIdx, err := getTableIndex(r, queryParam, len(tables))
	if err != nil {
		return UnnamedAPIResource{}, err
	}
	tableRef := tables[tableIdx]

	return tableRef, nil
}

func getTableIndex(r *http.Request, queryParam QueryParam, maxID int) (int, error) {
	tableIdx, err := parseIntQuery(r, queryParam)
	if errExceptEmptyQuery(err) {
		return 0, err
	}

	if tableIdx <= 0 || tableIdx > maxID {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. '%s' can range from 1 to %d for this resource.", tableIdx, queryParam.Name, queryParam.Name, maxID), nil)
	}
	tableIdx -= 1

	return tableIdx, nil
}
