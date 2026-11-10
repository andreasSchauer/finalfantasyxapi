package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetPrimer(r *http.Request, i handlerinput.Main[seedtypes.Primer, apitypes.Primer, apiresource.Named], id int32) (apitypes.Primer, error) {
	primer, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Primer{}, err
	}

	rel, err := getPrimerRelationships(r, primer)
	if err != nil {
		return apitypes.Primer{}, err
	}

	response := apitypes.Primer{
		ID:            primer.ID,
		Name:          primer.Name,
		KeyItem:       rel.KeyItem,
		Description:   rel.Description,
		AlBhedLetter:  primer.AlBhedLetter,
		EnglishLetter: primer.EnglishLetter,
		Treasures:     rel.Treasures,
		Areas:         rel.Areas,
	}

	return response, nil
}

func getPrimerRelationships(r *http.Request, primer seedtypes.Primer) (apitypes.Primer, error) {
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
