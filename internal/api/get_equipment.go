package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getEquipment(r *http.Request, i handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList], id int32) (EquipmentName, error) {
	equipment, err := verifyParamsAndGet(cfg, r, i, id)
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

func (cfg *Config) retrieveEquipment(r *http.Request, i handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idListQuery(cfg, r, i, ids, "auto_abilities", cfg.l.AutoAbilities, cfg.db.GetEquipmentIDsByAutoAbility)),
		fidl(nameIdQuery(r, i, ids, "character", cfg.e.characters.resourceType, cfg.e.characters.objLookup, cfg.db.GetEquipmentIDsByCharacter)),
		fidl(enumQuery(r, i, cfg.t.EquipType, ids, "type", cfg.db.GetEquipmentIDsByEquipType)),
		fidl(boolQuery2(r, i, ids, "celestial_weapon", cfg.db.GetEquipmentIDsCelestialWeapon)),
	})
}
