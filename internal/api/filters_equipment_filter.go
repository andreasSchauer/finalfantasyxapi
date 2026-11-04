package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func filterShopsEquipment(cfg *Config, r *http.Request, ctx context.Context) ([]int32, error) {
	i := cfg.E.Shops

	autoAbilityID, err := nullID(r, cfg.E.AutoAbilities, typestrings.QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.queryLookup[typestrings.QpnEmptySlots].parseIntList(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := nullNameID(r, cfg.E.Characters, typestrings.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetShopIDsEquipmentFilter(ctx, database.GetShopIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.resTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}

func filterTreasuresEquipment(cfg *Config, r *http.Request, ctx context.Context) ([]int32, error) {
	i := cfg.E.Treasures

	autoAbilityID, err := nullID(r, cfg.E.AutoAbilities, typestrings.QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.queryLookup[typestrings.QpnEmptySlots].parseIntList(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := nullNameID(r, cfg.E.Characters, typestrings.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetTreasureIDsEquipmentFilter(ctx, database.GetTreasureIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.resTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}
