package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getEquipmentTable(r *http.Request, i handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList], id int32) (EquipmentTable, error) {
	table, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return EquipmentTable{}, err
	}

	response := EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       namePtrToNamedAPIResPtr(cfg, cfg.e.characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   namesToNamedAPIResources(cfg, cfg.e.autoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: convertObjSlice(cfg, table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               convertObjSlice(cfg, table.EquipmentNames, convertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := getResPtrDB(cfg, r, cfg.e.celestialWeapons, table, cfg.db.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func (cfg *Config) retrieveEquipmentTables(r *http.Request, i handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idListQuery(cfg, r, i, ids, "auto_abilities", cfg.l.AutoAbilities, cfg.db.GetEquipmentTableIDsByAutoAbility)),
		fidl(enumQuery(r, i, cfg.t.EquipType, ids, "type", cfg.db.GetEquipmentTableIDsEquipType)),
		fidl(boolQuery2(r, i, ids, "celestial_weapon", cfg.db.GetEquipmentTableIDsCelestialWeapon)),
	})
}
