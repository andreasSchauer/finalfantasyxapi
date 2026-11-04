package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getEquipmentTable(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource], id int32) (EquipmentTable, error) {
	table, err := i.queryParamsVerifyGet(r, id)
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
		classRes, err := cfg.E.CelestialWeapons.dbGetRelationPtr(cfg, r.Context(), table, cfg.db.GetEquipmentTableCelestialWeaponID)
		if err != nil {
			return EquipmentTable{}, err
		}
		response.CelestialWeapon = classRes
	}

	return response, nil
}

func (cfg *Config) retrieveEquipmentTables(r *http.Request, i handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryIDList(cfg, r, ids, typestrings.QpnAutoAbilities, lookup.L.AutoAbilities, cfg.db.GetEquipmentTableIDsByAutoAbility),
		i.queryEnum(r, enums.T.EquipType, ids, typestrings.QpnType, cfg.db.GetEquipmentTableIDsEquipType),
		i.queryBool2(r, ids, typestrings.QpnCelestialWeapon, cfg.db.GetEquipmentTableIDsCelestialWeapon),
	})
}
