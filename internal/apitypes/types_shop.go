package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type Shop struct {
	ID          int32                       `json:"id"`
	Area        apiresource.AreaAPIResource `json:"area"`
	Category    string                      `json:"category"`
	Notes       *string                     `json:"notes,omitempty"`
	PreAirship  *SubShop                    `json:"pre_airship"`
	PostAirship *SubShop                    `json:"post_airship"`
}

type SubShop struct {
	Items     []ShopItem      `json:"items"`
	Equipment []ShopEquipment `json:"equipment"`
}

type ShopItem struct {
	Item  apiresource.NamedAPIResource `json:"item"`
	Price int32                        `json:"price"`
}

type ShopEquipment struct {
	Equipment FoundEquipment `json:"equipment"`
	Price     int32          `json:"price"`
}
