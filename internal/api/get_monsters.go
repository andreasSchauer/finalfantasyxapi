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

	species, err := newNamedAPIResourceFromType(cfg, cfg.e.monsterSpecies.endpoint, monster.Species, cfg.t.MonsterSpecies)
	if err != nil {
		return Monster{}, err
	}

	ctbIconType, err := newNamedAPIResourceFromType(cfg, cfg.e.ctbIconType.endpoint, monster.CTBIconType, cfg.t.CTBIconType)
	if err != nil {
		return Monster{}, err
	}

	response := Monster{
		ID:                   monster.ID,
		Name:                 monster.Name,
		Version:              monster.Version,
		Specification:        monster.Specification,
		Notes:                monster.Notes,
		Species:              species,
		IsStoryBased:         monster.IsStoryBased,
		IsRepeatable:         monster.IsRepeatable,
		CanBeCaptured:        monster.CanBeCaptured,
		AreaConquestLocation: monster.AreaConquestLocation,
		CTBIconType:          ctbIconType,
		HasOverdrive:         monster.HasOverdrive,
		IsUnderwater:         monster.IsUnderwater,
		IsZombie:             monster.IsZombie,
		Distance:             monster.Distance,
		Properties:           rel.Properties,
		AutoAbilities:        rel.AutoAbilities,
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           rel.RonsoRages,
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            rel.BaseStats,
		Items:                convertObjPtr(cfg, monster.Items, convertMonsterItems),
		Equipment:            convertObjPtr(cfg, monster.Equipment, convertMonsterEquipment),
		ElemResists:          rel.ElemResists,
		StatusImmunities:     rel.StatusImmunities,
		StatusResists:        rel.StatusResists,
		AlteredStates:        rel.AlteredStates,
		Abilities:            rel.Abilities,
	}

	return completeMonsterResponse(cfg, r, response)
}


func (cfg *Config) retrieveMonsters(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(basicQueryWrapper(cfg, r, i, resources, "elemental_resists", getMonstersByElemResists)),
		frl(idListQueryWrapper(cfg, r, i, resources, "status_resists", len(cfg.l.StatusConditions), getMonstersByStatusResists)),
		
		frl(idQueryWrapper(cfg, r, i, resources, "item", len(cfg.l.Items), getMonstersByItem)),
		frl(idQuery(cfg, r, i, resources, "ronso_rage", len(cfg.l.RonsoRages), cfg.db.GetMonsterIDsByRonsoRage)),
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationMonsterIDs)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationMonsterIDs)),
		frl(idQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetAreaMonsterIDs)),
		frl(idQueryWrapper(cfg, r, i, resources, "auto_ability", len(cfg.l.AutoAbilities), getMonstersByAutoAbility)),

		frl(intQuery(cfg, r, i, resources, "distance", cfg.db.GetMonsterIDsByDistance)),
		frl(typeQueryWrapper(cfg, r, i, cfg.t.MonsterType, resources, "type", getMonstersByType)),
		frl(typeQuery(cfg, r, i, cfg.t.MonsterSpecies, resources, "species", cfg.db.GetMonsterIDsBySpecies)),
		frl(nullTypeQuery(cfg, r, i, cfg.t.CreationArea, resources, "creation_area", cfg.db.GetMonsterIDsByMaCreationArea)),

		frl(boolQuery(cfg, r, i, resources, "story_based", cfg.db.GetMonsterIDsByIsStoryBased)),
		frl(boolQuery(cfg, r, i, resources, "repeatable", cfg.db.GetMonsterIDsByIsRepeatable)),
		frl(boolQuery(cfg, r, i, resources, "capture", cfg.db.GetMonsterIDsByCanBeCaptured)),
		frl(boolQuery(cfg, r, i, resources, "has_overdrive", cfg.db.GetMonsterIDsByHasOverdrive)),
		frl(boolQuery(cfg, r, i, resources, "underwater", cfg.db.GetMonsterIDsByIsUnderwater)),
		frl(boolQuery(cfg, r, i, resources, "zombie", cfg.db.GetMonsterIDsByIsZombie)),
	})
}
