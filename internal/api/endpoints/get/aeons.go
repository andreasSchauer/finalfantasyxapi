package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconv"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetAeon(r *http.Request, i handlerinput.Main[seedtypes.Aeon, apitypes.Aeon, apiresource.Named], id int32) (apitypes.Aeon, error) {
	aeon, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	rel, err := getAeonRelationships(r, aeon)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	response := apitypes.Aeon{
		ID:                     aeon.ID,
		Name:                   aeon.Name,
		UntypedUnit:            handlerinput.IDToTypedAPIResource(apilookups.E.PlayerUnits, aeon.PlayerUnit.ID),
		Area:                   handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, aeon.LocationArea),
		UnlockCondition:        aeon.UnlockCondition,
		IsOptional:             aeon.IsOptional,
		BattlesToRegenerate:    aeon.BattlesToRegenerate,
		PhysAtkDmgConstant:     aeon.PhysAtkDmgConstant,
		PhysAtkRange:           aeon.PhysAtkRange,
		PhysAtkShatterRate:     aeon.PhysAtkShatterRate,
		PhysAtkAccuracy:        ah.ConvertObjPtr(aeon.PhysAtkAccuracy, apiconv.ConvertAccuracy),
		CelestialWeapon:        rel.CelestialWeapon,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              th.GetAeonBaseStats(aeon),
		AeonCommands:           rel.AeonCommands,
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		Overdrives:             rel.Overdrives,
		WeaponAbilities:        ah.ConvertObjSlice(aeon.Weapon, apiconv.ConvertAeonEquipment),
		ArmorAbilities:         ah.ConvertObjSlice(aeon.Armor, apiconv.ConvertAeonEquipment),
	}

	return CompleteAeonResponse(r, response)
}

func getAeonRelationships(r *http.Request, aeon seedtypes.Aeon) (apitypes.Aeon, error) {
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

func CompleteAeonResponse(r *http.Request, aeon apitypes.Aeon) (apitypes.Aeon, error) {
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
