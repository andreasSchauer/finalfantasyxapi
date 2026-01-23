package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// still need to assemble Treasures

type AreaSub struct {
	ID                	int32           	`json:"id"`
	URL					string				`json:"url"`
	ParentLocation		SubName				`json:"parent_location"`
	ParentSublocation	SubName				`json:"parent_sublocation"`
	Name              	string          	`json:"name"`
	Version           	*int32          	`json:"version,omitempty"`
	Specification     	*string         	`json:"specification,omitempty"`
	Shops				[]ShopLocSub		`json:"shops"`
	Treasures			*TreasuresLocSub	`json:"treasures"`
	Monsters			[]SubName			`json:"monsters"`
}

func (a AreaSub) GetSectionName() string {
	return "areas"
}

type SubName struct {
	ID				int32 	`json:"id"`
	Name			string	`json:"name"`
	Version			*int32	`json:"version,omitempty"`
	Specification	*string	`json:"specification,omitempty"`
}

func createSubName(id int32, name string, version *int32, spec *string) SubName {
	return SubName{
		ID: 			id,
		Name: 			name,
		Version: 		version,
		Specification: 	spec,
	}
}

type ShopLocSub struct {
	Category		database.ShopCategory	`json:"category"`
	PreAirship		*ShopSubSummary			`json:"pre_airship"`
	PostAirship		*ShopSubSummary			`json:"post_airship"`
}

type ShopSubSummary struct {
	HasItems		bool	`json:"has_items"`
	HasEquipment	bool	`json:"has_equipment"`
}

type TreasuresLocSub struct {
	TreasureCount	int				`json:"treasure_count"`
	TotalGil		int32			`json:"total_gil"`
	Items			[]ItemAmountSub	`json:"items"`
	Equipment		[]EquipmentSub	`json:"equipment"`
}

type EquipmentSub struct {
	Name				string		`json:"name"`
	AutoAbilities		[]string	`json:"auto_abilities"`
	EmptySlotsAmount	int32		`json:"empty_slots_amount"`
}


func getSubAreas(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.areas
	areas := []AreaSub{}

	for _, areaID := range dbIDs {
		area, _ := seeding.GetResourceByID(areaID, i.objLookupID)
		monsters, err := getSubAreaMonsters(cfg, r, areaID)
		if err != nil {
			return nil, err
		}

		shops, err := getSubAreaShops(cfg, r, areaID)
		if err != nil {
			return nil, err
		}

		treasures, err := getSubAreaTreasures(cfg, r, areaID)
		if err != nil {
			return nil, err
		}

		areaSub := AreaSub{
			ID: area.ID,
			URL: createResourceURL(cfg, i.endpoint, areaID),
			ParentLocation: createSubName(area.SubLocation.Location.ID, area.SubLocation.Location.Name, nil, nil),
			ParentSublocation: createSubName(area.SubLocation.ID, area.SubLocation.Name, nil, nil),
			Name: area.Name,
			Version: area.Version,
			Specification: area.Specification,
			Shops: shops,
			Treasures: treasures,
			Monsters: monsters,
		}

		areas = append(areas, areaSub)
	}

	return toSubResourceSlice(areas), nil
}

func getSubAreaTreasures(cfg *Config, r *http.Request, areaID int32) (*TreasuresLocSub, error) {
	treasureIDs, err := cfg.db.GetAreaTreasureIDs(r.Context(), areaID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve treasures of area with id '%d'", areaID), err)
	}

	if len(treasureIDs) == 0 {
		return nil, nil
	}

	treasures := populateSubAreaTreasures(cfg, treasureIDs)
	return &treasures, nil
}


func populateSubAreaTreasures(cfg *Config, treasureIDs []int32) TreasuresLocSub {
	treasures := TreasuresLocSub{
		TreasureCount: 	len(treasureIDs),
		Items: 			[]ItemAmountSub{},
	}

	for _, treasureID := range treasureIDs {
		treasure, _ := seeding.GetResourceByID(treasureID, cfg.l.TreasuresID)

		switch treasure.LootType {
		case string(database.LootTypeGil):
			treasures.TotalGil += *treasure.GilAmount

		case string(database.LootTypeItem):
			for _, itemAmount := range treasure.Items {
				ia := createSubItemAmount(cfg, itemAmount)
				treasures.Items = append(treasures.Items, ia)
			}

		case string(database.LootTypeEquipment):
			equipment := treasure.Equipment
			es := EquipmentSub{
				Name: 				equipment.Name,
				AutoAbilities: 		sortNamesByID(equipment.Abilities, cfg.l.AutoAbilities),
				EmptySlotsAmount: 	equipment.EmptySlotsAmount,
			}
			treasures.Equipment = append(treasures.Equipment, es)
		}
	}

	treasures.Items = sortItemAmountSubsByID(cfg, treasures.Items)
	return treasures
}

func sortNamesByID[T h.HasID](s []string, lookup map[string]T) []string {
	slices.SortStableFunc(s, func (a, b string) int{
		A, _ := seeding.GetResource(a, lookup)
		B, _ := seeding.GetResource(b, lookup)

		if A.GetID() < B.GetID() {
			return -1
		}

		if A.GetID() > B.GetID() {
			return -1
		}

		return 0
	})

	return s
}


func getSubAreaShops(cfg *Config, r *http.Request, areaID int32) ([]ShopLocSub, error) {
	shops := []ShopLocSub{}

	shopIDs, err := cfg.db.GetAreaShopIDs(r.Context(), areaID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve shops of area with id '%d'", areaID), err)
	}

	for _, shopID := range shopIDs {
		shopLookup, _ := seeding.GetResourceByID(shopID, cfg.l.ShopsID)
		shop := ShopLocSub{
			Category: 		database.ShopCategory(shopLookup.Category),
			PreAirship: 	createShopLocSub(shopLookup.PreAirship),
			PostAirship: 	createShopLocSub(shopLookup.PostAirship),
		}
		shops = append(shops, shop)
	}

	return shops, nil
}

func createShopLocSub(shop *seeding.SubShop) *ShopSubSummary {
	if shop == nil {
		return nil
	}
	shopLoc := ShopSubSummary{}

	if len(shop.Items) != 0 {
		shopLoc.HasItems = true
	}

	if len(shop.Equipment) != 0 {
		shopLoc.HasEquipment = true
	}

	return &shopLoc
}


func getSubAreaMonsters(cfg *Config, r *http.Request, areaID int32) ([]SubName, error) {
	monsters := []SubName{}
	
	monIDs, err := cfg.db.GetAreaMonsterIDs(r.Context(), areaID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve monsters of area with id '%d'", areaID), err)
	}

	for _, monID := range monIDs {
		mon, _ := seeding.GetResourceByID(monID, cfg.l.MonstersID)
		subName := createSubName(mon.ID, mon.Name, mon.Version, mon.Specification)
		monsters = append(monsters, subName)
	}

	return monsters, nil
}