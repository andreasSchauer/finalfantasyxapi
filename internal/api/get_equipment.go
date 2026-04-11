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
		ID:                     	equipment.ID,
		Name:                   	equipment.Name,
		Character: 					rel.Character,
		EquipmentTable: 			rel.EquipmentTable,
		Type: 						rel.Type,
		Classification: 			rel.Classification,
		Priority: 					rel.Priority,
		CelestialWeapon: 			rel.CelestialWeapon,
		RequiredAutoAbilities: 		rel.RequiredAutoAbilities,
		SelectableAutoAbilities: 	rel.SelectableAutoAbilities,
		EmptySlotsAmt: 				rel.EmptySlotsAmt,
		Treasures: 					rel.Treasures,
		Shops: 						rel.Shops,
	}

	return response, nil
}


func (cfg *Config) retrieveEquipment(r *http.Request, i handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(idListQuery(cfg, r, i, resources, "auto_abilities", len(cfg.l.AutoAbilities), cfg.db.GetEquipmentIDsByAutoAbilty)),
		frl(nameIdQuery(cfg, r, i, resources, "character", cfg.e.characters.resourceType, cfg.e.characters.objLookup, cfg.db.GetEquipmentIDsByCharacter)),
		frl(enumQuery(cfg, r, i, cfg.t.EquipType, resources, "type", cfg.db.GetEquipmentIDsByEquipType)),
		frl(boolQuery2(cfg, r, i, resources, "celestial_weapon", cfg.db.GetEquipmentIDsCelestialWeapon)),
	})
}
