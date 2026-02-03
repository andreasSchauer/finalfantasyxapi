package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ShopSub struct {
	ID          int32       `json:"id"`
	URL         string      `json:"url"`
	Area        string      `json:"area"`
	Category	string		`json:"category"`
	Notes       *string     `json:"notes,omitempty"`
	PreAirship	*SubShopSub	`json:"pre_airship"`
	PostAirship	*SubShopSub	`json:"post_airship"`
}

func (s ShopSub) GetSectionName() string {
	return "shops"
}

func (s ShopSub) GetURL() string {
	return s.URL
}

type SubShopSub struct {
	Items     []string  `json:"items"`
	Equipment []string 	`json:"equipment"`
}

func handleShopsSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.shops
	shops := []ShopSub{}

	for _, shopID := range dbIDs {
		shop, _ := seeding.GetResourceByID(shopID, i.objLookupID)

		shopSub := ShopSub{
			ID:            shop.ID,
			URL:           createResourceURL(cfg, i.endpoint, shopID),
			Area:          idToLocAreaString(cfg, shop.AreaID),
			Notes:         shop.Notes,
			
		}

		shops = append(shops, shopSub)
	}

	return toSubResourceSlice(shops), nil
}
