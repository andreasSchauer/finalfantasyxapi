package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAeon(r *http.Request, i handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource], id int32) (Aeon, error) {
	aeon, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Aeon{}, err
	}

	rel, err := getAeonRelationships(r, aeon)
	if err != nil {
		return Aeon{}, err
	}

	response := Aeon{
		ID:                     aeon.ID,
		Name:                   aeon.Name,
		UntypedUnit:            idToTypedAPIResource(E.PlayerUnits, aeon.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(E.Areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        convertObjPtr(aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              getAeonBaseStats(aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        convertObjSlice(aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:         convertObjSlice(aeon.Armor, convertAeonEquipment),
	}

	response, err = applyAeonStats(r, response, nil)
	if err != nil {
		return Aeon{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)

	response.AgilityParameters, err = getUnitAgilityParams(r, response.BaseStats)
	if err != nil {
		return Aeon{}, err
	}

	return response, nil
}

func retrieveAeons(r *http.Request, i handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnOptional, config.Cfg.DB.GetAeonIDsOptional),
	})
}
