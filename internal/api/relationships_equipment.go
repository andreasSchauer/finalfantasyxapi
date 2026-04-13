package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getEquipmentRelationships(cfg *Config, r *http.Request, equipment seeding.EquipmentName) (EquipmentName, error) {
	i := cfg.e.equipment

	availabilityParams, err := getAvailabilityParams(cfg, r, i, equipment.ID)
	if err != nil {
		return EquipmentName{}, err
	}

	treasures, err := runAvailabilityQuery(cfg, r, cfg.e.treasures, equipment, availabilityParams, convGetEquipmentTreasureIDs(cfg))
	if err != nil {
		return EquipmentName{}, err
	}

	shops, err := runAvailabilityQuery(cfg, r, cfg.e.shops, equipment, availabilityParams, convGetEquipmentShopIDs(cfg))
	if err != nil {
		return EquipmentName{}, err
	}

	tableRef, err := getEquipmentTableRef(cfg, r, i.queryLookup["table"], equipment)
	if err != nil {
		return EquipmentName{}, err
	}
	table, _ := seeding.GetResourceByID(tableRef.ID, cfg.l.EquipmentTablesID)

	rel := EquipmentName{
		Character:               nameToNamedAPIResource(cfg, cfg.e.characters, equipment.CharacterName, nil),
		EquipmentTable:          tableRef,
		Type:                    table.Type,
		Classification:          table.Classification,
		Priority:                table.Priority,
		RequiredAutoAbilities:   namesToNamedAPIResources(cfg, cfg.e.autoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: convertObjSlice(cfg, table.SelectableAutoAbilities, convertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Treasures:               treasures,
		Shops:                   shops,
	}

	if table.Classification == string(database.EquipClassCelestialWeapon) {
		char, _ := seeding.GetResource(equipment.CharacterName, cfg.l.Characters)
		cwRes, err := getResPtrDB(cfg, r, cfg.e.celestialWeapons, char, cfg.db.GetCharacterCelestialWeaponID)
		if err != nil {
			return EquipmentName{}, err
		}
		rel.CelestialWeapon = cwRes
	}

	return rel, nil
}

func getEquipmentTableRef(cfg *Config, r *http.Request, queryParam QueryParam, equipment seeding.EquipmentName) (UnnamedAPIResource, error) {
	tables, err := getResourcesDbItem(cfg, r, cfg.e.equipmentTables, equipment, cfg.db.GetEquipmentEquipmentTableIDs)
	if err != nil {
		return UnnamedAPIResource{}, err
	}

	tableIdx, err := getTableIndex(r, queryParam, len(tables))
	if err != nil {
		return UnnamedAPIResource{}, err
	}
	tableRef := tables[tableIdx]

	return tableRef, nil
}

func getTableIndex(r *http.Request, queryParam QueryParam, maxID int) (int, error) {
	tableIdx, err := parseIntQuery(r, queryParam)
	if errIsNotEmptyQuery(err) {
		return 0, err
	}

	if tableIdx <= 0 || tableIdx > maxID {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. '%s' can range from 1 to %d for this resource.", tableIdx, queryParam.Name, queryParam.Name, maxID), nil)
	}
	tableIdx -= 1

	return tableIdx, nil
}
