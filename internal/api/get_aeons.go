package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getAeon(r *http.Request, i handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource], id int32) (Aeon, error) {
	aeon, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Aeon{}, err
	}

	rel, err := getAeonRelationships(s, r, aeon)
	if err != nil {
		return Aeon{}, err
	}

	response := Aeon{
		ID:                     aeon.ID,
		Name:                   aeon.Name,
		UntypedUnit:            idToTypedAPIResource(s, E.PlayerUnits, aeon.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(s, E.Areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        convertObjPtr(s, aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              getAeonBaseStats(s, aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        convertObjSlice(s, aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:         convertObjSlice(s, aeon.Armor, convertAeonEquipment),
	}

	response, err = applyAeonStats(s, r, response, nil)
	if err != nil {
		return Aeon{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)

	response.AgilityParameters, err = getUnitAgilityParams(s, r, response.BaseStats)
	if err != nil {
		return Aeon{}, err
	}

	return response, nil
}

func (s *Server) retrieveAeons(r *http.Request, i handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnOptional, s.Cfg.DB.GetAeonIDsOptional),
	})
}
