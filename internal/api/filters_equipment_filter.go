package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func filterShopsEquipment(server *Server, r *http.Request, ctx context.Context) ([]int32, error) {
	i := E.Shops

	autoAbilityID, err := nullID(r, E.AutoAbilities, ts.QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.queryLookup[ts.QpnEmptySlots].parseIntList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := nullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetShopIDsEquipmentFilter(ctx, database.GetShopIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.resTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}

func filterTreasuresEquipment(server *Server, r *http.Request, ctx context.Context) ([]int32, error) {
	i := E.Treasures

	autoAbilityID, err := nullID(r, E.AutoAbilities, ts.QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.queryLookup[ts.QpnEmptySlots].parseIntList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := nullNameID(r, E.Characters, ts.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetTreasureIDsEquipmentFilter(ctx, database.GetTreasureIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.resTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}
