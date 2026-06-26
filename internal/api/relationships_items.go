package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getItemRelationships(cfg *Config, r *http.Request, item seeding.Item) (Item, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.items, item.ID)
	if err != nil {
		return Item{}, err
	}

	rel, err := runItemRelQueries(cfg, r, item, availabilityParams)
	if err != nil {
		return Item{}, err
	}

	if item.Category == string(database.ItemCategorySphere) {
		sphereRes := nameToNamedAPIResource(cfg, cfg.e.spheres, item.Name, nil)
		rel.Sphere = &sphereRes
	}

	return rel, nil
}

func runItemRelQueries(cfg *Config, r *http.Request, item seeding.Item, availabilityParams RelAvlParams) (Item, error) {
	var rel Item
	g, ctx := errgroup.WithContext(r.Context())
	
	g.Go(func() error{
		monsters, err := runRelAvailabilityQuery(cfg, ctx, cfg.e.monsters, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeMonster))
		if err != nil {
			return err
		}
		rel.Monsters = getMonItemAmts(cfg, monsters, item.Name)
		return nil
	})

	g.Go(func() error{
		treasures, err := runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeTreasure))
		if err != nil {
			return err
		}
		rel.Treasures = itemAmtsToChildResAmts2(cfg, cfg.e.treasures, treasures, item.ID)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error{
		quests, err := runRelAvailabilityQuery(cfg, ctx, cfg.e.quests, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeQuest))
		if err != nil {
			return err
		}
		rel.Quests = itemAmtsToChildResAmts(cfg.e.quests, quests)
		return nil
	})

	g.Go(func() error{
		blitzballPrizes, err := runRelAvailabilityQuery(cfg, ctx, cfg.e.blitzballPrizes, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeBlitzball))
		if err != nil {
			return err
		}
		rel.BlitzballPrizes = itemAmtsToChildResAmts2(cfg, cfg.e.blitzballPrizes, blitzballPrizes, item.ID)
		return nil
	})

	g.Go(func() error{
		playerAbilities, err := getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, item, cfg.db.GetItemPlayerAbilityIDs)
		if err != nil {
			return err
		}
		rel.AeonLearnAbilities = itemAmtsToChildResAmts(cfg.e.playerAbilities, playerAbilities)
		return nil
	})

	g.Go(func() error{
		autoAbilities, err := getResourcesDbItem(cfg, ctx, cfg.e.autoAbilities, item, cfg.db.GetItemAutoAbilityIDs)
		if err != nil {
			return err
		}
		rel.AutoAbilities = itemAmtsToChildResAmts(cfg.e.autoAbilities, autoAbilities)
		return nil
	})

	g.Go(func() error{
		var err error
		rel.Mixes, err = getResourcesDbItem(cfg, ctx, cfg.e.mixes, item, cfg.db.GetItemMixIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Item{}, err
	}

	return rel, nil
}

func getMonItemAmts(cfg *Config, monsters []NamedAPIResource, itemName string) []MonItemAmts {
	monItemAmts := []MonItemAmts{}

	for _, monster := range monsters {
		monItemAmt := createItemMonster(cfg, itemName, monster)
		monItemAmts = append(monItemAmts, monItemAmt)
	}

	return monItemAmts
}
