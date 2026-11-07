package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSongRelationships(r *http.Request, song seedtypes.Song) (apitypes.Song, error) {
	var rel apitypes.Song
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.BossFights, err = E.MonsterFormations.DBGetRelations(ctx, song, ToIntManyNull(config.Cfg.DB.GetSongMonsterFormationIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = E.FMVs.DBGetRelations(ctx, song, ToIntManyNull(config.Cfg.DB.GetSongFmvIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Song{}, err
	}

	return rel, nil
}
