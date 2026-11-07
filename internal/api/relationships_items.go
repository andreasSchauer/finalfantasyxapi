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

func getItemRelationships(r *http.Request, item seedtypes.Item) (apitypes.Item, error) {
	availabilityParams, err := E.Items.QueryAvlRelGetParams(r, item.ID)
	if err != nil {
		return apitypes.Item{}, err
	}

	rel, err := runItemRelQueries(r, item, availabilityParams)
	if err != nil {
		return apitypes.Item{}, err
	}

	if item.Category == string(database.ItemCategorySphere) {
		sphereRes := NameToNamedAPIResource(E.Spheres, item.Name, nil)
		rel.Sphere = &sphereRes
	}

	return rel, nil
}

func runItemRelQueries(r *http.Request, item seedtypes.Item, availabilityParams RelAvlParams) (apitypes.Item, error) {
	var rel apitypes.Item
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		monsters, err := E.Monsters.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ViewSourceTypeMonster))
		if err != nil {
			return err
		}
		rel.Monsters = getMonItemAmts(monsters, item.Name)
		return nil
	})

	g.Go(func() error {
		treasures, err := E.Treasures.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ViewSourceTypeTreasure))
		if err != nil {
			return err
		}
		rel.Treasures = ItemAmtsToChildResAmts2(E.Treasures, treasures, item.ID)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = E.Shops.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		quests, err := E.Quests.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ViewSourceTypeQuest))
		if err != nil {
			return err
		}
		rel.Quests = ItemAmtsToChildResAmts(E.Quests, quests)
		return nil
	})

	g.Go(func() error {
		blitzballPrizes, err := E.BlitzballPrizes.QueryAvlRelRun(ctx, item, availabilityParams, getItemSourceIDs(ViewSourceTypeBlitzball))
		if err != nil {
			return err
		}
		rel.BlitzballPrizes = ItemAmtsToChildResAmts2(E.BlitzballPrizes, blitzballPrizes, item.ID)
		return nil
	})

	g.Go(func() error {
		playerAbilities, err := E.PlayerAbilities.DBGetRelations(ctx, item, config.Cfg.DB.GetItemPlayerAbilityIDs)
		if err != nil {
			return err
		}
		rel.AeonLearnAbilities = ItemAmtsToChildResAmts(E.PlayerAbilities, playerAbilities)
		return nil
	})

	g.Go(func() error {
		autoAbilities, err := E.AutoAbilities.DBGetRelations(ctx, item, config.Cfg.DB.GetItemAutoAbilityIDs)
		if err != nil {
			return err
		}
		rel.AutoAbilities = ItemAmtsToChildResAmts(E.AutoAbilities, autoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Mixes, err = E.Mixes.DBGetRelations(ctx, item, config.Cfg.DB.GetItemMixIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Item{}, err
	}

	return rel, nil
}

func getMonItemAmts(monsters []apiresource.NamedAPIResource, itemName string) []apitypes.MonItemAmts {
	monItemAmts := []apitypes.MonItemAmts{}

	for _, monster := range monsters {
		monItemAmt := createItemMonster(itemName, monster)
		monItemAmts = append(monItemAmts, monItemAmt)
	}

	return monItemAmts
}
