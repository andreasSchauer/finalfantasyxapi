package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetSongRelationships(r *http.Request, song seedtypes.Song) (apitypes.Song, error) {
	var rel apitypes.Song
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.BossFights, err = apilookups.E.MonsterFormations.DBGetRelations(ctx, song, handlerinput.ToIntManyNull(config.Cfg.DB.GetSongMonsterFormationIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, song, handlerinput.ToIntManyNull(config.Cfg.DB.GetSongFmvIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Song{}, err
	}

	if song.Credits != nil {
		rel.Composer = song.Credits.Composer
		rel.Arranger = song.Credits.Arranger
		rel.Performer = song.Credits.Performer
		rel.Lyricist = song.Credits.Lyricist
	}

	return rel, nil
}
