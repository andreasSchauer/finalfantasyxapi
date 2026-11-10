package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetEquipmentTable(r *http.Request, i handlerinput.Main[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed], id int32) (apitypes.EquipmentTable, error) {
	table, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EquipmentTable{}, err
	}

	rel, err := getEquipmentTableRelationships(r, table)
	if err != nil {
		return apitypes.EquipmentTable{}, err
	}

	response := apitypes.EquipmentTable{
		ID:                      table.ID,
		Type:                    table.Type,
		Classification:          table.Classification,
		SpecificCharacter:       handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Characters, table.SpecificCharacter, nil),
		Priority:                table.Priority,
		CelestialWeapon:         rel.CelestialWeapon,
		RequiredAutoAbilities:   handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, table.RequiredAutoAbilities),
		SelectableAutoAbilities: ah.ConvertObjSlice(table.SelectableAutoAbilities, apiconv.ConvertAbilityPool),
		RequiredSlots:           table.RequiredSlots,
		Equipment:               ah.ConvertObjSlice(table.EquipmentNames, apiconv.ConvertEquipmentName),
	}

	return response, nil
}

func getEquipmentTableRelationships(r *http.Request, table seedtypes.EquipmentTable) (apitypes.EquipmentTable, error) {
	var rel apitypes.EquipmentTable
	var err error

	if table.SpecificCharacter == nil || table.Classification != string(database.EquipClassCelestialWeapon) {
		return rel, nil
	}

	rel.CelestialWeapon, err = apilookups.E.CelestialWeapons.DBGetRelationPtr(r.Context(), table, config.Cfg.DB.GetEquipmentTableCelestialWeaponID)
	if err != nil {
		return apitypes.EquipmentTable{}, err
	}

	return rel, nil
}
