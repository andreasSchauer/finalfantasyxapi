package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAeon(r *http.Request, i handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList], id int32) (Aeon, error) {
	aeon, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Aeon{}, err
	}

	rel, err := getAeonRelationships(cfg, r, aeon)
	if err != nil {
		return Aeon{}, err
	}

	baseStats := aeon.BaseStats.XVals[0].BaseStats

	response := Aeon{
		ID:                     aeon.ID,
		Name:                   aeon.Name,
		UntypedUnit:            idToTypedAPIResource(cfg, cfg.e.playerUnits, aeon.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(cfg, cfg.e.areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              toResAmtType(cfg, cfg.e.stats, baseStats, newBaseStat),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        convertObjSlice(cfg, aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:         convertObjSlice(cfg, aeon.Armor, convertAeonEquipment),
	}

	response, err = applyAeonStats(cfg, r, response)
	if err != nil {
		return Aeon{}, err
	}

	response.Stats = createStats(response.BaseStats)

	return response, nil
}

func (cfg *Config) retrieveAeons(r *http.Request, i handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		boolQuery(r, i, ids, qpnOptional, cfg.db.GetAeonIDsOptional),
	})
}
