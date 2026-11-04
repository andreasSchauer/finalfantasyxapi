package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getEquipmentRelationships(server *Server, r *http.Request, equipment seedtypes.EquipmentName) (EquipmentName, error) {
	i := E.Equipment
	var table seedtypes.EquipmentTable

	rel := EquipmentName{
		Character: nameToNamedAPIResource(server, E.Characters, equipment.CharacterName, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := i.queryAvlRelGetParams(server, r, equipment.ID)
	if err != nil {
		return EquipmentName{}, err
	}

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.queryAvlRelRun(server, ctx, equipment, availabilityParams, getEquipmentSourceIDs(server, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = E.Shops.queryAvlRelRun(server, ctx, equipment, availabilityParams, getEquipmentSourceIDs(server, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTable, err = getEquipmentTableRef(server, r, ctx, i.queryLookup[ts.QpnTable], equipment)
		if err != nil {
			return err
		}
		table, _ = lookup.GetResourceByID(rel.EquipmentTable.ID, lookup.L.EquipmentTablesID)
		rel.Type = table.Type
		rel.Classification = table.Classification
		rel.Priority = table.Priority
		rel.RequiredSlots = table.RequiredSlots
		rel.RequiredAutoAbilities = namesToNamedAPIResources(server, E.AutoAbilities, table.RequiredAutoAbilities)
		rel.SelectableAutoAbilities = convertObjSlice(server, table.SelectableAutoAbilities, convertAbilityPool)
		return nil
	})

	err = g.Wait()
	if err != nil {
		return EquipmentName{}, err
	}

	if table.Classification == string(database.EquipClassCelestialWeapon) {
		char, _ := lookup.GetResource(equipment.CharacterName, lookup.L.Characters)
		cwRes, err := E.CelestialWeapons.dbGetRelationPtr(server, r.Context(), char, server.Cfg.DB.GetCharacterCelestialWeaponID)
		if err != nil {
			return EquipmentName{}, err
		}
		rel.CelestialWeapon = cwRes
	}

	return rel, nil
}

func getEquipmentTableRef(server *Server, r *http.Request, ctx context.Context, queryParam QueryParam, equipment seedtypes.EquipmentName) (UnnamedAPIResource, error) {
	tables, err := E.EquipmentTables.dbGetRelations(server, ctx, equipment, server.Cfg.DB.GetEquipmentEquipmentTableIDs)
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
