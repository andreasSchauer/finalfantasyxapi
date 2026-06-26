package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getCharacterRelationships(cfg *Config, r *http.Request, char seeding.Character) (Character, error) {
	rel := Character{
		OverdriveModes: getCharacterModeAmts(cfg, char),
	}
	g, ctx := errgroup.WithContext(r.Context())
	
	g.Go(func() error{
		var err error
		rel.CelestialWeapon, err = getResPtrDB(cfg, ctx, cfg.e.celestialWeapons, char, cfg.db.GetCharacterCelestialWeaponID)
		return err
	})

	g.Go(func() error{
		var err error
		rel.OverdriveCommand, err = getResPtrDB(cfg, ctx, cfg.e.overdriveCommands, char, cfg.db.GetCharacterOverdriveCommandID)
		return err
	})

	g.Go(func() error{
		var err error
		rel.CharacterClasses, err = getResourcesDbItem(cfg, ctx, cfg.e.characterClasses, char, cfg.db.GetCharacterCharClassIDs)
		return err
	})

	g.Go(func() error{
		var err error
		rel.DefaultPlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, char, cfg.db.GetCharacterDefaultAbilityIDs)
		return err
	})

	g.Go(func() error{
		var err error
		rel.StdSgPlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, char, cfg.db.GetCharacterSgAbilityIDs)
		return err
	})

	g.Go(func() error{
		var err error
		rel.ExpSgPlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, char, cfg.db.GetCharacterEgAbilityIDs)
		return err
	})


	err := g.Wait()
	if err != nil {
		return Character{}, err
	}

	return rel, nil
}

func getCharacterModeAmts(cfg *Config, char seeding.Character) []ResourceAmount[NamedAPIResource] {
	resAmts := []ResourceAmount[NamedAPIResource]{}

	if char.IsStoryBased {
		return resAmts
	}

	for i := range len(cfg.l.OverdriveModes) {
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
	i := cfg.e.overdriveModes

	mode, _ := seeding.GetResourceByID(id, i.objLookupID)
	if len(mode.ActionsToLearn) == 0 {
		return ResourceAmount[NamedAPIResource]{}, errContinue
	}

	amount := mode.ActionsToLearn[char.GetID()-1].Amount
	modeAmount := idAmountToResourceAmount(cfg, i, id, amount)

	return modeAmount, nil
}
