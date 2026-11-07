package api

import (
	"errors"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getCharacterRelationships(r *http.Request, char seedtypes.Character) (apitypes.Character, error) {
	rel := apitypes.Character{
		OverdriveModes: getCharacterModeAmts(char),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = E.CelestialWeapons.DBGetRelationPtr(ctx, char, config.Cfg.DB.GetCharacterCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommand, err = E.OverdriveCommands.DBGetRelationPtr(ctx, char, config.Cfg.DB.GetCharacterOverdriveCommandID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = E.CharacterClasses.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = E.PlayerAbilities.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterDefaultAbilityIDs)
		return err
	})

	if !char.IsStoryBased {
		g.Go(func() error {
			var err error
			rel.StdSphereGrid = ah.ConvertObjPtr(char.StdSphereGrid, convertSphereGrid)
			rel.StdSphereGrid.PlayerAbilities, err = E.PlayerAbilities.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterSgAbilityIDs)
			return err
		})

		g.Go(func() error {
			var err error
			rel.ExpSphereGrid = ah.ConvertObjPtr(char.ExpSphereGrid, convertSphereGrid)
			rel.ExpSphereGrid.PlayerAbilities, err = E.PlayerAbilities.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterEgAbilityIDs)
			return err
		})
	}

	err := g.Wait()
	if err != nil {
		return apitypes.Character{}, err
	}

	return rel, nil
}

func getCharacterClstlFormula(char apitypes.Character) *string {
	if char.CelestialWeapon == nil {
		return nil
	}

	weapon, _ := lookup.GetResource(char.CelestialWeapon.Name, lookup.L.CelestialWeapons)
	return &weapon.Formula
}

func getCharacterModeAmts(char seedtypes.Character) []apiresource.ResourceAmount[apiresource.NamedAPIResource] {
	resAmts := []apiresource.ResourceAmount[apiresource.NamedAPIResource]{}

	if char.IsStoryBased {
		return resAmts
	}

	for i := range len(lookup.L.OverdriveModes) {
		id := int32(i + 1)
		resAmt, err := getCharModeAmount(char, id)
		if errors.Is(err, errs.ErrContinue) {
			continue
		}
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

func getCharModeAmount(char seedtypes.Character, id int32) (apiresource.ResourceAmount[apiresource.NamedAPIResource], error) {
	i := E.OverdriveModes

	mode, _ := lookup.GetResourceByID(id, i.ObjLookupID)
	if len(mode.ActionsToLearn) == 0 {
		return apiresource.ResourceAmount[apiresource.NamedAPIResource]{}, errs.ErrContinue
	}

	amount := mode.ActionsToLearn[char.GetID()-1].Amount
	modeAmount := i.ToResAmt(id, amount)

	return modeAmount, nil
}

func applyOsgStats(r *http.Request, char apitypes.Character) ([]apitypes.BaseStat, error) {
	if char.IsStoryBased {
		return char.BaseStats, nil
	}

	var statMap map[string]int32
	queryParam := Q.Characters[ts.QpnOsgStats]

	sgType, err := queryParam.ParseEnum(r, enums.T.SphereGridType)
	if errs.QueryIsEmpty(err) {
		return char.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	switch sgType.Name {
	case string(database.SphereGridTypeStandard):
		statMap = char.StdSphereGrid.ToMap()

	case string(database.SphereGridTypeExpert):
		statMap = char.ExpSphereGrid.ToMap()
	}

	newStats := addToBaseStats(char.BaseStats, statMap)

	return newStats, nil
}
