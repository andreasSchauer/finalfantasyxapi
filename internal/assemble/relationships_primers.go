package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetPrimerRelationships(r *http.Request, primer seedtypes.Primer) (apitypes.Primer, error) {
	keyItem, _ := lookup.GetResource(primer, lookup.L.KeyItems)

	rel := apitypes.Primer{
		KeyItem:     handlerinput.NameToNamedAPIResource(apilookups.E.KeyItems, keyItem.Name, nil),
		Description: keyItem.Description,
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Treasures, err = apilookups.E.Treasures.DBGetRelations(ctx, primer, config.Cfg.DB.GetPrimerTreasureIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = apilookups.E.Areas.DBGetRelations(ctx, primer, config.Cfg.DB.GetPrimerAreaIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Primer{}, err
	}

	return rel, nil
}
