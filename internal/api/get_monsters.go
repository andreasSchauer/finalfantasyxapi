package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMonster(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList], id int32) (Monster, error) {
	monster, err := verifyParamsAndGet(cfg, r, i, id)
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
		Species:              enumToNamedAPIResource(cfg, cfg.e.monsterSpecies.endpoint, monster.Species, cfg.t.MonsterSpecies),
		Availability:         enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, monster.Availability, cfg.t.AvailabilityType),
		IsRepeatable:         monster.IsRepeatable,
		CanBeCaptured:        monster.CanBeCaptured,
		AreaConquestLocation: monster.AreaConquestLocation,
		Category:             enumToNamedAPIResource(cfg, cfg.e.monsterCategory.endpoint, monster.Category, cfg.t.MonsterCategory),
		CTBIconType:          monster.CTBIconType,
		HasOverdrive:         monster.HasOverdrive,
		IsUnderwater:         monster.IsUnderwater,
		IsZombie:             monster.IsZombie,
		Distance:             monster.Distance,
		Properties:           namesToNamedAPIResources(cfg, cfg.e.properties, monster.Properties),
		AutoAbilities:        namesToNamedAPIResources(cfg, cfg.e.autoAbilities, monster.AutoAbilities),
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           namesToNamedAPIResources(cfg, cfg.e.ronsoRages, monster.RonsoRages),
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            toResAmtType(cfg, cfg.e.stats, monster.BaseStats, newBaseStat),
		Items:                convertObjPtr(cfg, monster.Items, convertMonsterItems),
		Equipment:            convertObjPtr(cfg, monster.Equipment, convertMonsterEquipment),
		ElemResists:          getMonsterElemResists(cfg, monster.ElemResists),
		StatusImmunities:     namesToNamedAPIResources(cfg, cfg.e.statusConditions, monster.StatusImmunities),
		StatusResists:        toResAmtType(cfg, cfg.e.statusConditions, monster.StatusResists, newStatusResist),
		Abilities:            convertObjSlice(cfg, monster.Abilities, convertMonsterAbility),
		AlteredStates:        getMonsterAlteredStates(cfg, r, monster),
	}

	return completeMonsterResponse(cfg, r, response)
}

func (cfg *Config) retrieveMonsters(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(basicQueryWrapper(cfg, r, i, ids, qpnElementalResists, getMonstersByElemResists)),
		fidl(idListQueryWrapper(cfg, r, i, ids, qpnStatusResists, cfg.l.StatusConditions, getMonstersByStatusResists)),

		fidl(idQueryWrapper(cfg, r, i, ids, qpnItem, cfg.l.Items, getMonstersByItem)),
		fidl(idQuery(r, i, ids, qpnRonsoRage, cfg.l.RonsoRages, cfg.db.GetMonsterIDsByRonsoRage)),
		fidl(idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetMonsterIDsByLocation)),
		fidl(idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetMonsterIDsBySublocation)),
		fidl(idQuery(r, i, ids, qpnArea, cfg.l.Areas, cfg.db.GetMonsterIDsByArea)),
		fidl(idQueryWrapper(cfg, r, i, ids, qpnAutoAbility, cfg.l.AutoAbilities, getMonstersByAutoAbility)),

		fidl(intListQuery(cfg, r, i, ids, qpnEmptySlots, cfg.db.GetMonsterIDsByEmptySlots)),
		fidl(intListQuery(cfg, r, i, ids, qpnDistance, cfg.db.GetMonsterIDsByDistance)),

		fidl(enumListQuery(cfg, r, i, cfg.t.MonsterCategory, ids, qpnCategory, cfg.db.GetMonsterIDsByCategory)),
		fidl(enumQuery(r, i, cfg.t.MonsterSpecies, ids, qpnSpecies, cfg.db.GetMonsterIDsBySpecies)),
		fidl(enumQuery(r, i, cfg.t.CreationArea, ids, qpnCreationArea, ToEnumQuery(cfg.t.CreationArea, cfg.db.GetMonsterIDsByMaCreationArea))),

		fidl(boolQuery(r, i, ids, qpnRepeatable, cfg.db.GetMonsterIDsByIsRepeatable)),
		fidl(boolQuery(r, i, ids, qpnCapture, cfg.db.GetMonsterIDsByCanBeCaptured)),
		fidl(boolQuery(r, i, ids, qpnHasOverdrive, cfg.db.GetMonsterIDsByHasOverdrive)),
		fidl(boolQuery(r, i, ids, qpnUnderwater, cfg.db.GetMonsterIDsByIsUnderwater)),
		fidl(boolQuery(r, i, ids, qpnZombie, cfg.db.GetMonsterIDsByIsZombie)),
	})
}
