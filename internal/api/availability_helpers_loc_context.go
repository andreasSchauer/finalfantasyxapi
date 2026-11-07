package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type locContextParams struct {
	AvlType string
	ID      sql.NullInt32
	Type    sql.NullString
}

func getLocContextParams[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i HandlerInput[T, R, A]) (locContextParams, error) {
	avlType := AvlTypeSelf
	var locContextID sql.NullInt32
	var locContextType string

	locID, err := NullID(r, E.Locations, ts.QpnLocation, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = locID
		locContextType = string(ViewSourceTypeLocation)
	}

	subLocID, err := NullID(r, E.Sublocations, ts.QpnSublocation, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = subLocID
		locContextType = string(ViewSourceTypeSublocation)
	}

	areaID, err := NullID(r, E.Areas, ts.QpnArea, i.queryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !errs.QueryIsEmpty(err) {
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
