package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getCelestialWeapon(r *http.Request, i handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedApiResourceList], id int32) (CelestialWeapon, error) {
	celestialWeapon, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return CelestialWeapon{}, err
	}

	crest, _ := seeding.GetResource(celestialWeapon.KeyItemBase + " crest", cfg.l.KeyItems)
	sigil, _ := seeding.GetResource(celestialWeapon.KeyItemBase + " sigil", cfg.l.KeyItems)
	equipment, _ := seeding.GetResource(celestialWeapon.Name, cfg.l.EquipmentNames)

	tables, err := getResourcesDbItem(cfg, r, cfg.e.equipmentTables, equipment, cfg.db.GetEquipmentEquipmentTableIDs)
	if err != nil {
		return CelestialWeapon{}, err
	}
	table, _ := seeding.GetResourceByID(tables[0].ID, cfg.l.EquipmentTablesID)

	wpnTreasure, err := getResDbItemOne(cfg, r, cfg.e.treasures, celestialWeapon, cfg.db.GetCelestialWeaponTreasureID)
	if err != nil {
		return CelestialWeapon{}, err
	}

	crestTreasures, err := getResourcesDbItem(cfg, r, cfg.e.treasures, crest, cfg.db.GetKeyItemTreasureIDs)
	if err != nil {
		return CelestialWeapon{}, err
	}

	sigilQuests, err := getResourcesDbItem(cfg, r, cfg.e.quests, sigil, cfg.db.GetKeyItemQuestIDs)
	if err != nil {
		return CelestialWeapon{}, err
	}


	response := CelestialWeapon{
		ID:             celestialWeapon.ID,
		Name:           celestialWeapon.Name,
		Formula: 		celestialWeapon.Formula,
		Character: 		nameToNamedAPIResource(cfg, cfg.e.characters, celestialWeapon.Character, nil),
		Aeon: 			namePtrToNamedAPIResPtr(cfg, cfg.e.aeons, celestialWeapon.Aeon, nil),
		Equipment: 		nameToNamedAPIResource(cfg, cfg.e.equipment, equipment.Name, nil),
		AutoAbilities: 	namesToNamedAPIResources(cfg, cfg.e.autoAbilities, table.RequiredAutoAbilities),
		Crest: 			nameToNamedAPIResource(cfg, cfg.e.keyItems, crest.Name, nil),
		Sigil: 			nameToNamedAPIResource(cfg, cfg.e.keyItems, sigil.Name, nil),
		WpnTreasure: 	wpnTreasure,
		CrestTreasure: 	crestTreasures[0],
		SigilQuest: 	sigilQuests[0],
	}

	return response, nil
}



func (cfg *Config) retrieveCelestialWeapons(r *http.Request, i handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumQuery(cfg, r, i, cfg.t.CelestialFormula, resources, "formula", cfg.db.GetCelestialWeaponIDsByFormula)),
	})
}