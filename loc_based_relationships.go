package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationArea struct {
	Location    string `json:"location"`
	Sublocation string `json:"sublocation"`
	Area        string `json:"area"`
	Version     *int32 `json:"version,omitempty"`
}


func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: '%s', sublocation: '%s', area: '%s', version: '%v'", la.Location, la.Sublocation, la.Area, h.DerefOrNil(la.Version))
}

type LocationMusic struct {
	BackgroundMusic []LocationSong     `json:"background_music"`
	Cues            []LocationSong     `json:"cues"`
	FMVs            []NamedAPIResource `json:"fmvs"`
	BossMusic       []NamedAPIResource `json:"boss_fights"`
}

func (m LocationMusic) IsZero() bool {
	return len(m.BackgroundMusic) == 0 &&
		len(m.Cues) == 0 &&
		len(m.FMVs) == 0 &&
		len(m.BossMusic) == 0
}

type LocationMusicQueries struct {
	FMVSongs  func(context.Context, int32) ([]int32, error)
	BossMusic func(context.Context, int32) ([]int32, error)
}

type LocationSong struct {
	Song                   NamedAPIResource `json:"song"`
	ReplacesEncounterMusic bool             `json:"replaces_encounter_music"`
}

func (ls LocationSong) GetAPIResource() APIResource {
	return ls.Song
}

func newLocationSong(cfg *Config, i handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList], songID int32, replEncMusic bool) LocationSong {
	return LocationSong{
		Song:                   i.idToResFunc(cfg, i, songID),
		ReplacesEncounterMusic: replEncMusic,
	}
}


func completeLocationMusic(cfg *Config, r *http.Request, i handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList], item seeding.LookupableID, cueSongs, bmSongs []LocationSong, queries LocationMusicQueries) (LocationMusic, error) {
	fmvSongs, err := getResourcesDB(cfg, r, i, item, queries.FMVSongs)
	if err != nil {
		return LocationMusic{}, err
	}

	bossSongs, err := getResourcesDB(cfg, r, i, item, queries.BossMusic)
	if err != nil {
		return LocationMusic{}, err
	}

	music := LocationMusic{
		Cues:            cueSongs,
		BackgroundMusic: bmSongs,
		FMVs:            fmvSongs,
		BossMusic:       bossSongs,
	}

	return music, nil
}


func getLocBasedSidequests(cfg *Config, r *http.Request, item seeding.LookupableID, dbQuery func(context.Context, int32) ([]int32, error)) ([]NamedAPIResource, error) {
	resources := []NamedAPIResource{}

	dbQuestIDs, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get quests of %s.", item), err)
	}
	if len(dbQuestIDs) == 0 {
		return resources, nil
	}

	for _, dbID := range dbQuestIDs {
		sidequest, err := findSidequest(cfg, dbID)
		if err != nil {
			return nil, err
		}

		resource := cfg.e.sidequests.idToResFunc(cfg, cfg.e.sidequests, sidequest.ID)

		if !resourcesContain(resources, resource) {
			resources = append(resources, resource)
		}

	}

	return resources, nil
}

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
