package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getCharacterRelationships(srv *Server, r *http.Request, char seedtypes.Character) (Character, error) {
	rel := Character{
		OverdriveModes: getCharacterModeAmts(srv, char),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = E.CelestialWeapons.dbGetRelationPtr(srv, ctx, char, srv.Cfg.DB.GetCharacterCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommand, err = E.OverdriveCommands.dbGetRelationPtr(srv, ctx, char, srv.Cfg.DB.GetCharacterOverdriveCommandID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = E.CharacterClasses.dbGetRelations(srv, ctx, char, srv.Cfg.DB.GetCharacterCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = E.PlayerAbilities.dbGetRelations(srv, ctx, char, srv.Cfg.DB.GetCharacterDefaultAbilityIDs)
		return err
	})

	if !char.IsStoryBased {
		g.Go(func() error {
			var err error
			rel.StdSphereGrid = convertObjPtr(srv, char.StdSphereGrid, convertSphereGrid)
			rel.StdSphereGrid.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(srv, ctx, char, srv.Cfg.DB.GetCharacterSgAbilityIDs)
			return err
		})

		g.Go(func() error {
			var err error
			rel.ExpSphereGrid = convertObjPtr(srv, char.ExpSphereGrid, convertSphereGrid)
			rel.ExpSphereGrid.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(srv, ctx, char, srv.Cfg.DB.GetCharacterEgAbilityIDs)
			return err
		})
	}

	err := g.Wait()
	if err != nil {
		return Character{}, err
	}

	return rel, nil
}

func getCharacterClstlFormula(char Character) *string {
	if char.CelestialWeapon == nil {
		return nil
	}

	weapon, _ := lookup.GetResource(char.CelestialWeapon.Name, lookup.L.CelestialWeapons)
	return &weapon.Formula
}

func getCharacterModeAmts(srv *Server, char seedtypes.Character) []ResourceAmount[NamedAPIResource] {
	resAmts := []ResourceAmount[NamedAPIResource]{}

	if char.IsStoryBased {
		return resAmts
	}

	for i := range len(lookup.L.OverdriveModes) {
		id := int32(i + 1)
		resAmt, err := getCharModeAmount(srv, char, id)
		if errors.Is(err, errs.ErrContinue) {
			continue
		}
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

func getCharModeAmount(srv *Server, char seedtypes.Character, id int32) (ResourceAmount[NamedAPIResource], error) {
	i := E.OverdriveModes

	mode, _ := lookup.GetResourceByID(id, i.objLookupID)
	if len(mode.ActionsToLearn) == 0 {
		return ResourceAmount[NamedAPIResource]{}, errs.ErrContinue
	}

	amount := mode.ActionsToLearn[char.GetID()-1].Amount
	modeAmount := i.toResAmt(srv, id, amount)

	return modeAmount, nil
}

func applyOsgStats(r *http.Request, char Character) ([]BaseStat, error) {
	if char.IsStoryBased {
		return char.BaseStats, nil
	}

	var statMap map[string]int32
	queryParam := Q.characters[ts.QpnOsgStats]

	sgType, err := queryParam.parseEnum(r, enums.T.SphereGridType)
	if queryIsEmpty(err) {
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
