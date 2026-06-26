package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getAutoAbilityRelationships(cfg *Config, r *http.Request, autoAbility seeding.AutoAbility) (AutoAbility, error) {
	var rel AutoAbility
	g, ctx := errgroup.WithContext(r.Context())
	var monsterItems []NamedAPIResource
	
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.autoAbilities, autoAbility.ID)
	if err != nil {
		return AutoAbility{}, err
	}

	g.Go(func() error{
		var err error
		monsterItems, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsters, autoAbility, availabilityParams, getAutoAbilityItemMonsterIDs(cfg))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.MonstersDrop, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsters, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeMonster, nil))
		return err
	})

	g.Go(func() error{
		var err error
		preAirship := database.ShopTypePreAirship
		rel.ShopsPreAirship, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeShop, &preAirship))
		return err
	})
	
	g.Go(func() error{
		var err error
		postAirship := database.ShopTypePostAirship
		rel.ShopsPostAirship, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeShop, &postAirship))
		return err
	})

	g.Go(func() error{
		var err error
		rel.Treasures, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeTreasure, nil))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.EquipmentTables, err = getResourcesDbItem(cfg, ctx, cfg.e.equipmentTables, autoAbility, cfg.db.GetAutoAbilityEquipmentTableIDs)
		return err
	})
	
	err = g.Wait()
	if err != nil {
		return AutoAbility{}, err
	}

	if autoAbility.RequiredItem != nil {
		rel.MonstersItems = getMonItemAmts(cfg, monsterItems, autoAbility.RequiredItem.ItemName)
	}

	return rel, nil
}
