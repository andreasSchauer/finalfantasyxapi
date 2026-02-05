package main

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SongSub struct {
	ID                int32   `json:"id"`
	URL               string  `json:"url"`
	Name              string  `json:"name"`
	Composer          *string `json:"composer,omitempty"`
	Arranger          *string `json:"arranger,omitempty"`
	Performer         *string `json:"performer,omitempty"`
	Lyricist          *string `json:"lyricist,omitempty"`
	DurationInSeconds int32   `json:"duration_in_seconds"`
	CanLoop           bool    `json:"can_loop"`
}

func (s SongSub) GetSectionName() string {
	return "songs"
}

func (s SongSub) GetURL() string {
	return s.URL
}

func handleSongsSection(cfg *Config, _ *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.songs
	songs := []SongSub{}

	for _, songID := range dbIDs {
		song, _ := seeding.GetResourceByID(songID, i.objLookupID)

		songSub := SongSub{
			ID:                song.ID,
			URL:               createResourceURL(cfg, i.endpoint, songID),
			Name:              song.Name,
			DurationInSeconds: song.DurationInSeconds,
			CanLoop:           song.CanLoop,
		}

		if song.Credits != nil {
			songSub.Composer = song.Credits.Composer
			songSub.Arranger = song.Credits.Arranger
			songSub.Performer = song.Credits.Performer
			songSub.Lyricist = song.Credits.Lyricist
		}

		songs = append(songs, songSub)
	}

	return toSubResourceSlice(songs), nil
}

func (cfg *Config) getLocationSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := cfg.db.GetLocationBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of location with id '%d'.", err)
	}

	cueSongIDs, err := cfg.db.GetLocationCueSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of location with id '%d'.", err)
	}

	bossSongIDs, err := cfg.db.GetLocationBossSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get boss song of location with id '%d'.", err)
	}

	fmvSongIDs, err := cfg.db.GetLocationFMVSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get fmv song of location with id '%d'.", err)
	}

	ids := combineIdSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}

func (cfg *Config) getSublocationSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := cfg.db.GetSublocationBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of sublocation with id '%d'.", err)
	}

	cueSongIDs, err := cfg.db.GetSublocationCueSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of sublocation with id '%d'.", err)
	}

	bossSongIDs, err := cfg.db.GetSublocationBossSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get boss song of sublocation with id '%d'.", err)
	}

	fmvSongIDs, err := cfg.db.GetSublocationFMVSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get fmv song of sublocation with id '%d'.", err)
	}

	ids := combineIdSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}

func (cfg *Config) getAreaSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := cfg.db.GetAreaBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of area with id '%d'.", err)
	}

	cueSongIDs, err := cfg.db.GetAreaCueSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of area with id '%d'.", err)
	}

	bossSongIDs, err := cfg.db.GetAreaBossSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get boss song of area with id '%d'.", err)
	}

	fmvSongIDs, err := cfg.db.GetAreaFMVSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get fmv song of area with id '%d'.", err)
	}

	ids := combineIdSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}
