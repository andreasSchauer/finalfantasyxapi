package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
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

func createSongSimple(srv *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.Songs
	song, _ := lookup.GetResourceByID(id, i.objLookupID)

	songSimple := SongSimple{
		ID:                song.ID,
		URL:               CreateResourceURL(srv, i.Endpoint, id),
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

func (srv *Server) getLocationSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := srv.Cfg.DB.GetLocationBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of location with id '%d'", err)
	}

	cueSongIDs, err := srv.Cfg.DB.GetLocationCueSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of location with id '%d'", err)
	}

	bossSongIDs, err := srv.Cfg.DB.GetLocationBossSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get boss song of location with id '%d'", err)
	}

	fmvSongIDs, err := srv.Cfg.DB.GetLocationFMVSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get fmv song of location with id '%d'", err)
	}

	ids := combineIDSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}

func (srv *Server) getSublocationSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := srv.Cfg.DB.GetSublocationBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of sublocation with id '%d'", err)
	}

	cueSongIDs, err := srv.Cfg.DB.GetSublocationCueSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of sublocation with id '%d'", err)
	}

	bossSongIDs, err := srv.Cfg.DB.GetSublocationBossSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get boss song of sublocation with id '%d'", err)
	}

	fmvSongIDs, err := srv.Cfg.DB.GetSublocationFMVSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get fmv song of sublocation with id '%d'", err)
	}

	ids := combineIDSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}

func (srv *Server) getAreaSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := srv.Cfg.DB.GetAreaBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of area with id '%d'", err)
	}

	cueSongIDs, err := ToIntManyNull(srv.Cfg.DB.GetAreaCueSongIDs)(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of area with id '%d'", err)
	}

	bossSongIDs, err := srv.Cfg.DB.GetAreaBossSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get boss song of area with id '%d'", err)
	}

	fmvSongIDs, err := srv.Cfg.DB.GetAreaFMVSongIDs(ctx, id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get fmv song of area with id '%d'", err)
	}

	ids := combineIDSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}
