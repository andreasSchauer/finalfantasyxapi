package main

import (
	//"database/sql"
	//"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type Area struct {
	ID					int32					`json:"id"`
	Name				string					`json:"name"`
	Version				*int32					`json:"version,omitempty"`
	Specification		*string					`json:"specification,omitempty"`
	ParentLocation		NamedAPIResource		`json:"parent_location"`
	ParentSublocation	NamedAPIResource		`json:"parent_sublocation"`
	StoryOnly			bool					`json:"story_only"`
	HasSaveSphere		bool					`json:"has_save_sphere"`
	AirshipDropOff		bool					`json:"airship_drop_off"`
	HasCompSphere		bool					`json:"has_comp_sphere"`
	CanRideChocobo		bool					`json:"can_ride_chocobo"`
	ConnectedAreas		[]AreaConnection		`json:"connected_areas"`
	Characters			[]NamedAPIResource		`json:"characters"`
	Aeons				[]NamedAPIResource		`json:"aeons"`
	Shops				[]UnnamedAPIResource	`json:"shops"`
	Treasures			[]UnnamedAPIResource	`json:"treasures"`
	Monsters			[]NamedAPIResource		`json:"monsters"`
	Formations			[]UnnamedAPIResource	`json:"formations"`
	Sidequest			*NamedAPIResource		`json:"sidequest"`
	Music				*Music					`json:"music"`
	FMVs				[]NamedAPIResource		`json:"fmvs"`
}


type AreaConnection struct {
	Area			LocationAPIResource			`json:"area"`
	ConnectionType	NamedAPIResource			`json:"connection_type"`
	StoryOnly		bool						`json:"story_only"`
	Notes			*string						`json:"notes,omitempty"`
}


type Music struct {
	BackgroundMusic		[]LocationSong			`json:"background_music"`
	Cues				[]LocationSong			`json:"cues"`
	FMVs				[]NamedAPIResource		`json:"fmvs"`
	BossFights			[]NamedAPIResource		`json:"boss_fights"`
}

func (m Music) IsZero() bool {
	return 	len(m.BackgroundMusic) == 0 &&
			len(m.Cues) == 0 &&
			len(m.FMVs) == 0 &&
			len(m.BossFights) == 0
}

type LocationSong struct {
	Song					NamedAPIResource	`json:"song"`
	ReplacesEncounterMusic 	bool				`json:"replaces_encounter_music"`
}


func (cfg *apiConfig) handleAreas(w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, "areas")
	
	// this whole thing can probably be generalized
	switch len(segments) {
	case 0:
		// /api/areas
		resourceList, err := cfg.retrieveAreas(r)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, resourceList)
		return
	case 1:
		// /api/areas/{id}
		idStr := segments[0]
		
		id, err := strconv.Atoi(idStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Wrong format. Usage: /api/areas/{id}.", err)
			return
		}

		area, err := cfg.getArea(r, int32(id))
		if handleHTTPError(w, err) {
			return
		}

		respondWithJSON(w, http.StatusOK, area)
		return

	case 2:
		// /api/areas/{id}/{subSection}
		// areaID := segments[0]
		subSection := segments[1]
		switch subSection {
		case "connected":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/connected")
			return
		case "monsters":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/monsters")
			return
		case "monster-formations":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/monster-formations")
			return
		case "shops":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/shops")
			return
		case "treasures":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/treasures")
			return
		default:
			fmt.Println(segments)
			fmt.Println("this should trigger an error: this sub section is not supported. Supported sub-sections: connected, monsters, monster-formations, shops, treasures.")
			return
		}

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/areas/{id}, or /api/areas/{id}/{sub-section}. Supported sub-sections: connected, monsters, monster-formations, shops, treasures.`, nil)
		return
	}
}


func (cfg *apiConfig) getArea(r *http.Request, id int32) (Area, error) {
	dbArea, err := cfg.db.GetArea(r.Context(), id)
	if err != nil {
		return Area{}, newHTTPError(http.StatusNotFound, "Couldn't get Area. Area with this ID doesn't exist.", err)
	}

	locArea := newLocationArea(h.NullStringToVal(dbArea.Location), h.NullStringToVal(dbArea.Sublocation), dbArea.Name, h.NullInt32ToPtr(dbArea.Version))

	connections, err := cfg.getAreaConnectedAreas(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	characters, err := cfg.getAreaCharacters(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	aeons, err := cfg.getAreaAeons(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	shops, err := cfg.getAreaShops(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	treasures, err := cfg.getAreaTreasures(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	monsters, err := cfg.getAreaMonsters(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	formations, err := cfg.getAreaFormations(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	sidequest, err := cfg.getAreaSidequest(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	music, err := cfg.getAreaMusic(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	fmvs, err := cfg.getAreaFMVs(r, dbArea, locArea)
	if err != nil {
		return Area{}, err
	}

	location  := cfg.newNamedAPIResourceSimple("locations", h.NullInt32ToVal(dbArea.LocationID), h.NullStringToVal(dbArea.Location))

	sublocation := cfg.newNamedAPIResourceSimple("sublocations", dbArea.SublocationID, h.NullStringToVal(dbArea.Sublocation))

	area := Area{
		ID: 				dbArea.ID,
		Name: 				dbArea.Name,
		Version: 			h.NullInt32ToPtr(dbArea.Version),
		Specification: 		h.NullStringToPtr(dbArea.Specification),
		ParentLocation: 	location,
		ParentSublocation: 	sublocation,
		StoryOnly: 			dbArea.StoryOnly,
		HasSaveSphere: 		dbArea.HasSaveSphere,
		AirshipDropOff: 	dbArea.AirshipDropOff,
		HasCompSphere: 		dbArea.HasCompilationSphere,
		CanRideChocobo: 	dbArea.CanRideChocobo,
		ConnectedAreas: 	connections,
		Characters: 		characters,
		Aeons: 				aeons,
		Shops: 				shops,
		Treasures: 			treasures,
		Monsters: 			monsters,
		Formations: 		formations,
		Sidequest: 			h.NilOrPtr(sidequest),
		Music: 				h.NilOrPtr(music),
		FMVs: 				fmvs,
	}

	return area, nil
}


func (cfg *apiConfig) getAreaCharacters(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbChars, err := cfg.db.GetAreaCharacters(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get Aeons of %s", locArea.Error()), err)
	}

	chars := createNamedAPIResourcesSimple(cfg, dbChars, "characters", func(char database.GetAreaCharactersRow) (int32, string) {
		return h.NullInt32ToVal(char.ID), char.Name
	})

	return chars, nil
}


func (cfg *apiConfig) getAreaAeons(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbAeons, err := cfg.db.GetAreaAeons(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get Aeons of %s", locArea.Error()), err)
	}

	aeons := createNamedAPIResourcesSimple(cfg, dbAeons, "aeons", func(aeon database.GetAreaAeonsRow) (int32, string) {
		return h.NullInt32ToVal(aeon.ID), aeon.Name
	})

	return aeons, nil
}


func (cfg *apiConfig) getAreaFMVs(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbFMVs, err := cfg.db.GetAreaFMVs(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get FMVs of %s", locArea.Error()), err)
	}

	fmvs := createNamedAPIResourcesSimple(cfg, dbFMVs, "fmvs", func(fmv database.GetAreaFMVsRow) (int32, string) {
		return fmv.ID, fmv.Name
	})

	return fmvs, nil
}


func (cfg *apiConfig) getAreaMusic(r *http.Request, area database.GetAreaRow, locArea LocationArea) (Music, error) {
	dbCues, err := cfg.db.GetAreaCues(r.Context(), area.ID)
	if err != nil {
		return Music{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get cues of %s", locArea.Error()), err)
	}

	songsCues := []LocationSong{}

	for _, cue := range dbCues {
		song := cfg.newNamedAPIResourceSimple("songs", h.NullInt32ToVal(cue.ID), h.NullStringToVal(cue.Name))

		locationSong := LocationSong{
			Song: 					song,
			ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
		}

		songsCues = append(songsCues, locationSong)
	}


	dbBm, err := cfg.db.GetAreaBackgroundMusic(r.Context(), area.ID)
	if err != nil {
		return Music{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get background music of %s", locArea.Error()), err)
	}

	songsBM := []LocationSong{}

	for _, bm := range dbBm {
		song := cfg.newNamedAPIResourceSimple("songs", h.NullInt32ToVal(bm.ID), h.NullStringToVal(bm.Name))

		locationSong := LocationSong{
			Song: 					song,
			ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
		}

		songsBM = append(songsBM, locationSong)
	}

	
	dbSongsFMVs, err := cfg.db.GetAreaFMVSongs(r.Context(), area.ID)
	if err != nil {
		return Music{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get fmv music of %s", locArea.Error()), err)
	}

	songsFMVs := createNamedAPIResourcesSimple(cfg, dbSongsFMVs, "songs", func(song database.GetAreaFMVSongsRow) (int32, string) {
		return song.ID, song.Name
	})


	dbSongsBossFights, err := cfg.db.GetAreaBossSongs(r.Context(), area.ID)
	if err != nil {
		return Music{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get boss fight songs of %s", locArea.Error()), err)
	}

	songsBossFights := createNamedAPIResourcesSimple(cfg, dbSongsBossFights, "songs", func(song database.GetAreaBossSongsRow) (int32, string) {
		return song.ID, song.Name
	})


	music := Music{
		Cues: 				songsCues,
		BackgroundMusic: 	songsBM,
		FMVs: 				songsFMVs,
		BossFights: 		songsBossFights,
	}

	return music, nil
}


func (cfg *apiConfig) getAreaSidequest(r *http.Request, area database.GetAreaRow, locArea LocationArea) (NamedAPIResource, error) {
	dbQuests, err := cfg.db.GetAreaQuests(r.Context(), area.ID)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get quests of %s", locArea.Error()), err)
	}
	if len(dbQuests) == 0 {
		return NamedAPIResource{}, nil
	}

	potentialSidequest := dbQuests[0]
	questName := h.NullStringToVal(potentialSidequest.Name)

	if potentialSidequest.Type.QuestType != database.QuestTypeSidequest {
		subquestName := questName
		subquest, err := seeding.GetResource(subquestName, cfg.l.Subquests)
		if err != nil {
			return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}

		dbSidequest, err := cfg.db.GetParentSidequest(r.Context(), subquest.ID)
		if err != nil {
			return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get parent sidequest of %s", h.NullStringToVal(potentialSidequest.Name)), err)
		}

		questName = h.NullStringToVal(dbSidequest.Name)
	} 

	sidequest, err := seeding.GetResource(questName, cfg.l.Sidequests)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	resource := cfg.newNamedAPIResourceSimple("sidequests", sidequest.ID, sidequest.Name)

	return resource, nil
}


func (cfg *apiConfig) getAreaFormations(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]UnnamedAPIResource, error) {
	dbFormations, err := cfg.db.GetAreaMonsterFormations(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get monster formations of %s", locArea.Error()), err)
	}

	formations := createUnnamedAPIResources(cfg, dbFormations, "monster-formations", func(formation database.GetAreaMonsterFormationsRow)(int32) {
		return formation.ID
	})

	return formations, err
}


func (cfg *apiConfig) getAreaMonsters(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbMons, err := cfg.db.GetAreaMonsters(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get monsters of %s", locArea.Error()), err)
	}

	mons := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetAreaMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return mons, nil
}


func (cfg *apiConfig) getAreaShops(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]UnnamedAPIResource, error) {
	dbShops, err := cfg.db.GetAreaShops(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get shops of %s", locArea.Error()), err)
	}

	shops := createUnnamedAPIResources(cfg, dbShops, "shops", func(shop database.Shop)(int32) {
		return shop.ID
	})

	return shops, nil
}


func (cfg *apiConfig) getAreaTreasures(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]UnnamedAPIResource, error) {
	dbTreasures, err := cfg.db.GetAreaTreasures(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get treasures of %s", locArea.Error()), err)
	}

	treasures := createUnnamedAPIResources(cfg, dbTreasures, "treasures", func(treasure database.Treasure)(int32) {
		return treasure.ID
	})

	return treasures, nil
}


func (cfg *apiConfig) getAreaConnectedAreas(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]AreaConnection, error) {
	dbConnAreas, err := cfg.db.GetAreaConnections(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't get connected areas of %s", locArea.Error()), err)
	}

	connectedAreas := []AreaConnection{}

	for _, dbConnArea := range dbConnAreas {
		locArea := newLocationArea(h.NullStringToVal(dbConnArea.Location), h.NullStringToVal(dbConnArea.Sublocation), h.NullStringToVal(dbConnArea.Area), h.NullInt32ToPtr(dbConnArea.Version))

		connType, err := cfg.newNamedAPIResourceFromType("connection-type", string(dbConnArea.ConnectionType), cfg.t.AreaConnectionType)
		if err != nil {
			return nil, err
		}

		connection := AreaConnection{
			Area: 			cfg.newLocationBasedAPIResource(locArea),
			ConnectionType: connType,
			StoryOnly: 		dbConnArea.StoryOnly,
			Notes: 			h.NullStringToPtr(dbConnArea.Notes),
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}


func (cfg *apiConfig) retrieveAreas(r *http.Request) (LocationdApiResourceList, error) {
	dbAreas, err := cfg.db.GetAreas(r.Context())
	if err != nil {
		return LocationdApiResourceList{}, newHTTPError(http.StatusInternalServerError, "Couldn't retrieve areas", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasRow) (string, string, string, *int32) {
		return h.NullStringToVal(area.Location), h.NullStringToVal(area.Sublocation), area.Name, h.NullInt32ToPtr(area.Version)
	})

	resourceList, err := cfg.newLocationAPIResourceList(r, resources)
	if err != nil {
		return LocationdApiResourceList{}, err
	}

	return resourceList, nil
}