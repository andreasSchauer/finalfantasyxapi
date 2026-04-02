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

	queryParamAutoAbility := i.queryLookup["auto_ability"]
	_, err := checkEmptyQuery(r, queryParamAutoAbility)
	if !errors.Is(err, errEmptyQuery) {
		return nil, errQueryRedirect
	}

	return filterShopsEquipment(cfg, r, nil)
}

func getShopsByAutoAbility(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	i := cfg.e.shops

	dbIDs, err := filterShopsEquipment(cfg, r, &id)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

func filterShopsEquipment(cfg *Config, r *http.Request, autoAbilityIdPtr *int32) ([]int32, error) {
	i := cfg.e.shops

	slotsPtr, err := getSlotsPtr(r, i.queryLookup)
	if err != nil {
		return nil, err
	}

	charIdPtr, err := getCharIdPtr(cfg, r, i.queryLookup)
	if err != nil {
		return nil, err
	}

	shopTypePtr, err := getTypePtr(r, "shop_type", cfg.e.shops.endpoint, cfg.t.ShopType, i.queryLookup)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetShopIDsEquipmentFilter(context.Background(), database.GetShopIDsEquipmentFilterParams{
		ShopType:      cfg.t.ShopType.nullConvFunc(shopTypePtr),
		AutoAbilityID: h.GetNullInt32(autoAbilityIdPtr),
		CharacterID:   h.GetNullInt32(charIdPtr),
		EmptySlots:    h.GetNullInt32(slotsPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by auto ability id '%s', character id '%s', shop type '%s', empty slots amount '%s'.", i.resourceType, h.PtrToString(autoAbilityIdPtr), h.PtrToString(charIdPtr), h.PtrToString(shopTypePtr), h.PtrToString(slotsPtr)), err)
	}

	return dbIDs, nil
}

// can generalize these functions pretty easliy, if needed

func getBoolPtr(r *http.Request, queryName string, queryLookup map[string]QueryType) (*bool, error) {
	queryParam := queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func getSlotsPtr(r *http.Request, queryLookup map[string]QueryType) (*int32, error) {
	var slotsPtr *int32
	queryParamEmptySlots := queryLookup["empty_slots"]
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

	return slotsPtr, nil
}

func getCharIdPtr(cfg *Config, r *http.Request, queryLookup map[string]QueryType) (*int32, error) {
	iChars := cfg.e.characters
	var charIdPtr *int32
	queryParamCharacter := queryLookup["character"]
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

	return charIdPtr, nil
}

func getTypePtr[E, N any](r *http.Request, queryName, endpoint string, et EnumType[E, N], queryLookup map[string]QueryType) (*string, error) {
	var typePtr *string

	queryParam := queryLookup[queryName]
	var isEmpty bool
	enumRes, err := parseEnumQuery(r, endpoint, queryParam, et)
	if err != nil {
		if errors.Is(err, errEmptyQuery) {
			isEmpty = true
		} else {
			return nil, err
		}
	}

	if !isEmpty {
		typePtr = &enumRes.Name
	}

	return typePtr, nil
}
