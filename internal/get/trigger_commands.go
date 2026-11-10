package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetTriggerCommand(r *http.Request, i handlerinput.Main[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.Named], id int32) (apitypes.TriggerCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}

	rel, err := getTriggerCommandRelationships(r, command)
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}

	response := apitypes.TriggerCommand{
		ID:                 command.ID,
		Name:               command.Name,
		Version:            command.Version,
		Specification:      command.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, command.Ability.ID),
		Description:        command.Description,
		Effect:             command.Effect,
		Rank:               command.Rank,
		AppearsInHelpBar:   command.AppearsInHelpBar,
		CanCopycat:         command.CanCopycat,
		UsedBy:             rel.UsedBy,
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, command.RelatedStats),
		Topmenu:            handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, command.Topmenu, nil),
		Cursor:             command.Cursor,
		MonsterFormations:  rel.MonsterFormations,
		BattleInteractions: ah.ConvertObjSlice(command.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	battleInteractions, err := ApplyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.TriggerCommand{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func getTriggerCommandRelationships(r *http.Request, command seedtypes.TriggerCommand) (apitypes.TriggerCommand, error) {
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
