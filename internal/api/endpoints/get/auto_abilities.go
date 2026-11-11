package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func AutoAbility(r *http.Request, i handlerinput.Main[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named], id int32) (apitypes.AutoAbility, error) {
	autoAbility, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	rel, err := getAutoAbilityRelationships(r, autoAbility)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	response := apitypes.AutoAbility{
		ID:                     autoAbility.ID,
		Name:                   autoAbility.Name,
		Description:            autoAbility.Description,
		Effect:                 autoAbility.Effect,
		EquipType:              autoAbility.Type,
		Category:               autoAbility.Category,
		RelatedStats:           handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           apilookups.E.AllItems.ToResAmtFromNamePtr(autoAbility.RequiredItem),
		LockedOutAutoAbilities: handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            handlerinput.NamesToNamedAPIResources(apilookups.E.Items, autoAbility.AutoItemUse),
		OnHitElement:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        ah.ConvertObjPtr(autoAbility.AddedElemResist, apiconvert.ElemResist),
		OnHitStatus:            ah.ConvertObjPtr(autoAbility.OnHitStatus, apiconvert.InflictedStatus),
		AddedStatusses:         handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     ah.ConvertObjSlice(autoAbility.AddedStatusResists, apiconvert.StatusResist),
		AddedProperty:          handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Properties, autoAbility.AddedProperty, nil),
		ConversionTo:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Modifiers, autoAbility.ConversionTo, nil),
		StatChanges:            ah.ConvertObjSlice(autoAbility.StatChanges, apiconvert.StatChange),
		ModifierChanges:        ah.ConvertObjSlice(autoAbility.ModifierChanges, apiconvert.ModifierChange),
		MonstersDrop:           rel.MonstersDrop,
		MonstersItems:          rel.MonstersItems,
		ShopsPreAirship:        rel.ShopsPreAirship,
		ShopsPostAirship:       rel.ShopsPostAirship,
		Treasures:              rel.Treasures,
		EquipmentTables:        rel.EquipmentTables,
	}

	return response, nil
}

func getAutoAbilityRelationships(r *http.Request, autoAbility seedtypes.AutoAbility) (apitypes.AutoAbility, error) {
	var rel apitypes.AutoAbility
	g, ctx := errgroup.WithContext(r.Context())
	var monsterItems []apiresource.Named

	availabilityParams, err := apilookups.E.AutoAbilities.QueryAvlRelGetParams(r, autoAbility.ID)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	g.Go(func() error {
		var err error
		monsterItems, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilityItemMonsterIDs())
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersDrop, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ts.VstMonster, nil))
		return err
	})

	g.Go(func() error {
		var err error
		preAirship := database.ShopTypePreAirship
		rel.ShopsPreAirship, err = apilookups.E.Shops.QueryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ts.VstShop, &preAirship))
		return err
	})

	g.Go(func() error {
		var err error
		postAirship := database.ShopTypePostAirship
		rel.ShopsPostAirship, err = apilookups.E.Shops.QueryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ts.VstShop, &postAirship))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ts.VstTreasure, nil))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTables, err = apilookups.E.EquipmentTables.DBGetRelations(ctx, autoAbility, config.Cfg.DB.GetAutoAbilityEquipmentTableIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	if autoAbility.RequiredItem != nil {
		rel.MonstersItems = getMonItemAmts(monsterItems, autoAbility.RequiredItem.ItemName)
	}

	return rel, nil
}
