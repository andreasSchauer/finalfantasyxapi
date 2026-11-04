package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getAeon(r *http.Request, i handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource], id int32) (Aeon, error) {
	aeon, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Aeon{}, err
	}

	rel, err := getAeonRelationships(srv, r, aeon)
	if err != nil {
		return Aeon{}, err
	}

	response := Aeon{
		ID:                     aeon.ID,
		Name:                   aeon.Name,
		UntypedUnit:            idToTypedAPIResource(srv, E.PlayerUnits, aeon.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(srv, E.Areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        convertObjPtr(srv, aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              getAeonBaseStats(srv, aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        convertObjSlice(srv, aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:         convertObjSlice(srv, aeon.Armor, convertAeonEquipment),
	}

	response, err = applyAeonStats(srv, r, response, nil)
	if err != nil {
		return Aeon{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)

	response.AgilityParameters, err = getUnitAgilityParams(srv, r, response.BaseStats)
	if err != nil {
		return Aeon{}, err
	}

	return response, nil
}

func (srv *Server) retrieveAeons(r *http.Request, i handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnOptional, srv.Cfg.DB.GetAeonIDsOptional),
	})
}
