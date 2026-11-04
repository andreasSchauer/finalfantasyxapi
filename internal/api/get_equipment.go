package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getEquipment(r *http.Request, i handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource], id int32) (EquipmentName, error) {
	equipment, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EquipmentName{}, err
	}

	rel, err := getEquipmentRelationships(cfg, r, equipment)
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

func (cfg *Config) retrieveEquipment(r *http.Request, i handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryIDList(cfg, r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, cfg.db.GetEquipmentIDsByAutoAbility),
		i.queryNameID(r, ids, ts.QpnCharacter, cfg.E.Characters.resTypeSingle, cfg.E.Characters.objLookup, cfg.db.GetEquipmentIDsByCharacter),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, cfg.db.GetEquipmentIDsByEquipType),
		i.queryBool2(r, ids, ts.QpnCelestialWeapon, cfg.db.GetEquipmentIDsCelestialWeapon),
	})
}
