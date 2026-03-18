package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func getShopsByEmptySlots(cfg *Config, r *http.Request, slots int32) ([]int32, error) {
	i := cfg.e.shops
	iChars := cfg.e.characters
	var charIdPtr *int32

	queryParamAutoAbility := i.queryLookup["auto_ability"]
	var slotsPtr *int32
	var shopTypePtr *string

	_, err := checkEmptyQuery(r, queryParamAutoAbility)
	if !errors.Is(err, errEmptyQuery) {
		return nil, errQueryUsedElsewhere
	}

	queryParamEmptySlots := i.queryLookup["empty_slots"]
	var slotsIsEmpty bool
	emptySlots, err := parseIntQuery(r, queryParamEmptySlots)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			slotsIsEmpty = true
		} else {
			return nil, err
		}
	}

	if !slotsIsEmpty {
		emptySlots32 := int32(emptySlots)
		slotsPtr = &emptySlots32
	}

	queryParamCharacter := i.queryLookup["character"]
	var charIsEmpty bool
	charID, err := parseNameOrIdQuery(r, queryParamCharacter, iChars.resourceType, iChars.objLookup)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			charIsEmpty = true
		} else {
			return nil, err
		}
	}

	if !charIsEmpty {
		charIdPtr = &charID
	}

	queryParamShopType := i.queryLookup["shop_type"]
	var shopTypeIsEmpty bool
	shopType, err := parseTypeQuery(r, i.endpoint, queryParamShopType, cfg.t.ShopType)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			shopTypeIsEmpty = true
		} else {
			return nil, err
		}
	}

	if !shopTypeIsEmpty {
		shopTypePtr = &shopType.Name
	}

	dbIDs, err := cfg.db.GetShopIDsByEmptySlots(r.Context(), database.GetShopIDsByEmptySlotsParams{
		ShopType:      cfg.t.ShopType.nullConvFunc(shopTypePtr),
		CharacterID:   h.GetNullInt32(charIdPtr),
		EmptySlots:    h.GetNullInt32(slotsPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by empty slots amount '%d'.", i.resourceType, slots), err)
	}

	return dbIDs, nil
}

func getShopsByAutoAbility(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	i := cfg.e.shops
	iChars := cfg.e.characters
	var charIdPtr *int32
	var slotsPtr *int32
	var shopTypePtr *string

	queryParamCharacter := i.queryLookup["character"]
	var charIsEmpty bool
	charID, err := parseNameOrIdQuery(r, queryParamCharacter, iChars.resourceType, iChars.objLookup)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			charIsEmpty = true
		} else {
			return nil, err
		}
	}

	if !charIsEmpty {
		charIdPtr = &charID
	}

	queryParamEmptySlots := i.queryLookup["empty_slots"]
	var slotsIsEmpty bool
	emptySlots, err := parseIntQuery(r, queryParamEmptySlots)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			slotsIsEmpty = true
		} else {
			return nil, err
		}
	}

	if !slotsIsEmpty {
		emptySlots32 := int32(emptySlots)
		slotsPtr = &emptySlots32
	}

	queryParamShopType := i.queryLookup["shop_type"]
	var shopTypeIsEmpty bool
	shopType, err := parseTypeQuery(r, i.endpoint, queryParamShopType, cfg.t.ShopType)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			shopTypeIsEmpty = true
		} else {
			return nil, err
		}
	}

	if !shopTypeIsEmpty {
		shopTypePtr = &shopType.Name
	}

	dbIDs, err := cfg.db.GetShopIDsEquipmentFilter(context.Background(), database.GetShopIDsEquipmentFilterParams{
		ShopType:      cfg.t.ShopType.nullConvFunc(shopTypePtr),
		AutoAbilityID: h.GetNullInt32(&id),
		CharacterID:   h.GetNullInt32(charIdPtr),
		EmptySlots:    h.GetNullInt32(slotsPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by auto-ability id '%d'.", i.resourceType, id), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}
