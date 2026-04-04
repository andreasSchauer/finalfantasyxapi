package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getShopsByEmptySlots(cfg *Config, r *http.Request, _ string, _ QueryParam) ([]int32, error) {
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

	emptySlots, err := parseIntListQuery(r, i.queryLookup["empty_slots"])
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return nil, err
	}

	charIdPtr, err := getQueryIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if err != nil {
		return nil, err
	}

	shopTypePtr, err := getQueryEnumPtr(r, "shop_type", i.endpoint, cfg.t.ShopType, i.queryLookup)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetShopIDsEquipmentFilter(context.Background(), database.GetShopIDsEquipmentFilterParams{
		ShopType:      cfg.t.ShopType.nullConvFunc(shopTypePtr),
		AutoAbilityID: h.GetNullInt32(autoAbilityIdPtr),
		CharacterID:   h.GetNullInt32(charIdPtr),
		EmptySlots:    emptySlots,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by auto ability id '%s', character id '%s', shop type '%s', empty slots amount '%s'.", i.resourceType, h.PtrToString(autoAbilityIdPtr), h.PtrToString(charIdPtr), h.PtrToString(shopTypePtr), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}
