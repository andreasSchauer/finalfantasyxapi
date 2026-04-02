package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getItemRelationships(cfg *Config, r *http.Request, item seeding.Item) (Item, error) {
	availabilityParams, err := getAvailabilityParams(cfg, r, cfg.e.items, item)
	if err != nil {
		return Item{}, err
	}

	monsters, err := runAvailabilityQuery(cfg, r, cfg.e.monsters, item, availabilityParams, convGetItemMonsterIDs(cfg))
	if err != nil {
		return Item{}, err
	}

	treasures, err := runAvailabilityQuery(cfg, r, cfg.e.treasures, item, availabilityParams, convGetItemTreasureIDs(cfg))
	if err != nil {
		return Item{}, err
	}

	shops, err := runAvailabilityQuery(cfg, r, cfg.e.shops, item, availabilityParams, convGetItemShopIDs(cfg))
	if err != nil {
		return Item{}, err
	}

	quests, err := runAvailabilityQuery(cfg, r, cfg.e.quests, item, availabilityParams, convGetItemQuestIDs(cfg))
	if err != nil {
		return Item{}, err
	}

	blitzballPrizes, err := getResourcesDbItem(cfg, r, cfg.e.blitzballPrizes, item, cfg.db.GetItemBlitzballPrizeIDs)
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
		Monsters:           getMonItemAmts(cfg, monsters, item),
		Treasures:          getForeignResAmts2(cfg, cfg.e.treasures, treasures, item.ID),
		Shops:              shops,
		Quests:             getForeignResAmts(cfg.e.quests, quests),
		BlitzballPrizes:    getForeignResAmts2(cfg, cfg.e.blitzballPrizes, blitzballPrizes, item.ID),
		AeonLearnAbilities: getForeignResAmts(cfg.e.playerAbilities, playerAbilities),
		AutoAbilities:      getForeignResAmts(cfg.e.autoAbilities, autoAbilities),
		Mixes:              mixes,
	}

	return rel, nil
}



func getMonItemAmts(cfg *Config, monsters []NamedAPIResource, item seeding.Item) []MonItemAmts {
	monItemAmts := []MonItemAmts{}

	for _, monster := range monsters {
		monItemAmt := createItemMonster(cfg, item, monster)
		monItemAmts = append(monItemAmts, monItemAmt)
	}

	return monItemAmts
}
