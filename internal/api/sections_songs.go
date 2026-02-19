package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SongSimple struct {
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

func (s SongSimple) GetURL() string {
	return s.URL
}

func createSongSimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.songs
	song, _ := seeding.GetResourceByID(id, i.objLookupID)

	songSimple := SongSimple{
		ID:                song.ID,
		URL:               createResourceURL(cfg, i.endpoint, id),
		Name:              song.Name,
		DurationInSeconds: song.DurationInSeconds,
		CanLoop:           song.CanLoop,
	}

	if song.Credits != nil {
		songSimple.Composer = song.Credits.Composer
		songSimple.Arranger = song.Credits.Arranger
		songSimple.Performer = song.Credits.Performer
		songSimple.Lyricist = song.Credits.Lyricist
	}

	return songSimple, nil
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
