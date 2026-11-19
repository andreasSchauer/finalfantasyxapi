package get

import (
	"context"
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func Equipment(r *http.Request, i handlerinput.Main[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.Named], id int32) (apitypes.EquipmentName, error) {
	equipment, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	response := apitypes.EquipmentName{
		ID:                      equipment.ID,
		Name:                    equipment.Name,
	}

	err = getEquipmentRelationships(r, equipment, &response)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}
	
	return response, nil
}

func getEquipmentRelationships(r *http.Request, equipment seedtypes.EquipmentName, response *apitypes.EquipmentName) error {
	i := apilookups.E.Equipment
	var table seedtypes.EquipmentTable

	response.Character = handlerinput.NameToNamedAPIResource(apilookups.E.Characters, equipment.CharacterName, nil)

	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := i.QueryAvlRelGetParams(r, equipment.ID)
	if err != nil {
		return err
	}

	g.Go(func() error {
		var err error
		response.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, equipment, availabilityParams, getEquipmentSourceIDs(ts.VstTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		response.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, equipment, availabilityParams, getEquipmentSourceIDs(ts.VstShop))
		return err
	})

	g.Go(func() error {
		var err error
		response.EquipmentTable, err = getEquipmentTableRef(r, ctx, i.QueryLookup[ts.QpnTable], equipment)
		if err != nil {
			return err
		}
		table, _ = lookup.GetResourceByID(response.EquipmentTable.ID, lookup.L.EquipmentTablesID)
		response.Type = table.Type
		response.Classification = table.Classification
		response.Priority = table.Priority
		response.RequiredSlots = table.RequiredSlots
		response.RequiredAutoAbilities = handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, table.RequiredAutoAbilities)
		response.SelectableAutoAbilities = ah.ConvertObjSlice(table.SelectableAutoAbilities, apiconvert.AbilityPool)
		return nil
	})

	err = g.Wait()
	if err != nil {
		return err
	}

	if table.Classification == string(database.EquipClassCelestialWeapon) {
		char, _ := lookup.GetResource(equipment.CharacterName, lookup.L.Characters)
		cwRes, err := apilookups.E.CelestialWeapons.DBGetRelationPtr(r.Context(), char, config.Cfg.DB.GetCharacterCelestialWeaponID)
		if err != nil {
			return err
		}
		response.CelestialWeapon = cwRes
	}

	return nil
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
