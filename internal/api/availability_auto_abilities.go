package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func filterAvlAutoAbilities(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.E.AutoAbilities

	avlParams, err := i.queryAvlRepGetParams(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	customize, err := i.queryLookup[typestrings.QpnCustomize].parseBool(r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	charID, err := nullID(r, cfg.E.Characters, typestrings.QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	methods, err := i.queryLookup[typestrings.QpnMethods].parseValueList(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterAutoAbilityIDsByAvailability(r.Context(), database.FilterAutoAbilityIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship:     avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		CharacterID:    charID,
		Methods:        methods,
		Customize:      customize,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
