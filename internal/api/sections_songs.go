package api

import (
	"context"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func CreateSongSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Songs
	song, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	songSimple := sectiontypes.SongSimple{
		ID:                song.ID,
		URL:               ah.CreateResourceURL(i.Endpoint, id),
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

func getLocationSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := config.Cfg.DB.GetLocationBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of location with id '%d'", err)
	}

	cueSongIDs, err := config.Cfg.DB.GetLocationCueSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of location with id '%d'", err)
	}

	bossSongIDs, err := config.Cfg.DB.GetLocationBossSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get boss song of location with id '%d'", err)
	}

	fmvSongIDs, err := config.Cfg.DB.GetLocationFMVSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get fmv song of location with id '%d'", err)
	}

	ids := ah.CombineIDSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}

func getSublocationSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := config.Cfg.DB.GetSublocationBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of sublocation with id '%d'", err)
	}

	cueSongIDs, err := config.Cfg.DB.GetSublocationCueSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of sublocation with id '%d'", err)
	}

	bossSongIDs, err := config.Cfg.DB.GetSublocationBossSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get boss song of sublocation with id '%d'", err)
	}

	fmvSongIDs, err := config.Cfg.DB.GetSublocationFMVSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get fmv song of sublocation with id '%d'", err)
	}

	ids := ah.CombineIDSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}

func getAreaSongIDs(ctx context.Context, id int32) ([]int32, error) {
	bgSongIDs, err := config.Cfg.DB.GetAreaBackgroundMusicSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get bg song ids of area with id '%d'", err)
	}

	cueSongIDs, err := handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaCueSongIDs)(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get cue song ids of area with id '%d'", err)
	}

	bossSongIDs, err := config.Cfg.DB.GetAreaBossSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get boss song of area with id '%d'", err)
	}

	fmvSongIDs, err := config.Cfg.DB.GetAreaFMVSongIDs(ctx, id)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get fmv song of area with id '%d'", err)
	}

	ids := ah.CombineIDSlices(bgSongIDs, cueSongIDs, bossSongIDs, fmvSongIDs)
	return ids, nil
}
