package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getSongRelationships(cfg *Config, r *http.Request, song seeding.Song) (Song, error) {
	var rel Song
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.BossFights, err = getResourcesDBItem(cfg, ctx, cfg.E.MonsterFormations, song, ToIntManyNull(cfg.db.GetSongMonsterFormationIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = getResourcesDBItem(cfg, ctx, cfg.E.FMVs, song, ToIntManyNull(cfg.db.GetSongFmvIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Song{}, err
	}

	return rel, nil
}
