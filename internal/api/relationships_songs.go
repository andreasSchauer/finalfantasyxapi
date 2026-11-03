package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSongRelationships(cfg *Config, r *http.Request, song seedtypes.Song) (Song, error) {
	var rel Song
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.BossFights, err = cfg.E.MonsterFormations.dbGetRelations(cfg, ctx, song, ToIntManyNull(cfg.db.GetSongMonsterFormationIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = cfg.E.FMVs.dbGetRelations(cfg, ctx, song, ToIntManyNull(cfg.db.GetSongFmvIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Song{}, err
	}

	return rel, nil
}
