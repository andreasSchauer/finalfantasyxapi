package apiconvert

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func SubShop(ss seedtypes.SubShop) apitypes.SubShop {
	return apitypes.SubShop{
		Items:     ah.ConvertObjSlice(ss.Items, ShopItem),
		Equipment: ah.ConvertObjSlice(ss.Equipment, ShopEquipment),
	}
}

func ShopItem(si seedtypes.ShopItem) apitypes.ShopItem {
	return apitypes.ShopItem{
		Item:  handlerinput.NameToNamedAPIResource(apilookups.E.Items, si.Name, nil),
		Price: si.Price,
	}
}

func ShopEquipment(se seedtypes.ShopEquipment) apitypes.ShopEquipment {
	return apitypes.ShopEquipment{
		Equipment: FoundEquipment(se.TreasureEquipment),
		Price:     se.Price,
	}
}
