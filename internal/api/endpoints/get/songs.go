package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Song(r *http.Request, i handlerinput.Main[seedtypes.Song, apitypes.Song, apiresource.Named], id int32) (apitypes.Song, error) {
	song, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Song{}, err
	}
	
	response := apitypes.Song{
		ID:                   song.ID,
		Name:                 song.Name,
		StreamingName:        song.StreamingName,
		InGameName:           song.InGameName,
		OstName:              song.OSTName,
		Translation:          song.Translation,
		DurationInSeconds:    song.DurationInSeconds,
		CanLoop:              song.CanLoop,
		SpecialUseCase:       song.SpecialUseCase,
		StreamingTrackNumber: song.StreamingTrackNumber,
		MusicSphereID:        song.MusicSphereID,
		OstDisc:              song.OSTDisc,
		OstTrackNumber:       song.OSTTrackNumber,
		BackgroundMusic:      ah.ConvertObjSlice(song.BackgroundMusic, apiconvert.BackgroundMusic),
		Cues:                 ah.ConvertObjSlice(song.Cues, apiconvert.Cue),
	}

	err = getSongRelationships(r, song, &response)
	if err != nil {
		return apitypes.Song{}, err
	}

	return response, nil
}

func getSongRelationships(r *http.Request, song seedtypes.Song, response *apitypes.Song) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.BossFights, err = apilookups.E.MonsterFormations.DBGetRelations(ctx, song, handlerinput.ToIntManyNull(config.Cfg.DB.GetSongMonsterFormationIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, song, handlerinput.ToIntManyNull(config.Cfg.DB.GetSongFmvIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	if song.Credits != nil {
		response.Composer = song.Credits.Composer
		response.Arranger = song.Credits.Arranger
		response.Performer = song.Credits.Performer
		response.Lyricist = song.Credits.Lyricist
	}

	return nil
}
