package retrieve

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func FilterShopsEquipment(r *http.Request, ctx context.Context) ([]int32, error) {
	i := apilookups.E.Shops

	autoAbilityID, err := handlerinput.NullID(r, apilookups.E.AutoAbilities, ts.QpnAutoAbility, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.QueryLookup[ts.QpnEmptySlots].ParseIntList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := handlerinput.NullNameID(r, apilookups.E.Characters, ts.QpnCharacter, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetShopIDsEquipmentFilter(ctx, database.GetShopIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.ResTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}

func FilterTreasuresEquipment(r *http.Request, ctx context.Context) ([]int32, error) {
	i := apilookups.E.Treasures

	autoAbilityID, err := handlerinput.NullID(r, apilookups.E.AutoAbilities, ts.QpnAutoAbility, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	emptySlots, err := i.QueryLookup[ts.QpnEmptySlots].ParseIntList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	charID, err := handlerinput.NullNameID(r, apilookups.E.Characters, ts.QpnCharacter, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetTreasureIDsEquipmentFilter(ctx, database.GetTreasureIDsEquipmentFilterParams{
		AutoAbilityID: autoAbilityID,
		CharacterID:   charID,
		EmptySlots:    h.SliceOrNil(emptySlots),
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%s', character id '%s', empty slots amount '%s'", i.ResTypePlural, h.NullInt32ToString(autoAbilityID), h.NullInt32ToString(charID), h.FormatIntSlice(emptySlots)), err)
	}

	return dbIDs, nil
}
