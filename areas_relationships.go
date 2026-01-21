package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaConnection struct {
	Area           LocationAPIResource `json:"area"`
	ConnectionType NamedAPIResource    `json:"connection_type"`
	StoryOnly      bool                `json:"story_only"`
	Notes          *string             `json:"notes,omitempty"`
}

func (ac AreaConnection) GetAPIResource() APIResource {
	return ac.Area
}

func getAreaRelationships(cfg *Config, r *http.Request, area seeding.Area) (Area, error) {
	connections, err := getAreaConnectedAreas(cfg, area)
	if err != nil {
		return Area{}, err
	}

	characters, err := getResourcesDB(cfg, r, cfg.e.characters, area, cfg.db.GetAreaCharacterIDs)
	if err != nil {
		return Area{}, err
	}

	aeons, err := getResourcesDB(cfg, r, cfg.e.aeons, area, cfg.db.GetAreaAeonIDs)
	if err != nil {
		return Area{}, err
	}

	shops, err := getResourcesDB(cfg, r, cfg.e.shops, area, cfg.db.GetAreaShopIDs)
	if err != nil {
		return Area{}, err
	}

	treasures, err := getResourcesDB(cfg, r, cfg.e.treasures, area, cfg.db.GetAreaTreasureIDs)
	if err != nil {
		return Area{}, err
	}

	monsters, err := getResourcesDB(cfg, r, cfg.e.monsters, area, cfg.db.GetAreaMonsterIDs)
	if err != nil {
		return Area{}, err
	}

	formations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, area, cfg.db.GetAreaMonsterFormationIDs)
	if err != nil {
		return Area{}, err
	}

	sidequest, err := getAreaSidequest(cfg, r, area)
	if err != nil {
		return Area{}, err
	}

	music, err := getAreaMusic(cfg, r, area)
	if err != nil {
		return Area{}, err
	}

	fmvs, err := getResourcesDB(cfg, r, cfg.e.fmvs, area, cfg.db.GetAreaFmvIDs)
	if err != nil {
		return Area{}, err
	}

	rel := Area{
		ConnectedAreas: connections,
		Characters:     characters,
		Aeons:          aeons,
		Shops:          shops,
		Treasures:      treasures,
		Monsters:       monsters,
		Formations:     formations,
		Sidequest:      h.ObjPtrOrNil(sidequest),
		Music:          h.ObjPtrOrNil(music),
		FMVs:           fmvs,
	}

	return rel, nil
}

func getAreaConnectedAreas(cfg *Config, area seeding.Area) ([]AreaConnection, error) {
	i := cfg.e.areas
	connectedAreas := []AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connType, err := newNamedAPIResourceFromType(cfg, cfg.e.connectionType.endpoint, connArea.ConnectionType, cfg.t.AreaConnectionType)
		if err != nil {
			return nil, err
		}

		connection := AreaConnection{
			Area:           locAreaToLocationAPIResource(cfg, i, connArea.LocationArea),
			ConnectionType: connType,
			StoryOnly:      connArea.StoryOnly,
			Notes:          connArea.Notes,
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}

func getAreaSidequest(cfg *Config, r *http.Request, area seeding.Area) (NamedAPIResource, error) {
	dbQuestIDs, err := cfg.db.GetAreaQuestIDs(r.Context(), area.ID)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get quests of %s.", area), err)
	}
	if len(dbQuestIDs) == 0 {
		return NamedAPIResource{}, nil
	}

	sidequest, err := findSidequest(cfg, dbQuestIDs[0])
	if err != nil {
		return NamedAPIResource{}, err
	}

	resource := cfg.e.sidequests.idToResFunc(cfg, cfg.e.sidequests, sidequest.ID)

	return resource, nil
}

// this is kind of scuffed for now. I will probably find a better way, once I've managed to implement proper parent type assertions
// then on the other hand, what else should I do here? it kind of is a special case, since these are not two equal types, but parent and child being categorized by the same master-type
// at least it's reusable
func findSidequest(cfg *Config, potentialSidequestID int32) (seeding.Sidequest, error) {
	potentialSidequest, _ := seeding.GetResourceByID(potentialSidequestID, cfg.l.QuestsID)
	sidequestID := potentialSidequestID

	if potentialSidequest.Type != database.QuestTypeSidequest {
		subquestName := potentialSidequest.Name
		subquest, err := seeding.GetResource(subquestName, cfg.l.Subquests)
		if err != nil {
			return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}

		sidequestID = subquest.SidequestID
	}

	sidequest, err := seeding.GetResourceByID(sidequestID, cfg.l.SidequestsID)
	if err != nil {
		return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	return sidequest, nil
}

func getAreaMusic(cfg *Config, r *http.Request, item seeding.LookupableID) (LocationMusic, error) {
	i := cfg.e.songs

	cueSongs, err := getAreaCueSongs(cfg, r, i, item)
	if err != nil {
		return LocationMusic{}, err
	}

	bmSongs, err := getAreaBMSongs(cfg, r, i, item)
	if err != nil {
		return LocationMusic{}, err
	}

	music, err := completeLocationMusic(cfg, r, i, item, cueSongs, bmSongs, LocationMusicQueries{
		FMVSongs:  cfg.db.GetAreaFMVSongIDs,
		BossMusic: cfg.db.GetAreaBossSongIDs,
	})
	if err != nil {
		return LocationMusic{}, err
	}

	return music, nil
}

func getAreaCueSongs(cfg *Config, r *http.Request, i handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList], item seeding.LookupableID) ([]LocationSong, error) {
	dbCueSongs, err := cfg.db.GetAreaCues(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get cues of %s", item), err)
	}

	var cueSongs []LocationSong
	for _, song := range dbCueSongs {
		cueSongs = append(cueSongs, newLocationSong(cfg, i, song.ID, song.ReplacesEncounterMusic))
	}

	return cueSongs, nil
}

func getAreaBMSongs(cfg *Config, r *http.Request, i handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList], item seeding.LookupableID) ([]LocationSong, error) {
	dbBMSongs, err := cfg.db.GetAreaBackgroundMusic(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get cues of %s", item), err)
	}

	var bmSongs []LocationSong
	for _, song := range dbBMSongs {
		bmSongs = append(bmSongs, newLocationSong(cfg, i, song.ID, song.ReplacesEncounterMusic))
	}

	return bmSongs, nil
}
