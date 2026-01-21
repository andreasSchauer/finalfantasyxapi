package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Monster struct {
	ID                   int32                 `json:"id"`
	Name                 string                `json:"name"`
	Version              *int32                `json:"version,omitempty"`
	Specification        *string               `json:"specification,omitempty"`
	AppliedState         *AppliedState         `json:"applied_state,omitempty"`
	AgilityParameters    *AgilityParams        `json:"agility_parameters"`
	Notes                *string               `json:"notes,omitempty"`
	Species              NamedAPIResource      `json:"species"`
	IsStoryBased         bool                  `json:"is_story_based"`
	IsRepeatable         bool                  `json:"is_repeatable"`
	CanBeCaptured        bool                  `json:"can_be_captured"`
	AreaConquestLocation *string               `json:"area_conquest_location,omitempty"`
	CTBIconType          NamedAPIResource      `json:"ctb_icon_type"`
	HasOverdrive         bool                  `json:"has_overdrive"`
	IsUnderwater         bool                  `json:"is_underwater"`
	IsZombie             bool                  `json:"is_zombie"`
	Distance             int32                 `json:"distance"`
	Properties           []NamedAPIResource    `json:"properties"`
	AutoAbilities        []NamedAPIResource    `json:"auto_abilities"`
	AP                   int32                 `json:"ap"`
	APOverkill           int32                 `json:"ap_overkill"`
	OverkillDamage       int32                 `json:"overkill_damage"`
	Gil                  int32                 `json:"gil"`
	StealGil             *int32                `json:"steal_gil"`
	RonsoRages           []NamedAPIResource    `json:"ronso_rages"`
	DoomCountdown        *int32                `json:"doom_countdown"`
	PoisonRate           *float32              `json:"poison_rate"`
	PoisonDamage         *int32                `json:"poison_damage,omitempty"`
	ThreatenChance       *int32                `json:"threaten_chance"`
	ZanmatoLevel         int32                 `json:"zanmato_level"`
	MonsterArenaPrice    *int32                `json:"monster_arena_price,omitempty"`
	SensorText           *string               `json:"sensor_text"`
	ScanText             *string               `json:"scan_text"`
	Locations            []LocationAPIResource `json:"locations"`
	Formations           []UnnamedAPIResource  `json:"formations"`
	BaseStats            []BaseStat            `json:"base_stats"`
	Items                *MonsterItems         `json:"items"`
	BribeChances         []BribeChance         `json:"bribe_chances,omitempty"`
	Equipment            *MonsterEquipment     `json:"equipment"`
	ElemResists          []ElementalResist     `json:"elem_resists"`
	StatusImmunities     []NamedAPIResource    `json:"status_immunities"`
	StatusResists        []StatusResist        `json:"status_resists"`
	Abilities            []MonsterAbility      `json:"abilities"`
	AlteredStates        []AlteredState        `json:"altered_states"`
}

func (m *Monster) Error() string {
	msg := fmt.Sprintf("monster '%s'", m.Name)

	if m.Version != nil {
		msg += fmt.Sprintf(", version '%d'", *m.Version)
	}

	return msg
}

func (cfg *Config) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.monsters

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointNameOrID(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubOrNameVer(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{name or id}', '/api/%s/{name}/{version}', or  '/api/%s/{id}/{subsection}'. available subsections: %s.", i.endpoint, i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}

func (cfg *Config) getMonster(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList], id int32) (Monster, error) {
	err := verifyQueryParams(r, i, &id)
	if err != nil {
		return Monster{}, err
	}

	monster, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Monster{}, err
	}

	rel, err := cfg.getMonsterRelationships(r, monster)
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
		Locations:            rel.Locations,
		Formations:           rel.Formations,
		BaseStats:            rel.BaseStats,
		Items:                cfg.getMonsterItems(monster.Items),
		Equipment:            cfg.getMonsterEquipment(monster.Equipment),
		ElemResists:          rel.ElemResists,
		StatusImmunities:     rel.StatusImmunities,
		StatusResists:        rel.StatusResists,
		AlteredStates:        rel.AlteredStates,
		Abilities:            rel.Abilities,
	}

	response, err = cfg.applyAlteredState(r, response, "altered_state")
	if err != nil {
		return Monster{}, err
	}

	response.BaseStats, err = cfg.applyAeonStats(r, response, "aeon_stats")
	if err != nil {
		return Monster{}, err
	}

	response.BaseStats, err = cfg.applyRonsoStats(r, response, "kimahri_stats")
	if err != nil {
		return Monster{}, err
	}

	response.ElemResists, err = cfg.applyOmnisElements(r, response, "omnis_elements")
	if err != nil {
		return Monster{}, err
	}

	response.BribeChances, err = cfg.getMonsterBribeChances(response)
	if err != nil {
		return Monster{}, err
	}

	response.PoisonDamage, err = cfg.getMonsterPoisonDamage(response)
	if err != nil {
		return Monster{}, err
	}

	response.AgilityParameters, err = cfg.getMonsterAgilityParams(r, response)
	if err != nil {
		return Monster{}, err
	}

	return response, nil
}

func (cfg *Config) retrieveMonsters(r *http.Request, i handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(basicQueryWrapper(cfg, r, i, resources, "elemental_resists", cfg.getMonstersByElemResists)),
		frl(idListQueryWrapper(cfg, r, i, resources, "status_resists", len(cfg.l.StatusConditions), cfg.getMonstersByStatusResists)),
		frl(idListQuery(cfg, r, i, resources, "auto_abilities", len(cfg.l.AutoAbilities), cfg.db.GetMonsterIDsByAutoAbilityIDs)),

		frl(idOnlyQueryWrapper(r, i, resources, "item", len(cfg.l.Items), cfg.getMonstersByItem)),
		frl(idOnlyQuery(cfg, r, i, resources, "ronso_rage", len(cfg.l.RonsoRages), cfg.db.GetMonsterIDsByRonsoRage)),
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationMonsterIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationMonsterIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetAreaMonsterIDs)),

		frl(intQuery(cfg, r, i, resources, "distance", cfg.db.GetMonsterIDsByDistance)),
		frl(typeQueryWrapper(cfg, r, i, cfg.t.CTBIconType, resources, "type", cfg.queryCTBIconType)),
		frl(typeQuery(cfg, r, i, cfg.t.MonsterSpecies, resources, "species", cfg.db.GetMonsterIDsBySpecies)),
		frl(nullTypeQuery(cfg, r, i, cfg.t.CreationArea, resources, "creation_area", cfg.db.GetMonsterIDsByMaCreationArea)),

		frl(boolQuery(cfg, r, i, resources, "story_based", cfg.db.GetMonsterIDsByIsStoryBased)),
		frl(boolQuery(cfg, r, i, resources, "repeatable", cfg.db.GetMonsterIDsByIsRepeatable)),
		frl(boolQuery(cfg, r, i, resources, "capture", cfg.db.GetMonsterIDsByCanBeCaptured)),
		frl(boolQuery(cfg, r, i, resources, "has_overdrive", cfg.db.GetMonsterIDsByHasOverdrive)),
		frl(boolQuery(cfg, r, i, resources, "underwater", cfg.db.GetMonsterIDsByIsUnderwater)),
		frl(boolQuery(cfg, r, i, resources, "zombie", cfg.db.GetMonsterIDsByIsZombie)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}
