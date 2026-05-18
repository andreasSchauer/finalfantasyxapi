package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getShop(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Shop, error) {
	shop, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Shop{}, err
	}

	response := Shop{
		ID:          shop.ID,
		Area:        idToAreaAPIResource(cfg, cfg.e.areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    enumToNamedAPIResource(cfg, cfg.e.shopCategory.endpoint, shop.Category, cfg.t.ShopCategory),
		PreAirship:  convertObjPtr(cfg, shop.PreAirship, convertSubShop),
		PostAirship: convertObjPtr(cfg, shop.PostAirship, convertSubShop),
	}

	return response, nil
}

func (cfg *Config) retrieveShops(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[UnnamedAPIResource]{
		frl(basicQueryWrapper(cfg, r, i, resources, "empty_slots", getShopsByEmptySlots)),
		frl(idQueryWrapper(cfg, r, i, resources, "auto_ability", len(cfg.l.AutoAbilities), getShopsByAutoAbility)),
		frl(idQueryWrapper(cfg, r, i, resources, "location", len(cfg.l.Locations), getShopIDsByLocation)),
		frl(idQueryWrapper(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), getShopIDsBySublocation)),
		frl(boolQueryWrapper(cfg, r, i, resources, "items", getShopIDsWithItems)),
		frl(boolQueryWrapper(cfg, r, i, resources, "equipment", getShopIDsWithEquipment)),
		frl(enumListQuery(cfg, r, i, cfg.t.AvailabilityType, resources, "availability", cfg.db.GetShopIDsByAvailability)),
		frl(enumListQuery(cfg, r, i, cfg.t.ShopCategory, resources, "category", cfg.db.GetShopIDsByCategory)),
	})
}
