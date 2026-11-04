package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getEquipmentTable(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource], id int32) (EquipmentTable, error) {
	table, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EquipmentTable{}, err
	}

	response := EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       namePtrToNamedAPIResPtr(srv, E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   namesToNamedAPIResources(srv, E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: convertObjSlice(srv, table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               convertObjSlice(srv, table.EquipmentNames, convertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := E.CelestialWeapons.dbGetRelationPtr(srv, r.Context(), table, srv.Cfg.DB.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func (srv *Server) retrieveEquipmentTables(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryIDList(srv, r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, srv.Cfg.DB.GetEquipmentTableIDsByAutoAbility),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, srv.Cfg.DB.GetEquipmentTableIDsEquipType),
		i.queryBool2(r, ids, ts.QpnCelestialWeapon, srv.Cfg.DB.GetEquipmentTableIDsCelestialWeapon),
	})
}
