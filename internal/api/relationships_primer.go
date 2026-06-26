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
		rel.Treasures, err = getResourcesDbItem(cfg, ctx, cfg.e.treasures, primer, cfg.db.GetPrimerTreasureIDs)
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.Areas, err = getResourcesDbItem(cfg, ctx, cfg.e.areas, primer, cfg.db.GetPrimerAreaIDs)
		return err
	})
	
	err := g.Wait()
	if err != nil {
		return Primer{}, err
	}

	return rel, nil
}
