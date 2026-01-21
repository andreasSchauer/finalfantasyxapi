package main

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

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

// can be used by areas, locations, and sublocations respectively. This is the closed to generic I can get with my DB setup
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
