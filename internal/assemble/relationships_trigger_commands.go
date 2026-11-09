package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetTriggerCommandRelationships(r *http.Request, command seedtypes.TriggerCommand) (apitypes.TriggerCommand, error) {
	var rel apitypes.TriggerCommand
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.MonsterFormations, err = apilookups.E.MonsterFormations.DBGetRelations(ctx, command, config.Cfg.DB.GetTriggerCommandMonsterFormationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.UsedBy, err = apilookups.E.CharacterClasses.DBGetRelations(ctx, command, config.Cfg.DB.GetTriggerCommandCharClassIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}

	return rel, nil
}
