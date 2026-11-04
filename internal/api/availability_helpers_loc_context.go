package api

import (
	"database/sql"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type locContextParams struct {
	AvlType string
	ID      sql.NullInt32
	Type    sql.NullString
}

func getLocContextParams[T seedtypes.Lookupable, R any, A APIResource](cfg *Config, r *http.Request, i handlerInput[T, R, A]) (locContextParams, error) {
	avlType := AvlTypeSelf
	var locContextID sql.NullInt32
	var locContextType string

	locID, err := nullID(r, cfg.E.Locations, typestrings.QpnLocation, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = locID
		locContextType = string(ViewSourceTypeLocation)
	}

	subLocID, err := nullID(r, cfg.E.Sublocations, typestrings.QpnSublocation, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = subLocID
		locContextType = string(ViewSourceTypeSublocation)
	}

	areaID, err := nullID(r, cfg.E.Areas, typestrings.QpnArea, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext2
		locContextID = areaID
		locContextType = string(ViewSourceTypeArea)
	}

	var locCtxTypePtr *string

	if locContextType != "" {
		locCtxTypePtr = &locContextType
	}

	params := locContextParams{
		AvlType: string(avlType),
		ID:      locContextID,
		Type:    h.GetNullString(locCtxTypePtr),
	}

	return params, nil
}
