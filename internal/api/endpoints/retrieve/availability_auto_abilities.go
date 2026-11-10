// Package retrieve adoiadjio
package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func FilterAvlAutoAbilities(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.AutoAbilities

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	customize, err := i.QueryLookup[ts.QpnCustomize].ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	charID, err := handlerinput.NullID(r, apilookups.E.Characters, ts.QpnCharacter, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	methods, err := i.QueryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterAutoAbilityIDsByAvailability(r.Context(), database.FilterAutoAbilityIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		IsRepeatable:   avlParams.IsRepeatable,
		PreAirship:     avlParams.PreAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		CharacterID:    charID,
		Methods:        methods,
		Customize:      customize,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
