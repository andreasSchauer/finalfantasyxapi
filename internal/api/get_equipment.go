package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getEquipment(r *http.Request, i handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList], id int32) (EquipmentName, error) {
	equipment, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveEquipment(r *http.Request, i handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idListQuery(cfg, r, i, ids, qpnAutoAbilities, cfg.l.AutoAbilities, cfg.db.GetEquipmentIDsByAutoAbility),
		nameIdQuery(r, i, ids, qpnCharacter, cfg.e.characters.resTypeSing, cfg.e.characters.objLookup, cfg.db.GetEquipmentIDsByCharacter),
		enumQuery(r, i, cfg.t.EquipType, ids, qpnType, cfg.db.GetEquipmentIDsByEquipType),
		boolQuery2(r, i, ids, qpnCelestialWeapon, cfg.db.GetEquipmentIDsCelestialWeapon),
	})
}
