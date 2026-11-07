package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func filterShopsEquipment(r *http.Request, ctx context.Context) ([]int32, error) {
	i := E.Shops

	autoAbilityID, err := NullID(r, E.AutoAbilities, ts.QpnAutoAbility, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.queryLookup[ts.QpnEmptySlots].ParseIntList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := NullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetShopIDsEquipmentFilter(ctx, database.GetShopIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.resTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}

func filterTreasuresEquipment(r *http.Request, ctx context.Context) ([]int32, error) {
	i := E.Treasures

	autoAbilityID, err := NullID(r, E.AutoAbilities, ts.QpnAutoAbility, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.queryLookup[ts.QpnEmptySlots].ParseIntList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := NullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetTreasureIDsEquipmentFilter(ctx, database.GetTreasureIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.resTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}
