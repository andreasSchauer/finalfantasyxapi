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

func (cfg *Config) retrieveItems(r *http.Request, i handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.ItemCategory, ids, qpnCategory, cfg.db.GetItemIDsCategory)),
		fidl(boolQuery2(r, i, ids, qpnHasAbility, cfg.db.GetItemIDsWithAbility)),
		fidl(nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSing, cfg.l.Stats, cfg.db.GetItemIDsByRelatedStat)),
		fidl(valueListQuery(cfg, r, i, ids, qpnMethods, cfg.db.GetItemIDsByMethods)),
		fidl(idQueryWrapper(cfg, r, i, ids, qpnLocation, cfg.e.locations.objLookup, getItemsByLocation)),
		fidl(idQueryWrapper(cfg, r, i, ids, qpnSublocation, cfg.e.sublocations.objLookup, getItemsBySublocation)),
		fidl(idQueryWrapper(cfg, r, i, ids, qpnArea, cfg.e.areas.objLookup, getItemsByArea)),
	})
}
