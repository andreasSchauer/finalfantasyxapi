package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getEquipmentTable(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource], id int32) (EquipmentTable, error) {
	table, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EquipmentTable{}, err
	}

	response := EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       namePtrToNamedAPIResPtr(E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   namesToNamedAPIResources(E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: ConvertObjSlice(table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               ConvertObjSlice(table.EquipmentNames, convertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := E.CelestialWeapons.dbGetRelationPtr(r.Context(), table, config.Cfg.DB.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func retrieveEquipmentTables(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryIDList(r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, config.Cfg.DB.GetEquipmentTableIDsByAutoAbility),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetEquipmentTableIDsEquipType),
		i.queryBool2(r, ids, ts.QpnCelestialWeapon, config.Cfg.DB.GetEquipmentTableIDsCelestialWeapon),
	})
}
