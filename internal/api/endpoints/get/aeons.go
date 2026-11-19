package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Aeon(r *http.Request, i handlerinput.Main[seedtypes.Aeon, apitypes.Aeon, apiresource.Named], id int32) (apitypes.Aeon, error) {
	aeon, err := i.VerifyQueriesGet(r, id)
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
		PhysAtkAccuracy:        ah.ConvertObjPtr(aeon.PhysAtkAccuracy, apiconvert.Accuracy),
		BaseStats:              th.GetAeonBaseStats(aeon),
		WeaponAbilities:        ah.ConvertObjSlice(aeon.Weapon, apiconvert.AeonEquipment),
		ArmorAbilities:         ah.ConvertObjSlice(aeon.Armor, apiconvert.AeonEquipment),
	}

	err = getAeonRelationships(r, aeon, &response)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	err = completeAeonResponse(r, &response)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	return response, nil
}

func getAeonRelationships(r *http.Request, aeon seedtypes.Aeon, response *apitypes.Aeon) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.CelestialWeapon, err = apilookups.E.CelestialWeapons.DBGetRelationPtr(ctx, aeon, config.Cfg.DB.GetAeonCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		response.CharacterClasses, err = apilookups.E.CharacterClasses.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.AeonCommands, err = apilookups.E.AeonCommands.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonAeonCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.DefaultPlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonDefaultAbilityIDs)
		return err
	})

	return g.Wait()
}

func completeAeonResponse(r *http.Request, aeon *apitypes.Aeon) error {
	var err error

	*aeon, err = applyAeonStats(r, *aeon, nil)
	if err != nil {
		return err
	}

	aeon.RegenAmounts = calcRegenAmounts(aeon.BaseStats)

	aeon.AgilityParameters, err = getUnitAgilityParams(r, aeon.BaseStats)
	if err != nil {
		return err
	}

	return nil
}
