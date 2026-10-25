package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMonster(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList], id int32) (Monster, error) {
	monster, err := verifyParamsAndGet(r, i, id)
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
		BaseStats:            toResAmtType(cfg, cfg.E.Stats, monster.BaseStats, newBaseStat),
		Items:                convertObjPtr(cfg, monster.Items, convertMonsterItems),
		Equipment:            convertObjPtr(cfg, monster.Equipment, convertMonsterEquipment),
		ElemResists:          getMonsterElemResists(cfg, monster.ElemResists),
		StatusImmunities:     namesToNamedAPIResources(cfg, cfg.E.StatusConditions, monster.StatusImmunities),
		StatusResists:        toResAmtType(cfg, cfg.E.StatusConditions, monster.StatusResists, newStatusResist),
		Abilities:            convertObjSlice(cfg, monster.Abilities, convertMonsterAbility),
		AlteredStates:        getMonsterAlteredStates(cfg, r, monster),
	}

	return completeMonsterResponse(cfg, r, response)
}

func (cfg *Config) retrieveMonsters(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		basicQueryWrapper(cfg, r, i, ids, QpnElementalResists, getMonstersByElemResists),
		idListQueryWrapper(cfg, r, i, ids, QpnStatusResists, cfg.L.StatusConditions, getMonstersByStatusResists),

		idQueryWrapper(cfg, r, i, ids, QpnItem, cfg.L.Items, getMonstersByItem),
		idQuery(r, i, ids, QpnRonsoRage, cfg.L.RonsoRages, cfg.db.GetMonsterIDsByRonsoRage),
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetMonsterIDsByLocation),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetMonsterIDsBySublocation),
		idQuery(r, i, ids, QpnArea, cfg.L.Areas, cfg.db.GetMonsterIDsByArea),
		idQueryWrapper(cfg, r, i, ids, QpnAutoAbility, cfg.L.AutoAbilities, getMonstersByAutoAbility),

		intListQuery(cfg, r, i, ids, QpnEmptySlots, cfg.db.GetMonsterIDsByEmptySlots),
		intListQuery(cfg, r, i, ids, QpnDistance, cfg.db.GetMonsterIDsByDistance),

		enumListQuery(cfg, r, i, cfg.T.MonsterCategory, ids, QpnCategory, cfg.db.GetMonsterIDsByCategory),
		enumQuery(r, i, cfg.T.MonsterSpecies, ids, QpnSpecies, cfg.db.GetMonsterIDsBySpecies),
		enumQuery(r, i, cfg.T.CreationArea, ids, QpnCreationArea, ToEnumQuery(cfg.T.CreationArea, cfg.db.GetMonsterIDsByMaCreationArea)),

		boolQuery(r, i, ids, QpnRepeatable, cfg.db.GetMonsterIDsByIsRepeatable),
		boolQuery(r, i, ids, QpnCapture, cfg.db.GetMonsterIDsByCanBeCaptured),
		boolQuery(r, i, ids, QpnHasOverdrive, cfg.db.GetMonsterIDsByHasOverdrive),
		boolQuery(r, i, ids, QpnUnderwater, cfg.db.GetMonsterIDsByIsUnderwater),
		boolQuery(r, i, ids, QpnZombie, cfg.db.GetMonsterIDsByIsZombie),
	})
}
