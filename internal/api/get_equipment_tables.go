package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getEquipmentTable(r *http.Request, i handlerinput.Main[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed], id int32) (apitypes.EquipmentTable, error) {
	table, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EquipmentTable{}, err
	}

	response := apitypes.EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: ah.ConvertObjSlice(table.SelectableAutoAbilities, ConvertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               ah.ConvertObjSlice(table.EquipmentNames, ConvertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := apilookups.E.CelestialWeapons.DBGetRelationPtr(r.Context(), table, config.Cfg.DB.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return apitypes.EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func retrieveEquipmentTables(r *http.Request, i handlerinput.Main[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIDList(r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, config.Cfg.DB.GetEquipmentTableIDsByAutoAbility),
		i.QueryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetEquipmentTableIDsEquipType),
		i.QueryBool2(r, ids, ts.QpnCelestialWeapon, config.Cfg.DB.GetEquipmentTableIDsCelestialWeapon),
	})
}
