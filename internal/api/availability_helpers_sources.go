package api

import (
	"database/sql"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type locBasedSources struct {
	RequiredSources []string
	ExcludedSources []string
	MonsterID       sql.NullInt32
	ItemID          sql.NullInt32
	KeyItemID       sql.NullInt32
	Methods         []string
}


func getLocBasedSources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (locBasedSources, error) {
	reqs := []string{}
	excls := []string{}

	monID, err := getQueryIdPtr(r, cfg.e.monsters, "monster", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, "monster-single")
	}
	
	itemID, err := getQueryIdPtr(r, cfg.e.items, "item", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeItem))
	}

	keyItemID, err := getQueryIdPtr(r, cfg.e.keyItems, "key_item", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeKeyItem))
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup["method"])
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[string]string{
		"monsters": 	string(ViewSourceTypeMonster),
		"boss_fights": 	string(ViewSourceTypeBoss),
		"shops": 		string(ViewSourceTypeShop),
		"treasures": 	string(ViewSourceTypeTreasure),
		"sidequests": 	string(ViewSourceTypeQuest),
	})
	if err != nil {
		return locBasedSources{}, err
	}

	sources := locBasedSources{
		RequiredSources: h.SliceOrNil(reqs),
		ExcludedSources: h.SliceOrNil(excls),
		MonsterID: 		 h.GetNullInt32(monID),
		ItemID: 		 h.GetNullInt32(itemID),
		KeyItemID: 		 h.GetNullInt32(keyItemID),
		Methods: 		 h.SliceOrNil(methods),
	}

	return sources, nil
}



type shopSources struct {
	AvlType				string
	RequiredSources 	[]string
	ExcludedSources 	[]string
	AutoAbilityID       sql.NullInt32
	CharacterID       	sql.NullInt32
	EmptySlots        	[]int32
}


func getShopSources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (shopSources, error) {
	avlType := AvlTypeSelf
	reqs := []string{}
	excls := []string{}

	autoAbilityID, err := getQueryIdPtr(r, cfg.e.autoAbilities, "auto_ability", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	emptySlots, err := parseIntListQuery(cfg, r, i.queryLookup["empty_slots"])
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	charID, err := getQueryNameIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[string]string{
		"items": 		string(ViewSourceTypeItem),
		"equipment": 	string(ViewSourceTypeEquipment),
	})
	if err != nil {
		return shopSources{}, err
	}

	sources := shopSources{
		AvlType:      		string(avlType),
		RequiredSources:    h.SliceOrNil(reqs),
		ExcludedSources: 	h.SliceOrNil(excls),
		AutoAbilityID: 		h.GetNullInt32(autoAbilityID),
		CharacterID: 		h.GetNullInt32(charID),
		EmptySlots: 		h.SliceOrNil(emptySlots),
	}

	return sources, nil
}



func parseBoolSources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], reqs, excls []string, sourceMap map[string]string) ([]string, []string, error) {
	for queryParam := range sourceMap {
		b, err := parseBooleanQuery(r, i.queryLookup[queryParam])
		if errExceptEmptyQuery(err) {
			return nil, nil, err
		}
		if !queryIsEmpty(err) {
			if b {
				reqs = append(reqs, sourceMap[queryParam])
			} else {
				excls = append(excls, sourceMap[queryParam])
			}
		}
	}

	return reqs, excls, nil
}