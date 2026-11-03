package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getAutoAbilityRelationships(cfg *Config, r *http.Request, autoAbility seedtypes.AutoAbility) (AutoAbility, error) {
	var rel AutoAbility
	g, ctx := errgroup.WithContext(r.Context())
	var monsterItems []NamedAPIResource

	availabilityParams, err := cfg.E.AutoAbilities.queryAvlRelGetParams(cfg, r, autoAbility.ID)
	if err != nil {
		return AutoAbility{}, err
	}

	g.Go(func() error {
		var err error
		monsterItems, err = cfg.E.Monsters.queryAvlRelRun(cfg, ctx, autoAbility, availabilityParams, getAutoAbilityItemMonsterIDs(cfg))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersDrop, err = cfg.E.Monsters.queryAvlRelRun(cfg, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeMonster, nil))
		return err
	})

	g.Go(func() error {
		var err error
		preAirship := database.ShopTypePreAirship
		rel.ShopsPreAirship, err = cfg.E.Shops.queryAvlRelRun(cfg, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeShop, &preAirship))
		return err
	})

	g.Go(func() error {
		var err error
		postAirship := database.ShopTypePostAirship
		rel.ShopsPostAirship, err = cfg.E.Shops.queryAvlRelRun(cfg, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeShop, &postAirship))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = cfg.E.Treasures.queryAvlRelRun(cfg, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeTreasure, nil))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTables, err = cfg.E.EquipmentTables.dbGetRelations(cfg, ctx, autoAbility, cfg.db.GetAutoAbilityEquipmentTableIDs)
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
