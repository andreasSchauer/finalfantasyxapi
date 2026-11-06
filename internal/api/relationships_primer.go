package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getPrimerRelationships(r *http.Request, primer seedtypes.Primer) (apitypes.Primer, error) {
	var rel apitypes.Primer
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.dbGetRelations(ctx, primer, config.Cfg.DB.GetPrimerTreasureIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.dbGetRelations(ctx, primer, config.Cfg.DB.GetPrimerAreaIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Primer{}, err
	}

	return rel, nil
}
