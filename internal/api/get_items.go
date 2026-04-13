package api

import (
	"net/http"

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
		UntypedItem:        idToTypedAPIResource(cfg, cfg.e.allItems, item.MasterItem.ID),
		Category:           enumToNamedAPIResource(cfg, cfg.e.itemCategory.endpoint, item.Category, cfg.t.ItemCategory),
		Description:        item.Description,
		SgDescription:      item.SphereGridDescription,
		Effect:             item.Effect,
		Usability:          item.Usability,
		BasePrice:          item.BasePrice,
		SellValue:          item.SellValue,
		ItemAbility:        namePtrToNamedAPIResPtr(cfg, cfg.e.itemAbilities, itemAbilityNamePtr, nil),
		Sphere:             rel.Sphere,
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

func (cfg *Config) retrieveItems(r *http.Request, i handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.ItemCategory, resources, "category", cfg.db.GetItemIDsCategory)),
		frl(boolQuery2(cfg, r, i, resources, "has_ability", cfg.db.GetItemIDsWithAbility)),
		frl(nameIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetItemIDsByRelatedStat)),
		frl(basicQueryWrapper(cfg, r, i, resources, "method", getItemsByMethod)),
	})
}
