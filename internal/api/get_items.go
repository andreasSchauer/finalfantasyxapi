package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getItem(r *http.Request, i handlerInput[seedtypes.Item, Item, NamedAPIResource], id int32) (Item, error) {
	item, err := i.queryParamsVerifyGet(r, id)
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
		UntypedItem:        idToTypedAPIResource(cfg, cfg.E.AllItems, item.MasterItem.ID),
		Category:           item.Category,
		Description:        item.Description,
		SgDescription:      item.SphereGridDescription,
		Effect:             item.Effect,
		Usability:          item.Usability,
		BasePrice:          item.BasePrice,
		SellValue:          item.SellValue,
		ItemAbility:        namePtrToNamedAPIResPtr(cfg, cfg.E.ItemAbilities, itemAbilityNamePtr, nil),
		Sphere:             rel.Sphere,
		AvailableMenus:     namesToNamedAPIResources(cfg, cfg.E.Submenus, item.AvailableMenus),
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.E.Stats, item.RelatedStats),
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

func (cfg *Config) retrieveItems(r *http.Request, i handlerInput[seedtypes.Item, Item, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, cfg.T.ItemCategory, ids, QpnCategory, cfg.db.GetItemIDsCategory),
		i.queryBool2(r, ids, QpnHasAbility, cfg.db.GetItemIDsWithAbility),
		i.queryNameID(r, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, lookup.L.Stats, cfg.db.GetItemIDsByRelatedStat),
		i.queryValueList(cfg, r, ids, QpnMethods, cfg.db.GetItemIDsByMethods),
		i.queryIDWrapper(cfg, r, ids, QpnLocation, cfg.E.Locations.objLookup, getItemsByLocation),
		i.queryIDWrapper(cfg, r, ids, QpnSublocation, cfg.E.Sublocations.objLookup, getItemsBySublocation),
		i.queryIDWrapper(cfg, r, ids, QpnArea, cfg.E.Areas.objLookup, getItemsByArea),
	})
}
