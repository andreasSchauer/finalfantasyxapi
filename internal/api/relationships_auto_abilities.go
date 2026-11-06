package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getAutoAbilityRelationships(r *http.Request, autoAbility seedtypes.AutoAbility) (apitypes.AutoAbility, error) {
	var rel apitypes.AutoAbility
	g, ctx := errgroup.WithContext(r.Context())
	var monsterItems []apiresource.NamedAPIResource

	availabilityParams, err := E.AutoAbilities.queryAvlRelGetParams(r, autoAbility.ID)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	g.Go(func() error {
		var err error
		monsterItems, err = E.Monsters.queryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilityItemMonsterIDs())
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersDrop, err = E.Monsters.queryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ViewSourceTypeMonster, nil))
		return err
	})

	g.Go(func() error {
		var err error
		preAirship := database.ShopTypePreAirship
		rel.ShopsPreAirship, err = E.Shops.queryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ViewSourceTypeShop, &preAirship))
		return err
	})

	g.Go(func() error {
		var err error
		postAirship := database.ShopTypePostAirship
		rel.ShopsPostAirship, err = E.Shops.queryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ViewSourceTypeShop, &postAirship))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.queryAvlRelRun(ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(ViewSourceTypeTreasure, nil))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTables, err = E.EquipmentTables.dbGetRelations(ctx, autoAbility, config.Cfg.DB.GetAutoAbilityEquipmentTableIDs)
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
