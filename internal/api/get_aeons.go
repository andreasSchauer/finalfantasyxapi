package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAeon(r *http.Request, i handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedAPIResourceList], id int32) (Aeon, error) {
	aeon, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Aeon{}, err
	}

	rel, err := getAeonRelationships(cfg, r, aeon)
	if err != nil {
		return Aeon{}, err
	}

	response := Aeon{
		ID:                     aeon.ID,
		Name:                   aeon.Name,
		UntypedUnit:            idToTypedAPIResource(cfg, cfg.E.PlayerUnits, aeon.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(cfg, cfg.E.Areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              getAeonBaseStats(cfg, aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        convertObjSlice(cfg, aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:         convertObjSlice(cfg, aeon.Armor, convertAeonEquipment),
	}

	response, err = applyAeonStats(cfg, r, response, nil)
	if err != nil {
		return Aeon{}, err
	}

	response.RegenAmounts = calcRegenAmounts(cfg, response.BaseStats)

	response.AgilityParameters, err = getUnitAgilityParams(cfg, r, response.BaseStats)
	if err != nil {
		return Aeon{}, err
	}

	return response, nil
}

func (cfg *Config) retrieveAeons(r *http.Request, i handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		boolQuery(r, i, ids, QpnOptional, cfg.db.GetAeonIDsOptional),
	})
}
