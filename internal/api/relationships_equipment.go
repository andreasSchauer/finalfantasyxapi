package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getEquipmentRelationships(r *http.Request, equipment seedtypes.EquipmentName) (apitypes.EquipmentName, error) {
	i := apilookups.E.Equipment
	var table seedtypes.EquipmentTable

	rel := apitypes.EquipmentName{
		Character: handlerinput.NameToNamedAPIResource(apilookups.E.Characters, equipment.CharacterName, nil),
	}
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := i.QueryAvlRelGetParams(r, equipment.ID)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	g.Go(func() error {
		var err error
		rel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, equipment, availabilityParams, getEquipmentSourceIDs(ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, equipment, availabilityParams, getEquipmentSourceIDs(ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTable, err = getEquipmentTableRef(r, ctx, i.QueryLookup[ts.QpnTable], equipment)
		if err != nil {
			return err
		}
		table, _ = lookup.GetResourceByID(rel.EquipmentTable.ID, lookup.L.EquipmentTablesID)
		rel.Type = table.Type
		rel.Classification = table.Classification
		rel.Priority = table.Priority
		rel.RequiredSlots = table.RequiredSlots
		rel.RequiredAutoAbilities = handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, table.RequiredAutoAbilities)
		rel.SelectableAutoAbilities = ah.ConvertObjSlice(table.SelectableAutoAbilities, apiconv.ConvertAbilityPool)
		return nil
	})

	err = g.Wait()
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	if table.Classification == string(database.EquipClassCelestialWeapon) {
		char, _ := lookup.GetResource(equipment.CharacterName, lookup.L.Characters)
		cwRes, err := apilookups.E.CelestialWeapons.DBGetRelationPtr(r.Context(), char, config.Cfg.DB.GetCharacterCelestialWeaponID)
		if err != nil {
			return apitypes.EquipmentName{}, err
		}
		rel.CelestialWeapon = cwRes
	}

	return rel, nil
}

func getEquipmentTableRef(r *http.Request, ctx context.Context, queryParam queryparam.QueryParam, equipment seedtypes.EquipmentName) (apiresource.Unnamed, error) {
	tables, err := apilookups.E.EquipmentTables.DBGetRelations(ctx, equipment, config.Cfg.DB.GetEquipmentEquipmentTableIDs)
	if err != nil {
		return apiresource.Unnamed{}, err
	}

	tableIdx, err := getTableIndex(r, queryParam, len(tables))
	if err != nil {
		return apiresource.Unnamed{}, err
	}
	tableRef := tables[tableIdx]

	return tableRef, nil
}

func getTableIndex(r *http.Request, queryParam queryparam.QueryParam, maxID int) (int, error) {
	tableIdx, err := queryParam.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return 0, err
	}

	if tableIdx <= 0 || tableIdx > maxID {
		return 0, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. '%s' can range from 1 to %d for this resource", tableIdx, queryParam.Name, queryParam.Name, maxID), nil)
	}
	tableIdx -= 1

	return tableIdx, nil
}
