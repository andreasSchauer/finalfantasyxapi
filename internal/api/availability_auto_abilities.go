package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func filterAvlAutoAbilities(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.autoAbilities

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	reqItem, err := parseBooleanQuery(r, i.queryLookup[qpnReqItem])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	charID, err := getQueryIdPtr(r, cfg.e.characters, qpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup[qpnMethods])
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
		CharacterID:    h.GetNullInt32(charID),
		Methods:        methods,
		ReqItem:        reqItem,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
