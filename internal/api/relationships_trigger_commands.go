package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getTriggerCommandRelationships(cfg *Config, r *http.Request, command seeding.TriggerCommand) (TriggerCommand, error) {
	var rel TriggerCommand
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.MonsterFormations, err = getResourcesDbItem(cfg, ctx, cfg.e.monsterFormations, command, cfg.db.GetTriggerCommandMonsterFormationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.UsedBy, err = getResourcesDbItem(cfg, ctx, cfg.e.characterClasses, command, cfg.db.GetTriggerCommandCharClassIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return TriggerCommand{}, err
	}

	return rel, nil
}