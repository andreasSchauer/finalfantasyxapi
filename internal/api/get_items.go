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
		itemAbilityNamePtr = &item.Name
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
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.ItemCategory, ids, "category", cfg.db.GetItemIDsCategory)),
		fidl(boolQuery2(r, i, ids, "has_ability", cfg.db.GetItemIDsWithAbility)),
		fidl(nameIdQuery(r, i, ids, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetItemIDsByRelatedStat)),
		fidl(valueListQuery(cfg, r, i, ids, "methods", cfg.db.GetItemIDsByMethods)),
		fidl(idQueryWrapper(cfg, r, i, ids, "location", cfg.e.locations.objLookup, getItemsByLocation)),
		fidl(idQueryWrapper(cfg, r, i, ids, "sublocation", cfg.e.sublocations.objLookup, getItemsBySublocation)),
		fidl(idQueryWrapper(cfg, r, i, ids, "area", cfg.e.areas.objLookup, getItemsByArea)),
	})
}
