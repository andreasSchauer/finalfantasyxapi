package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type LocRel struct {
	Characters []NamedAPIResource   `json:"characters"`
	Aeons      []NamedAPIResource   `json:"aeons"`
	Shops      []UnnamedAPIResource `json:"shops"`
	Treasures  []UnnamedAPIResource `json:"treasures"`
	Monsters   []NamedAPIResource   `json:"monsters"`
	Formations []UnnamedAPIResource `json:"formations"`
	Sidequests []NamedAPIResource   `json:"sidequests"`
	Music      *LocBasedMusic       `json:"music"`
	FMVs       []NamedAPIResource   `json:"fmvs"`
}

type LocationArea struct {
	Location    string `json:"location"`
	Sublocation string `json:"sublocation"`
	Area        string `json:"area"`
	Version     *int32 `json:"version,omitempty"`
}

func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: '%s', sublocation: '%s', area: '%s', version: '%v'", la.Location, la.Sublocation, la.Area, h.DerefOrNil(la.Version))
}

type LocBasedMusic struct {
	BackgroundMusic []NamedAPIResource `json:"background_music"`
	Cues            []NamedAPIResource `json:"cues"`
	FMVs            []NamedAPIResource `json:"fmvs"`
	BossMusic       []NamedAPIResource `json:"boss_fights"`
}

func (m LocBasedMusic) IsZero() bool {
	return len(m.BackgroundMusic) == 0 &&
		len(m.Cues) == 0 &&
		len(m.FMVs) == 0 &&
		len(m.BossMusic) == 0
}

type LocBasedMusicQueries struct {
	CueSongs  func(context.Context, int32) ([]int32, error)
	BmSongs   func(context.Context, int32) ([]int32, error)
	FMVSongs  func(context.Context, int32) ([]int32, error)
	BossMusic func(context.Context, int32) ([]int32, error)
}

func getMusicLocBased(cfg *Config, r *http.Request, item seeding.LookupableID, queries LocBasedMusicQueries) (*LocBasedMusic, error) {
	i := cfg.e.songs

	cueSongs, err := getResourcesDB(cfg, r, i, item, queries.CueSongs)
	if err != nil {
		return nil, err
	}

	bmSongs, err := getResourcesDB(cfg, r, i, item, queries.BmSongs)
	if err != nil {
		return nil, err
	}

	fmvSongs, err := getResourcesDB(cfg, r, i, item, queries.FMVSongs)
	if err != nil {
		return nil, err
	}

	bossSongs, err := getResourcesDB(cfg, r, i, item, queries.BossMusic)
	if err != nil {
		return nil, err
	}

	music := LocBasedMusic{
		Cues:            cueSongs,
		BackgroundMusic: bmSongs,
		FMVs:            fmvSongs,
		BossMusic:       bossSongs,
	}

	return &music, nil
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


func findSidequest(cfg *Config, questID int32) (seeding.Sidequest, error) {
	questLookup, _ := seeding.GetResourceByID(questID, cfg.l.QuestsID)

	if questLookup.Type == database.QuestTypeSidequest {
		sidequest, err := seeding.GetResource(questLookup.Name, cfg.l.Sidequests)
		if err != nil {
			return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}

		return sidequest, nil
	}

	subquest, err := seeding.GetResource(questLookup.Name, cfg.l.Subquests)
	if err != nil {
		return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	sidequest, err := seeding.GetResourceByID(subquest.SidequestID, cfg.l.SidequestsID)
	if err != nil {
		return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	return sidequest, nil
}
