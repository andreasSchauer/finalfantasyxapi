package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getEquipment(r *http.Request, i handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource], id int32) (EquipmentName, error) {
	equipment, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EquipmentName{}, err
	}

	rel, err := getEquipmentRelationships(srv, r, equipment)
	if err != nil {
		return EquipmentName{}, err
	}

	response := EquipmentName{
		ID:                      equipment.ID,
		Name:                    equipment.Name,
		Character:               rel.Character,
		EquipmentTable:          rel.EquipmentTable,
		Type:                    rel.Type,
		Classification:          rel.Classification,
		Priority:                rel.Priority,
		CelestialWeapon:         rel.CelestialWeapon,
		RequiredAutoAbilities:   rel.RequiredAutoAbilities,
		SelectableAutoAbilities: rel.SelectableAutoAbilities,
		RequiredSlots:           rel.RequiredSlots,
		Treasures:               rel.Treasures,
		Shops:                   rel.Shops,
	}

	return response, nil
}

func (srv *Server) retrieveEquipment(r *http.Request, i handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryIDList(srv, r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, srv.Cfg.DB.GetEquipmentIDsByAutoAbility),
		i.queryNameID(r, ids, ts.QpnCharacter, E.Characters.resTypeSingle, E.Characters.objLookup, srv.Cfg.DB.GetEquipmentIDsByCharacter),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, srv.Cfg.DB.GetEquipmentIDsByEquipType),
		i.queryBool2(r, ids, ts.QpnCelestialWeapon, srv.Cfg.DB.GetEquipmentIDsCelestialWeapon),
	})
}
