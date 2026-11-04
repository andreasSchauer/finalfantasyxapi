package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getTriggerCommandRelationships(server *Server, r *http.Request, command seedtypes.TriggerCommand) (TriggerCommand, error) {
	var rel TriggerCommand
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.MonsterFormations, err = E.MonsterFormations.dbGetRelations(server, ctx, command, server.Cfg.DB.GetTriggerCommandMonsterFormationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.UsedBy, err = E.CharacterClasses.dbGetRelations(server, ctx, command, server.Cfg.DB.GetTriggerCommandCharClassIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return TriggerCommand{}, err
	}

	return rel, nil
}
