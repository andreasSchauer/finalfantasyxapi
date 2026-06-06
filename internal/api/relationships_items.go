package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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
	monsters, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeMonster))
	if err != nil {
		return Item{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeTreasure))
	if err != nil {
		return Item{}, err
	}

	shops, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeShop))
	if err != nil {
		return Item{}, err
	}

	quests, err := runRelAvailabilityQuery(cfg, r, cfg.e.quests, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeQuest))
	if err != nil {
		return Item{}, err
	}

	blitzballPrizes, err := runRelAvailabilityQuery(cfg, r, cfg.e.blitzballPrizes, item, availabilityParams, getItemSourceIDs(cfg, ViewSourceTypeBlitzball))
	if err != nil {
		return Item{}, err
	}

	playerAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, item, cfg.db.GetItemPlayerAbilityIDs)
	if err != nil {
		return Item{}, err
	}

	autoAbilities, err := getResourcesDbItem(cfg, r, cfg.e.autoAbilities, item, cfg.db.GetItemAutoAbilityIDs)
	if err != nil {
		return Item{}, err
	}

	mixes, err := getResourcesDbItem(cfg, r, cfg.e.mixes, item, cfg.db.GetItemMixIDs)
	if err != nil {
		return Item{}, err
	}

	rel := Item{
		Monsters:           getMonItemAmts(cfg, monsters, item.Name),
		Treasures:          itemAmtsToChildResAmts2(cfg, cfg.e.treasures, treasures, item.ID),
		Shops:              shops,
		Quests:             itemAmtsToChildResAmts(cfg.e.quests, quests),
		BlitzballPrizes:    itemAmtsToChildResAmts2(cfg, cfg.e.blitzballPrizes, blitzballPrizes, item.ID),
		AeonLearnAbilities: itemAmtsToChildResAmts(cfg.e.playerAbilities, playerAbilities),
		AutoAbilities:      itemAmtsToChildResAmts(cfg.e.autoAbilities, autoAbilities),
		Mixes:              mixes,
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
