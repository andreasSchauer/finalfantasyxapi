package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getEquipmentRelationships(cfg *Config, r *http.Request, equipment seedtypes.EquipmentName) (EquipmentName, error) {
	i := cfg.E.Equipment
	var table seedtypes.EquipmentTable

	rel := EquipmentName{
		Character: nameToNamedAPIResource(cfg, cfg.E.Characters, equipment.CharacterName, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := i.queryAvlRelGetParams(cfg, r, equipment.ID)
	if err != nil {
		return EquipmentName{}, err
	}

	g.Go(func() error {
		var err error
		rel.Treasures, err = cfg.E.Treasures.queryAvlRelRun(cfg, ctx, equipment, availabilityParams, getEquipmentSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = cfg.E.Shops.queryAvlRelRun(cfg, ctx, equipment, availabilityParams, getEquipmentSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTable, err = getEquipmentTableRef(cfg, r, ctx, i.queryLookup[typestrings.QpnTable], equipment)
		if err != nil {
			return err
		}
		table, _ = lookup.GetResourceByID(rel.EquipmentTable.ID, lookup.L.EquipmentTablesID)
		rel.Type = table.Type
		rel.Classification = table.Classification
		rel.Priority = table.Priority
		rel.RequiredSlots = table.RequiredSlots
		rel.RequiredAutoAbilities = namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, table.RequiredAutoAbilities)
		rel.SelectableAutoAbilities = convertObjSlice(cfg, table.SelectableAutoAbilities, convertAbilityPool)
		return nil
	})

	err = g.Wait()
	if err != nil {
		return EquipmentName{}, err
	}

	if table.Classification == string(database.EquipClassCelestialWeapon) {
		char, _ := lookup.GetResource(equipment.CharacterName, lookup.L.Characters)
		cwRes, err := cfg.E.CelestialWeapons.dbGetRelationPtr(cfg, r.Context(), char, cfg.db.GetCharacterCelestialWeaponID)
		if err != nil {
			return EquipmentName{}, err
		}
		rel.CelestialWeapon = cwRes
	}

	return rel, nil
}

func getEquipmentTableRef(cfg *Config, r *http.Request, ctx context.Context, queryParam QueryParam, equipment seedtypes.EquipmentName) (UnnamedAPIResource, error) {
	tables, err := cfg.E.EquipmentTables.dbGetRelations(cfg, ctx, equipment, cfg.db.GetEquipmentEquipmentTableIDs)
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
	tableIdx, err := queryParam.parseInt(r)
	if errExceptEmptyQuery(err) {
		return 0, err
	}

	if tableIdx <= 0 || tableIdx > maxID {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. '%s' can range from 1 to %d for this resource", tableIdx, queryParam.Name, queryParam.Name, maxID), nil)
	}
	tableIdx -= 1

	return tableIdx, nil
}
