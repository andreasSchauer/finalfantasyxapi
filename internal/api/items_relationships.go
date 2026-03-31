package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// one or both of the query functions return errEmptyQuery
func getItemRelationships(cfg *Config, r *http.Request, item seeding.Item) (Item, error) {
	queryParamAvailability := cfg.e.items.queryLookup["availability"]
	availabilitySlice, err := parseEnumSliceQuery(r, cfg.e.items.endpoint, queryParamAvailability, cfg.t.AvailabilityType)
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return Item{}, err
	}

	repeatable, err := getBoolPtr(r, "repeatable", cfg.e.monsters.queryLookup)
	if err != nil {
		return Item{}, err
	}

	monsterIDs, err := cfg.db.GetItemMonsterIDs(r.Context(), database.GetItemMonsterIDsParams{
		ItemID:       item.ID,
		Repeatable:   h.GetNullBool(repeatable),
		Availability: availabilitySlice,
	})
	if err != nil {
		return Item{}, newHTTPErrorDB(cfg.e.monsters.resourceType, item, err)
	}
	monsters := idsToAPIResources(cfg, cfg.e.monsters, monsterIDs)

	treasureIDs, err := cfg.db.GetItemTreasureIDs(r.Context(), database.GetItemTreasureIDsParams{
		ItemID:       item.ID,
		Availability: availabilitySlice,
	})
	if err != nil {
		return Item{}, newHTTPErrorDB(cfg.e.treasures.resourceType, item, err)
	}
	treasures := idsToAPIResources(cfg, cfg.e.treasures, treasureIDs)

	shopIDs, err := cfg.db.GetItemShopIDs(r.Context(), database.GetItemShopIDsParams{
		ItemID:       item.ID,
		Availability: availabilitySlice,
	})
	if err != nil {
		return Item{}, newHTTPErrorDB(cfg.e.shops.resourceType, item, err)
	}
	shops := idsToAPIResources(cfg, cfg.e.shops, shopIDs)

	questIDs, err := cfg.db.GetItemQuestIDs(r.Context(), database.GetItemQuestIDsParams{
		ItemID:       item.ID,
		Repeatable:   h.GetNullBool(repeatable),
		Availability: availabilitySlice,
	})
	if err != nil {
		return Item{}, newHTTPErrorDB(cfg.e.quests.resourceType, item, err)
	}
	quests := idsToAPIResources(cfg, cfg.e.quests, questIDs)

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
