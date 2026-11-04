package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getItemRelationships(srv *Server, r *http.Request, item seedtypes.Item) (Item, error) {
	availabilityParams, err := E.Items.queryAvlRelGetParams(srv, r, item.ID)
	if err != nil {
		return Item{}, err
	}

	rel, err := runItemRelQueries(srv, r, item, availabilityParams)
	if err != nil {
		return Item{}, err
	}

	if item.Category == string(database.ItemCategorySphere) {
		sphereRes := nameToNamedAPIResource(srv, E.Spheres, item.Name, nil)
		rel.Sphere = &sphereRes
	}

	return rel, nil
}

func runItemRelQueries(srv *Server, r *http.Request, item seedtypes.Item, availabilityParams RelAvlParams) (Item, error) {
	var rel Item
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		monsters, err := E.Monsters.queryAvlRelRun(srv, ctx, item, availabilityParams, getItemSourceIDs(srv, ViewSourceTypeMonster))
		if err != nil {
			return err
		}
		rel.Monsters = getMonItemAmts(monsters, item.Name)
		return nil
	})

	g.Go(func() error {
		treasures, err := E.Treasures.queryAvlRelRun(srv, ctx, item, availabilityParams, getItemSourceIDs(srv, ViewSourceTypeTreasure))
		if err != nil {
			return err
		}
		rel.Treasures = itemAmtsToChildResAmts2(E.Treasures, treasures, item.ID)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = E.Shops.queryAvlRelRun(srv, ctx, item, availabilityParams, getItemSourceIDs(srv, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		quests, err := E.Quests.queryAvlRelRun(srv, ctx, item, availabilityParams, getItemSourceIDs(srv, ViewSourceTypeQuest))
		if err != nil {
			return err
		}
		rel.Quests = itemAmtsToChildResAmts(E.Quests, quests)
		return nil
	})

	g.Go(func() error {
		blitzballPrizes, err := E.BlitzballPrizes.queryAvlRelRun(srv, ctx, item, availabilityParams, getItemSourceIDs(srv, ViewSourceTypeBlitzball))
		if err != nil {
			return err
		}
		rel.BlitzballPrizes = itemAmtsToChildResAmts2(E.BlitzballPrizes, blitzballPrizes, item.ID)
		return nil
	})

	g.Go(func() error {
		playerAbilities, err := E.PlayerAbilities.dbGetRelations(srv, ctx, item, srv.Cfg.DB.GetItemPlayerAbilityIDs)
		if err != nil {
			return err
		}
		rel.AeonLearnAbilities = itemAmtsToChildResAmts(E.PlayerAbilities, playerAbilities)
		return nil
	})

	g.Go(func() error {
		autoAbilities, err := E.AutoAbilities.dbGetRelations(srv, ctx, item, srv.Cfg.DB.GetItemAutoAbilityIDs)
		if err != nil {
			return err
		}
		rel.AutoAbilities = itemAmtsToChildResAmts(E.AutoAbilities, autoAbilities)
		return nil
	})

	g.Go(func() error {
		var err error
		rel.Mixes, err = E.Mixes.dbGetRelations(srv, ctx, item, srv.Cfg.DB.GetItemMixIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Item{}, err
	}

	return rel, nil
}

func getMonItemAmts(monsters []NamedAPIResource, itemName string) []MonItemAmts {
	monItemAmts := []MonItemAmts{}

	for _, monster := range monsters {
		monItemAmt := createItemMonster(itemName, monster)
		monItemAmts = append(monItemAmts, monItemAmt)
	}

	return monItemAmts
}
