package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMonster(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource], id int32) (Monster, error) {
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

func (cfg *Config) retrieveMonsters(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryBasicWrapper(cfg, r, ids, QpnElementalResists, getMonstersByElemResists),
		i.queryIDListWrapper(cfg, r, ids, QpnStatusResists, cfg.L.StatusConditions, getMonstersByStatusResists),

		i.queryIDWrapper(cfg, r, ids, QpnItem, cfg.L.Items, getMonstersByItem),
		i.queryID(r, ids, QpnRonsoRage, cfg.L.RonsoRages, cfg.db.GetMonsterIDsByRonsoRage),
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetMonsterIDsByLocation),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetMonsterIDsBySublocation),
		i.queryID(r, ids, QpnArea, cfg.L.Areas, cfg.db.GetMonsterIDsByArea),
		i.queryIDWrapper(cfg, r, ids, QpnAutoAbility, cfg.L.AutoAbilities, getMonstersByAutoAbility),

		i.queryIntList(cfg, r, ids, QpnEmptySlots, cfg.db.GetMonsterIDsByEmptySlots),
		i.queryIntList(cfg, r, ids, QpnDistance, cfg.db.GetMonsterIDsByDistance),

		i.queryEnumList(cfg, r, cfg.T.MonsterCategory, ids, QpnCategory, cfg.db.GetMonsterIDsByCategory),
		i.queryEnum(r, cfg.T.MonsterSpecies, ids, QpnSpecies, cfg.db.GetMonsterIDsBySpecies),
		i.queryEnum(r, cfg.T.CreationArea, ids, QpnCreationArea, ToEnumQuery(cfg.T.CreationArea, cfg.db.GetMonsterIDsByMaCreationArea)),

		i.queryBool(r, ids, QpnRepeatable, cfg.db.GetMonsterIDsByIsRepeatable),
		i.queryBool(r, ids, QpnCapture, cfg.db.GetMonsterIDsByCanBeCaptured),
		i.queryBool(r, ids, QpnHasOverdrive, cfg.db.GetMonsterIDsByHasOverdrive),
		i.queryBool(r, ids, QpnUnderwater, cfg.db.GetMonsterIDsByIsUnderwater),
		i.queryBool(r, ids, QpnZombie, cfg.db.GetMonsterIDsByIsZombie),
	})
}
