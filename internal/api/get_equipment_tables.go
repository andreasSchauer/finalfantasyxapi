package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getEquipmentTable(r *http.Request, i handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList], id int32) (EquipmentTable, error) {
	table, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return EquipmentTable{}, err
	}

	response := EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       namePtrToNamedAPIResPtr(cfg, cfg.E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		RequiredAutoAbilities:   namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: convertObjSlice(cfg, table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               convertObjSlice(cfg, table.EquipmentNames, convertEquipmentName),
	}

	if table.SpecificCharacter != nil && table.Classification == string(database.EquipClassCelestialWeapon) {
		classRes, err := getResPtrDB(cfg, r.Context(), cfg.E.CelestialWeapons, table, cfg.db.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func (cfg *Config) retrieveEquipmentTables(r *http.Request, i handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idListQuery(cfg, r, i, ids, QpnAutoAbilities, cfg.L.AutoAbilities, cfg.db.GetEquipmentTableIDsByAutoAbility),
		enumQuery(r, i, cfg.T.EquipType, ids, QpnType, cfg.db.GetEquipmentTableIDsEquipType),
		boolQuery2(r, i, ids, QpnCelestialWeapon, cfg.db.GetEquipmentTableIDsCelestialWeapon),
	})
}
