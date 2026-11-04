package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

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

func convertSubShop(srv *Server, ss seedtypes.SubShop) SubShop {
	return SubShop{
		Items:     convertObjSlice(srv, ss.Items, convertShopItem),
		Equipment: convertObjSlice(srv, ss.Equipment, convertShopEquipment),
	}
}

type ShopItem struct {
	Item  NamedAPIResource `json:"item"`
	Price int32            `json:"price"`
}

func convertShopItem(srv *Server, si seedtypes.ShopItem) ShopItem {
	return ShopItem{
		Item:  nameToNamedAPIResource(srv, E.Items, si.Name, nil),
		Price: si.Price,
	}
}

type ShopEquipment struct {
	Equipment FoundEquipment `json:"equipment"`
	Price     int32          `json:"price"`
}

func convertShopEquipment(srv *Server, se seedtypes.ShopEquipment) ShopEquipment {
	return ShopEquipment{
		Equipment: convertFoundEquipment(srv, se.TreasureEquipment),
		Price:     se.Price,
	}
}
