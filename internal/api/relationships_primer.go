package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getPrimerRelationships(cfg *Config, r *http.Request, primer seeding.Primer) (Primer, error) {
	var rel Primer
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Treasures, err = cfg.E.Treasures.dbGetRelations(cfg, ctx, primer, cfg.db.GetPrimerTreasureIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = cfg.E.Areas.dbGetRelations(cfg, ctx, primer, cfg.db.GetPrimerAreaIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Primer{}, err
	}

	return rel, nil
}
