package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getPropertyRelationships(cfg *Config, r *http.Request, property seedtypes.Property) (Property, error) {
	var rel Property
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = cfg.E.AutoAbilities.dbGetRelations(cfg, ctx, property, ToIntManyNull(cfg.db.GetPropertyAutoAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = cfg.E.Monsters.dbGetRelations(cfg, ctx, property, cfg.db.GetPropertyMonsterIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Property{}, err
	}

	return rel, nil
}
