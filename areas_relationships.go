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

func (cfg *Config) getAreaRelationships(r *http.Request, areaLookup seeding.Area) (Area, error) {
	connections, err := cfg.getAreaConnectedAreas(areaLookup)
	if err != nil {
		return Area{}, err
	}

	characters, err := getNamedResources(cfg, r, cfg.e.characters, areaLookup, cfg.db.GetAreaCharacterIDs)
	if err != nil {
		return Area{}, err
	}

	aeons, err := getNamedResources(cfg, r, cfg.e.aeons, areaLookup, cfg.db.GetAreaAeonIDs)
	if err != nil {
		return Area{}, err
	}

	shops, err := getUnnamedResources(cfg, r, cfg.e.shops, areaLookup, cfg.db.GetAreaShopIDs)
	if err != nil {
		return Area{}, err
	}

	treasures, err := getUnnamedResources(cfg, r, cfg.e.treasures, areaLookup, cfg.db.GetAreaTreasureIDs)
	if err != nil {
		return Area{}, err
	}

	monsters, err := getNamedResources(cfg, r, cfg.e.monsters, areaLookup, cfg.db.GetAreaMonsterIDs)
	if err != nil {
		return Area{}, err
	}

	formations, err := getUnnamedResources(cfg, r, cfg.e.monsterFormations, areaLookup, cfg.db.GetAreaMonsterFormationIDs)
	if err != nil {
		return Area{}, err
	}

	sidequest, err := cfg.getAreaSidequest(r, areaLookup)
	if err != nil {
		return Area{}, err
	}

	music, err := cfg.getAreaMusic(r, areaLookup)
	if err != nil {
		return Area{}, err
	}

	fmvs, err := getNamedResources(cfg, r, cfg.e.fmvs, areaLookup, cfg.db.GetAreaFmvIDs)
	if err != nil {
		return Area{}, err
	}

	area := Area{
		ConnectedAreas: connections,
		Characters:     characters,
		Aeons:          aeons,
		Shops:          shops,
		Treasures:      treasures,
		Monsters:       monsters,
		Formations:     formations,
		Sidequest:      h.NilOrPtr(sidequest),
		Music:          h.NilOrPtr(music),
		FMVs:           fmvs,
	}

	return area, nil
}

func (cfg *Config) getAreaConnectedAreas(area seeding.Area) ([]AreaConnection, error) {
	i := cfg.e.areas
	connectedAreas := []AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		locArea := connArea.LocationArea

		connType, err := cfg.newNamedAPIResourceFromType(cfg.e.connectionType.endpoint, string(connArea.ConnectionType), cfg.t.AreaConnectionType)
		if err != nil {
			return nil, err
		}

		connection := AreaConnection{
			Area:           locAreaToLocationAPIResource(cfg, i, locArea),
			ConnectionType: connType,
			StoryOnly:      connArea.StoryOnly,
			Notes:          connArea.Notes,
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}



func (cfg *Config) getAreaSidequest(r *http.Request, area seeding.Area) (NamedAPIResource, error) {
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

	resource := cfg.newNamedAPIResourceSimple(cfg.e.sidequests.endpoint, sidequest.ID, sidequest.Name)

	return resource, nil
}

// this is kind of scuffed for now. I will probably find a better way, once I've managed to implement proper parent type assertions
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


func (cfg *Config) getAreaMusic(r *http.Request, item seeding.LookupableID) (LocationMusic, error) {
	i := cfg.e.songs

	dbCues, err := cfg.db.GetAreaCues(r.Context(), item.GetID())
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get cues of %s.", item), err)
	}

	dbBm, err := cfg.db.GetAreaBackgroundMusic(r.Context(), item.GetID())
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get background music of %s.", item), err)
	}

	dbFMVSongIDs, err := cfg.db.GetAreaFMVSongIDs(r.Context(), item.GetID())
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get fmv songs of %s.", item), err)
	}

	dbBossSongIDs, err := cfg.db.GetAreaBossSongIDs(r.Context(), item.GetID())
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get boss fight songs of %s.", item), err)
	}

	music := LocationMusic{
		Cues:            getAreaCues(cfg, i, dbCues),
		BackgroundMusic: getAreaBM(cfg, i, dbBm),
		FMVs:            idsToNamedAPIResources(cfg, i, dbFMVSongIDs),
		BossFights:      idsToNamedAPIResources(cfg, i, dbBossSongIDs),
	}

	return music, nil
}