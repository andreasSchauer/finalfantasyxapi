package get

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func CharacterClass(r *http.Request, i handlerinput.Main[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named], id int32) (apitypes.CharacterClass, error) {
	class, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	response := apitypes.CharacterClass{
		ID:                  class.ID,
		Name:                class.Name,
		Category:            class.Category,
	}

	err = getCharClassRelationships(r, class, &response)
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	return response, nil
}

func getCharClassRelationships(r *http.Request, class seedtypes.CharacterClass, response *apitypes.CharacterClass) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.Members, err = apilookups.E.PlayerUnits.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.DefaultAbilities, err = apilookups.E.Abilities.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		response.LearnableAbilities, err = getClassLearnableAbilities(ctx, class, response.DefaultAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		response.DefaultOverdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.LearnableOverdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Submenus, err = apilookups.E.Submenus.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassSubmenuIDs)
		return err
	})

	return g.Wait()
}

func getClassLearnableAbilities(ctx context.Context, class seedtypes.CharacterClass, defaultAbilities []apiresource.Typed) ([]apiresource.Typed, error) {
	allAbilities, err := apilookups.E.Abilities.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := apiresource.RemoveResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
