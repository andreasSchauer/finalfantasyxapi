package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type locContextParams struct {
	AvlType string
	ID      sql.NullInt32
	Type    sql.NullString
}

func getLocContextParams[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.Main[T, R, A]) (locContextParams, error) {
	avlType := ts.AvlTypeSelf
	var locContextID sql.NullInt32
	var locContextType string

	locID, err := handlerinput.NullID(r, apilookups.E.Locations, ts.QpnLocation, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = ts.AvlTypeContext
		locContextID = locID
		locContextType = string(ts.VstLocation)
	}

	subLocID, err := handlerinput.NullID(r, apilookups.E.Sublocations, ts.QpnSublocation, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = ts.AvlTypeContext
		locContextID = subLocID
		locContextType = string(ts.VstSublocation)
	}

	areaID, err := handlerinput.NullID(r, apilookups.E.Areas, ts.QpnArea, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = ts.AvlTypeContext2
		locContextID = areaID
		locContextType = string(ts.VstArea)
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
