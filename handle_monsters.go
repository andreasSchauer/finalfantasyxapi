package main

import (
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type Monster struct {
	ID                   int32				`json:"id"`
	Name                 string            	`json:"name"`
	Version              *int32           	`json:"version,omitempty"`
	Specification        *string          	`json:"specification,omitempty"`
	Notes                *string          	`json:"notes,omitempty"`
	Species              string           	`json:"species"`
	IsStoryBased         bool             	`json:"is_story_based"`
	CanBeCaptured        bool              	`json:"can_be_captured"`
	AreaConquestLocation *string           	`json:"area_conquest_location,omitempty"`
	CTBIconType          string            	`json:"ctb_icon_type"`
	HasOverdrive         bool              	`json:"has_overdrive"`
	IsUnderwater         bool              	`json:"is_underwater"`
	IsZombie             bool              	`json:"is_zombie"`
	Distance             int32             	`json:"distance"`
	Properties           []NamedAPIResource `json:"properties"`
	AutoAbilities        []NamedAPIResource `json:"auto_abilities"`
	AP                   int32             	`json:"ap"`
	APOverkill           int32             	`json:"ap_overkill"`
	OverkillDamage       int32             	`json:"overkill_damage"`
	Gil                  int32             	`json:"gil"`
	StealGil             *int32            	`json:"steal_gil"`
	RonsoRages           []NamedAPIResource `json:"ronso_rages"`
	DoomCountdown        *int32            	`json:"doom_countdown"`
	PoisonRate           *float32          	`json:"poison_rate"`
	ThreatenChance       *int32            	`json:"threaten_chance"`
	ZanmatoLevel         int32             	`json:"zanmato_level"`
	MonsterArenaPrice    *int32            	`json:"monster_arena_price,omitempty"`
	SensorText           string            	`json:"sensor_text"`
	ScanText             *string           	`json:"scan_text"`
	BaseStats            []BaseStat        	`json:"base_stats"`
	//Items                *MonsterItems     `json:"items"`
	//Equipment            *MonsterEquipment `json:"equipment"`
	//ElemResists          []ElementalResist `json:"elem_resists"`
	StatusImmunities     []NamedAPIResource `json:"status_immunities"`
	//StatusResists        []StatusResist    `json:"status_resists"`
	//AlteredStates        []AlteredState    `json:"altered_states"`
	//Abilities            []MonsterAbility  `json:"abilities"`
}


type BaseStat struct {
	Stat	NamedAPIResource	`json:"stat"`
	Value	int32
}


type MonsterItems struct{

}


func (cfg *apiConfig) handleMonsters(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/monsters/")
	segments := strings.Split(path, "/")

	if path == "" {
		cfg.handleMonstersRetrieve(w, r)
		return
	}

	switch len(segments) {
	case 1:
		// /api/monsters/{name or id}
		segment := segments[0]

		input, err := parseSingleSegmentResource(segment, cfg.l.Monsters)
		if handleHTTPError(w, err) {
			return
		}

		cfg.handleMonsterGet(w, r, input)
		return
	case 2:
		// /api/monsters/{name}/{version}

		name := segments[0]
		versionStr := segments[1]

		input, err := parseNameVersionResource(name, versionStr, cfg.l.Monsters)
		if handleHTTPError(w, err) {
			return
		}

		cfg.handleMonsterGet(w, r, input)
		return
	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/monsters/{name or id}, or /api/monsters/{name}/{version}`, nil)
		return
	}
}


func (cfg *apiConfig) handleMonsterGet(w http.ResponseWriter, r *http.Request, input parseResponse) {
	if input.Name != "" {
		dbMons, err := cfg.db.GetMonstersByName(r.Context(), input.Name)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't get multiple Monsters", err)
			return
		}

		resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
			return mon.ID, mon.Name, &mon.Version.Int32, &mon.Specification.String
		})

		resourceList, err := cfg.newNamedAPIResourceList(r, resources)
		if handleHTTPError(w, err) {
			return
		}

		respondWithJSON(w, http.StatusMultipleChoices, resourceList)
		return
	}

	dbMonster, err := cfg.db.GetMonster(r.Context(), input.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get Monster. Monster with this ID doesn't exist.", err)
		return
	}



	response := Monster{
		ID: 					dbMonster.ID,
		Name: 					dbMonster.Name,
		Version: 				h.ConvertNullInt32(dbMonster.Version),
		Specification: 			h.ConvertNullString(dbMonster.Specification),
		Notes: 					h.ConvertNullString(dbMonster.Notes),
		Species: 				string(dbMonster.Species),
		IsStoryBased: 			dbMonster.IsStoryBased,
		CanBeCaptured: 			dbMonster.CanBeCaptured,
		AreaConquestLocation: 	h.ConvertNullMaCreationArea(dbMonster.AreaConquestLocation),
		CTBIconType: 			string(dbMonster.CtbIconType),
		HasOverdrive: 			dbMonster.HasOverdrive,
		IsUnderwater: 			dbMonster.IsUnderwater,
		IsZombie: 				dbMonster.IsZombie,
		Distance: 				anyToInt32(dbMonster.Distance),
		AP: 					dbMonster.Ap,
		APOverkill: 			dbMonster.ApOverkill,
		OverkillDamage: 		dbMonster.OverkillDamage,
		Gil: 					dbMonster.Gil,
		StealGil: 				h.ConvertNullInt32(dbMonster.StealGil),
		DoomCountdown: 			anyToInt32Ptr(dbMonster.DoomCountdown),
		PoisonRate: 			anyToFloat32Ptr(dbMonster.PoisonRate),
		ThreatenChance: 		anyToInt32Ptr(dbMonster.ThreatenChance),
		ZanmatoLevel: 			anyToInt32(dbMonster.ZanmatoLevel),
		MonsterArenaPrice: 		h.ConvertNullInt32(dbMonster.MonsterArenaPrice),
		SensorText: 			dbMonster.SensorText,
		ScanText: 				h.ConvertNullString(dbMonster.ScanText),
	}

	respondWithJSON(w, http.StatusOK, response)
}


func (cfg *apiConfig) handleMonstersRetrieve(w http.ResponseWriter, r *http.Request) {
	dbMons, err := cfg.db.GetMonsters(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve monsters", err)
		return
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.ConvertNullInt32(mon.Version), h.ConvertNullString(mon.Specification)
	})

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resourceList)
}