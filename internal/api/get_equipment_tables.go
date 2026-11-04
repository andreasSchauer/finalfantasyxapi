package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getEquipmentTable(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource], id int32) (EquipmentTable, error) {
	table, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EquipmentTable{}, err
	}

	response := EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       namePtrToNamedAPIResPtr(s, E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   namesToNamedAPIResources(s, E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: convertObjSlice(s, table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               convertObjSlice(s, table.EquipmentNames, convertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := E.CelestialWeapons.dbGetRelationPtr(s, r.Context(), table, s.Cfg.DB.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func (s *Server) retrieveEquipmentTables(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryIDList(s, r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, s.Cfg.DB.GetEquipmentTableIDsByAutoAbility),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, s.Cfg.DB.GetEquipmentTableIDsEquipType),
		i.queryBool2(r, ids, ts.QpnCelestialWeapon, s.Cfg.DB.GetEquipmentTableIDsCelestialWeapon),
	})
}
