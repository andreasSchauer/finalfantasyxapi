package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAeon(r *http.Request, i handlerinput.Main[seedtypes.Aeon, apitypes.Aeon, apiresource.Named], id int32) (apitypes.Aeon, error) {
	aeon, err := i.VerifyQueriesGet(r, id)
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
		UntypedUnit:            handlerinput.IDToTypedAPIResource(apilookups.E.PlayerUnits, aeon.PlayerUnit.ID),
		Area:                   handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        ah.ConvertObjPtr(aeon.PhysAtkAccuracy, apiconv.ConvertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              th.GetAeonBaseStats(aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        ah.ConvertObjSlice(aeon.Weapon, apiconv.ConvertAeonEquipment),
		ArmorAbilities:         ah.ConvertObjSlice(aeon.Armor, apiconv.ConvertAeonEquipment),
	}

	return completeAeonResponse(r, response)
}

func retrieveAeons(r *http.Request, i handlerinput.Main[seedtypes.Aeon, apitypes.Aeon, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryBool(r, ids, ts.QpnOptional, config.Cfg.DB.GetAeonIDsOptional),
	})
}
