package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSongRelationships(srv *Server, r *http.Request, song seedtypes.Song) (Song, error) {
	var rel Song
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.BossFights, err = E.MonsterFormations.dbGetRelations(srv, ctx, song, ToIntManyNull(config.Cfg.DB.GetSongMonsterFormationIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = E.FMVs.dbGetRelations(srv, ctx, song, ToIntManyNull(config.Cfg.DB.GetSongFmvIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Song{}, err
	}

	return rel, nil
}
