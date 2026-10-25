package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func filterShopsEquipment(cfg *Config, r *http.Request, ctx context.Context) ([]int32, error) {
	i := cfg.E.Shops

	autoAbilityID, err := getQueryIdPtr(r, cfg.E.AutoAbilities, QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := parseIntListQuery(cfg, r, i.queryLookup[QpnEmptySlots])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := getQueryNameIdPtr(r, cfg.E.Characters, QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetShopIDsEquipmentFilter(ctx, database.GetShopIDsEquipmentFilterParams{
		AutoAbilityID: h.GetNullInt32(autoAbilityID),
		CharacterID:   h.GetNullInt32(charID),
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'.", i.resTypePlural, h.PtrToString(autoAbilityID), h.PtrToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}

func filterTreasuresEquipment(cfg *Config, r *http.Request, ctx context.Context) ([]int32, error) {
	i := cfg.E.Treasures

	autoAbilityID, err := getQueryIdPtr(r, cfg.E.AutoAbilities, QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := parseIntListQuery(cfg, r, i.queryLookup[QpnEmptySlots])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := getQueryNameIdPtr(r, cfg.E.Characters, QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetTreasureIDsEquipmentFilter(ctx, database.GetTreasureIDsEquipmentFilterParams{
		AutoAbilityID: h.GetNullInt32(autoAbilityID),
		CharacterID:   h.GetNullInt32(charID),
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'.", i.resTypePlural, h.PtrToString(autoAbilityID), h.PtrToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}
