package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Shop struct {
	ID          int32           `json:"id"`
	Area        AreaAPIResource `json:"area"`
	Category    string          `json:"category"`
	Notes       *string         `json:"notes,omitempty"`
	PreAirship  *SubShop        `json:"pre_airship"`
	PostAirship *SubShop        `json:"post_airship"`
}

type SubShop struct {
	Items     []ShopItem      `json:"items"`
	Equipment []ShopEquipment `json:"equipment"`
}

func convertSubShop(cfg *Config, ss seeding.SubShop) SubShop {
	return SubShop{
		Items:     convertObjSlice(cfg, ss.Items, convertShopItem),
		Equipment: convertObjSlice(cfg, ss.Equipment, convertShopEquipment),
	}
}

type ShopItem struct {
	Item  NamedAPIResource `json:"item"`
	Price int32            `json:"price"`
}

func convertShopItem(cfg *Config, si seeding.ShopItem) ShopItem {
	return ShopItem{
		Item:  nameToNamedAPIResource(cfg, cfg.e.items, si.Name, nil),
		Price: si.Price,
	}
}

type ShopEquipment struct {
	Equipment FoundEquipment `json:"equipment"`
	Price     int32          `json:"price"`
}

func convertShopEquipment(cfg *Config, se seeding.ShopEquipment) ShopEquipment {
	return ShopEquipment{
		Equipment: convertFoundEquipment(cfg, se.FoundEquipment),
		Price:     se.Price,
	}
}

func (cfg *Config) HandleShops(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.shops

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointIDOnly(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

func (cfg *Config) getShop(r *http.Request, i handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Shop, error) {
	shop, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Shop{}, err
	}

	response := Shop{
		ID:          shop.ID,
		Area:        idToAreaAPIResource(cfg, cfg.e.areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
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
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationShopIDs)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationShopIDs)),
		frl(idQuery(cfg, r, i, resources, "auto_ability", len(cfg.l.AutoAbilities), cfg.db.GetShopIDsByAutoAbility)),
		frl(boolQuery2(cfg, r, i, resources, "items", cfg.db.GetShopIDsWithItems)),
		frl(boolQuery2(cfg, r, i, resources, "equipment", cfg.db.GetShopIDsWithEquipment)),
		frl(boolQuery2(cfg, r, i, resources, "pre_airship", cfg.db.GetShopIDsPreAirship)),
		frl(boolQuery2(cfg, r, i, resources, "post_airship", cfg.db.GetShopIDsPostAirship)),
		frl(typeQuery(cfg, r, i, cfg.t.ShopCategory, resources, "category", cfg.db.GetShopIDsByCategory)),
	})
}
