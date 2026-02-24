package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type Shop struct {
	ID          int32           `json:"id"`
	Area        AreaAPIResource `json:"area"`
	Category    NamedAPIResource          `json:"category"`
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