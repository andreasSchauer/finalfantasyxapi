package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getPropertyRelationships(cfg *Config, r *http.Request, property seeding.Property) (Property, error) {
	var rel Property
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.autoAbilities, property, ToIntManyNull(cfg.db.GetPropertyAutoAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = getResourcesDbItem(cfg, ctx, cfg.e.monsters, property, cfg.db.GetPropertyMonsterIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Property{}, err
	}

	return rel, nil
}