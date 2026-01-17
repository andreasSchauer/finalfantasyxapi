package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
	AlteredStates        []AlteredState        `json:"altered_states"`
	Abilities            []MonsterAbility      `json:"abilities"`
}

func (m *Monster) Error() string {
	msg := fmt.Sprintf("monster '%s'", m.Name)

	if m.Version != nil {
		msg += fmt.Sprintf(", version '%d'", *m.Version)
	}

	return msg
}

func getMonsterName(mon database.Monster) string {
	monster := Monster{
		Name:    mon.Name,
		Version: h.NullInt32ToPtr(mon.Version),
	}

	return monster.Error()
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
		handleEndpointSubOrNameVer(w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{name or id}', '/api/%s/{name}/{version}', or  '/api/%s/{id}/{subsection}'. available subsections: %s.", i.endpoint, i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}

func (cfg *Config) getMonster(r *http.Request, id int32) (Monster, error) {
	i := cfg.e.monsters

	err := verifyQueryParams(r, i, &id)
	if err != nil {
		return Monster{}, err
	}

	dbMonster, err := cfg.db.GetMonster(r.Context(), id)
	if err != nil {
		return Monster{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("monster with id '%d' doesn't exist.", id), err)
	}

	monsterItems, err := cfg.getMonsterItems(r, dbMonster)
	if err != nil {
		return Monster{}, err
	}

	monsterEquipment, err := cfg.getMonsterEquipment(r, dbMonster)
	if err != nil {
		return Monster{}, err
	}

	rel, err := cfg.getMonsterRelationships(r, dbMonster)
	if err != nil {
		return Monster{}, err
	}

	species, err := newNamedAPIResourceFromType(cfg, cfg.e.monsterSpecies.endpoint, string(dbMonster.Species), cfg.t.MonsterSpecies)
	if err != nil {
		return Monster{}, err
	}

	ctbIconType, err := newNamedAPIResourceFromType(cfg, cfg.e.ctbIconType.endpoint, string(dbMonster.CtbIconType), cfg.t.CTBIconType)
	if err != nil {
		return Monster{}, err
	}

	monster := Monster{
		ID:                   dbMonster.ID,
		Name:                 dbMonster.Name,
		Version:              h.NullInt32ToPtr(dbMonster.Version),
		Specification:        h.NullStringToPtr(dbMonster.Specification),
		Notes:                h.NullStringToPtr(dbMonster.Notes),
		Species:              species,
		IsStoryBased:         dbMonster.IsStoryBased,
		IsRepeatable:         dbMonster.IsRepeatable,
		CanBeCaptured:        dbMonster.CanBeCaptured,
		AreaConquestLocation: h.ConvertNullMaCreationArea(dbMonster.AreaConquestLocation),
		CTBIconType:          ctbIconType,
		HasOverdrive:         dbMonster.HasOverdrive,
		IsUnderwater:         dbMonster.IsUnderwater,
		IsZombie:             dbMonster.IsZombie,
		Distance:             anyToInt32(dbMonster.Distance),
		Properties:           rel.Properties,
		AutoAbilities:        rel.AutoAbilities,
		AP:                   dbMonster.Ap,
		APOverkill:           dbMonster.ApOverkill,
		OverkillDamage:       dbMonster.OverkillDamage,
		Gil:                  dbMonster.Gil,
		StealGil:             h.NullInt32ToPtr(dbMonster.StealGil),
		RonsoRages:           rel.RonsoRages,
		DoomCountdown:        anyToInt32Ptr(dbMonster.DoomCountdown),
		PoisonRate:           anyToFloat32Ptr(dbMonster.PoisonRate),
		ThreatenChance:       anyToInt32Ptr(dbMonster.ThreatenChance),
		ZanmatoLevel:         anyToInt32(dbMonster.ZanmatoLevel),
		MonsterArenaPrice:    h.NullInt32ToPtr(dbMonster.MonsterArenaPrice),
		SensorText:           h.NullStringToPtr(dbMonster.SensorText),
		ScanText:             h.NullStringToPtr(dbMonster.ScanText),
		Locations:            rel.Locations,
		Formations:           rel.Formations,
		BaseStats:            rel.BaseStats,
		Items:                h.NilOrPtr(monsterItems),
		Equipment:            h.NilOrPtr(monsterEquipment),
		ElemResists:          rel.ElemResists,
		StatusImmunities:     rel.StatusImmunities,
		StatusResists:        rel.StatusResists,
		AlteredStates:        rel.AlteredStates,
		Abilities:            rel.Abilities,
	}

	monster, err = cfg.applyAlteredState(r, monster)
	if err != nil {
		return Monster{}, err
	}

	monster.BaseStats, err = cfg.applyRonsoStats(r, monster)
	if err != nil {
		return Monster{}, err
	}

	monster.ElemResists, err = cfg.applyOmnisElemResists(r, monster)
	if err != nil {
		return Monster{}, err
	}

	monster.BribeChances, err = cfg.getMonsterBribeChances(monster)
	if err != nil {
		return Monster{}, err
	}

	monster.PoisonDamage, err = cfg.getMonsterPoisonDamage(monster)
	if err != nil {
		return Monster{}, err
	}

	monster.AgilityParameters, err = cfg.getMonsterAgilityVals(r, monster)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func (cfg *Config) getMultipleMonsters(r *http.Request, monsterName string) (NamedApiResourceList, error) {
	i := cfg.e.monsters
	return getMultipleAPIResources(cfg, r, i, monsterName)
}

func (cfg *Config) retrieveMonsters(r *http.Request) (NamedApiResourceList, error) {
	i := cfg.e.monsters

	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(basicQueryWrapper(cfg, r, i, resources, "elemental-resists", cfg.getMonstersByElemResists)),
		frl(idListQueryWrapper(cfg, r, i, resources, "status-resists", len(cfg.l.StatusConditions), cfg.getMonstersByStatusResists)),
		frl(idListQuery(cfg, r, i, resources, "auto-abilities", len(cfg.l.AutoAbilities), cfg.db.GetMonsterIDsByAutoAbilityIDs)),
		
		frl(idOnlyQueryWrapper(r, i, resources, "item", len(cfg.l.Items), cfg.getMonstersByItem)),
		frl(idOnlyQuery(cfg, r, i, resources, "ronso-rage", len(cfg.l.RonsoRages), cfg.db.GetMonsterIDsByRonsoRage)),
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationMonsterIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationMonsterIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetAreaMonsterIDs)),

		frl(intQuery(cfg, r, i, resources, "distance", cfg.db.GetMonsterIDsByDistance)),
		frl(typeQueryWrapper(cfg, r, i, cfg.t.CTBIconType, resources, "type", cfg.queryCTBIconType)),
		frl(typeQuery(cfg, r, i, cfg.t.MonsterSpecies, resources, "species", cfg.db.GetMonsterIDsBySpecies)),
		frl(nullTypeQuery(cfg, r, i, cfg.t.CreationArea, resources, "creation-area", cfg.db.GetMonsterIDsByMaCreationArea)),

		frl(boolQuery(cfg, r, i, resources, "story-based", cfg.db.GetMonsterIDsByIsStoryBased)),
		frl(boolQuery(cfg, r, i, resources, "repeatable", cfg.db.GetMonsterIDsByIsRepeatable)),
		frl(boolQuery(cfg, r, i, resources, "capture", cfg.db.GetMonsterIDsByCanBeCaptured)),
		frl(boolQuery(cfg, r, i, resources, "has-overdrive", cfg.db.GetMonsterIDsByHasOverdrive)),
		frl(boolQuery(cfg, r, i, resources, "underwater", cfg.db.GetMonsterIDsByIsUnderwater)),
		frl(boolQuery(cfg, r, i, resources, "zombie", cfg.db.GetMonsterIDsByIsZombie)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}
