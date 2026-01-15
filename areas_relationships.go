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

func (cfg *Config) getAreaRelationships(r *http.Request, dbArea database.GetAreaRow) (Area, error) {
	locArea := newLocationArea(dbArea.Location, dbArea.Sublocation, dbArea.Name, h.NullInt32ToPtr(dbArea.Version))

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

func (cfg *Config) getAreaConnectedAreas(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]AreaConnection, error) {
	dbConnAreas, err := cfg.db.GetAreaConnections(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get connected areas of %s.", locArea.Error()), err)
	}

	connectedAreas := []AreaConnection{}

	for _, dbConnArea := range dbConnAreas {
		locArea := newLocationArea(dbConnArea.Location, dbConnArea.Sublocation, dbConnArea.Area, h.NullInt32ToPtr(dbConnArea.Version))

		connType, err := cfg.newNamedAPIResourceFromType(cfg.e.connectionType.endpoint, string(dbConnArea.ConnectionType), cfg.t.AreaConnectionType)
		if err != nil {
			return nil, err
		}

		connection := AreaConnection{
			Area:           cfg.newLocationBasedAPIResource(locArea),
			ConnectionType: connType,
			StoryOnly:      dbConnArea.StoryOnly,
			Notes:          h.NullStringToPtr(dbConnArea.Notes),
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}

// all of these can be generalized
// they are essentially: get ids, create resources
// exceptions within areas: sidequest and music
func (cfg *Config) getAreaCharacters(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbChars, err := cfg.db.GetAreaCharacters(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get characters of %s.", locArea.Error()), err)
	}

	chars := createNamedAPIResourcesSimple(cfg, dbChars, cfg.e.characters.endpoint, func(char database.GetAreaCharactersRow) (int32, string) {
		return char.ID, char.Name
	})

	return chars, nil
}

func (cfg *Config) getAreaAeons(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbAeons, err := cfg.db.GetAreaAeons(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get aeons of %s.", locArea.Error()), err)
	}

	aeons := createNamedAPIResourcesSimple(cfg, dbAeons, cfg.e.aeons.endpoint, func(aeon database.GetAreaAeonsRow) (int32, string) {
		return aeon.ID, aeon.Name
	})

	return aeons, nil
}

func (cfg *Config) getAreaShops(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]UnnamedAPIResource, error) {
	dbShops, err := cfg.db.GetAreaShops(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get shops of %s.", locArea.Error()), err)
	}

	shops := createUnnamedAPIResources(cfg, dbShops, cfg.e.shops.endpoint, func(shop database.Shop) int32 {
		return shop.ID
	})

	return shops, nil
}

func (cfg *Config) getAreaTreasures(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]UnnamedAPIResource, error) {
	dbTreasures, err := cfg.db.GetAreaTreasures(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get treasures of %s.", locArea.Error()), err)
	}

	treasures := createUnnamedAPIResources(cfg, dbTreasures, cfg.e.treasures.endpoint, func(treasure database.Treasure) int32 {
		return treasure.ID
	})

	return treasures, nil
}

func (cfg *Config) getAreaMonsters(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbMons, err := cfg.db.GetAreaMonsters(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get monsters of %s.", locArea.Error()), err)
	}

	mons := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.GetAreaMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return mons, nil
}

func (cfg *Config) getAreaFormations(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]UnnamedAPIResource, error) {
	dbFormations, err := cfg.db.GetAreaMonsterFormations(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get monster formations of %s.", locArea.Error()), err)
	}

	formations := createUnnamedAPIResources(cfg, dbFormations, cfg.e.monsterFormations.endpoint, func(formation database.GetAreaMonsterFormationsRow) int32 {
		return formation.ID
	})

	return formations, err
}

func (cfg *Config) getAreaSidequest(r *http.Request, area database.GetAreaRow, locArea LocationArea) (NamedAPIResource, error) {
	dbQuests, err := cfg.db.GetAreaQuests(r.Context(), area.ID)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get quests of %s.", locArea.Error()), err)
	}
	if len(dbQuests) == 0 {
		return NamedAPIResource{}, nil
	}

	// turn this block into a helper function
	potentialSidequest := dbQuests[0]
	questName := potentialSidequest.Name

	if potentialSidequest.Type != database.QuestTypeSidequest {
		subquestName := questName
		subquest, err := seeding.GetResource(subquestName, cfg.l.Subquests)
		if err != nil {
			return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}

		dbSidequest, err := cfg.db.GetParentSidequest(r.Context(), subquest.ID)
		if err != nil {
			return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get parent sidequest of '%s'.", potentialSidequest.Name), err)
		}

		questName = h.NullStringToVal(dbSidequest.Name)
	}

	sidequest, err := seeding.GetResource(questName, cfg.l.Sidequests)
	if err != nil {
		return NamedAPIResource{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	resource := cfg.newNamedAPIResourceSimple(cfg.e.sidequests.endpoint, sidequest.ID, sidequest.Name)

	return resource, nil
}

func (cfg *Config) getAreaMusic(r *http.Request, area database.GetAreaRow, locArea LocationArea) (LocationMusic, error) {
	dbCues, err := cfg.db.GetAreaCues(r.Context(), area.ID)
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get cues of %s.", locArea.Error()), err)
	}

	dbBm, err := cfg.db.GetAreaBackgroundMusic(r.Context(), area.ID)
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get background music of %s.", locArea.Error()), err)
	}

	dbSongsFMVs, err := cfg.db.GetAreaFMVSongs(r.Context(), area.ID)
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get fmv songs of %s.", locArea.Error()), err)
	}

	dbSongsBossFights, err := cfg.db.GetAreaBossSongs(r.Context(), area.ID)
	if err != nil {
		return LocationMusic{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get boss fight songs of %s.", locArea.Error()), err)
	}

	// these two can be generalized
	// also put the entire assembly into a helper function and into location_music.go
	songsCues := cfg.getAreaCues(dbCues)
	songsBM := cfg.getAreaBM(dbBm)

	songsFMVs := createNamedAPIResourcesSimple(cfg, dbSongsFMVs, cfg.e.songs.endpoint, func(song database.GetAreaFMVSongsRow) (int32, string) {
		return song.ID, song.Name
	})

	songsBossFights := createNamedAPIResourcesSimple(cfg, dbSongsBossFights, cfg.e.songs.endpoint, func(song database.GetAreaBossSongsRow) (int32, string) {
		return song.ID, song.Name
	})

	music := LocationMusic{
		Cues:            songsCues,
		BackgroundMusic: songsBM,
		FMVs:            songsFMVs,
		BossFights:      songsBossFights,
	}

	return music, nil
}

func (cfg *Config) getAreaFMVs(r *http.Request, area database.GetAreaRow, locArea LocationArea) ([]NamedAPIResource, error) {
	dbFMVs, err := cfg.db.GetAreaFMVs(r.Context(), area.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get fmvs of %s.", locArea.Error()), err)
	}

	fmvs := createNamedAPIResourcesSimple(cfg, dbFMVs, cfg.e.fmvs.endpoint, func(fmv database.GetAreaFMVsRow) (int32, string) {
		return fmv.ID, fmv.Name
	})

	return fmvs, nil
}
