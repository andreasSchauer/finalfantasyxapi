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

func Primer(r *http.Request, i handlerinput.Main[seedtypes.Primer, apitypes.Primer, apiresource.Named], id int32) (apitypes.Primer, error) {
	primer, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Primer{}, err
	}
	
	response := apitypes.Primer{
		ID:            primer.ID,
		Name:          primer.Name,
		AlBhedLetter:  primer.AlBhedLetter,
		EnglishLetter: primer.EnglishLetter,
	}

	err = getPrimerRelationships(r, primer, &response)
	if err != nil {
		return apitypes.Primer{}, err
	}

	return response, nil
}

func getPrimerRelationships(r *http.Request, primer seedtypes.Primer, response *apitypes.Primer) error {
	keyItem, _ := lookup.GetResource(primer, lookup.L.KeyItems)

	response.KeyItem = handlerinput.NameToNamedAPIResource(apilookups.E.KeyItems, keyItem.Name, nil)
	response.Description = keyItem.Description

	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.Treasures, err = apilookups.E.Treasures.DBGetRelations(ctx, primer, config.Cfg.DB.GetPrimerTreasureIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Areas, err = apilookups.E.Areas.DBGetRelations(ctx, primer, config.Cfg.DB.GetPrimerAreaIDs)
		return err
	})

	return g.Wait()
}
