package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetAeonRelationships(r *http.Request, aeon seedtypes.Aeon) (apitypes.Aeon, error) {
	var rel apitypes.Aeon
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = apilookups.E.CelestialWeapons.DBGetRelationPtr(ctx, aeon, config.Cfg.DB.GetAeonCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = apilookups.E.CharacterClasses.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = apilookups.E.AeonCommands.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonAeonCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonDefaultAbilityIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Aeon{}, err
	}

	return rel, nil
}

func CompleteAeon2Response(r *http.Request, aeon apitypes.Aeon) (apitypes.Aeon, error) {
	aeon, err := applyAeonStats(r, aeon, nil)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	aeon.RegenAmounts = calcRegenAmounts(aeon.BaseStats)

	aeon.AgilityParameters, err = getUnitAgilityParams(r, aeon.BaseStats)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	return aeon, nil
}
