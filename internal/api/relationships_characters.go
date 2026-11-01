package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getCharacterRelationships(cfg *Config, r *http.Request, char seeding.Character) (Character, error) {
	rel := Character{
		OverdriveModes: getCharacterModeAmts(cfg, char),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = cfg.E.CelestialWeapons.dbGetRelationPtr(cfg, ctx, char, cfg.db.GetCharacterCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommand, err = cfg.E.OverdriveCommands.dbGetRelationPtr(cfg, ctx, char, cfg.db.GetCharacterOverdriveCommandID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = cfg.E.CharacterClasses.dbGetRelations(cfg, ctx, char, cfg.db.GetCharacterCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = cfg.E.PlayerAbilities.dbGetRelations(cfg, ctx, char, cfg.db.GetCharacterDefaultAbilityIDs)
		return err
	})

	if !char.IsStoryBased {
		g.Go(func() error {
			var err error
			rel.StdSphereGrid = convertObjPtr(cfg, char.StdSphereGrid, convertSphereGrid)
			rel.StdSphereGrid.PlayerAbilities, err = cfg.E.PlayerAbilities.dbGetRelations(cfg, ctx, char, cfg.db.GetCharacterSgAbilityIDs)
			return err
		})

		g.Go(func() error {
			var err error
			rel.ExpSphereGrid = convertObjPtr(cfg, char.ExpSphereGrid, convertSphereGrid)
			rel.ExpSphereGrid.PlayerAbilities, err = cfg.E.PlayerAbilities.dbGetRelations(cfg, ctx, char, cfg.db.GetCharacterEgAbilityIDs)
			return err
		})
	}

	err := g.Wait()
	if err != nil {
		return Character{}, err
	}

	return rel, nil
}

func getCharacterClstlFormula(cfg *Config, char Character) *string {
	if char.CelestialWeapon == nil {
		return nil
	}

	weapon, _ := seeding.GetResource(char.CelestialWeapon.Name, cfg.L.CelestialWeapons)
	return &weapon.Formula
}

func getCharacterModeAmts(cfg *Config, char seeding.Character) []ResourceAmount[NamedAPIResource] {
	resAmts := []ResourceAmount[NamedAPIResource]{}

	if char.IsStoryBased {
		return resAmts
	}

	for i := range len(cfg.L.OverdriveModes) {
		id := int32(i + 1)
		resAmt, err := getCharModeAmount(cfg, char, id)
		if errors.Is(err, errContinue) {
			continue
		}
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

func getCharModeAmount(cfg *Config, char seeding.Character, id int32) (ResourceAmount[NamedAPIResource], error) {
	i := cfg.E.OverdriveModes

	mode, _ := seeding.GetResourceByID(id, i.objLookupID)
	if len(mode.ActionsToLearn) == 0 {
		return ResourceAmount[NamedAPIResource]{}, errContinue
	}

	amount := mode.ActionsToLearn[char.GetID()-1].Amount
	modeAmount := i.toResAmt(cfg, id, amount)

	return modeAmount, nil
}

func applyOsgStats(cfg *Config, r *http.Request, char Character) ([]BaseStat, error) {
	if char.IsStoryBased {
		return char.BaseStats, nil
	}

	var statMap map[string]int32
	queryParam := cfg.Q.characters[QpnOsgStats]

	sgType, err := parseEnumQuery(r, epCharacters, queryParam, cfg.T.SphereGridType)
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
