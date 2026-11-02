package api

import (
	"database/sql"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type locContextParams struct {
	AvlType string
	ID      sql.NullInt32
	Type    sql.NullString
}

func getLocContextParams[T seeding.Lookupable, R any, A APIResource](cfg *Config, r *http.Request, i handlerInput[T, R, A]) (locContextParams, error) {
	avlType := AvlTypeSelf
	var locContextID sql.NullInt32
	var locContextType string

	locID, err := nullInt32FromIDQuery(r, cfg.E.Locations, QpnLocation, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = locID
		locContextType = string(ViewSourceTypeLocation)
	}

	subLocID, err := nullInt32FromIDQuery(r, cfg.E.Sublocations, QpnSublocation, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = subLocID
		locContextType = string(ViewSourceTypeSublocation)
	}

	areaID, err := nullInt32FromIDQuery(r, cfg.E.Areas, QpnArea, i.queryLookup)
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
