package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocRel struct {
	Characters []NamedAPIResource   `json:"characters"`
	Aeons      []NamedAPIResource   `json:"aeons"`
	Shops      []UnnamedAPIResource `json:"shops"`
	Treasures  []UnnamedAPIResource `json:"treasures"`
	Monsters   []NamedAPIResource   `json:"monsters"`
	Formations []UnnamedAPIResource `json:"monster_formations"`
	Quests     []QuestAPIResource   `json:"quests"`
	Music      *LocBasedMusic       `json:"music"`
	FMVs       []NamedAPIResource   `json:"fmvs"`
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

func getMusicLocBased(cfg *Config, r *http.Request, item seeding.Lookupable, queries LocBasedMusicQueries) (*LocBasedMusic, error) {
	i := cfg.e.songs

	cueSongs, err := getResourcesDbItem(cfg, r, i, item, queries.CueSongs)
	if err != nil {
		return nil, err
	}

	bmSongs, err := getResourcesDbItem(cfg, r, i, item, queries.BmSongs)
	if err != nil {
		return nil, err
	}

	fmvSongs, err := getResourcesDbItem(cfg, r, i, item, queries.FMVSongs)
	if err != nil {
		return nil, err
	}

	bossSongs, err := getResourcesDbItem(cfg, r, i, item, queries.BossMusic)
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

func getLocBasedSidequests(cfg *Config, r *http.Request, item seeding.Lookupable, p RelAvlParams, dbQuery RelAvailabilityDbQuery) ([]QuestAPIResource, error) {
	resources := []QuestAPIResource{}

	dbQuestIDs, err := dbQuery(r.Context(), p)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %s of %s.", cfg.e.quests.resTypePlural, item), err)
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
	quest, _ := seeding.GetResourceByID(questID, cfg.l.QuestsID)

	if quest.Type == database.QuestTypeSidequest {
		sidequest, err := seeding.GetResource(quest, cfg.l.Sidequests)
		if err != nil {
			return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
		}

		return sidequest, nil
	}

	subquest, err := seeding.GetResource(quest, cfg.l.Subquests)
	if err != nil {
		return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	sidequest, err := seeding.GetResourceByID(subquest.SidequestID, cfg.l.SidequestsID)
	if err != nil {
		return seeding.Sidequest{}, newHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	return sidequest, nil
}
