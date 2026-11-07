package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAeon(r *http.Request, i HandlerInput[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource], id int32) (apitypes.Aeon, error) {
	aeon, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	rel, err := getAeonRelationships(r, aeon)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	response := apitypes.Aeon{
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
		PhysAtkAccuracy:        ah.ConvertObjPtr(aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              getAeonBaseStats(aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        ah.ConvertObjSlice(aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:         ah.ConvertObjSlice(aeon.Armor, convertAeonEquipment),
	}

	response, err = applyAeonStats(r, response, nil)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)

	response.AgilityParameters, err = getUnitAgilityParams(r, response.BaseStats)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	return response, nil
}

func retrieveAeons(r *http.Request, i HandlerInput[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryBool(r, ids, ts.QpnOptional, config.Cfg.DB.GetAeonIDsOptional),
	})
}
