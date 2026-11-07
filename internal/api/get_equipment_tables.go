package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getEquipmentTable(r *http.Request, i HandlerInput[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource], id int32) (apitypes.EquipmentTable, error) {
	table, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.EquipmentTable{}, err
	}

	response := apitypes.EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       NamePtrToNamedAPIResPtr(E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   NamesToNamedAPIResources(E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: ah.ConvertObjSlice(table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               ah.ConvertObjSlice(table.EquipmentNames, convertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := E.CelestialWeapons.DBGetRelationPtr(r.Context(), table, config.Cfg.DB.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return apitypes.EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func retrieveEquipmentTables(r *http.Request, i HandlerInput[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryIDList(r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, config.Cfg.DB.GetEquipmentTableIDsByAutoAbility),
		i.QueryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetEquipmentTableIDsEquipType),
		i.QueryBool2(r, ids, ts.QpnCelestialWeapon, config.Cfg.DB.GetEquipmentTableIDsCelestialWeapon),
	})
}
