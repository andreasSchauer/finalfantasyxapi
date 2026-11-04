package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getMonster(r *http.Request, i handlerInput[seedtypes.Monster, Monster, NamedAPIResource], id int32) (Monster, error) {
	monster, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Monster{}, err
	}

	rel, err := getMonsterRelationships(cfg, r, monster)
	if err != nil {
		return Monster{}, err
	}

	response := Monster{
		ID:                   monster.ID,
		Name:                 monster.Name,
		Version:              monster.Version,
		Specification:        monster.Specification,
		Notes:                monster.Notes,
		Species:              monster.Species,
		Availability:         monster.Availability,
		IsRepeatable:         monster.IsRepeatable,
		CanBeCaptured:        monster.CanBeCaptured,
		AreaConquestLocation: monster.AreaConquestLocation,
		Category:             monster.Category,
		CTBIconType:          monster.CTBIconType,
		HasOverdrive:         monster.HasOverdrive,
		IsUnderwater:         monster.IsUnderwater,
		IsZombie:             monster.IsZombie,
		Distance:             monster.Distance,
		Properties:           namesToNamedAPIResources(cfg, cfg.E.Properties, monster.Properties),
		AutoAbilities:        namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, monster.AutoAbilities),
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           namesToNamedAPIResources(cfg, cfg.E.RonsoRages, monster.RonsoRages),
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            cfg.E.Stats.toResAmtType(cfg, monster.BaseStats, newBaseStat),
		Items:                convertObjPtr(cfg, monster.Items, convertMonsterItems),
		Equipment:            convertObjPtr(cfg, monster.Equipment, convertMonsterEquipment),
		ElemResists:          getMonsterElemResists(cfg, monster.ElemResists),
		StatusImmunities:     namesToNamedAPIResources(cfg, cfg.E.StatusConditions, monster.StatusImmunities),
		StatusResists:        cfg.E.StatusConditions.toResAmtType(cfg, monster.StatusResists, newStatusResist),
		Abilities:            convertObjSlice(cfg, monster.Abilities, convertMonsterAbility),
		AlteredStates:        getMonsterAlteredStates(cfg, r, monster),
	}

	return completeMonsterResponse(cfg, r, response)
}

func (cfg *Config) retrieveMonsters(r *http.Request, i handlerInput[seedtypes.Monster, Monster, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryBasicWrapper(cfg, r, ids, typestrings.QpnElementalResists, getMonstersByElemResists),
		i.queryIDListWrapper(cfg, r, ids, typestrings.QpnStatusResists, lookup.L.StatusConditions, getMonstersByStatusResists),

		i.queryIDWrapper(cfg, r, ids, typestrings.QpnItem, lookup.L.Items, getMonstersByItem),
		i.queryID(r, ids, typestrings.QpnRonsoRage, lookup.L.RonsoRages, cfg.db.GetMonsterIDsByRonsoRage),
		i.queryID(r, ids, typestrings.QpnLocation, lookup.L.Locations, cfg.db.GetMonsterIDsByLocation),
		i.queryID(r, ids, typestrings.QpnSublocation, lookup.L.Sublocations, cfg.db.GetMonsterIDsBySublocation),
		i.queryID(r, ids, typestrings.QpnArea, lookup.L.Areas, cfg.db.GetMonsterIDsByArea),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnAutoAbility, lookup.L.AutoAbilities, getMonstersByAutoAbility),

		i.queryIntList(cfg, r, ids, typestrings.QpnEmptySlots, cfg.db.GetMonsterIDsByEmptySlots),
		i.queryIntList(cfg, r, ids, typestrings.QpnDistance, cfg.db.GetMonsterIDsByDistance),

		i.queryEnumList(cfg, r, enums.T.MonsterCategory, ids, typestrings.QpnCategory, cfg.db.GetMonsterIDsByCategory),
		i.queryEnum(r, enums.T.MonsterSpecies, ids, typestrings.QpnSpecies, cfg.db.GetMonsterIDsBySpecies),
		i.queryEnum(r, enums.T.CreationArea, ids, typestrings.QpnCreationArea, ToEnumQuery(enums.T.CreationArea, cfg.db.GetMonsterIDsByMaCreationArea)),

		i.queryBool(r, ids, typestrings.QpnRepeatable, cfg.db.GetMonsterIDsByIsRepeatable),
		i.queryBool(r, ids, typestrings.QpnCapture, cfg.db.GetMonsterIDsByCanBeCaptured),
		i.queryBool(r, ids, typestrings.QpnHasOverdrive, cfg.db.GetMonsterIDsByHasOverdrive),
		i.queryBool(r, ids, typestrings.QpnUnderwater, cfg.db.GetMonsterIDsByIsUnderwater),
		i.queryBool(r, ids, typestrings.QpnZombie, cfg.db.GetMonsterIDsByIsZombie),
	})
}
