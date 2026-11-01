package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getItemRelationships(cfg *Config, r *http.Request, item seeding.Item) (Item, error) {
	availabilityParams, err := cfg.E.Items.queryAvlRelGetParams(cfg, r, item.ID)
	if err != nil {
		return Item{}, err
	}

	rel, err := runItemRelQueries(cfg, r, item, availabilityParams)
	if err != nil {
		return Item{}, err
	}

	if item.Category == string(database.ItemCategorySphere) {
		sphereRes := nameToNamedAPIResource(cfg, cfg.E.Spheres, item.Name, nil)
		rel.Sphere = &sphereRes
	}

	return rel, nil
}

func runItemRelQueries(cfg *Config, r *http.Request, item seeding.Item, availabilityParams RelAvlParams) (Item, error) {
	var rel Item
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		monsters, err := cfg.E.Monsters.queryAvlRelRun(cfg, ctx, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeMonster))
		if err != nil {
			return err
		}
		rel.Monsters = getMonItemAmts(cfg, monsters, item.Name)
		return nil
	})

	g.Go(func() error {
		treasures, err := cfg.E.Treasures.queryAvlRelRun(cfg, ctx, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeTreasure))
		if err != nil {
			return err
		}
		rel.Treasures = itemAmtsToChildResAmts2(cfg, cfg.E.Treasures, treasures, item.ID)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = cfg.E.Shops.queryAvlRelRun(cfg, ctx, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		quests, err := cfg.E.Quests.queryAvlRelRun(cfg, ctx, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeQuest))
		if err != nil {
			return err
		}
		rel.Quests = itemAmtsToChildResAmts(cfg.E.Quests, quests)
		return nil
	})

	g.Go(func() error {
		blitzballPrizes, err := cfg.E.BlitzballPrizes.queryAvlRelRun(cfg, ctx, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeBlitzball))
		if err != nil {
			return err
		}
		rel.BlitzballPrizes = itemAmtsToChildResAmts2(cfg, cfg.E.BlitzballPrizes, blitzballPrizes, item.ID)
		return nil
	})

	g.Go(func() error {
		playerAbilities, err := cfg.E.PlayerAbilities.dbGetRelations(cfg, ctx, item, cfg.db.GetItemPlayerAbilityIDs)
		if err != nil {
			return err
		}
		rel.AeonLearnAbilities = itemAmtsToChildResAmts(cfg.E.PlayerAbilities, playerAbilities)
		return nil
	})

	g.Go(func() error {
		autoAbilities, err := cfg.E.AutoAbilities.dbGetRelations(cfg, ctx, item, cfg.db.GetItemAutoAbilityIDs)
		if err != nil {
			return err
		}
		rel.AutoAbilities = itemAmtsToChildResAmts(cfg.E.AutoAbilities, autoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Mixes, err = cfg.E.Mixes.dbGetRelations(cfg, ctx, item, cfg.db.GetItemMixIDs)
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
