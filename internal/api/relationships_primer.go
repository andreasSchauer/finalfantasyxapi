package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getPrimerRelationships(srv *Server, r *http.Request, primer seedtypes.Primer) (Primer, error) {
	var rel Primer
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.dbGetRelations(srv, ctx, primer, config.Cfg.DB.GetPrimerTreasureIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.dbGetRelations(srv, ctx, primer, config.Cfg.DB.GetPrimerAreaIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Primer{}, err
	}

	return rel, nil
}
