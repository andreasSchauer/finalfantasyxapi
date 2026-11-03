package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getTriggerCommandRelationships(cfg *Config, r *http.Request, command seedtypes.TriggerCommand) (TriggerCommand, error) {
	var rel TriggerCommand
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.MonsterFormations, err = cfg.E.MonsterFormations.dbGetRelations(cfg, ctx, command, cfg.db.GetTriggerCommandMonsterFormationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.UsedBy, err = cfg.E.CharacterClasses.dbGetRelations(cfg, ctx, command, cfg.db.GetTriggerCommandCharClassIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return TriggerCommand{}, err
	}

	return rel, nil
}
