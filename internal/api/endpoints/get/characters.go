package get

import (
	"errors"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"golang.org/x/sync/errgroup"

	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Character(r *http.Request, i handlerinput.Main[seedtypes.Character, apitypes.Character, apiresource.Named], id int32) (apitypes.Character, error) {
	character, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Character{}, err
	}

	rel, err := getCharacterRelationships(r, character)
	if err != nil {
		return apitypes.Character{}, err
	}

	response := apitypes.Character{
		ID:                     character.ID,
		Name:                   character.Name,
		UntypedUnit:            handlerinput.IDToTypedAPIResource(apilookups.E.PlayerUnits, character.PlayerUnit.ID),
		Area:                   handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, character.LocationArea),
		IsStoryBased:           character.IsStoryBased,
		CanFightUnderwater:     character.CanFightUnderwater,
		PhysAtkRange:           character.PhysAtkRange,
		WeaponType:             character.WeaponType,
		ArmorType:              character.ArmorType,
		CelestialWeapon:        rel.CelestialWeapon,
		OverdriveCommand:       rel.OverdriveCommand,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              apilookups.E.Stats.ToResAmtType(character.BaseStats, th.NewBaseStat),
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		StdSphereGrid:          rel.StdSphereGrid,
		ExpSphereGrid:          rel.ExpSphereGrid,
		OverdriveModes:         rel.OverdriveModes,
	}

	return completeCharacterResponse(r, response)
}

func getCharacterRelationships(r *http.Request, char seedtypes.Character) (apitypes.Character, error) {
	rel := apitypes.Character{
		OverdriveModes: getCharacterModeAmts(char),
	}
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = apilookups.E.CelestialWeapons.DBGetRelationPtr(ctx, char, config.Cfg.DB.GetCharacterCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommand, err = apilookups.E.OverdriveCommands.DBGetRelationPtr(ctx, char, config.Cfg.DB.GetCharacterOverdriveCommandID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = apilookups.E.CharacterClasses.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterDefaultAbilityIDs)
		return err
	})

	if !char.IsStoryBased {
		g.Go(func() error {
			var err error
			rel.StdSphereGrid = ah.ConvertObjPtr(char.StdSphereGrid, apiconvert.SphereGrid)
			rel.StdSphereGrid.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterSgAbilityIDs)
			return err
		})

		g.Go(func() error {
			var err error
			rel.ExpSphereGrid = ah.ConvertObjPtr(char.ExpSphereGrid, apiconvert.SphereGrid)
			rel.ExpSphereGrid.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, char, config.Cfg.DB.GetCharacterEgAbilityIDs)
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

func getCharacterModeAmts(char seedtypes.Character) []apiresource.ResourceAmount[apiresource.Named] {
	resAmts := []apiresource.ResourceAmount[apiresource.Named]{}

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

func getCharModeAmount(char seedtypes.Character, id int32) (apiresource.ResourceAmount[apiresource.Named], error) {
	i := apilookups.E.OverdriveModes

	mode, _ := lookup.GetResourceByID(id, i.ObjLookupID)
	if len(mode.ActionsToLearn) == 0 {
		return apiresource.ResourceAmount[apiresource.Named]{}, errs.ErrContinue
	}

	amount := mode.ActionsToLearn[char.GetID()-1].Amount
	modeAmount := i.ToResAmt(id, amount)

	return modeAmount, nil
}

func completeCharacterResponse(r *http.Request, char apitypes.Character) (apitypes.Character, error) {
	char, err := applyOsgStats(r, char)
	if err != nil {
		return apitypes.Character{}, err
	}

	char.RegenAmounts = calcRegenAmounts(char.BaseStats)
	char.CelestialFormula = getCharacterClstlFormula(char)

	char.AgilityParameters, err = getUnitAgilityParams(r, char.BaseStats)
	if err != nil {
		return apitypes.Character{}, err
	}

	return char, nil
}
