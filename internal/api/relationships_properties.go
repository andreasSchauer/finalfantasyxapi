package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getPropertyRelationships(r *http.Request, property seedtypes.Property) (apitypes.Property, error) {
	var rel apitypes.Property
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, property, handlerinput.ToIntManyNull(config.Cfg.DB.GetPropertyAutoAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(ctx, property, config.Cfg.DB.GetPropertyMonsterIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Property{}, err
	}

	return rel, nil
}
