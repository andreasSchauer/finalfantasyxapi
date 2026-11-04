package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getPropertyRelationships(server *Server, r *http.Request, property seedtypes.Property) (Property, error) {
	var rel Property
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(server, ctx, property, ToIntManyNull(server.Cfg.DB.GetPropertyAutoAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = E.Monsters.dbGetRelations(server, ctx, property, server.Cfg.DB.GetPropertyMonsterIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Property{}, err
	}

	return rel, nil
}
