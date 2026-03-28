package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getItem(r *http.Request, i handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList], id int32) (Item, error) {
	item, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Item{}, err
	}

	var itemAbilityNamePtr *string
	if len(item.BattleInteractions) > 0 {
		itemAbility, _ := seeding.GetResource(item.Name, cfg.l.ItemAbilities)
		itemAbilityNamePtr = &itemAbility.Name
	}

	rel, err := getItemRelationships(cfg, r, item)
	if err != nil {
		return Item{}, err
	}

	response := Item{
		ID:                 item.ID,
		Name:               item.Name,
		Description:        item.Description,
		SgDescription:      item.SphereGridDescription,
		Effect:             item.Effect,
		Category:           newNamedAPIResourceFromEnum(cfg, cfg.e.itemCategory.endpoint, item.Category, cfg.t.ItemCategory),
		Usability:          item.Usability,
		BasePrice:          item.BasePrice,
		SellValue:          item.SellValue,
		ItemAbility:        namePtrToNamedAPIResPtr(cfg, cfg.e.itemAbilities, itemAbilityNamePtr, nil),
		AvailableMenus:     namesToNamedAPIResources(cfg, cfg.e.submenus, item.AvailableMenus),
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.e.stats, item.RelatedStats),
		Monsters:           rel.Monsters,
		Treasures:          rel.Treasures,
		Shops:              rel.Shops,
		Quests:             rel.Quests,
		BlitzballPrizes:    rel.BlitzballPrizes,
		AeonLearnAbilities: rel.AeonLearnAbilities,
		AutoAbilities:      rel.AutoAbilities,
		Mixes:              rel.Mixes,
	}

	return response, nil
}

func getItemRelationships(cfg *Config, r *http.Request, item seeding.Item) (Item, error) {
	monsters, err := getMonItemAmts(cfg, r, item)
	if err != nil {
		return Item{}, err
	}

	treasures, err := getResourcesDbItem(cfg, r, cfg.e.treasures, item, cfg.db.GetTreasureIDsByItem)
	if err != nil {
		return Item{}, err
	}

	shops, err := getResourcesDbItem(cfg, r, cfg.e.shops, item, cfg.db.GetItemShopIDs)
	if err != nil {
		return Item{}, err
	}

	quests, err := getResourcesDbItem(cfg, r, cfg.e.quests, item, cfg.db.GetItemQuestIDs)
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
		Monsters:           monsters,
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


func getMonItemAmts(cfg *Config, r *http.Request, item seeding.Item) ([]MonItemAmts, error) {
	i := cfg.e.monsters
	monItemAmts := []MonItemAmts{}

	postAirship, err := getBoolPtr(r, "post_airship", i.queryLookup)
	if err != nil {
		return nil, err
	}
	
	storyBased, err := getBoolPtr(r, "story_based", i.queryLookup)
	if err != nil {
		return nil, err
	}

	repeatable, err := getBoolPtr(r, "repeatable", i.queryLookup)
	if err != nil {
		return nil, err
	}

	dbIds, err := cfg.db.GetItemMonsterIDs(r.Context(), database.GetItemMonsterIDsParams{
		ItemID: 		item.GetID(),
		PostAirship: 	h.GetNullBool(postAirship),
		StoryBased: 	h.GetNullBool(storyBased),
		Repeatable: 	h.GetNullBool(repeatable),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %ss of %s.", i.resourceType, item), err)
	}

	monsters := idsToAPIResources(cfg, i, dbIds)

	for _, monster := range monsters {
		monItemAmt := createItemMonster(cfg, item, monster)
		monItemAmts = append(monItemAmts, monItemAmt)
	}

	return monItemAmts, nil
}

func (cfg *Config) retrieveItems(r *http.Request, i handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{})
}
