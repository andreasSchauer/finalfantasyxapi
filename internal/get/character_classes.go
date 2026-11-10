package get

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetCharacterClass(r *http.Request, i handlerinput.Main[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named], id int32) (apitypes.CharacterClass, error) {
	class, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	rel, err := getCharClassRelationships(r, class)
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	response := apitypes.CharacterClass{
		ID:                  class.ID,
		Name:                class.Name,
		Category:            class.Category,
		Members:             rel.Members,
		DefaultAbilities:    rel.DefaultAbilities,
		DefaultOverdrives:   rel.DefaultOverdrives,
		LearnableAbilities:  rel.LearnableAbilities,
		LearnableOverdrives: rel.LearnableOverdrives,
		Submenus:            rel.Submenus,
	}

	return response, nil
}

func getCharClassRelationships(r *http.Request, class seedtypes.CharacterClass) (apitypes.CharacterClass, error) {
	var rel apitypes.CharacterClass
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Members, err = apilookups.E.PlayerUnits.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultAbilities, err = apilookups.E.Abilities.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		rel.LearnableAbilities, err = getClassLearnableAbilities(ctx, class, rel.DefaultAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultOverdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.LearnableOverdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Submenus, err = apilookups.E.Submenus.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassSubmenuIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	return rel, nil
}

func getClassLearnableAbilities(ctx context.Context, class seedtypes.CharacterClass, defaultAbilities []apiresource.Typed) ([]apiresource.Typed, error) {
	allAbilities, err := apilookups.E.Abilities.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := handlerinput.RemoveResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
