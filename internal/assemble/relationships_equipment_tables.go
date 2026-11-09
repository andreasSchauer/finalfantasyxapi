package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetEquipmentTableRelationships(r *http.Request, table seedtypes.EquipmentTable) (apitypes.EquipmentTable, error) {
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
