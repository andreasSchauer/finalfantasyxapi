package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getEquipment(r *http.Request, i handlerInput[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource], id int32) (apitypes.EquipmentName, error) {
	equipment, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	rel, err := getEquipmentRelationships(r, equipment)
	if err != nil {
		return apitypes.EquipmentName{}, err
	}

	response := apitypes.EquipmentName{
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

func retrieveEquipment(r *http.Request, i handlerInput[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryIDList(r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, config.Cfg.DB.GetEquipmentIDsByAutoAbility),
		i.queryNameID(r, ids, ts.QpnCharacter, E.Characters.resTypeSingle, E.Characters.objLookup, config.Cfg.DB.GetEquipmentIDsByCharacter),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetEquipmentIDsByEquipType),
		i.queryBool2(r, ids, ts.QpnCelestialWeapon, config.Cfg.DB.GetEquipmentIDsCelestialWeapon),
	})
}
