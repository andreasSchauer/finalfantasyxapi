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
		ID:                  aeon.ID,
		Name:                aeon.Name,
		Area:                locAreaToAreaAPIResource(cfg, cfg.e.areas, aeon.LocationArea),
		UnlockCondition:     aeon.UnlockCondition,
		IsOptional:          aeon.IsOptional,
		BattlesToRegenerate: aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:  aeon.PhysAtkDmgConstant,
		PhysAtkRange:        aeon.PhysAtkRange,
		PhysAtkShatterRate:  aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:     convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy),
		CelestialWeapon:     rel.CelestialWeapon,
		CharacterClasses:    rel.CharacterClasses,
		BaseStats:           namesToResourceAmounts(cfg, cfg.e.stats, baseStats, newBaseStat),
		AeonCommands:        rel.AeonCommands,
		DefaultAbilities:    rel.DefaultAbilities,
		Overdrives:          rel.Overdrives,
		WeaponAbilities:     convertObjSlice(cfg, aeon.Weapon, convertAeonEquipment),
		ArmorAbilities:      convertObjSlice(cfg, aeon.Armor, convertAeonEquipment),
	}

	response.BaseStats, err = applyAeonStatsBattles(cfg, r, response, "battles")
	if err != nil {
		return Aeon{}, err
	}

	return response, nil
}

func (cfg *Config) retrieveAeons(r *http.Request, i handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(boolQuery(cfg, r, i, resources, "optional", cfg.db.GetAeonIDsOptional)),
	})
}
