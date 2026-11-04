package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getAutoAbilityRelationships(srv *Server, r *http.Request, autoAbility seedtypes.AutoAbility) (AutoAbility, error) {
	var rel AutoAbility
	g, ctx := errgroup.WithContext(r.Context())
	var monsterItems []NamedAPIResource

	availabilityParams, err := E.AutoAbilities.queryAvlRelGetParams(srv, r, autoAbility.ID)
	if err != nil {
		return AutoAbility{}, err
	}

	g.Go(func() error {
		var err error
		monsterItems, err = E.Monsters.queryAvlRelRun(srv, ctx, autoAbility, availabilityParams, getAutoAbilityItemMonsterIDs(srv))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersDrop, err = E.Monsters.queryAvlRelRun(srv, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(srv, ViewSourceTypeMonster, nil))
		return err
	})

	g.Go(func() error {
		var err error
		preAirship := database.ShopTypePreAirship
		rel.ShopsPreAirship, err = E.Shops.queryAvlRelRun(srv, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(srv, ViewSourceTypeShop, &preAirship))
		return err
	})

	g.Go(func() error {
		var err error
		postAirship := database.ShopTypePostAirship
		rel.ShopsPostAirship, err = E.Shops.queryAvlRelRun(srv, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(srv, ViewSourceTypeShop, &postAirship))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.queryAvlRelRun(srv, ctx, autoAbility, availabilityParams, getAutoAbilitySourceIDs(srv, ViewSourceTypeTreasure, nil))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EquipmentTables, err = E.EquipmentTables.dbGetRelations(srv, ctx, autoAbility, srv.Cfg.DB.GetAutoAbilityEquipmentTableIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return AutoAbility{}, err
	}

	if autoAbility.RequiredItem != nil {
		rel.MonstersItems = getMonItemAmts(monsterItems, autoAbility.RequiredItem.ItemName)
	}

	return rel, nil
}
